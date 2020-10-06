## POST product
```sh
curl -i -H "Content-Type: application/json" -X POST \
-d '{"activate_date":"2020-04-24T02:27:37Z","brand":"blockbit","detail":"eyJjYXRlZ29yeSI6IndoaXRlIHdpbmUiLCJjb250ZW50IjoiMjQlIiwib3JpZ24iOiJBdXN0cmFsaWEifQ==","name":"blockbit wine","organization":"blockbit","product_code":"20123124436","product_date":"2019-06-19","uni_code":"4GIGj93Ghr"}' \
https://rest-api-id.execeast-1.amazonaws.com/staging/products
```

## GET product
```sh
curl https://rest-api-id.execute-api.us-east-1.amazonaws.com/staging/products?uni_code=4GIGj93Ghr
```