package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"

	"gorm.io/gorm"
)

type BillDetailRepository interface {
	GroupBy(result interface{}) error 
}
type billDetailRepository struct {
	db *gorm.DB
}

func (m *billDetailRepository) GroupBy(result interface{}) error {
	var BillDetailSlice []model.BillDetail
	m.db.Find(BillDetailSlice)
	// res := m.db.Model(BillDetailSlice).Select("cast(trans_date) as date, qty as total").Group("date").Find(result)
	res := m.db.Raw("")
	// res := m.db.Model("t_bill_detail").Select("cast(trans_date) as date, sum(t_bill_detail.qty*menu_price.price) as total").Joins("join menu_price on menu_price.id=menu_price_id").Group("date").Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func NewBillDetailRepository(db *gorm.DB) BillDetailRepository {
	repo := new(billDetailRepository)
	repo.db = db
	return repo
}
