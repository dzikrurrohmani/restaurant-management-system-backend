package manager

import (
	"livecode-wmb-rest-api/repository"
)

type RepositoryManager interface {
	// ProductRepo Disini kumpulan semua repo dalam 1 project yang dibuat
	BillRepo() repository.BillRepository
	CustomerRepo() repository.CustomerRepository
	DiscountRepo() repository.DiscountRepository
	MenuPriceRepo() repository.MenuPriceRepository
	MenuRepo() repository.MenuRepository
	TableRepo() repository.TableRepository
	TransTypeRepo() repository.TransTypeRepository
	BillDetailRepo() repository.BillDetailRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func (r *repositoryManager) DiscountRepo() repository.DiscountRepository {
	return repository.NewDiscountRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuPriceRepo() repository.MenuPriceRepository {
	return repository.NewMenuPriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuRepo() repository.MenuRepository {
	return repository.NewMenuRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TableRepo() repository.TableRepository {
	return repository.NewTableRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TransTypeRepo() repository.TransTypeRepository {
	return repository.NewTransTypeRepository(r.infra.SqlDb())
}

func (r *repositoryManager) BillDetailRepo() repository.BillDetailRepository {
	return repository.NewBillDetailRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
