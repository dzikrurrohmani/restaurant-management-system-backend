package repository

import (
	"context"
	"errors"
	"fmt"
	"livecode-wmb-2/model"
	"livecode-wmb-2/service"
	"livecode-wmb-2/utils"
	"log"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(bill *model.Bill, table *model.Table) error
	FindBy(by map[string]interface{}, preloads []string) ([]model.Bill, error)
	FindAll() ([]model.Bill, error)
	UpdateBy(bill *model.Bill, by map[string]interface{}) error
	LopeiPayment(lopeiId int32, amount float32) error 
}
type billRepository struct {
	db     *gorm.DB
	client service.LopeiPaymentClient
}

func (m *billRepository) Create(bill *model.Bill, table *model.Table) error {
	// Transactional
	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := m.db.Create(bill).Error
	if err != nil {
		if e := tx.Rollback().Error; e != nil {
			return e
		}
		return err
	}
	if table != nil {
		err = m.db.Model(table).Updates(map[string]interface{}{"is_available": "false"}).Error
		if err != nil {
			if e := tx.Rollback().Error; e != nil {
				return e
			}
			return err
		}
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (m *billRepository) FindBy(by map[string]interface{}, preloads []string) ([]model.Bill, error) {
	Stmt := m.db
	for _, preload := range preloads {
		Stmt = Stmt.Preload(preload)
	}
	var bills []model.Bill
	result := Stmt.Where(by).Find(&bills)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return bills, nil
}

func (m *billRepository) FindAll() ([]model.Bill, error) {
	var bills []model.Bill
	result := m.db.Find(&bills)
	if err := result.Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (m *billRepository) UpdateBy(bill *model.Bill, by map[string]interface{}) error {
	result := m.db.Model(bill).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (m *billRepository) LopeiPayment(lopeiId int32, amount float32) error {
	payment, err := m.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeiId: lopeiId,
		Amount:  amount,
	})
	if err != nil {
		log.Fatalln("Error when calling do payment..", err)
	}
	if payment.Error != nil {
		return utils.InsufficientBallanceError()
	}
	fmt.Println(amount, payment.Error)
	return nil
}

func NewBillRepository(db *gorm.DB, client service.LopeiPaymentClient) BillRepository {
	repo := new(billRepository)
	repo.db = db
	repo.client = client
	return repo
}
