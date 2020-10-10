package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gufranmirza/microservice-proto/proto/v1/product/v1product"
)

func main() {
	resp, err := http.Get("http://localhost:8001/provider-api/v1/products/1000")
	if err != nil {
		fmt.Printf("Failed to reterive product data from API with err: %v \n", err)
		os.Exit(1)
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	product := &v1product.Product{}
	err = jsonpb.Unmarshal(resp.Body, product)
	if err != nil {
		fmt.Printf("Failed to deserialize product data %s from API with err: %v \n", resp.Body, err)
		os.Exit(1)
	}

	fmt.Println("#ProductID: ", product.Id)
	fmt.Println("#Name: ", product.Name)
	fmt.Println("#Description: ", product.Description)
	fmt.Println("#Manufacturer: ", product.Manufacturer)
	fmt.Println("#Price: ", product.Price)
	fmt.Println("#InStock: ", product.InStock)
	fmt.Println("#Category: ", product.Category)
}
