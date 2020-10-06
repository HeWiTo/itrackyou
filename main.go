package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type product struct {
	ActivateDate string `json:"activate_date"`
	Brand        string `json:"brand"`
	Detail       string `json:"detail"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	ProductCode  string `json:"product_code"`
	ProductDate  string `json:"product_date"`
	UniCode      string `json:"uni_code"`
}

func show() (*product, error) {
	prd := &product{
		ActivateDate: "2020-04-24T02:27:37Z",
		Brand:        "blockbit",
		Detail:       "eyJjYXRlZ29yeSI6IndoaXRlIHdpbmUiLCJjb250ZW50IjoiMjQlIiwib3JpZ24iOiJBdXN0cmFsaWEifQ==",
		Name:         "blockbit wine",
		Organization: "blockbit",
		ProductCode:  "20123124436",
		ProductDate:  "2019-06-19",
		UniCode:      "4GIGj93Ghr",
	}

	return prd, nil
}

func main() {
	lambda.Start(show)
}
