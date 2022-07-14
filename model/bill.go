package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Bill struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"billId"`
	TransDate   time.Time    `json:"transDate"`
	CustomerID  uint         `json:"customerId"`
	TableID     *uint        `json:"tableId"`
	TransTypeID string       `json:"transtypeId"`
	BillDetails []BillDetail `json:"billBillDetails"`
	Customer    Customer     `json:"-"`
}

func (Bill) TableName() string {
	return "t_bill"
}

func (c *Bill) ToString() (string, error) {
	bill, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(bill), nil
}
