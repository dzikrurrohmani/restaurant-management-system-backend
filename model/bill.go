package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Bill struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	TransDate   time.Time
	CustomerID  uint
	TableID     uint
	TransTypeID string
	BillDetails []BillDetail
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
