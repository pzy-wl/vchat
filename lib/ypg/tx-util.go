package ypg

import "github.com/vhaoran/vchat/lib/ylog"

func Tx(callback func() error) error {
	tx := X.Begin()
	defer func() {
		if err := recover(); err != nil {
			ylog.Error(err)
			tx.Rollback()
		}
	}()

	err := callback()
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}
