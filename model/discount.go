package model

import (
	"encoding/json"
	"errors"
)

type Discount struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Description string
	Pct         uint
	// Customers   []Customer `gorm:"many2many:m_customer_discount"`
}

func (Discount) TableName() string {
	return "m_discount"
}

func (c *Discount) ToString() (string, error) {
	discount, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(discount), nil
}
