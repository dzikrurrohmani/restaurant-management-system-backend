package model

import (
	"encoding/json"
	"errors"
)

type MenuPrice struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"menupriceId"`
	MenuID uint    `json:"menuId"`
	Price  float32 `gorm:"not null" json:"price"`
}

func (MenuPrice) TableName() string {
	return "m_menu_price"
}

func (c *MenuPrice) ToString() (string, error) {
	menuPrice, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(menuPrice), nil
}
