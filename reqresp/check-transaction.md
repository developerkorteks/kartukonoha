curl 'https://putri-veronica.my.id/api/user/limited/xl/check-transaction.json' \
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
  --data-raw '{"trx_id":"136b9b37-32ef-45a7-8d1f-750d5d194585"}'



{
    "statusCode": 200,
    "message": "SUCCESS",
    "success": true,
    "data": {
        "trx_id": "136b9b37-32ef-45a7-8d1f-750d5d194585",
        "name": "[Method E-Wallet] Bonus Flex Kuota Utama 7GB",
        "code": "BONUS_FLEX_KUOTAUTAMA_7GB_E",
        "price": 500,
        "admin_fee": 0,
        "total_price": 500,
        "destination_msisdn": "6287786388052",
        "time_date": "2025-08-25 08:03 WIB",
        "sn_only": "136b9b37-32ef-45a7-8d1f-750d5d194585",
        "sn_and_info": "00020101021226650013CO.XENDIT.WWW01189360084800000039760215qtQFhXV3crrxttp0303UME51370014ID.CO.QRIS.WWW0215ID2025392088550520450995303360540410005802ID5902XL6015JAKARTA SELATAN61051295062290525TlMORY4QSWi4KerHF0L8qu2fN63045F94",
        "status": 1,
        "is_refunded": 0,
        "has_parsial_refund": false,
        "refund_amount": 0,
        "refund_reason": "",
        "balance_before_transaction": 2637079,
        "balance_after_transaction": 2636579,
        "rc": "00",
        "rc_message": "Transaksi Sukses",
        "channel": "self",
        "channel_transaction_code": "1414226",
        "have_deeplink": false,
        "deeplink_url": "",
        "is_qris": true,
        "qris_data": {
            "qr_code": "00020101021226650013CO.XENDIT.WWW01189360084800000039760215qtQFhXV3crrxttp0303UME51370014ID.CO.QRIS.WWW0215ID2025392088550520450995303360540410005802ID5902XL6015JAKARTA SELATAN61051295062290525TlMORY4QSWi4KerHF0L8qu2fN63045F94",
            "payment_expired_at": 1756084119,
            "remaining_time": 204
        }
    }
}
