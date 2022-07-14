package model

import (
	"encoding/json"
	"errors"
)

type TransType struct {
	ID          string `gorm:"primaryKey" json:"transtypeId"`
	Description string `json:"transtypeDescription"`
	Bills       []Bill `json:"-"`
}

func (TransType) TableName() string {
	return "m_trans_type"
}

func (c *TransType) ToString() (string, error) {
	transType, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return "", errors.New("error in converting to json form")
	}
	return string(transType), nil
}
