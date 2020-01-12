package yamq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/weihaoranW/vchat/lib/yconfig"
	"go.uber.org/atomic"
	"log"
	"sync"
	"time"
)

type (
	AMqWorker struct {
		sync.RWMutex
		//当前已连接数量
		count atomic.Int64
		//最小数量
		min atomic.Int64

		//最大数量
		max atomic.Int64
		//chan,消息传递队列
		queue chan *AMqData

		config *yconfig.RabbitConfig
	}

	AMqData struct {
		Queue string
		//0/1/2
		Qos byte
		//must be bytes or string
		Data interface{}
	}
)

const (
	chanLength = 2000000
)

func (r *AMqWorker) NewPoolClient(config *yconfig.RabbitConfig) error {
	r.Lock()
	r.config = config
	r.queue = make(chan *AMqData, chanLength)
	r.Unlock()

	r.min.Store(int64(config.PoolMin))
	r.max.Store(int64(config.PoolMax))
	r.count.Store(0)

	go r.scan(*config)
	return r.createOne(*config)
}

func (r *AMqWorker) scan(cfg yconfig.RabbitConfig) {
	for {
		count := r.count.Load()
		min := r.min.Load()
		max := r.max.Load()

		//log.Println("##### emq pool count:", count, " max:", max, " len(chan)", len(r.queue))
		if count < min {
			r.createOne(cfg)
			continue
		}

		//如果长度过半，则加一个长度
		len := len(r.queue)
		if int64(len) > count*5 && count < max {
			//ylog.SysLogN().WarnF("emq连接池数量需要增加 count: %d len: %d  max: %d", count, len, max) //log.Printf("emq连接池数量需要增加 count: %d len: %d, max: %d", count, len, max)
			r.createOne(cfg)
			continue
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func (r *AMqWorker) createOne(cfg yconfig.RabbitConfig) error {
	cnt, err := r.GetCntDirect(cfg)
	//
	if err != nil {
		return errors.New("没有获取到emq连接")
	}
	ch, err := cnt.Channel()
	if err != nil {
		return errors.New("没有获取到rabbit连接")
	}

	r.count.Inc()
	log.Println("mq pool added,len (", r.count.Load())

	go func() {
		//出错时，释放连接数量
		defer func() {
			r.count.Dec()
			ch.Close()
			cnt.Close()
			log.Println("##### mq pool close,len(", r.count.Load(), ")")
		}()

		//start := time.Now()
		endLoop := false
		for !endLoop {
			select {
			case bean, ok := <-r.queue:
				if ok && bean != nil {
					success := false
					for i := 0; i < 3; i++ {
						if err := publishWrap(ch, bean); err == nil {
							success = true
							break
						}
						time.Sleep(50 * time.Millisecond)
					}

					// 重新放入对列中，进行发送
					if !success {
						r.Publish(bean.Queue, bean.Qos, bean.Data)
					}
					continue
				}
			case <-time.After(time.Minute * 1):
				log.Println("mq wait timeout-->")
				if r.count.Load() > r.min.Load() {
					endLoop = true
				}
			}
		}
	}()
	return nil
}

func (r *AMqWorker) getSleepDuration(connCount int64) time.Duration {
	if connCount <= 50 {
		return time.Millisecond * 500
	}

	if connCount <= 100 {
		return time.Second * 1000
	}
	if connCount > 1000 {
		return time.Second * 10000
	}
	return time.Duration(connCount/100) * time.Second
}

func (r *AMqWorker) Publish(topic string, qos byte, msg interface{}) error {
	bean := &AMqData{
		Queue: topic,
		Qos:   qos,
		Data:  msg,
	}

	select {
	case r.queue <- bean:
	case <-time.After(100 * time.Millisecond):
		return errors.New("MqWorker) PublishQos发送消息超时")
	}
	return nil
}

func (r *AMqWorker) Subscribe(topic string, handler func(body *amqp.Delivery) error) (cnt *amqp.Connection, err error) {
	//---------从池中扑出来一个，不归还的mq client------------
	cnt, err = r.GetCntDirect(*r.config)
	if cnt == nil {
		return nil, fmt.Errorf("emqx Subscribe 无法再建立连接")
	}

	//订阅并返回，有多个订阅时，可以直接在返回结果上操作
	err = consumeWrap(topic, handler)
	return cnt, err
}

// 根据配置，直接生成一个新的连接
func (r *AMqWorker) GetCntDirect(cfg yconfig.RabbitConfig) (cnt *amqp.Connection, err error) {
	// lock ??
	cnt, err = getCntOfRabbitMQ(cfg.Url)
	return
}
