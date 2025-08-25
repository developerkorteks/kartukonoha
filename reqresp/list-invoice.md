curl 'https://putri-veronica.my.id/api/user/invoice/list.json' \
  -H 'accept: */*' \
  -H 'accept-language: en-US,en;q=0.6' \
  -H 'content-type: application/json' \
  -H 'origin: https://putri-veronica.my.id' \
  -H 'priority: u=1, i' \
  -H 'referer: https://putri-veronica.my.id/dashboard/my-invoices' \
  -H 'sec-ch-ua: "Brave";v="137", "Chromium";v="137", "Not/A)Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-gpc: 1' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36' \
  -H 'x-token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJuYWJpbCIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJBPT1JNU82TkRaWElXMWlwdFoxTURNelkiLCJpYXQiOjE3NTYxMTMzNjgsImV4cCI6MTc1NjE0MjE2OH0.sGDPK1uX-n498VrDXT7VnPbUfNz8GY7GNQYoAJKCqNTOtTwG0LUepvAY28e4ltZNLoM4Y2S65Ou7rYNlamwHuFHcmA09pLczOGwpikMj8wSpNKeUwgkoAvVAs7Iua0Et5EEQUve7Q84z-8g5YkNbUSIAGbMqXG8N7c1HCB8LzHgge7hyD0b0vC4bSo-vyPipi4Ut6x2kh3p0Gh55W-TaSEjEpq79jO8egx6dDGfsp32-NS1oy9P1l-la741uhCq9QXQ0Ht_-fU2H5657N2TM6-j2oxdfSxSxvwjT7rrsv3zI1OKRzXDsv-DXbFeaSKmSgHMxEN3Pl-CyG8Dcnwv2pA' \
  --data-raw '{"search":"","last_created_at":0,"limit":20}'


{
    "statusCode": 200,
    "message": "Get invoice list success!",
    "success": true,
    "pagination": {
        "totalRecord": 2,
        "limit": 20,
        "loadmore": false,
        "last_created_at": 0
    },
    "data": [
        {
            "amount": 25000,
            "id": "2dc60519-2d4e-4522-b631-dfd6002bea88",
            "invoice_id": "INV-MEQCFKC5-VHNPF29",
            "title": "Pembelian Topup Balance Credits",
            "due_date": "1756338764980",
            "paid_at": "1756079626129",
            "metadata": {
                "item": [
                    {
                        "name": "Pembelian Topup Balance Credits",
                        "quantity": 1,
                        "price": "25000.00"
                    }
                ],
                "customer": {
                    "firstName": "nabil",
                    "lastName": "-",
                    "email": "kortekslol@gmail.com",
                    "phone": "+6287817739901"
                }
            },
            "invoice_status_id": "paid",
            "created_at": "1756079564980",
            "modified_at": "1756079626129",
            "created_at_month": "August",
            "created_at_year": 2025,
            "username": "kortkeks",
            "fullname": "nabil",
            "email": "kortekslol@gmail.com",
            "phone": "+6287817739901",
            "payments": [
                {
                    "amount": 25000,
                    "id": "2e0bee9c-2454-4a09-86c1-ce797317ddde",
                    "payment_status_id": "success",
                    "payment_method_id": "duitku",
                    "reference_type": "invoice",
                    "reference_id": "INV-MEQCFKC5-VHNPF29",
                    "external_id": "D1959225C57ELCQYRVR6QMQ",
                    "publisher_order_id": "SQ252ILXXVBL7XZVV16",
                    "payment_link": "https://app-prod.duitku.com/redirect_checkout?reference=D1959225C57ELCQYRVR6QMQ",
                    "created_at": "1756079565209",
                    "expired_at": "1756165965112",
                    "completed_at": "1756079626121"
                }
            ]
        },
        {
            "amount": 24000,
            "id": "1c4bf925-51ef-4aae-8568-615a0495f9af",
            "invoice_id": "INV-MEQCEV3M-2JRF1H1",
            "title": "Pembelian Topup Balance Credits",
            "due_date": "1756338732274",
            "paid_at": null,
            "metadata": {
                "item": [
                    {
                        "name": "Pembelian Topup Balance Credits",
                        "quantity": 1,
                        "price": "24000.00"
                    }
                ],
                "customer": {
                    "firstName": "nabil",
                    "lastName": "-",
                    "email": "kortekslol@gmail.com",
                    "phone": "+6287817739901"
                }
            },
            "invoice_status_id": "unpaid",
            "created_at": "1756079532273",
            "modified_at": null,
            "created_at_month": "August",
            "created_at_year": 2025,
            "username": "kortkeks",
            "fullname": "nabil",
            "email": "kortekslol@gmail.com",
            "phone": "+6287817739901",
            "payments": [
                {
                    "amount": 24000,
                    "id": "b18740b5-7590-4177-94e6-a4e86c85f8e0",
                    "payment_status_id": "pending",
                    "payment_method_id": "duitku",
                    "reference_type": "invoice",
                    "reference_id": "INV-MEQCEV3M-2JRF1H1",
                    "external_id": "D1959225KFDMQWRJP7F7J5G",
                    "publisher_order_id": null,
                    "payment_link": "https://app-prod.duitku.com/redirect_checkout?reference=D1959225KFDMQWRJP7F7J5G",
                    "created_at": "1756079532489",
                    "expired_at": "1756165932403",
                    "completed_at": null
                }
            ]
        }
    ]
}
