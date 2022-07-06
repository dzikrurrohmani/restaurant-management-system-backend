package manager

import (
	bill_usecase "livecode-wmb-rest-api/usecase/bill"
	customer_usecase "livecode-wmb-rest-api/usecase/customer"
	discount_usecase "livecode-wmb-rest-api/usecase/discount"
	menu_usecase "livecode-wmb-rest-api/usecase/menu"
	menuprice_usecase "livecode-wmb-rest-api/usecase/menu_price"
	table_usecase "livecode-wmb-rest-api/usecase/table"
	transtype_usecase "livecode-wmb-rest-api/usecase/trans_type"
)

type UseCaseManager interface {
	CreateDiscountUseCase() discount_usecase.CreateDiscountUseCase
	ReadDiscountUseCase() discount_usecase.ReadDiscountUseCase
	UpdateDiscountUseCase() discount_usecase.UpdateDiscountUseCase
	DeleteDiscountUseCase() discount_usecase.DeleteDiscountUseCase
	CreateMenuUseCase() menu_usecase.CreateMenuUseCase
	ReadMenuUseCase() menu_usecase.ReadMenuUseCase
	UpdateMenuUseCase() menu_usecase.UpdateMenuUseCase
	DeleteMenuUseCase() menu_usecase.DeleteMenuUseCase
	CreateMenuPriceUseCase() menuprice_usecase.CreateMenuPriceUseCase
	ReadMenuPriceUseCase() menuprice_usecase.ReadMenuPriceUseCase
	UpdateMenuPriceUseCase() menuprice_usecase.UpdateMenuPriceUseCase
	DeleteMenuPriceUseCase() menuprice_usecase.DeleteMenuPriceUseCase
	CreateTableUseCase() table_usecase.CreateTableUseCase
	ReadTableUseCase() table_usecase.ReadTableUseCase
	UpdateTableUseCase() table_usecase.UpdateTableUseCase
	DeleteTableUseCase() table_usecase.DeleteTableUseCase
	CreateTransTypeUseCase() transtype_usecase.CreateTransTypeUseCase
	ReadTransTypeUseCase() transtype_usecase.ReadTransTypeUseCase
	UpdateTransTypeUseCase() transtype_usecase.UpdateTransTypeUseCase
	DeleteTransTypeUseCase() transtype_usecase.DeleteTransTypeUseCase
	CustomerRegistrationUseCase() customer_usecase.CustomerRegistrationUseCase
	MemberActivationUseCase() customer_usecase.MemberActivationUseCase
	CustomerOrderUseCase() bill_usecase.CustomerOrderUseCase
	CustomerPaymentUseCase() bill_usecase.CustomerPaymentUseCase 
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CreateDiscountUseCase() discount_usecase.CreateDiscountUseCase {
	return discount_usecase.NewCreateDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) ReadDiscountUseCase() discount_usecase.ReadDiscountUseCase {
	return discount_usecase.NewReadDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) UpdateDiscountUseCase() discount_usecase.UpdateDiscountUseCase {
	return discount_usecase.NewUpdateDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) DeleteDiscountUseCase() discount_usecase.DeleteDiscountUseCase {
	return discount_usecase.NewDeleteDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) CreateMenuUseCase() menu_usecase.CreateMenuUseCase {
	return menu_usecase.NewCreateMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) ReadMenuUseCase() menu_usecase.ReadMenuUseCase {
	return menu_usecase.NewReadMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) UpdateMenuUseCase() menu_usecase.UpdateMenuUseCase {
	return menu_usecase.NewUpdateMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) DeleteMenuUseCase() menu_usecase.DeleteMenuUseCase {
	return menu_usecase.NewDeleteMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) CreateMenuPriceUseCase() menuprice_usecase.CreateMenuPriceUseCase {
	return menuprice_usecase.NewCreateMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) ReadMenuPriceUseCase() menuprice_usecase.ReadMenuPriceUseCase {
	return menuprice_usecase.NewReadMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) UpdateMenuPriceUseCase() menuprice_usecase.UpdateMenuPriceUseCase {
	return menuprice_usecase.NewUpdateMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) DeleteMenuPriceUseCase() menuprice_usecase.DeleteMenuPriceUseCase {
	return menuprice_usecase.NewDeleteMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) CreateTableUseCase() table_usecase.CreateTableUseCase {
	return table_usecase.NewCreateTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) ReadTableUseCase() table_usecase.ReadTableUseCase {
	return table_usecase.NewReadTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) UpdateTableUseCase() table_usecase.UpdateTableUseCase {
	return table_usecase.NewUpdateTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) DeleteTableUseCase() table_usecase.DeleteTableUseCase {
	return table_usecase.NewDeleteTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) CreateTransTypeUseCase() transtype_usecase.CreateTransTypeUseCase {
	return transtype_usecase.NewCreateTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) ReadTransTypeUseCase() transtype_usecase.ReadTransTypeUseCase {
	return transtype_usecase.NewReadTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) UpdateTransTypeUseCase() transtype_usecase.UpdateTransTypeUseCase {
	return transtype_usecase.NewUpdateTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) DeleteTransTypeUseCase() transtype_usecase.DeleteTransTypeUseCase {
	return transtype_usecase.NewDeleteTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) CustomerRegistrationUseCase() customer_usecase.CustomerRegistrationUseCase {
	return customer_usecase.NewCustomerRegistrationUseCase(u.repoManager.CustomerRepo())
}
func (u *useCaseManager) MemberActivationUseCase() customer_usecase.MemberActivationUseCase {
	return customer_usecase.NewMemberActivationUseCase(u.repoManager.CustomerRepo(), u.repoManager.DiscountRepo())
}
func (u *useCaseManager) CustomerOrderUseCase() bill_usecase.CustomerOrderUseCase {
	return bill_usecase.NewCustomerOrderUseCase(u.repoManager.BillRepo(), u.repoManager.TableRepo())
}
func (u *useCaseManager) CustomerPaymentUseCase() bill_usecase.CustomerPaymentUseCase  {
	return bill_usecase.NewCustomerPaymentUseCase(u.repoManager.BillRepo(), u.repoManager.TableRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
