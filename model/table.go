package model

import (
	"encoding/json"
	"errors"
)

type Table struct {
	BaseModel       BaseModel `gorm:"embedded"`
	TableDescription string
	IsAvailable     bool `gorm:"default:true"`
	Bills           []Bill
}

func (Table) TableName() string {
	return "m_table"
}

func (c *Table) ToString() (string, error) {
	table, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(table), nil
}
