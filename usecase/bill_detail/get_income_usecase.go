package billdetail_usecase

import (
	"livecode-wmb-2/repository"
)

type GetIncomeUseCase interface {
	GetDailyIncome(result interface{}) error
}

type readBillDetailUseCase struct {
	billDetailRepo repository.BillDetailRepository
}

func (m *readBillDetailUseCase) GetDailyIncome(result interface{}) error {
	return m.billDetailRepo.GroupBy(result)
}

func NewGetIncomeUseCase(billDetailRepo repository.BillDetailRepository) GetIncomeUseCase {
	usecase := new(readBillDetailUseCase)
	usecase.billDetailRepo = billDetailRepo
	return usecase
}
