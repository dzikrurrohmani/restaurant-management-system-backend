package model

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type MenuPrice struct {
	BaseModel BaseModel `gorm:"embedded"`
	MenuID    sql.NullInt64
	Price     float32 `gorm:"not null"`
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
