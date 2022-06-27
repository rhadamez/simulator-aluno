package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rhadamez/simulator-aluno/application/route"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route.Route{
		ID:       "1",
		ClientId: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
	fmt.Println(stringjson[1])
	fmt.Println(stringjson[2])
	fmt.Println(stringjson[3])
	fmt.Println(stringjson[4])
}
