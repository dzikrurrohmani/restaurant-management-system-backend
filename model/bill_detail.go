package model

import (
	"encoding/json"
	"errors"
)

type BillDetail struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"billdetailId"`
	BillID      uint      `json:"billId"`
	MenuPriceID uint      `json:"menupriceId"`
	Qty         float32   `json:"billdetailQty"`
	MenuPrice   MenuPrice `json:"-"`
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
