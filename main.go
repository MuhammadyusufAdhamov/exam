package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	PostgresUser     = "postgres"
	PostgresDatabase = "exam"
	PostgresPassword = "7"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s database=%s password=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresDatabase,
		PostgresPassword,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}

	dbManager := NewDBProductManager(db)

	//_, err = dbManager.CreateProducts(&Products{
	//	Name:     "Redmi 9T",
	//	Price:    200,
	//	ImageUrl: "mi.com/redmi9t",
	//	Images: []*ProductImages{
	//		{
	//			ImageUrl:       "mi.com/redmi9t1",
	//			SequenceNumber: 1,
	//		},
	//		{
	//			ImageUrl:       "mi.com/redmi9t2",
	//			SequenceNumber: 2,
	//		},
	//		{
	//			ImageUrl:       "mi.com/redmi9t3",
	//			SequenceNumber: 3,
	//		},
	//	},
	//})
	//if err != nil {
	//	log.Fatalf("failed to create product: %v", err)
	//}
	//
	//product, err := dbManager.GetProducts(1)
	//if err != nil {
	//	log.Fatalf("failed to get product: %v", err)
	//}
	//
	//PrintProducts(product)

	//resp, err := dbManager.GetAllProducts(&GetProductsParams{
	//	Limit: 10,
	//	Page:  1,
	//})
	//
	//if err != nil {
	//	log.Fatalf("failed to get product: %v", err)
	//}
	//
	//fmt.Printf("%v", resp)
	//
	//err = dbManager.UpdateProduct(&Products{
	//	Id:       2,
	//	Name:     "malibu",
	//	Price:    40000.0,
	//	ImageUrl: "malibu.com",
	//	Images: []*ProductImages{
	//		{
	//			ImageUrl:       "malibu.com/image1",
	//			SequenceNumber: 1,
	//		},
	//		{
	//			ImageUrl:       "malibu.com/image2",
	//			SequenceNumber: 2,
	//		},
	//	},
	//})
	//if err != nil {
	//	log.Fatalf("failed to update product: %v", err)
	//}

	err = dbManager.DeleteProduct(1)
	if err != nil {
		log.Fatalf("failed to delete product: %v", err)
	}
}

func PrintProducts(product *Products) {
	fmt.Println("---------Product--------")
	fmt.Println("Id: ", product.Id)
	fmt.Println("Name: ", product.Name)
	fmt.Println("Price: ", product.Price)
	fmt.Println("Image url: ", product.ImageUrl)
	fmt.Println("Created at: ", product.CreatedAt)
}
