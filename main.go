package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
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

var unicodeRegexp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the `uni_code` query string parameter from the request and validate it.
	unicode := req.QueryStringParameters["uni_code"]
	if !unicodeRegexp.MatchString(unicode) {
		return clientError(http.StatusBadRequest)
	}

	// Fetch the product record from the dynamodb based on the uni_code value.
	prod, err := getItem(unicode)
	if err != nil {
		return serverError(err)
	}
	if prod == nil {
		return clientError(http.StatusNotFound)
	}

	// marshal the product record into JSON.
	js, err := json.Marshal(prod)
	if err != nil {
		return serverError(err)
	}

	// Return a response with a 200 OK status and the JSON product record as the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	prd := new(product)
	err := json.Unmarshal([]byte(req.Body), prd)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if !unicodeRegexp.MatchString(prd.UniCode) {
		return clientError(http.StatusBadRequest)
	}
	if prd.Detail == "" || prd.Brand == "" || prd.Name == "" || prd.Organization == "" || prd.ProductCode == "" || prd.ProductDate == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(prd)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/products?uni_code=%s", prd.UniCode)},
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}
