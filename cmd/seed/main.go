package main

import (
	"kaco/model"
	"os"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../../.env")
	model.Database(os.Getenv("MYSQL_DSN"))

	product1 := model.Product{
		Name:        "香菇",
		Price:       200,
		Category:    "vegetable",
		ChCategory:  "蔬菜",
		CreatedAt:   "2023-1-05 21:13",
		HotSpot:     "天然香菇",
		Description: "xxxxxxxx",
	}

	model.DB.Create(&product1)

	product2 := model.Product{
		Name:        "香菜",
		Price:       200,
		Category:    "vegetable",
		ChCategory:  "蔬菜",
		CreatedAt:   "2023-1-05 21:13",
		HotSpot:     "天然香菜",
		Description: "xxxxxxxx",
	}

	model.DB.Create(&product2)

}
