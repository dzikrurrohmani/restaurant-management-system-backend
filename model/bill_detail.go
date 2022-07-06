package model

import (
	"encoding/json"
	"errors"
)

type BillDetail struct {
	BaseModel   BaseModel `gorm:"embedded"`
	BillID      uint
	MenuPriceID uint
	Qty         float32
	MenuPrice   MenuPrice
}

func (BillDetail) TableName() string {
	return "t_bill_detail"
}

func (c *BillDetail) ToString() (string, error) {
	billDetail, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(billDetail), nil
}
