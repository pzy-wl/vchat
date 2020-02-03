package ypg

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/common/ypage"
)

func Page(db *gorm.DB, l interface{}, src *ypage.PageBean) (*ypage.PageBean, error) {
	if src == nil {
		return nil, errors.New("condition is nil")
	}
	if !reflectUtils.IsSlice(l) {
		return nil, errors.New("passed l is not slice")
	}
	ptr, err := reflectUtils.MakeSliceElemPtr(l)
	if err != nil {
		return nil, errors.New("passed l is not slice")
	}

	//-------- -----------------------------
	src.ValidateAdjust()

	exp, p := new(ypage.SqlWhere).GetWhere(src.Where)

	err = db.Order(ypage.GetSort(src.Sort)).
		Limit(src.RowsPerPage).
		Offset(src.GetSkip()).
		Where(exp, p...).
		Find(l).Error
	if err != nil {
		return nil, err
	}

	//
	if src.RowsCount <= 0 {
		count := 0
		err = db.Model(ptr).Where(exp, p...).Count(&count).Error
		src.RowsCount = int64(count)
	}

	src.PagesCount = src.GetPagesCount()
	src.Data = l

	return src, nil
}