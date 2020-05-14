package ypg

import (
	"errors"

	"github.com/jinzhu/gorm"
)

func DBErr(db *gorm.DB) error {
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected <= 0 {
		return errors.New("新增失败，示知原因")
	}
	return nil
}
