package main

import (
	"livecode-wmb-rest-api/delivery"
)

func main() {
	delivery.Server().Run()
	// config := config.NewConfigDB()
	// db := config.DbConn()
	// defer config.DbClose()

	// // migration.MigrateDb(db)

	// // // SOAL 1
	// // menuRepo := repository.NewMenuRepository(db)
	// // command.MenuCreate(menuRepo)
	// // command.MenuRead(menuRepo)
	// // command.MenuUpdate(menuRepo)
	// // command.MenuDelete(menuRepo)

	// // // SOAL 2
	// // menuPriceRepo := repository.NewMenuPriceRepository(db)
	// // command.MenuPriceCreate(menuPriceRepo)
	// // command.MenuPriceRead(menuPriceRepo)
	// // command.MenuPriceUpdate(menuPriceRepo)
	// // command.MenuPriceDelete(menuPriceRepo)

	// // // SOAL 3
	// tableRepo := repository.NewTableRepository(db)
	// // command.TableCreate(tableRepo)
	// // command.TableRead(tableRepo)
	// // command.TableUpdate(tableRepo)
	// // command.TableDelete(tableRepo)

	// // // SOAL 4
	// // transTypeRepo := repository.NewTransTypeRepository(db)
	// // command.TransTypeCreate(transTypeRepo)
	// // command.TransTypeRead(transTypeRepo)
	// // command.TransTypeUpdate(transTypeRepo)
	// // command.TransTypeDelete(transTypeRepo)

	// // // SOAL 5
	// discountRepo := repository.NewDiscountRepository(db)
	// // command.DiscountCreate(discountRepo)
	// // command.DiscountRead(discountRepo)
	// // command.DiscountUpdate(discountRepo)
	// // command.DiscountDelete(discountRepo)

	// // // SOAL 6
	// customerRepo := repository.NewCustomerRepository(db)
	// // command.RegistCustomer(customerRepo)

	// // // SOAL 7
	// command.ActivateMember(customerRepo, discountRepo)

	// // // SOAL 8
	// billRepo := repository.NewBillRepository(db)
	// // command.BillCreate(billRepo, tableRepo)

	// // // SOAL 9
	// command.BillPayment(billRepo, tableRepo)
}

// DROP SCHEMA public CASCADE;
// CREATE SCHEMA public;
