package model

import (
	"encoding/json"
	"errors"
)

type Table struct {
	ID               uint `gorm:"primaryKey;autoIncrement" json:"tableId"`
	TableDescription string `json:"tableDescription"`
	IsAvailable      bool `gorm:"default:false" json:"tableAvailability"`
	Bills            []Bill `json:"-"`
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
