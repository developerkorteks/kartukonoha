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
  -H 'x-token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJuYWJpbCIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJRPT1Bek02TkRaWElXMWlwdFozTmpPVEEiLCJpYXQiOjE3NTYwODMwNTcsImV4cCI6MTc1NjExMTg1N30.StZXYHfBrNKYDfgxAFvArE6xcqYI2E1DdUeZHmHdjLP5FR6OKxDAkH6yEGGYyzP_9_qegLk740t8NILBp8jE9xminroTfB3b87vk3rvFFPktxwta-LJaCWedDrYUgygiukKE6yRjMuFI11EST6UKzTSaw8TzueBYUp0ZA2dtaR51uDuKfM_zCqBxRW-k_ZpNxz3hZ1EfjIQ0Vc1gvdofGKf9iD5CnWHGookWYwVJcw1s_UqrjqrW6HyOv51buyImBZZnuuRmH_xpGt5jLm4aI6pm6vnQcKdK0shxBPPcltxePe2QnugHlYOLZDENr7tX9ov35pstOF9_swa1JGoWEw' \
  --data-raw '{"package_code":"BONUS_FLEX_KUOTAUTAMA_7GB_E","phone":"087786388052","access_token":"1097690:ec321a89-6593-4b01-9a9e-383744301283","payment_method":"QRIS"}'


{
    "statusCode": 200,
    "message": "Silakan scan/upload QR di bawah ke E-Wallet/Mbanking Anda dan segera lakukan pembayaran! Abaikan jika sudah melakukan pembayaran.",
    "success": true,
    "data": {
        "msisdn": "6287786388052",
        "package_code": "BONUS_FLEX_KUOTAUTAMA_7GB_E",
        "package_name": "[Method E-Wallet] Bonus Flex Kuota Utama 7GB",
        "package_processing_fee": 500,
        "trx_id": "136b9b37-32ef-45a7-8d1f-750d5d194585",
        "have_deeplink": false,
        "deeplink_data": {
            "payment_method": "QRIS",
            "deeplink_url": ""
        },
        "is_qris": true,
        "qris_data": {
            "qr_code": "00020101021226650013CO.XENDIT.WWW01189360084800000039760215qtQFhXV3crrxttp0303UME51370014ID.CO.QRIS.WWW0215ID2025392088550520450995303360540410005802ID5902XL6015JAKARTA SELATAN61051295062290525TlMORY4QSWi4KerHF0L8qu2fN63045F94",
            "payment_expired_at": 1756084119,
            "remaining_time": 300
        }
    }
}
