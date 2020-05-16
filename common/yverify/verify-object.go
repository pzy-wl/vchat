package yverify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/common/reflectUtils"
)

type (
	VerifyOBJ struct {
		onErrStop bool
		Errs      []string
	}
)

//输入为整形值 或浮点娄，必须大于0
func NewObj(onErrStop ...bool) *VerifyOBJ {
	b := true
	if len(onErrStop) > 0 {
		b = onErrStop[0]
	}

	return &VerifyOBJ{
		onErrStop: b,
		Errs:      make([]string, 0),
	}
}

func (r *VerifyOBJ) hasErr() bool {
	return r.Errs != nil && len(r.Errs) > 0
}

func (r *VerifyOBJ) NotZeroI(name string, src interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsInt(src) {
		return r
	}
	//
	if ii, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if ii == 0 {
			msg := fmt.Sprintf("<%s>值<%d>必须不为0", name, ii)

			r.push(msg)
			return r
		}
	}
	return r
}

func (r *VerifyOBJ) Gt(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		if f2, err := strconv.ParseFloat(fmt.Sprint(dst), 64); err == nil {
			if !(f1 > f2) {
				msg := fmt.Sprintf("<%s>值<%f>必须大于<%f>", name, f1, f2)

				r.push(msg)
				return r
			}
		}
	}
	return r
}
func (r *VerifyOBJ) GtI(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if i1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if i2, err := strconv.ParseInt(fmt.Sprint(dst), 10, 64); err == nil {
			if !(i1 > i2) {
				msg := fmt.Sprintf("<%s>值<%d>必须大于<%d>", name, i1, i2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) Gte(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		if f2, err := strconv.ParseFloat(fmt.Sprint(dst), 64); err == nil {
			if !(f1 >= f2) {
				msg := fmt.Sprintf("<%s>值<%f>必须大于或等于<%f>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) GteI(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if f2, err := strconv.ParseInt(fmt.Sprint(dst), 10, 64); err == nil {
			if !(f1 >= f2) {
				msg := fmt.Sprintf("<%s>值<%d>必须大于或等于<%d>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点，src<dst
func (r *VerifyOBJ) Lt(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		if f2, err := strconv.ParseFloat(fmt.Sprint(dst), 64); err == nil {
			if !(f1 < f2) {
				msg := fmt.Sprintf("<%s>值<%f>必须小于<%f>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点，src<dst
func (r *VerifyOBJ) LtI(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if f2, err := strconv.ParseInt(fmt.Sprint(dst), 10, 64); err == nil {
			if !(f1 < f2) {
				msg := fmt.Sprintf("<%s>值<%d>必须小于<%d>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) Lte(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		if f2, err := strconv.ParseFloat(fmt.Sprint(dst), 64); err == nil {
			if !(f1 <= f2) {
				msg := fmt.Sprintf("<%s>值<%f>必须小于或等于<%f>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) LteI(name string, src interface{}, dst interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil || dst == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if f2, err := strconv.ParseInt(fmt.Sprint(dst), 10, 64); err == nil {
			if !(f1 <= f2) {
				msg := fmt.Sprintf("<%s>值<%d>必须小于或等于<%d>", name, f1, f2)
				r.push(msg)
				return r
			}
		}
	}
	return r
}

func (r *VerifyOBJ) push(s string) {
	if r.Errs == nil {
		r.Errs = make([]string, 0)
	}
	r.Errs = append(r.Errs, s)
}

func (r *VerifyOBJ) NotEmptyPtr(name string, ptr interface{}) *VerifyOBJ {
	if ptr == nil {
		msg := fmt.Sprintf("<%s>不能为空", name)
		r.push(msg)
		return r
	}

	if !reflectUtils.IsPointer(ptr) {
		return r
	}
	//
	v := reflect.ValueOf(ptr)
	if v.IsNil() {
		msg := fmt.Sprintf("<%s>不能为空", name)
		r.push(msg)
		return r
	}

	return r
}

// l is string,array,slice,amp
func (r *VerifyOBJ) NotEmpty(name string, l interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if l == nil {
		msg := fmt.Sprintf("<%s>不能为空", name)
		r.push(msg)
		return r
	}

	v := reflect.Indirect(reflect.ValueOf(l))
	i := 0
	switch v.Kind() {
	case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
		i = v.Len()
	default:
		i = 0
	}

	if i <= 0 {
		msg := fmt.Sprintf("<%s>不能为空,长度必须大于0", name)
		r.push(msg)
	}
	return r
}

func (r *VerifyOBJ) InSlice(name string, src, l interface{}) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if g.InSlice(src, l) {
		return r
	}
	msg := fmt.Sprintf("<%s>元素不在列表中", name)
	r.push(msg)
	return r
}

//这是自定义的fn，用户可以传入自定义验证结果
func (r *VerifyOBJ) Fn(l ...error) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	for _, err := range l {
		if err != nil {
			r.push(err.Error())
		}
	}
	return r
}

func (r *VerifyOBJ) FnBool(name string, b bool) *VerifyOBJ {
	if r.hasErr() && r.onErrStop {
		return r
	}

	if !b {
		r.push(fmt.Sprintf("<%s>值必须为真值", name))
	}
	return r
}

//这是链式语法验证的结果
func (r *VerifyOBJ) Err() error {
	if len(r.Errs) > 0 {
		s := strings.Join(r.Errs, ";  ")
		return errors.New(s)
	}
	return nil
}
