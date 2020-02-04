package ctl

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/vhaoran/vchat/demo/unit/intf"
	"github.com/vhaoran/vchat/lib/ykit"
	"log"
)

type PBImpl struct {
}

func (r *PBImpl) Exec(ctx context.Context, in *intf.PBIn) (*ykit.Result, error) {
	log.Println("------PBImpl) Exec----", "------------")
	spew.Dump(in)
	log.Println("------end----", "------------")
	return ykit.ROK(in), nil
}
