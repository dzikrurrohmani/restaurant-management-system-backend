package model

import (
	"encoding/json"
	"errors"
)

type Menu struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"menuId"`
	MenuName   string      `gorm:"size:50;not null" json:"menuName"`
	Price      int         `json:"menuPrice"`
	Category   string      `json:"menuCategory"`
	MenuPrices []MenuPrice `json:"-"`
}

func (Menu) TableName() string {
	return "m_menu"
}

func (c *Menu) ToString() (string, error) {
	menu, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(menu), nil
}
