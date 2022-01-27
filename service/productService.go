package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	os2 "os"
	"products/model"
)

func GetProducts(filename string) []model.Product {
	jsonFile, err := os2.Open(filename)

	if err != nil {
		fmt.Println(err)
		os2.Exit(1)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//bs, err := ioutil.ReadFile(filename)

	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(1)
	//}
	var products []model.Product

	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		return nil
	}
	fmt.Println(products)
	fmt.Printf("%+v", products)
	defer func(jsonFile *os2.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)
	return products
}
