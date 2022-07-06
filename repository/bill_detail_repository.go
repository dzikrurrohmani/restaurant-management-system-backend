package repository

import (
	"errors"

	"gorm.io/gorm"
)

type BillDetailRepository interface {
	GroupBy(result interface{}) error
}
type billDetailRepository struct {
	db *gorm.DB
}

func (m *billDetailRepository) GroupBy(result interface{}) error {
	res := m.db.Raw("select date(b.trans_date) as date, sum(tb.qty*mp.price) as total from t_bill b join t_bill_detail tb on b.id=tb.bill_id join m_menu_price mp on tb.menu_price_id=mp.id group by date").Scan(result)
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
