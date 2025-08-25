curl 'https://putri-veronica.my.id/api/user/limited/xl/status-kartu.json' \
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
  -H 'x-token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJrb3J0a2VrcyIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJ3PT1NNU02TlRaWElXMWlwdFozT1ROREEiLCJpYXQiOjE3NTYwNTM1MzMsImV4cCI6MTc1NjA4MjMzM30.FFtPvvJos5Wz-rcNNdIUj2plb_7RHW8baSVl299fw25A2ROZOC-q00MgcEJJBnJCq_KcskUF6AEnPMkVFv8cyieYWEeFELCHmBCLJi2jqkzQVkrPgR7DNmjbpzFnwKK_iMWFiej15o9TbK8gDNLL3b4S7zuU5PZBnraqiL1tm9Zx--K2kotNSOQ225ggI_7ca3BwpDjeKu1nF-iZMbMlgwd69gOkmonnIHHQeU8C2C2CQxfj3KazAv95Ho4_BT2yniptczzYI_xEg_Rk_yt51N1hdFd1LcLaZNJ1zhRCN9wnr6t_tKZRbstZdfWpfk03tdRaCKsNAZlpBMHC_CMBSA' \
  --data-raw '{"access_token":"1097498:7c17952a-343a-49b1-9e2e-4a52b7682b7a"}'



{
    "statusCode": 200,
    "message": "Berhasil mendapatkan info status kartu!",
    "success": true,
    "data": {
        "msisdn": "6287817739901",
        "subscription_status": "Aktif",
        "pulsa": 20,
        "pulsa_real": "Rp20",
        "active_until": "2025-09-18 23:59",
        "location": "Kab. Demak"
    }
}
