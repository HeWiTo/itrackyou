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
	prd, err := getItem("4GIGj93Ghr")
	if err != nil {
		return nil, err
	}

	return prd, nil
}

func main() {
	lambda.Start(show)
}
