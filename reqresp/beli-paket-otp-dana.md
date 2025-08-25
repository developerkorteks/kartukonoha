curl 'https://putri-veronica.my.id/api/user/limited/xl/beli-paket-otp.json' \
  -H 'accept: */*' \
  -H 'accept-language: en-US,en;q=0.6' \
  -H 'content-type: application/json' \
  -H 'origin: https://putri-veronica.my.id' \
  -H 'priority: u=1, i' \
  -H 'referer: https://putri-veronica.my.id/limited/tembak-paket-xl' \
  -H 'sec-ch-ua: "Brave";v="137", "Chromium";v="137", "Not/A)Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-gpc: 1' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36' \
  -H 'x-token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJuYWJpbCIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJBPT1JNU82TkRaWElXMWlwdFoxTURNelkiLCJpYXQiOjE3NTYxMTMzNjgsImV4cCI6MTc1NjE0MjE2OH0.sGDPK1uX-n498VrDXT7VnPbUfNz8GY7GNQYoAJKCqNTOtTwG0LUepvAY28e4ltZNLoM4Y2S65Ou7rYNlamwHuFHcmA09pLczOGwpikMj8wSpNKeUwgkoAvVAs7Iua0Et5EEQUve7Q84z-8g5YkNbUSIAGbMqXG8N7c1HCB8LzHgge7hyD0b0vC4bSo-vyPipi4Ut6x2kh3p0Gh55W-TaSEjEpq79jO8egx6dDGfsp32-NS1oy9P1l-la741uhCq9QXQ0Ht_-fU2H5657N2TM6-j2oxdfSxSxvwjT7rrsv3zI1OKRzXDsv-DXbFeaSKmSgHMxEN3Pl-CyG8Dcnwv2pA' \
  --data-raw '{"package_code":"XL_EDU_2GB_1K_DANA","phone":"6287817739901","access_token":"1097751:9ce75b8d-79bd-4dc0-855a-6c14267e7f5e","payment_method":"DANA"}'



{
    "statusCode": 200,
    "message": "Pastikan kamu punya aplikasi DANA dan segera bayar dengan klik tombol BAYAR di bawah ini melalui DANA.",
    "success": true,
    "data": {
        "msisdn": "6287817739901",
        "package_code": "XL_EDU_2GB_1K_DANA",
        "package_name": "[Metode E-Wallet] Xtra Edukasi 2GB Rp1.000 1 Hari",
        "package_processing_fee": 0,
        "trx_id": "ad77e6ec-a22b-4536-bb08-589c6c9eb891",
        "have_deeplink": true,
        "deeplink_data": {
            "payment_method": "DANA",
            "deeplink_url": "https://m.dana.id/n/cashier/new/checkout?bizNo=20250825111212800110166730365280603&timestamp=1756113836620&originSourcePlatform=IPG&mid=216620000035833566172&did=216650000204712460170&sid=216660000204589345173&sign=kXLYff0cYCwFYSrI9CX3NCmQA96SzKOeEu1wsoNBSZfubebp3RygvwGnkpq5QlrSBYsa%2BvkP6Av4hIjIyS0kyJucduXkvclUxbcPeKgpcJIMyiwEy8971t2iVCPWPOfGh7CY4H47nN2Xch3Ft3HrYtoskby0WooSgl3g01Gw%2BkgIcH4RXhnUKXrz1GB35UAdrLj9Lk6kzDx0Fl9NjAk296UPEmMdaC20vqw%2FmkamUn2HomMTE0G0xTQgAGSobnKPVtgUpFY8%2FRlJHRUaFMS4v9FO%2FU9jpT6iQ6WmrCUpsoeIFGHYvyCEHeCs7Jin8TgjtwP8pF20CzAwRQXLZKJSBQ%3D%3D&forceToH5=false&newRegistrationPage=true"
        },
        "is_qris": false,
        "qris_data": []
    }
}
