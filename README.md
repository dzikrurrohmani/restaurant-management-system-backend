## HOW TO RUN

### SET ENV
```bash
source config/.mig_env.sh # untuk set environment migration
OR
source config/.dev_env.sh # untuk set environment development
```
### RUN PROJECT
```go
go run .
```

## API Spec

### Create Menu
- Request: POST
- Endpoint : `/menu`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body : 
```json
{
  "menuId": 1,
  "menuName": "Nasi Goreng"
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menuId": 1,
    "menuName": "Nasi Goreng",
    "menuPrices": null
  }
}
  ```

### Get Menu
- Request: GET
- Endpoint : `/menu`
- Header :
    - Content-Type: application/json; charset=utf-8

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": [
    {
      "menuId": 1,
      "menuName": "Nasi Goreng",
      "menuPrices": null
    }
  ]
}
  ```

### Update Menu
- Request: PUT
- Endpoint : `/menu`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
  "menuId": 1,
  "menuName": "Nasi Goreng Sedang"
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menuId": 1,
    "menuName": "Nasi Goreng Sedang",
    "menuPrices": null
  }
}
  ```

### Delete Menu
- Request: DELETE
- Endpoint : `/menu`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
    "menuId": 1
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menuId": 1,
    "menuName": "",
    "menuPrices": null
  }
}
  ```

### Create MenuPrice
- Request: POST
- Endpoint : `/menuprice`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body : 
```json
{
  "MenuID": 1,
  "Price": 7000
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menupriceId": 1,
    "menuId": 1,
    "price": 7000
  }
}
  ```

### Get MenuPrice
- Request: GET
- Endpoint : `/menuprice`
- Header :
    - Content-Type: application/json; charset=utf-8

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": [
    {
      "menupriceId": 1,
      "menuId": 1,
      "price": 7000
    }
  ]
}
  ```

### Update MenuPrice
- Request: PUT
- Endpoint : `/menuprice`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
  "menupriceId": 1,
  "menuId": 1,
  "price": 12000
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menupriceId": 1,
    "menuId": 1,
    "price": 12000
  }
}
  ```

### Delete MenuPrice
- Request: DELETE
- Endpoint : `/menuprice`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
  "menupriceId": 1
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "menupriceId": 1,
    "menuId": 0,
    "price": 0
  }
}
  ```

### Customer Registration
- Request: POST
- Endpoint : `/customer`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
	"CustomerName":  "Anak Pertama",
	"MobilePhoneNo": "011111111111"
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "ID": 1,
    "CustomerName": "Anak Pertama",
    "MobilePhoneNo": "011111111111",
    "IsMember": false,
    "Bills": null,
    "Discounts": null
  }
}
  ```

### Member Activation
- Request: PUT
- Endpoint : `/customer`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
  "MobilePhoneNo": "011111111111",
  "DiscountId": 1
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "ID": 1,
    "CustomerName": "Anak Pertama",
    "MobilePhoneNo": "011111111111",
    "IsMember": true,
    "Bills": null,
    "Discounts": null
  }
}
  ```

### Customer Order Eat In
- Request: POST
- Endpoint : `/bill`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
	"CustomerID":  1,
	"TableID":     {
      "Int64": 2,
      "Valid": true
    },
	"TransTypeID": "EI",
	"BillDetails": [
    {
      "MenuPriceID": 1,
	   "Qty": 2
    },
    {
      "MenuPriceID": 2,
	   "Qty": 3
    }
  ]
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "ID": 4,
    "TransDate": "2022-07-06T20:46:21.486750907+07:00",
    "CustomerID": 1,
    "TableID": {
      "Int64": 2,
      "Valid": true
    },
    "TransTypeID": "EI",
    "BillDetails": [
      {
        "ID": 7,
        "BillID": 4,
        "MenuPriceID": 1,
        "Qty": 2,
        "MenuPrice": {
          "menupriceId": 0,
          "menuId": 0,
          "price": 0
        }
      },
      {
        "ID": 8,
        "BillID": 4,
        "MenuPriceID": 2,
        "Qty": 3,
        "MenuPrice": {
          "menupriceId": 0,
          "menuId": 0,
          "price": 0
        }
      }
    ]
  }
}
  ```

### Customer Order Take Away
- Request: POST
- Endpoint : `/bill`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
	"CustomerID":  1,
	"TransTypeID": "TA",
	"BillDetails": [
    {
      "MenuPriceID": 1,
	   "Qty": 2
    },
    {
      "MenuPriceID": 2,
	   "Qty": 3
    }
  ]
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "ID": 1,
    "TransDate": "2022-07-06T20:35:49.32361166+07:00",
    "CustomerID": 1,
    "TableID": {
      "Int64": 0,
      "Valid": false
    },
    "TransTypeID": "TA",
    "BillDetails": [
      {
        "ID": 1,
        "BillID": 1,
        "MenuPriceID": 1,
        "Qty": 2,
        "MenuPrice": {
          "menupriceId": 0,
          "menuId": 0,
          "price": 0
        }
      },
      {
        "ID": 2,
        "BillID": 1,
        "MenuPriceID": 2,
        "Qty": 3,
        "MenuPrice": {
          "menupriceId": 0,
          "menuId": 0,
          "price": 0
        }
      }
    ]
  }
}
  ```

### Customer Payment
- Request: GET
- Endpoint : `/bill`
- Header :
    - Content-Type: application/json; charset=utf-8
- Body :
```json
{
	"BillID":  2
}
```

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": {
    "ID": 2,
    "TransDate": "2022-07-06T20:40:32.049991+07:00",
    "CustomerID": 1,
    "TableID": 2,
    "TransTypeID": "EI",
    "BillDetails": [
      {
        "ID": 3,
        "BillID": 2,
        "MenuPriceID": 1,
        "Qty": 2,
        "MenuPrice": {
          "menupriceId": 1,
          "menuId": 1,
          "price": 12000
        }
      },
      {
        "ID": 4,
        "BillID": 2,
        "MenuPriceID": 2,
        "Qty": 3,
        "MenuPrice": {
          "menupriceId": 2,
          "menuId": 1,
          "price": 17000
        }
      }
    ]
  }
}
  ```

### Get Daily Income
- Request: GET
- Endpoint : `/billdetail`
- Header :
    - Content-Type: application/json; charset=utf-8

Response:
```json
{
  "responseCode": "00",
  "responseMessage": "Success",
  "data": [
    {
      "Date": "2022-07-06T00:00:00Z",
      "Total": 75000
    },
    {
      "Date": "2022-07-05T00:00:00Z",
      "Total": 150000
    },
    {
      "Date": "2022-07-04T00:00:00Z",
      "Total": 75000
    }
  ]
}
  ```