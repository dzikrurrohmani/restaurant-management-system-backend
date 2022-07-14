package model

import (
	"encoding/json"
	"errors"
)

type Customer struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"customerId"`
	CustomerName  string `gorm:"size:50;not null" json:"customerName"`
	MobilePhoneNo string `gorm:"unique; size:13" json:"customerPhone"`
	IsMember      bool   `gorm:"default:false" json:"customerMember"`
	// Bills         []Bill `json:"customerBills"`
	Discounts []Discount `gorm:"many2many:m_customer_discount" json:"-"`
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
