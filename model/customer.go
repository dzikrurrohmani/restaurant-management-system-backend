package model

import (
	"encoding/json"
	"errors"
)

type Customer struct {
	BaseModel     BaseModel `gorm:"embedded"`
	CustomerName  string    `gorm:"size:50;not null"`
	MobilePhoneNo string    `gorm:"unique; size:13"`
	IsMember      bool      `gorm:"default:false"`
	Bills         []Bill
	Discounts     []Discount `gorm:"many2many:m_customer_discount"`
}

func (Customer) TableName() string {
	return "m_customer"
}

func (c *Customer) ToString() (string, error) {
	customer, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(customer), nil
}
