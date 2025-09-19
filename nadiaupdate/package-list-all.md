curl 'https://putri-veronica.my.id/api/user/limited/xl/package-list-all.json' \
  -X 'POST' \
  -H 'accept: */*' \
  -H 'accept-language: en-US,en;q=0.9' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 0' \
  -H 'content-type: application/json' \
  -H 'origin: https://putri-veronica.my.id' \
  -H 'pragma: no-cache' \
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
  -H 'x-token: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJuYWJpbCIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJBPT1rMU42TlRaWElXMWlwdFp4TXpORFEiLCJpYXQiOjE3NTgxNjQ4NzksImV4cCI6MTc1ODE5MzY3OX0.yxj5UfLa1-B1XZg_xp5-6XEGzrA9bvb44U629RecsrTlE6AK1f8uWnqVu15ZQVGopMhiQsS_CuL9oZ_eckkzMUaMnWpfW_qaOi-98_z2NCPsUo6IqFC7iJmcyHWKyCZHAC8uLa6y5AGmNf9H2sd-HQ18SLB1PjJ3ALVDCqxZfV8-k5YB0CA1TA5RPnWlYlx4zvUwQszqElOXMcRccW_WFvcyPnrORohfvCPSL2pZpYCUkg8cBLPu9wY1beU69zPz7Nr0jE5jTU0noF8yHGAJUAV9Ov-pw2D_IjtTo-bjiJuxipeu1KENPGo5Mpu-_WGc0I-GfXBRTSrBw2aPPIao2A'



{
    "statusCode": 200,
    "message": "Berhasil mendapatkan daftar paket!",
    "success": true,
    "data": [
        {
            "package_code": "TEST_PRODUCT_REST_API_SUCCESS",
            "package_name": "Kode Produk untuk Keperluan Test Integrasi REST API (SUKSES RESPONSE)",
            "package_name_alias_short": "Kode Produk untuk Keperluan Test Integrasi REST API",
            "package_description": "Kode Produk untuk Keperluan Test Integrasi REST API (SUKSES RESPONSE)\r\n\r\nSetelah 1 menit status akan berubah otomatis menjadi SUKSES\r\n\r\nCek update status trx melalui Endpoint API Check Transaction",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TEST_PRODUCT_REST_API_FAILED",
            "package_name": "Kode Produk untuk Keperluan Test Integrasi REST API (FAILED RESPONSE)",
            "package_name_alias_short": "Kode Produk untuk Keperluan Test Integrasi REST API",
            "package_description": "Kode Produk untuk Keperluan Test Integrasi REST API (FAILED RESPONSE)\r\n\r\nSetelah 1 menit status akan berubah otomatis menjadi GAGAL\r\n\r\nCek update status trx melalui Endpoint API Check Transaction",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTYOUTUBEXC_V1_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo YouTube untuk Xtra Combo (Versi 1)",
            "package_name_alias_short": "Unlimited Turbo YouTube untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo YouTube untuk akses YouTube secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp25.000 untuk XUT Addon ini. Akan terpotong Rp25.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTYOUTUBEXC_V2_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo YouTube untuk Xtra Combo (Versi 2)",
            "package_name_alias_short": "Unlimited Turbo YouTube untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo YouTube untuk akses YouTube secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp15.000 untuk XUT Addon ini. Akan terpotong Rp15.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTYOUTUBEXC_V3_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo YouTube untuk Xtra Combo (Versi 3)",
            "package_name_alias_short": "Unlimited Turbo YouTube untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo YouTube untuk akses YouTube secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp20.000 untuk XUT Addon ini. Akan terpotong Rp20.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTYOUTUBEXC_V4_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo YouTube untuk Xtra Combo (Versi 4)",
            "package_name_alias_short": "Unlimited Turbo YouTube untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo YouTube untuk akses YouTube secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp30.000 untuk XUT Addon ini. Akan terpotong Rp30.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTTIKTOKXC_V1_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo (Versi 1)",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp25.000 untuk XUT Addon ini. Akan terpotong Rp25.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTTIKTOKXC_V2_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo (Versi 2)",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp15.000 untuk XUT Addon ini. Akan terpotong Rp15.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTTIKTOKXC_V3_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo (Versi 3)",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp20.000 untuk XUT Addon ini. Akan terpotong Rp20.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XUTTIKTOKXC_V4_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo (Versi 4)",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "Jika ada notif \"Harga paket di sistem XL kemungkinan berubah atau tidak cocok dengan nomor kamu\" saat pembelian, silakan coba VERSI lainnya.\r\n\r\n- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Silakan gandeng xtra combo special (xcs)/xcp/xcvip plus (kecuali xtra combo vip double youtube tidak bisa). Tapi, yang combo special saja sih yang murah.\r\n- Beberapa nomor ada yang bisa tanpa paket gandengan (langsung masuk). Jadi coba-coba dulu saja.\r\n- Sediakan pulsa sebesar Rp30.000 untuk XUT Addon ini. Akan terpotong Rp30.000. Tapi juga jangan sedia pulsa banyak-banyak ya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXCVIP5PLUS10DBLYT_PREV",
            "package_name": "[RESMI] XL Xtra Combo 5GB + 10GB VIP Double YouTube 30 Hari (Unlimited untuk SSH/VPN, BACA DESKRIPSI)",
            "package_name_alias_short": "XL Xtra Combo 5GB + 10GB VIP Double YouTube 30 Hari",
            "package_description": "✅ KUOTA FULL 24 JAM NASIONAL\r\n✅ 5GB KUOTA UTAMA (REGULER)\r\n✅ 10GB KUOTA YOUTUBE\r\n✅ 20 MENIT TELPON ALL OPERATOR\r\n✅ AKTIF 30 HARI\r\n\r\nBISA UNTUK ALTERNATIF INJECT SSH/VPN (UNLIMITED).\r\n\r\nDILARANG TEMBAK KE KARTU XL KAMU YANG BEKAS/PERNAH TEMBAK XUT VIDIO/IFLIX 25K/SEJENISNYA. SPEED JADI SUPER LEMOT.\r\n\r\nKUOTA YOUTUBE NYA GAK BOLEH DIHABISKAN SAMPAI HABIS AGAR INJECTNYA TETAP JALAN (GAK TAU KENAPA). TIDAK BOLEH DIRECT KONEKSI KE YOUTUBE. HARUS FULL TRAFFIC DARI SSH/VPN.\r\n\r\nKONFIGURASI AMBIL DI V2RAY HELPER, SEHARUSNYA KUOTA UTAMA TIDAK TERPOTONG/TERPOTONG DIKIT-DIKIT SAJA JIKA INJECTNYA WORK.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMA1BLN_JAMU_CIRCLE_PREV",
            "package_name": "JAMU XL CIRCLE Jika Kuota Tidak Bisa Digunakan Walau Kartu Tidak Tenggang",
            "package_name_alias_short": "JAMU XL CIRCLE Jika Kuota Tidak Bisa Digunakan Walau Kartu Tidak Tenggang",
            "package_description": "JAMU XL CIRCLE Jika Kuota Tidak Bisa Digunakan Walau Kartu Tidak Tenggang\r\n\r\nPastikan paket lainnya kuota utama sudah habis terlebih dahulu\r\n\r\nCek kuota circle:\r\nhttps://leonskennedy.my.id/cekcircle/\r\n\r\nCek kuota selain circle:\r\nhttps://sidompul.google.com/",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_MASTIF_30D_P_V2",
            "package_name": "Masa Aktif XL 30 Hari",
            "package_name_alias_short": "Masa Aktif XL 30 Hari",
            "package_description": "Masa Aktif Kartu XL Prabayar 1 Bulan atau 30 Hari\r\n\r\nPENTING:\r\n- Hanya mengisi kembali masa aktif kartu XL prabayar agar balik ke 30 hari, bukan untuk mengakumulasinya (Contoh ada di bawah).\r\n- TIDAK BISA UNTUK KARTU SELAIN XL PRABAYAR/PREPAID.\r\n\r\nCONTOH\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH, MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF\r\n\r\nUntuk kartu XL yang sudah terlanjur di masa tenggang belum kami test, silakan coba saja.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLIFLIXXCL_E",
            "package_name": "[Method E-Wallet] XL IFLIX UNLIMITED KHUSUS XTRA COMBO LITE",
            "package_name_alias_short": "XL IFLIX UNLIMITED KHUSUS XTRA COMBO LITE",
            "package_description": "WAJIB DIGANDENG PAKET XTRA COMBO LITE/XCL (BAGI YANG MASIH PUNYA).\r\n\r\nUNLIMITED BUAT INJECT SSH/VPN. ADA DI MENU V2RAY HELPER, PRESET TRIK: XL IFLIX\r\n\r\nHANYA BISA DIINJECT UNTUK TKP JABAR, BANTEN, JAKARTA, SUMATERA MENGGUNAKAN BUG live.iflix.com metode SNI (SSL/TLS).\r\n\r\nMASA MENGIKUTI SISA MASA AKTIF PAKET XTRA COMBO LITE/XCL KAMU.\r\n\r\nSEDIAKAN SALDO E-WALLET RP6.000 UNTUK BAYAR KE XL (AKAN TERPOTONG SEGINI).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLIFLIXXCL_P",
            "package_name": "[Method Pulsa] XL IFLIX UNLIMITED KHUSUS XTRA COMBO LITE",
            "package_name_alias_short": "XL IFLIX UNLIMITED KHUSUS XTRA COMBO LITE",
            "package_description": "WAJIB DIGANDENG PAKET XTRA COMBO LITE/XCL (BAGI YANG MASIH PUNYA).\r\n\r\nUNLIMITED BUAT INJECT SSH/VPN. ADA DI MENU V2RAY HELPER, PRESET TRIK: XL IFLIX\r\n\r\nHANYA BISA DIINJECT UNTUK TKP JABAR, BANTEN, JAKARTA, SUMATERA MENGGUNAKAN BUG live.iflix.com metode SNI (SSL/TLS).\r\n\r\nMASA MENGIKUTI SISA MASA AKTIF PAKET XTRA COMBO LITE/XCL KAMU.\r\n\r\nSEDIAKAN PULSA RP5.000 UNTUK BAYAR KE XL (AKAN TERPOTONG SEGINI).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIZTAYO_P",
            "package_name": "[Method Pulsa] XL BIZ Tayo Unlimited",
            "package_name_alias_short": "XL BIZ Tayo Unlimited",
            "package_description": "GUNAKAN NOMOR KHUSUS/TUMBAL!! JANGAN NOMOR PRIBADI UNTUK MENGHINDARI HAL TIDAK DIINGINKAN, MISAL AKSES DATA INTERNET KENA BAN OLEH PROVIDER KARENA INI PADA UMUMNYA UNTUK KARTU XL BIZ BUKAN UNTUK PRABAYAR. KAMI TIDAK BERTANGGUNG JAWAB JIKA ANDA MASIH NEKAT MENCOBANYA DI NOMOR PRIBADI.\r\n\r\nTIDAK ADA GARANSI\r\nTIDAK ADA REFUND\r\n\r\nUNLIMITED BUAT INJECT SSH/VPN. NANTI AKAN KAMI UPDATE DI MENU V2RAY HELPER.\r\n\r\nMASA AKTIF 30 HARI.\r\n\r\nSEDIAKAN PULSA RP49.000 (AKAN TERPOTONG SEGINI)",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIZTAYO_E",
            "package_name": "[Method E-Wallet] XL BIZ Tayo Unlimited",
            "package_name_alias_short": "XL BIZ Tayo Unlimited",
            "package_description": "GUNAKAN METODE PEMBAYARAN DANA UNTUK BAYAR KE XL. METODE QRIS GAK BISA. ANE JUGA GAK TAU KENAPA.\r\n\r\nGUNAKAN NOMOR KHUSUS/TUMBAL!! JANGAN NOMOR PRIBADI UNTUK MENGHINDARI HAL TIDAK DIINGINKAN, MISAL AKSES DATA INTERNET KENA BAN OLEH PROVIDER KARENA INI PADA UMUMNYA UNTUK KARTU XL BIZ BUKAN UNTUK PRABAYAR. KAMI TIDAK BERTANGGUNG JAWAB JIKA ANDA MASIH NEKAT MENCOBANYA DI NOMOR PRIBADI.\r\n\r\nTIDAK ADA GARANSI\r\nTIDAK ADA REFUND\r\n\r\nUNLIMITED BUAT INJECT SSH/VPN. NANTI AKAN KAMI UPDATE DI MENU V2RAY HELPER.\r\n\r\nMASA AKTIF 30 HARI.\r\n\r\nSEDIAKAN SALDO E-WALLET RP49.000 (AKAN TERPOTONG SEGINI). GUNAKAN METODE PEMBAYARAN DANA UNTUK BAYAR KE XL. METODE QRIS GAK BISA. ANE JUGA GAK TAU KENAPA.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXCPADDONMULTIRITUAL_PREVIEW",
            "package_name": "RITUAL MULTI ADDON XCP RP0 BACA ATURAN PAKAI SAMPAI SELESAI :v",
            "package_name_alias_short": "RITUAL MULTI ADDON XCP RP0 BACA ATURAN PAKAI SAMPAI SELESAI :v",
            "package_description": "HANYA DISPLAY TRIK.\r\nSHARE INI = COID\r\nDARI KEJADIAN SEBELUMNYA SELAMA INI HARUSNYA SUDAH PAHAM DAN KALIAN GAK BLUNDER MENGULANGI HAL ITU.\r\n\r\nTRIK:\r\n\r\nSemua bahan ada di panel, cari sendiri.. Ini garis besarnya:\r\n\r\nInget gak boleh ada xcp/xcvip plus. Unreg bersih sampai akar akarnya di dial juga *808# > Info > Info Kartu XL-Ku > Stop Langganan. Walaupun di MyXL nampak bersih.\r\n\r\nPulsanya di bawah 30K dlu (WAJIB).\r\nTembak xcp 0kb di panel. WAJIB XCP 0KB LEWAT PANEL INI YA GAK BOLEH LAINNYA. KARENA UDAH ANE KASIH JIMAT..\r\n\r\nSetelah tembak xcp 0kb nanti biasanya sukses tapi xcp 0kb gak akan mau masuk. ok, it's okay memang begitu.\r\n\r\nLalu, tembak addon xcp seperti kuota youtube 15GB, streaming 15GB, dll dengan metode pulsa dengan masing masing percobaan interval jeda 20 detik. Tenang aja walau ada fee di addon 500 perak abaikan aja, gak bakal kena potong.. Nanti akan muncul notif gagal \"Failed call ipaass purchase...\" Wajib muncul kata ini di masing masing addon. WAJIB LEWAT PANEL INI JUGA, KARENA UDAH ANE KASIH JIMAT TERBARU..\r\n\r\nSetelah itu, isi pulsa 30K dan tembak xcp 0kb via metode pulsa (disarankan). Nanti akses sukses, xcp 0KB akan masuk dalam 3-5 menit. Dalam 15-30 menitan addon xcp baru akan masuk. Selama proses tunggu addon, jangan belikan paket apa-apa dulu lagi.\r\n\r\nDan boom...\r\nAddon semuanya akan masuk\r\nJika hoki xcp 0kb dapat kuota utama 3GB\r\n\r\nUntuk implementasi ke Addon XUT Netflix, Viu, dsb belum ane coba.. cobain sendiri ya. pegel ngoprek -_-",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "21:30",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "cae84e917d1da32a00d20b388b1b883b",
            "package_name": "[Metode Pulsa] Xtra Combo Plus 5GB (XCP 0KB atau Tanpa Kuota Utama)",
            "package_name_alias_short": "Xtra Combo Plus 5GB",
            "package_description": "- Masa Aktif 28 Hari.\r\n- Pastikan tidak ada paket Xtra Combo Plus yang sedang aktif di kartu XL kamu. Jika ada, unreg via dial *808#\r\n- Cocok untuk syarat sikat bonus Addon Xtra Combo Plus.\r\n- Sediakan Pulsa Rp30.000 (Akan terpotong segini). Harganya jadi segini, maklum guyss paket pada naik dari sananya hahaha...\r\n- Terdapat fee/biaya Admin pembelian paket (tertera).\r\n- Beda dengan XCP 10GB di atas, yang ini gak bisa diinject untuk SSH/VPN, kecuali untuk Addon XCP..",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "62a49f0f7585e542fc195914ee11d2de",
            "package_name": "[Method E-Wallet] Xtra Combo Plus 5GB (XCP 0KB atau Tanpa Kuota Utama)",
            "package_name_alias_short": "Xtra Combo Plus 5GB",
            "package_description": "- Masa Aktif 28 Hari.\r\n- Kartu jangan posisi masa tenggang saat didor.\r\n- Cocok untuk syarat sikat bonus Addon Xtra Combo Plus.\r\n- Pastikan tidak ada paket Xtra Combo Plus yang sedang aktif di kartu XL kamu. Jika ada, unreg via dial *808#\r\n- Sediakan saldo E-wallet Dana Rp30.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket. Harganya jadi segini, maklum guyss paket pada naik dari sananya hahaha...\r\n- Terdapat fee/biaya Admin (tertera) untuk mendukung sewa server kami.\r\n- Beda dengan XCP 10GB di atas, yang ini gak bisa diinject untuk SSH/VPN, kecuali untuk mainan Addon XCP.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 15GB",
            "package_description": "- Bonus Kuota 15GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_description": "- Bonus Kuota 15GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_SOCMED_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok",
            "package_description": "- Bonus Kuota 15GB Social Media (Spotify, Instagram, TikTok Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_SOCMED_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok",
            "package_description": "- Bonus Kuota 15GB Social Media (Spotify, Instagram, TikTok Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_STREAMING_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube",
            "package_description": "- Bonus Kuota 15GB Streaming (Iflix, Viu, YouTube Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_STREAMING_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube",
            "package_description": "- Bonus Kuota 15GB Streaming (Iflix, Viu, YouTube Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_GAMES_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG",
            "package_description": "- Bonus Kuota 15GB Games (Mobile Legends, Free Fire, PUBG Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_GAMES_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG",
            "package_description": "- Bonus Kuota 15GB Games (Mobile Legends, Free Fire, PUBG Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_YT_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_description": "- Bonus Kuota 15GB YouTube Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_description": "- Bonus Kuota 15GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_description": "- Bonus Kuota 15GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_SOCMED_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok",
            "package_description": "- Bonus Kuota 15GB Social Media (Spotify, Instagram, TikTok Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_SOCMED_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Social Media Spotify, Instagram, TikTok",
            "package_description": "- Bonus Kuota 15GB Social Media (Spotify, Instagram, TikTok Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_STREAMING_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube",
            "package_description": "- Bonus Kuota 15GB Streaming (Iflix, Viu, YouTube Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_STREAMING_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Streaming Iflix, Viu, YouTube",
            "package_description": "- Bonus Kuota 15GB Streaming (Iflix, Viu, YouTube Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_GAMES_15GB_7K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG",
            "package_description": "- Bonus Kuota 15GB Games (Mobile Legends, Free Fire, PUBG Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_GAMES_15GB_7K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG (Masing-Masing 5GB)",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus 15GB Games Mobile Legends, Free Fire, PUBG",
            "package_description": "- Bonus Kuota 15GB Games (Mobile Legends, Free Fire, PUBG Masing-Masing 5GB) Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus. Berikan jeda 10 menit setelah tembak paket Xtra Combo Plus/Xtra Combo VIP Plus baru tembak ini.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke posisi semula lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.7.500 (AKAN TERPOTONG RP7.500).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_YT_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_description": "- Bonus Kuota 15GB YouTube Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_description": "- Bonus Kuota 15GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 15GB",
            "package_description": "- Bonus Kuota 15GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 15GB",
            "package_description": "- Bonus Kuota 15GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_15GB_5K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 15GB",
            "package_description": "- Bonus Kuota 15GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_YT_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus YouTube 15GB",
            "package_description": "- Bonus Kuota 15GB YouTube Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 15GB",
            "package_description": "- Bonus Kuota 15GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_15GB_5K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 15GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 15GB",
            "package_description": "- Bonus Kuota 15GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 15GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.5.000 (AKAN TERPOTONG RP5.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE5GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 5GB - 9GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 5GB - 9GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 5GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE10GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 10GB - 14GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 10GB - 14GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 10GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE15GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 15-19GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 15-19GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 15GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE20GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 20-24GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 20-24GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 20GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE25GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 25-29GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 25-29GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 25GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE30GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 30-34GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 30-34GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 30GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE35GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 35-39GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 35-39GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 35GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder.\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE40GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 40-44GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 40-44GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 40GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder.\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE45GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 45-49GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 45-49GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 45GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE50GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 50-54GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 50-54GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 50GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE10GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 10GB - 14GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 10GB - 14GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 10GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE15GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 15-19GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 15-19GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 15GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE20GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 20-24GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 20-24GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 20GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE25GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 25-29GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 25-29GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 25GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE30GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 30-34GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 30-34GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 30GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE35GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 35-39GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 35-39GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 35GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder.\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE40GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 40-44GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 40-44GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 40GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder.\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE45GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 45-49GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 45-49GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 45GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE50GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 50-54GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 50-54GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 50GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE55GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 55-59GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 55-59GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 55GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "CIRCLE60GB",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] XL CIRCLE 60-70GB FULL KUOTA UTAMA REGULER",
            "package_name_alias_short": "XL CIRCLE 60-70GB FULL KUOTA UTAMA REGULER",
            "package_description": "XL Kuota Circle 60GB-70GB Full Kuota Utama\r\n\r\nPERHATIAN!!! PENTING:\r\n# Cek kuota circle HANYA DAPAT DILAKUKAN dengan membuka aplikasi MyXL yang terbaru -> Klik menu XL CIRCLE di bawah layar aplikasi MyXL. BUKAN dicek lewat \"Lihat Paket Saya\" atau Sidompul.\r\nIni karena Paket Circle berbeda dengan paket pada umumnya.\r\n# Jika dalam bulan ini pernah order/gabung Circle dari manapun, maka tidak bisa order paket ini, harus menunggu pada tanggal 1 bulan depan.\r\n\r\n- Hanya bisa di kartu XL Prabayar.\r\n- Pastikan nomor tujuan umurnya paling tidak sudah ada 60 hari ya. Kurang dari itu, tidak bisa diorder. Bisa cek umur kartu di sini:\r\nhttps://sidompul.google.com/\r\n- Jika beruntung, bisa dapat kuota lebih dari ini.\r\n- Masa aktif 28 hari atau hingga kuota circlenya habis. Jika kuota circle habis sebelum 28 hari maka kamu akan otomatis kami tendang dan masuk ke kondisi FREEZE/BEKU. Order ulang bisa dilakukan di bulan depan.\r\n- Tidak menambah masa aktif kartu XL.\r\n- Tidak dapat digunakan jika masa aktif kartu XL dalam masa tenggang.\r\n- Untuk reload masa aktif kartu sebulan, silakan tembak paket masa aktif XL 30 hari di panel kami terlebih dahulu.\r\n- Jika ada kuota utama selain circle, mereka akan dikonsumsi terlebih dahulu baru kuota circle.\r\n- Cek kuota circle kamu lewat aplikasi MyXL terbaru > Klik menu XL CIRCLE.\r\n- Tidak menggunakan OTP.\r\n- Dilarang unreg atau keluar dari keanggotaan XL CIRCLE. Keluar = garansi hangus tanpa refund. Tidak bisa diinvite ulang.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": true,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLSUPERMININEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota SUPER MINI V2 (13GB - 30GB)",
            "package_name_alias_short": "Akrab Anggota SUPER MINI V2",
            "package_description": "Akrab Anggota SUPER MINI V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 13GB - 15GB\r\nAREA 2 = 15GB - 17GB\r\nAREA 3 = 20GB - 22GB\r\nAREA 4 = 30GB - 32GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMININEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MINI V2 (31GB - 48GB)",
            "package_name_alias_short": "Akrab Anggota MINI V2",
            "package_description": "Akrab Anggota MINI V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 31GB - 33GB\r\nAREA 2 = 33GB - 35GB\r\nAREA 3 = 38GB - 40GB\r\nAREA 4 = 48GB - 50GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG V2 (38GB - 55GB)",
            "package_name_alias_short": "Akrab Anggota BIG V2",
            "package_description": "Akrab Anggota BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 38GB - 40GB\r\nAREA 2 = 40GB - 42GB\r\nAREA 3 = 45GB - 47GB\r\nAREA 4 = 55GB - 57GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGEXTRANEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG EXTRA V2 (33GB - 71GB)",
            "package_name_alias_short": "Akrab Anggota BIG EXTRA V2",
            "package_description": "Akrab Anggota BIG EXTRA V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 33GB\r\nAREA 2 = 36GB\r\nAREA 3 = 47GB\r\nAREA 4 = 71GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBOLITENEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO LITE V2 (47GB - 64GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO LITE V2",
            "package_description": "Akrab Anggota JUMBO LITE V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 47GB\r\nAREA 2 = 49GB\r\nAREA 3 = 54GB\r\nAREA 4 = 64GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBONEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO V2 (65GB - 123GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO V2",
            "package_description": "Akrab Anggota JUMBO V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 65GB\r\nAREA 2 = 70GB\r\nAREA 3 = 83GB\r\nAREA 4 = 123GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMEGABIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MEGA BIG V2 (88GB - 105GB)",
            "package_name_alias_short": "Akrab Anggota MEGA BIG V2",
            "package_description": "Akrab Anggota MEGA BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 88GB - 90GB\r\nAREA 2 = 90GB - 92GB\r\nAREA 3 = 95GB - 97GB\r\nAREA 4 = 105GB - 107GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_1D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 1 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 1 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 1 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 1 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_2D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 2 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 2 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 2 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 2 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_3D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 3 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 3 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 3 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 3 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_5D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 5 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 5 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 5 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 5 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_7D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 7 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 7 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 7 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 7 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_10D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 10 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 10 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 10 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 10 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_15D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 15 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 15 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 15 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 15 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMININEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MINI V2 (31GB - 48GB)",
            "package_name_alias_short": "Akrab Anggota MINI V2",
            "package_description": "Akrab Anggota MINI V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 31GB - 33GB\r\nAREA 2 = 33GB - 35GB\r\nAREA 3 = 38GB - 40GB\r\nAREA 4 = 48GB - 50GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG V2 (38GB - 55GB)",
            "package_name_alias_short": "Akrab Anggota BIG V2",
            "package_description": "Akrab Anggota BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 38GB - 40GB\r\nAREA 2 = 40GB - 42GB\r\nAREA 3 = 45GB - 47GB\r\nAREA 4 = 55GB - 57GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGEXTRANEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG EXTRA V2 (33GB - 71GB)",
            "package_name_alias_short": "Akrab Anggota BIG EXTRA V2",
            "package_description": "Akrab Anggota BIG EXTRA V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 33GB\r\nAREA 2 = 36GB\r\nAREA 3 = 47GB\r\nAREA 4 = 71GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBOLITENEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO LITE V2 (47GB - 64GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO LITE V2",
            "package_description": "Akrab Anggota JUMBO LITE V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 47GB\r\nAREA 2 = 49GB\r\nAREA 3 = 54GB\r\nAREA 4 = 64GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBONEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO V2 (65GB - 123GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO V2",
            "package_description": "Akrab Anggota JUMBO V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 65GB\r\nAREA 2 = 70GB\r\nAREA 3 = 83GB\r\nAREA 4 = 123GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMEGABIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MEGA BIG V2 (88GB - 105GB)",
            "package_name_alias_short": "Akrab Anggota MEGA BIG V2",
            "package_description": "Akrab Anggota MEGA BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 88GB - 90GB\r\nAREA 2 = 90GB - 92GB\r\nAREA 3 = 95GB - 97GB\r\nAREA 4 = 105GB - 107GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_1D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 1 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 1 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 1 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 1 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_2D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 2 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 2 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 2 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 2 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_3D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 3 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 3 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 3 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 3 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_5D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 5 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 5 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 5 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 5 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_7D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 7 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 7 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 7 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 7 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_10D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 10 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 10 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 10 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 10 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_15D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 15 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 15 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 15 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 15 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_20D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 20 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 20 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 20 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 20 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLEXTRAMEGABIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota EXTRA MEGA BIG V2 (105GB - 163GB)",
            "package_name_alias_short": "Akrab Anggota EXTRA MEGA BIG V2",
            "package_description": "Akrab Anggota EXTRA MEGA BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 105GB\r\nAREA 2 = 110GB\r\nAREA 3 = 123GB\r\nAREA 4 = 163GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGULTIMATENEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG ULTIMATE V2 (54GB - 71GB)",
            "package_name_alias_short": "Akrab Anggota BIG ULTIMATE V2",
            "package_description": "Akrab Anggota BIG ULTIMATE V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 54GB\r\nAREA 2 = 56GB\r\nAREA 3 = 61GB\r\nAREA 4 = 71GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXTRAON8GB28D_P",
            "package_name": "[Metode Pulsa] Xtra ON 8GB 28 Hari",
            "package_name_alias_short": "Xtra ON 8GB 28 Hari",
            "package_description": "Nikmati kuota utama 8GB untuk akses ke semua aplikasi favoritmu dengan masa aktif 28 hari.\r\n\r\n- Berlaku untuk pelanggan XL Prabayar.\r\n- Masa aktif 28 hari, tidak ada perpanjangan otomatis.\r\n- Kuota utama berlaku nasional (seluruh Indonesia).\r\n\r\nSediakan Pulsa Rp29.000 (akan terpotong Rp29.000).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXTRAON8GB28D_E",
            "package_name": "[Metode E-Wallet] Xtra ON 8GB 28 Hari",
            "package_name_alias_short": "Xtra ON 8GB 28 Hari",
            "package_description": "Nikmati kuota utama 8GB untuk akses ke semua aplikasi favoritmu dengan masa aktif 28 hari.\r\n\r\n- Berlaku untuk pelanggan XL Prabayar.\r\n- Masa aktif 28 hari, tidak ada perpanjangan otomatis.\r\n- Kuota utama berlaku nasional (seluruh Indonesia).\r\n\r\nSediakan Saldo E-Wallet Rp29.000 (akan terpotong Rp29.000).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BEBASPUAS15GB30HR_PREV",
            "package_name": "[OTOMATIS][RESMI][GARANSI] XL Data Bebas Puas 30 Hari 75GB-90GB Kuota Utama 24 Jam, Promo Triple Quota (Anti Hangus Sisa Kuota Jika Perpanjangan Terus), Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_name_alias_short": "XL Data Bebas Puas 30 Hari 75GB-90GB Kuota Utama 24 Jam, Promo Triple Quota , Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_description": "- Bisa ditimpa dengan variant Bebas Puas lainnya, misalnya timpa Bebas Puas 234GB atau Variant yang sama.\r\n- TIDAK PERLU KODE OTP.\r\n- Paket Resmi (OFFICIAL) jalur Sidompul.\r\n- MASA AKTIF 30 HARI.\r\n- GARANSI 30 HARI.\r\n- Bisa untuk semua Area.\r\n- Hanya bisa untuk XL Prepaid/Prabayar.\r\n- Kuota bersifat full reguler 24 jam.\r\n- Triple Quota Utama: 75GB.\r\n- Tidak wajib ada pulsa.\r\n- Jika ada paket bebas puas, maka akan keakumulasi sisa kuota dan masa aktifnya.\r\n- Kalau mau nimpa, bonusnya claim dlu di myxl (kalau dapat). Nanti rugi klo gk diambil dulu sebelum timpa, soalnya bakal dapet bonus lgi. Ada bonus: Kuota YouTube, TikTok, Utama (pilih salah satu).\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLDBEPU5430D_PREV",
            "package_name": "[OTOMATIS][RESMI][GARANSI] XL Data Bebas Puas 30 Hari 234GB Kuota Utama 24 Jam - Promo Triple Quota (Anti Hangus Sisa Kuota Jika Perpanjangan Terus), Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_name_alias_short": "XL Data Bebas Puas 30 Hari 234GB Kuota Utama 24 Jam - Promo Triple Quota , Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_description": "- Bisa ditimpa dengan variant Bebas Puas lainnya, misalnya timpa Bebas Puas 75GB atau Variant yang sama.\r\n- TIDAK PERLU KODE OTP.\r\n- Paket Resmi (OFFICIAL) jalur Sidompul.\r\n- MASA AKTIF 30 HARI.\r\n- GARANSI 30 HARI.\r\n- Bisa untuk semua Area.\r\n- Hanya bisa untuk XL Prepaid/Prabayar.\r\n- Kuota bersifat full reguler 24 jam.\r\n- Triple Quota Utama: 234GB.\r\n- Tidak wajib ada pulsa.\r\n- Jika ada paket bebas puas, maka akan keakumulasi sisa kuota dan masa aktifnya.\r\n- Kalau mau nimpa, bonusnya claim dlu di myxl. Nanti rugi klo gk diambil dulu sebelum timpa, soalnya bakal dapet bonus lgi. Ada bonus: Kuota YouTube, TikTok, Utama (pilih salah satu).\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBOV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO LITE 50.5-67.5GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota JUMBO LITE 50.5-67.5GB",
            "package_description": "Akrab Anggota JUMBO LITE\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n- AREA 1 = 50.5 - 52.5 GB\r\n- AREA 2 = 52.5 - 54.5 GB\r\n- AREA 3 = 57.5 - 59.5 GB\r\n- AREA 4 = 67.5 - 69.5 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGMEGA",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MEGA BIG 88-105GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota MEGA BIG 88-105GB",
            "package_description": "Akrab Anggota MEGA BIG\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 88 GB - 90 GB\r\n~ AREA 2 : 90 GB - 92 GB\r\n~ AREA 3 : 95 GB - 97 GB\r\n~ AREA 4 : 105 GB - 107 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMINISUPERV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota SUPER MINI PROMO 13-30GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota SUPER MINI PROMO 13-30GB",
            "package_description": "Akrab Anggota SUPER MINI PROMO\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n- AREA 1 = 13 - 15 GB\r\n- AREA 2 = 15 - 17 GB\r\n- AREA 3 = 20 - 22 GB\r\n- AREA 4 = 30 - 32 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF -+28 HARI. UDAH DARI XLNYA GITU.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMINI",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MINI 31.75-48.75GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota MINI 31.75-48.75GB",
            "package_description": "Akrab Anggota MINI\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 31.75 GB - 33.75 GB\r\n~ AREA 2 : 33.75 - 35.75 GB\r\n~ AREA 3 : 38.75 - 40.75 GB\r\n~ AREA 4 : 48.75 - 50.75 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIG",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG 38-55GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota BIG 38-55GB",
            "package_description": "Akrab Anggota BIG\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 38 GB - 40 GB\r\n~ AREA 2 : 40 GB - 42 GB\r\n~ AREA 3 : 45 GB - 47 GB\r\n~ AREA 4 : 55 GB - 57 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBO",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota JUMBO",
            "package_description": "Akrab Anggota JUMBO\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ Area 1 = 65 GB\r\n~ Area 2 = 70 GB\r\n~ Area 3 = 83 GB\r\n~ Area 4 = 123 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "f2fa05f28c146fdb67488ec53ce5c2cc",
            "package_name": "[METODE PULSA] XL DATA REGULER 2.8GB 27 JAM",
            "package_name_alias_short": "XL DATA REGULER 2.8GB 27 JAM",
            "package_description": "MASA AKTIF PAKET GAK 2 HARI FULL. PATOKANNYA EXPIRED PAKETNYA ADA DI APLIKASI MYXL (CEK KLO UDAH MASUK).\r\n\r\nSediakan pulsa Rp500 (akan terpotong Rp500) di kartu XL kamu (dan kunci/lock) untuk setiap kali tembak/eksekusi.\r\n\r\nHanya bisa untuk XL Prabayar/Prepaid.\r\n\r\nSISA KUOTA GAK BISA AKUMULASI (BERSIFAT REFILLABLE/ISI ULANG SAJA)\r\nKALAU MASA AKTIF PAKET BISA AKUMULASI.\r\n\r\nJIKA TIDAK BISA DITEMBAK ULANG, SILAKAN TUNGGU 2 HARI ATAU LEBIH KEMUDIAN COBA TEMBAK ULANG NANTI...",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLEDU2GB1D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Edukasi 2GB 1 Hari",
            "package_name_alias_short": "XL Xtra Edukasi 2GB 1 Hari",
            "package_description": "XL Xtra Edukasi 2GB 1 Hari\r\n\r\nBisa ditumpuk kuotanya. Bisa ditumpuk juga dengan Xtra edukasi yang 15GB 7 Hari sehingga jadi 7 hari (setelah 2GB nya ditumpuk).\r\n\r\nBisa buat inject VPN. Confignya cek dan ambil/generate di bagian menu khusus di panel ini.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG V2 (38GB - 55GB)",
            "package_name_alias_short": "Akrab Anggota BIG V2",
            "package_description": "Akrab Anggota BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 38GB - 40GB\r\nAREA 2 = 40GB - 42GB\r\nAREA 3 = 45GB - 47GB\r\nAREA 4 = 55GB - 57GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGEXTRANEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG EXTRA V2 (33GB - 71GB)",
            "package_name_alias_short": "Akrab Anggota BIG EXTRA V2",
            "package_description": "Akrab Anggota BIG EXTRA V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 33GB\r\nAREA 2 = 36GB\r\nAREA 3 = 47GB\r\nAREA 4 = 71GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBOLITENEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO LITE V2 (47GB - 64GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO LITE V2",
            "package_description": "Akrab Anggota JUMBO LITE V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 47GB\r\nAREA 2 = 49GB\r\nAREA 3 = 54GB\r\nAREA 4 = 64GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBONEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO V2 (65GB - 123GB)",
            "package_name_alias_short": "Akrab Anggota JUMBO V2",
            "package_description": "Akrab Anggota JUMBO V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 65GB\r\nAREA 2 = 70GB\r\nAREA 3 = 83GB\r\nAREA 4 = 123GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMEGABIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MEGA BIG V2 (88GB - 105GB)",
            "package_name_alias_short": "Akrab Anggota MEGA BIG V2",
            "package_description": "Akrab Anggota MEGA BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 88GB - 90GB\r\nAREA 2 = 90GB - 92GB\r\nAREA 3 = 95GB - 97GB\r\nAREA 4 = 105GB - 107GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_1D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 1 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 1 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 1 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 1 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_2D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 2 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 2 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 2 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 2 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_3D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 3 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 3 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 3 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 3 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_5D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 5 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 5 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 5 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 5 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_7D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 7 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 7 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 7 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 7 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_10D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 10 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 10 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 10 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 10 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_15D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 15 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 15 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 15 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 15 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBEKASANLV2_20D",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BEKASAN L V2 20 Hari (8GB - 25GB)",
            "package_name_alias_short": "Akrab Anggota BEKASAN L V2 20 Hari",
            "package_description": "Akrab Anggota BEKASAN L V2 20 Hari\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 8GB\r\nAREA 2 = 10GB\r\nAREA 3 = 15GB\r\nAREA 4 = 25GB\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF 20 HARI.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:59"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLEXTRAMEGABIGNEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota EXTRA MEGA BIG V2 (105GB - 163GB)",
            "package_name_alias_short": "Akrab Anggota EXTRA MEGA BIG V2",
            "package_description": "Akrab Anggota EXTRA MEGA BIG V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 105GB\r\nAREA 2 = 110GB\r\nAREA 3 = 123GB\r\nAREA 4 = 163GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGULTIMATENEWV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG ULTIMATE V2 (54GB - 71GB)",
            "package_name_alias_short": "Akrab Anggota BIG ULTIMATE V2",
            "package_description": "Akrab Anggota BIG ULTIMATE V2\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\nAREA 1 = 54GB\r\nAREA 2 = 56GB\r\nAREA 3 = 61GB\r\nAREA 4 = 71GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXTRAON8GB28D_P",
            "package_name": "[Metode Pulsa] Xtra ON 8GB 28 Hari",
            "package_name_alias_short": "Xtra ON 8GB 28 Hari",
            "package_description": "Nikmati kuota utama 8GB untuk akses ke semua aplikasi favoritmu dengan masa aktif 28 hari.\r\n\r\n- Berlaku untuk pelanggan XL Prabayar.\r\n- Masa aktif 28 hari, tidak ada perpanjangan otomatis.\r\n- Kuota utama berlaku nasional (seluruh Indonesia).\r\n\r\nSediakan Pulsa Rp29.000 (akan terpotong Rp29.000).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXTRAON8GB28D_E",
            "package_name": "[Metode E-Wallet] Xtra ON 8GB 28 Hari",
            "package_name_alias_short": "Xtra ON 8GB 28 Hari",
            "package_description": "Nikmati kuota utama 8GB untuk akses ke semua aplikasi favoritmu dengan masa aktif 28 hari.\r\n\r\n- Berlaku untuk pelanggan XL Prabayar.\r\n- Masa aktif 28 hari, tidak ada perpanjangan otomatis.\r\n- Kuota utama berlaku nasional (seluruh Indonesia).\r\n\r\nSediakan Saldo E-Wallet Rp29.000 (akan terpotong Rp29.000).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BEBASPUAS15GB30HR_PREV",
            "package_name": "[OTOMATIS][RESMI][GARANSI] XL Data Bebas Puas 30 Hari 75GB-90GB Kuota Utama 24 Jam, Promo Triple Quota (Anti Hangus Sisa Kuota Jika Perpanjangan Terus), Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_name_alias_short": "XL Data Bebas Puas 30 Hari 75GB-90GB Kuota Utama 24 Jam, Promo Triple Quota , Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_description": "- Bisa ditimpa dengan variant Bebas Puas lainnya, misalnya timpa Bebas Puas 234GB atau Variant yang sama.\r\n- TIDAK PERLU KODE OTP.\r\n- Paket Resmi (OFFICIAL) jalur Sidompul.\r\n- MASA AKTIF 30 HARI.\r\n- GARANSI 30 HARI.\r\n- Bisa untuk semua Area.\r\n- Hanya bisa untuk XL Prepaid/Prabayar.\r\n- Kuota bersifat full reguler 24 jam.\r\n- Triple Quota Utama: 75GB.\r\n- Tidak wajib ada pulsa.\r\n- Jika ada paket bebas puas, maka akan keakumulasi sisa kuota dan masa aktifnya.\r\n- Kalau mau nimpa, bonusnya claim dlu di myxl (kalau dapat). Nanti rugi klo gk diambil dulu sebelum timpa, soalnya bakal dapet bonus lgi. Ada bonus: Kuota YouTube, TikTok, Utama (pilih salah satu).\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLDBEPU5430D_PREV",
            "package_name": "[OTOMATIS][RESMI][GARANSI] XL Data Bebas Puas 30 Hari 234GB Kuota Utama 24 Jam - Promo Triple Quota (Anti Hangus Sisa Kuota Jika Perpanjangan Terus), Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_name_alias_short": "XL Data Bebas Puas 30 Hari 234GB Kuota Utama 24 Jam - Promo Triple Quota , Biasanya Dapat Bonus Claim -> Segera Claim di MyXL",
            "package_description": "- Bisa ditimpa dengan variant Bebas Puas lainnya, misalnya timpa Bebas Puas 75GB atau Variant yang sama.\r\n- TIDAK PERLU KODE OTP.\r\n- Paket Resmi (OFFICIAL) jalur Sidompul.\r\n- MASA AKTIF 30 HARI.\r\n- GARANSI 30 HARI.\r\n- Bisa untuk semua Area.\r\n- Hanya bisa untuk XL Prepaid/Prabayar.\r\n- Kuota bersifat full reguler 24 jam.\r\n- Triple Quota Utama: 234GB.\r\n- Tidak wajib ada pulsa.\r\n- Jika ada paket bebas puas, maka akan keakumulasi sisa kuota dan masa aktifnya.\r\n- Kalau mau nimpa, bonusnya claim dlu di myxl. Nanti rugi klo gk diambil dulu sebelum timpa, soalnya bakal dapet bonus lgi. Ada bonus: Kuota YouTube, TikTok, Utama (pilih salah satu).\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBOV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO LITE 50.5-67.5GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota JUMBO LITE 50.5-67.5GB",
            "package_description": "Akrab Anggota JUMBO LITE\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n- AREA 1 = 50.5 - 52.5 GB\r\n- AREA 2 = 52.5 - 54.5 GB\r\n- AREA 3 = 57.5 - 59.5 GB\r\n- AREA 4 = 67.5 - 69.5 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIGMEGA",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MEGA BIG 88-105GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota MEGA BIG 88-105GB",
            "package_description": "Akrab Anggota MEGA BIG\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 88 GB - 90 GB\r\n~ AREA 2 : 90 GB - 92 GB\r\n~ AREA 3 : 95 GB - 97 GB\r\n~ AREA 4 : 105 GB - 107 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMINISUPERV2",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota SUPER MINI PROMO 13-30GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota SUPER MINI PROMO 13-30GB",
            "package_description": "Akrab Anggota SUPER MINI PROMO\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n- AREA 1 = 13 - 15 GB\r\n- AREA 2 = 15 - 17 GB\r\n- AREA 3 = 20 - 22 GB\r\n- AREA 4 = 30 - 32 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF -+28 HARI. UDAH DARI XLNYA GITU.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMINI",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota MINI 31.75-48.75GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota MINI 31.75-48.75GB",
            "package_description": "Akrab Anggota MINI\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 31.75 GB - 33.75 GB\r\n~ AREA 2 : 33.75 - 35.75 GB\r\n~ AREA 3 : 38.75 - 40.75 GB\r\n~ AREA 4 : 48.75 - 50.75 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLBIG",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota BIG 38-55GB (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota BIG 38-55GB",
            "package_description": "Akrab Anggota BIG\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ AREA 1 : 38 GB - 40 GB\r\n~ AREA 2 : 40 GB - 42 GB\r\n~ AREA 3 : 45 GB - 47 GB\r\n~ AREA 4 : 55 GB - 57 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLJUMBO",
            "package_name": "[PROSES OTOMATIS] [RESMI] [GARANSI FULL] Akrab Anggota JUMBO (BACA DESKRIPSI)",
            "package_name_alias_short": "Akrab Anggota JUMBO",
            "package_description": "Akrab Anggota JUMBO\r\n\r\n- WAKTU PROSES OTOMATIS 1-5 MENIT BY SYSTEM.\r\n- Paket Resmi (OFFICIAL).\r\n- GARANSI FULL.\r\n- Bisa untuk semua Area.\r\n- Bisa untuk XL / AXIS / LIVEON.\r\n- Cara Unreg paket akrab yang benar:\r\nhttps://google.com/cara-unreg-paket-akrab-yang-benar\r\n- Kuota bersifat full reguler 24 jam. Total kuota yang akan kamu dapat adalah sebagai berikut:\r\n~ Area 1 = 65 GB\r\n~ Area 2 = 70 GB\r\n~ Area 3 = 83 GB\r\n~ Area 4 = 123 GB\r\n- Apabila Kuota MyRewards belum masuk (belum masuk full sesuai deskripsi di atas), tunggu 1x24 jam dulu ya baru laporan ke Admin.\r\n- Cek area kamu di:\r\nhttp://bit.ly/area_akrab\r\n- Pastikan Anda telah memasukkan kartu perdana yang ingin ditembak ke Handphone/Modem/Mifi agar dapat mendeteksi lokasi BTS terdekat dan mendapatkan bonus kuota lokal.\r\n- MASA AKTIF KURANG LEBIH 28 HARI. SUDAH DARI KETENTUAN PIHAK XL.\r\n- Tidak wajib ada pulsa.\r\n- Jika transaksi status GAGAL maka saldo panel akan otomatis di-refund.\r\n- Memerlukan waktu 10-30 detik. Jangan tinggalkan halaman atau memutuskan koneksi internet.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": true,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "f2fa05f28c146fdb67488ec53ce5c2cc",
            "package_name": "[METODE PULSA] XL DATA REGULER 2.8GB 27 JAM",
            "package_name_alias_short": "XL DATA REGULER 2.8GB 27 JAM",
            "package_description": "MASA AKTIF PAKET GAK 2 HARI FULL. PATOKANNYA EXPIRED PAKETNYA ADA DI APLIKASI MYXL (CEK KLO UDAH MASUK).\r\n\r\nSediakan pulsa Rp500 (akan terpotong Rp500) di kartu XL kamu (dan kunci/lock) untuk setiap kali tembak/eksekusi.\r\n\r\nHanya bisa untuk XL Prabayar/Prepaid.\r\n\r\nSISA KUOTA GAK BISA AKUMULASI (BERSIFAT REFILLABLE/ISI ULANG SAJA)\r\nKALAU MASA AKTIF PAKET BISA AKUMULASI.\r\n\r\nJIKA TIDAK BISA DITEMBAK ULANG, SILAKAN TUNGGU 2 HARI ATAU LEBIH KEMUDIAN COBA TEMBAK ULANG NANTI...",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLEDU2GB1D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Edukasi 2GB 1 Hari",
            "package_name_alias_short": "XL Xtra Edukasi 2GB 1 Hari",
            "package_description": "XL Xtra Edukasi 2GB 1 Hari\r\n\r\nBisa ditumpuk kuotanya. Bisa ditumpuk juga dengan Xtra edukasi yang 15GB 7 Hari sehingga jadi 7 hari (setelah 2GB nya ditumpuk).\r\n\r\nBisa buat inject VPN. Confignya cek dan ambil/generate di bagian menu khusus di panel ini.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLEDU5GB3D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Edukasi 5GB 3 Hari",
            "package_name_alias_short": "XL Xtra Edukasi 5GB 3 Hari",
            "package_description": "XL Xtra Edukasi 5GB 3 Hari\r\n\r\nBisa ditumpuk kuotanya. Bisa ditumpuk juga dengan Xtra edukasi yang 15GB 7 Hari sehingga jadi 7 hari (setelah 2GB/5GB nya ditumpuk).\r\n\r\nBisa buat inject VPN. Confignya cek dan ambil/generate di bagian menu khusus di panel ini.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLEDU15GB7D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Edukasi 15GB 7 Hari",
            "package_name_alias_short": "XL Xtra Edukasi 15GB 7 Hari",
            "package_description": "XL Xtra Edukasi 15GB 7 Hari\r\n\r\nBisa ditumpuk kuotanya.\r\n\r\nBisa buat inject VPN. Confignya cek dan ambil/generate di bagian menu khusus di panel ini.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLCONF2GB1D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Conference 2GB 1 Hari",
            "package_name_alias_short": "XL Xtra Conference 2GB 1 Hari",
            "package_name_alias_short": "XL Xtra Conference 2GB 1 Hari",
            "package_description": "XL Xtra Conference 2GB 1 Hari\r\n\r\nBisa ditumpuk kuotanya. Bisa ditumpuk juga dengan Xtra conference yang 15GB 7 Hari sehingga jadi 7 hari (setelah 2GB nya ditumpuk).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLCONF5GB3D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Conference 5GB 3 Hari",
            "package_name_alias_short": "XL Xtra Conference 5GB 3 Hari",
            "package_description": "XL Xtra Conference 5GB 3 Hari\r\n\r\nBisa ditumpuk kuotanya. Bisa ditumpuk juga dengan Xtra conference yang 15GB 7 Hari sehingga jadi 7 hari (setelah 2GB/5GB nya ditumpuk).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "SXLCONF15GB7D_PREV",
            "package_name": "[RESMI & NO OTP] XL Xtra Conference 15GB 7 Hari",
            "package_name_alias_short": "XL Xtra Conference 15GB 7 Hari",
            "package_description": "XL Xtra Conference 15GB 7 Hari\r\n\r\nBisa ditumpuk kuotanya.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLXCS8GB_PREV",
            "package_name": "XL Xtra Combo Special 8GB (1GB-8GB) 30 Hari Non Discount (Sedia Saldo Panel)",
            "package_name_alias_short": "XL Xtra Combo Special 8GB  30 Hari Non Discount",
            "package_description": "- XL Xtra Combo Special Up to (hingga) 8 GB tergantung Area / 30 Hari harga normal. Kuota yang didapat akan bervariasi tergantung TKP/Area. Karena beda dengan paket akrab, maka kami tidak bisa memastikan dapat berapa GB, bisa dari 1GB hingga 8GB. Jangan kaget apabila nanti ada yang dapat kuota 1GB/2GB/3GB/4GB/5GB/6GB/7GB/8GB. Membeli = SETUJU ya.\r\n- Dapat digunakan untuk syarat aktivasi Paket XL Unlimited Turbo di aplikasi MyXL untuk kebutuhan inject/trik internet murah menggunakan SSH/VPN.\r\n- Kartu jangan posisi masa tenggang saat ditembak.\r\n- Terdapat fee/biaya Admin (tertera). Sediakan saldo panel untuk menembak, tidak perlu repot sediakan pulsa.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOHPREMIUM7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Premium bayar ke XL Rp13.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Premium bayar ke XL Rp13.500",
            "package_description": "- Xtra Unlimited Turbo Premium untuk akses LINE, Gojek, Facebook, Instagram, dan YouTube secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp13.500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSUPER7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Super bayar ke XL Rp8.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Super bayar ke XL Rp8.500",
            "package_description": "- Xtra Unlimited Turbo Super untuk akses LINE, Gojek, Facebook, dan Instagram secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp8.500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSTANDARD7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Standard bayar ke XL Rp3.000 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Standard bayar ke XL Rp3.000",
            "package_description": "- Xtra Unlimited Turbo Standard untuk akses LINE, Gojek, dan Facebook secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp3.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHBASIC7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Basic bayar ke XL Rp500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Basic bayar ke XL Rp500",
            "package_description": "- Xtra Unlimited Turbo Basic untuk akses LINE dan Gojek secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 10-20 an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHBASIC7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Basic bayar ke XL Rp500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Basic bayar ke XL Rp500",
            "package_description": "- Xtra Unlimited Turbo Basic untuk akses LINE dan Gojek secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 10-20 an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSTANDARD7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Standard bayar ke XL Rp3.000 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Standard bayar ke XL Rp3.000",
            "package_description": "- Xtra Unlimited Turbo Standard untuk akses LINE, Gojek, dan Facebook secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp3.000 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSUPER7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Super bayar ke XL Rp8.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Super bayar ke XL Rp8.500",
            "package_description": "- Xtra Unlimited Turbo Super untuk akses LINE, Gojek, Facebook, dan Instagram secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp8.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHPREMIUM7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Premium bayar ke XL Rp13.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Premium bayar ke XL Rp13.500",
            "package_description": "- Xtra Unlimited Turbo Premium untuk akses LINE, Gojek, Facebook, Instagram, dan YouTube secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp13.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLFLEXS_PREV",
            "package_name": "Xtra Combo Flex S",
            "package_name_alias_short": "Xtra Combo Flex S",
            "package_description": "- 3.5GB kuota utama dan tambahan kuota lokal (sesuai area jika ada) 28 Hari.\r\n- Digunakan untuk syarat tembak Bonus Flex.\r\n- Kartu jangan posisi masa tenggang saat ditembak.\r\n- Terdapat fee/biaya Admin (tertera). Sediakan saldo panel untuk menembak, tidak perlu repot sediakan pulsa.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "b0012edd21983678eb7ebc08d8f04ecd",
            "package_name": "Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "1447d54e21d581c9fb340e1cbf4e8fca",
            "package_name": "[SPECIAL FOR YOU] Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "REINVIT",
            "package_name": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_name_alias_short": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_description": "Kini tidak perlu lapor Admin !!! Apabila paket Anggota Akrab Anda dari kami tidak masuk atau terkena bug kick/tendang dari pengelola lama (pengelola ghoib) karena diunreg paksa paket akrabnya, Anda bisa melakukan reinvite di sini.\r\n\r\nJika tidak berfungsi/work, silakan hubungi Admin.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMA1BLN_PREV",
            "package_name": "[RESMI dan NO OTP] Tambah Masa Aktif XL 1 Bulan (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif XL 1 Bulan",
            "package_description": "Masa Aktif Kartu XL Prabayar 1 Bulan atau 30 Hari\r\n\r\nPENTING:\r\nHanya mengisi kembali masa aktif kartu XL prabayar agar balik ke 30 hari, bukan untuk mengakumulasinya.\r\n\r\nCONTOH\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH, MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF\r\n\r\nTidak bisa untuk kartu XL yang sudah terlanjur di masa tenggang/hangus",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TRIMA2_PREV",
            "package_name": "[RESMI] Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_name_alias_short": "Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF THREE :v\r\n\r\nAMAN DI BOM TRX\r\n\r\n1 nomor max 2-3 kali\r\n\r\nAkumulasi maksimal 365 hari ya\r\nJika pengisian produk dilakukan di kartu Three dengan masa aktif lebih dari 365 hari maka transaksi tetap akan sukses namun masa aktif tidak bertambah",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA90HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)\r\n\r\nAkumulasi masa aktif pastikan tidak melebihi 360 hari.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA30HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)\r\n\r\n- Akumulasi masa aktif pastikan tidak melebihi 360 hari",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ISATMA30D_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu INDOSAT 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF INDOSAT :v\r\n\r\nTambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)\r\n\r\nSEBELUM ORDER, BACA PESAN PENTING BERIKUT:\r\n\r\n✅ PRODUK AKTIF 24 JAM SELAMA GAK MAINTENANCE/PEMELIHARAAN SISTEM DARI INDOSAT\r\n✅ AKUMULASI MAKSIMAL 30 HARI\r\n\r\nSEBAGAI CONTOH:\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIUXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOKXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp13.500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSUPER7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Super bayar ke XL Rp8.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Super bayar ke XL Rp8.500",
            "package_description": "- Xtra Unlimited Turbo Super untuk akses LINE, Gojek, Facebook, dan Instagram secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp8.500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSTANDARD7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Standard bayar ke XL Rp3.000 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Standard bayar ke XL Rp3.000",
            "package_description": "- Xtra Unlimited Turbo Standard untuk akses LINE, Gojek, dan Facebook secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp3.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHBASIC7H",
            "package_name": "[Method E-Wallet] Unlimited Turbo Basic bayar ke XL Rp500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Basic bayar ke XL Rp500",
            "package_description": "- Xtra Unlimited Turbo Basic untuk akses LINE dan Gojek secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 10-20 an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp500 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHBASIC7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Basic bayar ke XL Rp500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Basic bayar ke XL Rp500",
            "package_description": "- Xtra Unlimited Turbo Basic untuk akses LINE dan Gojek secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 10-20 an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSTANDARD7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Standard bayar ke XL Rp3.000 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Standard bayar ke XL Rp3.000",
            "package_description": "- Xtra Unlimited Turbo Standard untuk akses LINE, Gojek, dan Facebook secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp3.000 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSUPER7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Super bayar ke XL Rp8.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Super bayar ke XL Rp8.500",
            "package_description": "- Xtra Unlimited Turbo Super untuk akses LINE, Gojek, Facebook, dan Instagram secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp8.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHPREMIUM7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Premium bayar ke XL Rp13.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Premium bayar ke XL Rp13.500",
            "package_description": "- Xtra Unlimited Turbo Premium untuk akses LINE, Gojek, Facebook, Instagram, dan YouTube secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp13.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLFLEXS_PREV",
            "package_name": "Xtra Combo Flex S",
            "package_name_alias_short": "Xtra Combo Flex S",
            "package_description": "- 3.5GB kuota utama dan tambahan kuota lokal (sesuai area jika ada) 28 Hari.\r\n- Digunakan untuk syarat tembak Bonus Flex.\r\n- Kartu jangan posisi masa tenggang saat ditembak.\r\n- Terdapat fee/biaya Admin (tertera). Sediakan saldo panel untuk menembak, tidak perlu repot sediakan pulsa.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "b0012edd21983678eb7ebc08d8f04ecd",
            "package_name": "Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "1447d54e21d581c9fb340e1cbf4e8fca",
            "package_name": "[SPECIAL FOR YOU] Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "REINVIT",
            "package_name": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_name_alias_short": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_description": "Kini tidak perlu lapor Admin !!! Apabila paket Anggota Akrab Anda dari kami tidak masuk atau terkena bug kick/tendang dari pengelola lama (pengelola ghoib) karena diunreg paksa paket akrabnya, Anda bisa melakukan reinvite di sini.\r\n\r\nJika tidak berfungsi/work, silakan hubungi Admin.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMA1BLN_PREV",
            "package_name": "[RESMI dan NO OTP] Tambah Masa Aktif XL 1 Bulan (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif XL 1 Bulan",
            "package_description": "Masa Aktif Kartu XL Prabayar 1 Bulan atau 30 Hari\r\n\r\nPENTING:\r\nHanya mengisi kembali masa aktif kartu XL prabayar agar balik ke 30 hari, bukan untuk mengakumulasinya.\r\n\r\nCONTOH\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH, MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF\r\n\r\nTidak bisa untuk kartu XL yang sudah terlanjur di masa tenggang/hangus",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TRIMA2_PREV",
            "package_name": "[RESMI] Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_name_alias_short": "Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF THREE :v\r\n\r\nAMAN DI BOM TRX\r\n\r\n1 nomor max 2-3 kali\r\n\r\nAkumulasi maksimal 365 hari ya\r\nJika pengisian produk dilakukan di kartu Three dengan masa aktif lebih dari 365 hari maka transaksi tetap akan sukses namun masa aktif tidak bertambah",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA90HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)\r\n\r\nAkumulasi masa aktif pastikan tidak melebihi 360 hari.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA30HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)\r\n\r\n- Akumulasi masa aktif pastikan tidak melebihi 360 hari",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ISATMA30D_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu INDOSAT 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF INDOSAT :v\r\n\r\nTambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)\r\n\r\nSEBELUM ORDER, BACA PESAN PENTING BERIKUT:\r\n\r\n✅ PRODUK AKTIF 24 JAM SELAMA GAK MAINTENANCE/PEMELIHARAAN SISTEM DARI INDOSAT\r\n✅ AKUMULASI MAKSIMAL 30 HARI\r\n\r\nSEBAGAI CONTOH:\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIUXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOKXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIU_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu ketika semua kuota utama paket habis.\r\n- Masa Aktif mengikut paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini (wajib ditembak dulu paket Xtra Combo Special 8GB/14GB nya).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 100-150GB an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan Pulsa minimal Rp25.000 (Akan terpotong segini).\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOK_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok ketika semua kuota utama paket habis.\r\n\r\n- Khusus untuk trik XUT Addon, silakan tembak paket ini sampai muncul notif \"Paket berhasil dibeli. Silakan cek kuotanya....\". JIKA NOTIF PERTAMA \"Paket berhasil dibeli. Silakan cek kuotanya....\" MAKA BERIKAN JEDA SEKITAR 10 MENIT.\r\n\r\nSEMBARI MENUNGGU 10 MENIT, KAMU MUNGKIN BISA LANJUT KE ADDON LAINNYA AGAR TIDAK SALING TUNGGU, INTINYA JIKA PERTAMA MUNCUL NOTIF \"Paket berhasil dibeli...\" maka tunggu juga selama 10 menit.\r\n\r\nSETELAH ITU, BELI LAGI HINGGA MUNCUL NOTIF \"422 -> Failed call ipaas purchase, with status code:422 : null,\" (DIPAKSA) DI MASING MASING ADDON. INI TUJUAN YANG KITA CARI-CARI, HARUS ADA RESPON 422 SEPERTI INI.\r\n\r\nDI PALING AKHIR JIKA SEMUA ADDON SUDAH DIRITUALKAN DAN DAPAT RESPON \"422 -> Failed call ipaas purchase, with status code:422 : null,\", TINGGAL GANDENG DENGAN XTRA COMBO SPECIAL 8GB, DAN TUNGGU SELAMA MAX 1-2 JAM NANTI SEMUA ADDON AKAN MASUK. SELAMA MASA TUNGGU 1-2 JAM, SEBAIKNYA JANGAN BELI PAKET APA-APA LAGI DULU.\r\n\r\nSEBELUM MELAKUKAN RITUAL DI ATAS, PASTIKAN TIDAK ADA XTRA COMBO APAPUN, WAJIB CEK VIAL DIAL *808# > INFO > INFO Kartu XL-Ku > Stop Langganan. Unreg sampai bersih dari berbau Xtra Combo 100GB, Xtra Combo Plus, Xtra Combo VIP Plus, Xtra Hotrod, dan sejenisnya (kalau Xtra Combo Flex biarkan saja).\r\n\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
"is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSTANDARD7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Standard bayar ke XL Rp3.000 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Standard bayar ke XL Rp3.000",
            "package_description": "- Xtra Unlimited Turbo Standard untuk akses LINE, Gojek, dan Facebook secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp3.000 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHSUPER7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Super bayar ke XL Rp8.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Super bayar ke XL Rp8.500",
            "package_description": "- Xtra Unlimited Turbo Super untuk akses LINE, Gojek, Facebook, dan Instagram secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp8.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLUNLITURBOHPREMIUM7H_P",
            "package_name": "[Method Pulsa] Unlimited Turbo Premium bayar ke XL Rp13.500 (Untuk Xtra Hotrod/Xtra Combo)",
            "package_name_alias_short": "Unlimited Turbo Premium bayar ke XL Rp13.500",
            "package_description": "- Xtra Unlimited Turbo Premium untuk akses LINE, Gojek, Facebook, Instagram, dan YouTube secara Unlimited dengan Kuota WhatsApp Harian 50MB/hari.\r\n- Masa Aktif 7 Hari mengikuti paket utama Xtra Hotrod 7 Hari ATAU menjadi <b>30 Hari</b> (jika digabung dengan Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan pulsa minimal Rp13.500 (Akan terpotong segini).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": true
        },
        {
            "package_code": "XLFLEXS_PREV",
            "package_name": "Xtra Combo Flex S",
            "package_name_alias_short": "Xtra Combo Flex S",
            "package_description": "- 3.5GB kuota utama dan tambahan kuota lokal (sesuai area jika ada) 28 Hari.\r\n- Digunakan untuk syarat tembak Bonus Flex.\r\n- Kartu jangan posisi masa tenggang saat ditembak.\r\n- Terdapat fee/biaya Admin (tertera). Sediakan saldo panel untuk menembak, tidak perlu repot sediakan pulsa.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "b0012edd21983678eb7ebc08d8f04ecd",
            "package_name": "Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "1447d54e21d581c9fb340e1cbf4e8fca",
            "package_name": "[SPECIAL FOR YOU] Masa Aktif XL 1 Tahun (BACA DESKRIPSI)",
            "package_name_alias_short": "Masa Aktif XL 1 Tahun",
            "package_description": "<b>JANGAN BARBAR ATAU SAYA TUTUP, SOALNYA COCOK TERUTAMA BUAT JAGA MASA AKTIF PENGELOLA PAKET AKRAB</b>\r\n\r\n- Masa aktif selama 1 tahun (12 bulan).\r\n- Wajib tembak masa aktif sebulan dulu biar Aman. Ada di panel ini. Cari saja.\r\n- Jangan kaget jika masa aktif kartu tidak langsung bertambah. Masa aktif akan diperpanjang setiap bulan, tidak langsung dapat 1 tahun (12 bulan). Pemicu perpanjangan masa aktif kartu nanti saat paket setahun perpanjangan Rp0 pada tanggal xx xx xxxx 00.00 WIB. Akan tambah 30 hari tiap bulannya.\r\nJika kartunya mendekati masa tenggang, sebaiknya bisa tembak masa aktif sebulan dulu yang ada di panel.\r\n- Tidak dapat diakumulasi.\r\n- Tidak bisa dalam keadaan masa tenggang. Jika sudah terlanjur masa tenggang, tembak masa aktif sebulan di panel.\r\n- Sisa pulsa boleh dalam posisi Rp0 (kosongan).\r\n- Terdapat fee/biaya pembelian paket ini untuk mendukung biaya infrastruktur server kami (tertera).\r\n- NO GARANSI.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": true,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "06:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "REINVIT",
            "package_name": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_name_alias_short": "REINVITE ANGGOTA PAKET AKRAB DARI KAMI",
            "package_description": "Kini tidak perlu lapor Admin !!! Apabila paket Anggota Akrab Anda dari kami tidak masuk atau terkena bug kick/tendang dari pengelola lama (pengelola ghoib) karena diunreg paksa paket akrabnya, Anda bisa melakukan reinvite di sini.\r\n\r\nJika tidak berfungsi/work, silakan hubungi Admin.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": true,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLMA1BLN_PREV",
            "package_name": "[RESMI dan NO OTP] Tambah Masa Aktif XL 1 Bulan (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif XL 1 Bulan",
            "package_description": "Masa Aktif Kartu XL Prabayar 1 Bulan atau 30 Hari\r\n\r\nPENTING:\r\nHanya mengisi kembali masa aktif kartu XL prabayar agar balik ke 30 hari, bukan untuk mengakumulasinya.\r\n\r\nCONTOH\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH, MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF\r\n\r\nTidak bisa untuk kartu XL yang sudah terlanjur di masa tenggang/hangus",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TRIMA2_PREV",
            "package_name": "[RESMI] Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_name_alias_short": "Tri Tambah Masa Aktif Kartu 4 Bulan",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF THREE :v\r\n\r\nAMAN DI BOM TRX\r\n\r\n1 nomor max 2-3 kali\r\n\r\nAkumulasi maksimal 365 hari ya\r\nJika pengisian produk dilakukan di kartu Three dengan masa aktif lebih dari 365 hari maka transaksi tetap akan sukses namun masa aktif tidak bertambah",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA90HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 3 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 3 BULAN (90 Hari)\r\n\r\nAkumulasi masa aktif pastikan tidak melebihi 360 hari.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "TSELMA30HARI_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu TELKOMSEL 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF TELKOMSEL :v\r\n\r\nTambah Masa Aktif Kartu TELKOMSEL 1 BULAN (30 Hari)\r\n\r\n- Akumulasi masa aktif pastikan tidak melebihi 360 hari",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ISATMA30D_PREV",
            "package_name": "[RESMI] Tambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)",
            "package_name_alias_short": "Tambah Masa Aktif Kartu INDOSAT 1 BULAN",
            "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF INDOSAT :v\r\n\r\nTambah Masa Aktif Kartu INDOSAT 1 BULAN (30 Hari)\r\n\r\nSEBELUM ORDER, BACA PESAN PENTING BERIKUT:\r\n\r\n✅ PRODUK AKTIF 24 JAM SELAMA GAK MAINTENANCE/PEMELIHARAAN SISTEM DARI INDOSAT\r\n✅ AKUMULASI MAKSIMAL 30 HARI\r\n\r\nSEBAGAI CONTOH:\r\nMASA AKTIF NOMOR SAAT MEMBELI TERSISA 25 HARI MAKA YANG DITERIMA HANYA 5 HARI, JIKA MASA AKTIF TERSISA 30 HARI ATAU LEBIH MAKA TRANSAKSI TETAP SUKSES DAN ANDA TIDAK MENDAPATKAN TAMBAHAN MASA AKTIF",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": true,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "00:26",
                "prohibited_hour_endtime": "23:36"
            },
            "need_check_stock": false,
            "is_show_payment_method": false,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIUXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOKXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIU_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu ketika semua kuota utama paket habis.\r\n- Masa Aktif mengikut paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini (wajib ditembak dulu paket Xtra Combo Special 8GB/14GB nya).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 100-150GB an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan Pulsa minimal Rp25.000 (Akan terpotong segini).\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOK_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok ketika semua kuota utama paket habis.\r\n\r\n- Khusus untuk trik XUT Addon, silakan tembak paket ini sampai muncul notif \"Paket berhasil dibeli. Silakan cek kuotanya....\". JIKA NOTIF PERTAMA \"Paket berhasil dibeli. Silakan cek kuotanya....\" MAKA BERIKAN JEDA SEKITAR 10 MENIT.\r\n\r\nSEMBARI MENUNGGU 10 MENIT, KAMU MUNGKIN BISA LANJUT KE ADDON LAINNYA AGAR TIDAK SALING TUNGGU, INTINYA JIKA PERTAMA MUNCUL NOTIF \"Paket berhasil dibeli...\" maka tunggu juga selama 10 menit.\r\n\r\nSETELAH ITU, BELI LAGI HINGGA MUNCUL NOTIF \"422 -> Failed call ipaas purchase, with status code:422 : null,\" (DIPAKSA) DI MASING MASING ADDON. INI TUJUAN YANG KITA CARI-CARI, HARUS ADA RESPON 422 SEPERTI INI.\r\n\r\nDI PALING AKHIR JIKA SEMUA ADDON SUDAH DIRITUALKAN DAN DAPAT RESPON \"422 -> Failed call ipaas purchase, with status code:422 : null,\", TINGGAL GANDENG DENGAN XTRA COMBO SPECIAL 8GB, DAN TUNGGU SELAMA MAX 1-2 JAM NANTI SEMUA ADDON AKAN MASUK. SELAMA MASA TUNGGU 1-2 JAM, SEBAIKNYA JANGAN BELI PAKET APA-APA LAGI DULU.\r\n\r\nSEBELUM MELAKUKAN RITUAL DI ATAS, PASTIKAN TIDAK ADA XTRA COMBO APAPUN, WAJIB CEK VIAL DIAL *808# > INFO > INFO Kartu XL-Ku > Stop Langganan. Unreg sampai bersih dari berbau Xtra Combo 100GB, Xtra Combo Plus, Xtra Combo VIP Plus, Xtra Hotrod, dan sejenisnya (kalau Xtra Combo Flex biarkan saja).\r\n\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BONUS_XCP_2GB_FB_1000_P",
            "package_name": "[Method Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_description": "- Bonus/Addon Kuota 2GB akses Facebook untuk Paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Paket akan memotong Pulsa Rp1.000.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 2GB lagi.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BONUS_XCP_2GB_FB_1000_D",
            "package_name": "[Method E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_description": "- Bonus/Addon Kuota 2GB akses Facebook untuk Paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Sediakan Saldo E-wallet Dana kamu minimal Rp1.000 (akan terpotong segini), klik tombol BAYAR yang muncul setelah beli paket untuk memproses pembelian paket ke laman payment gateway Offical Dana.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 2GB lagi.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_5GB_2C5K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_name_alias_short": "Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_description": "Xtra Edukasi 5GB Rp2.500 3 Hari\r\nButuh Pulsa Rp2.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_15GB_7C5K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_name_alias_short": "Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_description": "Xtra Edukasi 15GB Rp7.500 7 Hari\r\nButuh Pulsa Rp7.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_2GB_1K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_name_alias_short": "Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_description": "Xtra Edukasi 2GB Rp1.000 1 Hari\r\nButuh Pulsa Rp1.000",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_2GB_1K_DANA",
            "package_name": "[Metode E-Wallet] Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_name_alias_short": "Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_description": "Xtra Edukasi 2GB Rp1.000 1 Hari\r\nButuh Saldo E-Wallet DANA Rp1.000",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_5GB_2C5K_DANA",
            "package_name": "[Metode E-Wallet] Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_name_alias_short": "Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_description": "Xtra Edukasi 5GB Rp2.500 3 Hari\r\nButuh Saldo E-Wallet DANA Rp2.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO","payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_D",
            "package_name": "[Metode E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN SALDO E-WALLET DANA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000). SEGERA KLIK TOMBOL BAYAR YANG MUNCUL.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_WA_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus WhatsApp 10GB",
            "package_description": "- Bonus Kuota 10GB WhatsApp Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_TT_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus TikTok 10GB",
            "package_description": "- Bonus Kuota 10GB TikTok Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_FB_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 10GB",
            "package_description": "- Bonus Kuota 10GB Facebook Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "ADN_XCP_IG_10GB_3K_P",
            "package_name": "[Metode Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Instagram 10GB",
            "package_description": "- Bonus Kuota 10GB Instagram Paket Xtra Combo Plus/Xtra Combo VIP Plus. Cocok untuk Mix Paket XCP/XCVIP Plus ke pelanggan.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 10GB lagi.\r\n- SEDIAKAN PULSA MINIMAL RP.3.000 (AKAN TERPOTONG RP3.000).\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIUXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOKXC",
            "package_name": "[Method E-Wallet] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok secara Unlimited.\r\n- Masa Aktif mengikuti paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini.\r\n- FUP paket random/acak tergantung kartu XL dan lokasi kamu. Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan saldo E-wallet Dana Rp25.000 (Akan terpotong segini). Klik tombol Bayar ketika sudah Beli Paket.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOVIU_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo Viu untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo Viu untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo Viu untuk akses Viu ketika semua kuota utama paket habis.\r\n- Masa Aktif mengikut paket Xtra Combo Special 8GB/14GB yang dapat ditembak di panel tembak XL Reborn ini (wajib ditembak dulu paket Xtra Combo Special 8GB/14GB nya).\r\n- Dapat dipakai untuk Inject SSH/VPN (bagi yang tau-tau saja).\r\n- FUP paket sekitar 100-150GB an (random/acak tergantung kartu XL dan lokasi kamu). Jika terkena FUP, saat diinject untuk VPN akan terasa lemot/tidak bisa konek. Solusinya adalah bisa tembak ulang paket Xtra Unlimited Turbo ini maka FUP akan reset.\r\n- Sediakan Pulsa minimal Rp25.000 (Akan terpotong segini).\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XLUNLITURBOTIKTOK_PULSA",
            "package_name": "[Method Pulsa] Unlimited Turbo TikTok untuk Xtra Combo",
            "package_name_alias_short": "Unlimited Turbo TikTok untuk Xtra Combo",
            "package_description": "- Xtra Unlimited Turbo TikTok untuk akses TikTok ketika semua kuota utama paket habis.\r\n\r\n- Khusus untuk trik XUT Addon, silakan tembak paket ini sampai muncul notif \"Paket berhasil dibeli. Silakan cek kuotanya....\". JIKA NOTIF PERTAMA \"Paket berhasil dibeli. Silakan cek kuotanya....\" MAKA BERIKAN JEDA SEKITAR 10 MENIT.\r\n\r\nSEMBARI MENUNGGU 10 MENIT, KAMU MUNGKIN BISA LANJUT KE ADDON LAINNYA AGAR TIDAK SALING TUNGGU, INTINYA JIKA PERTAMA MUNCUL NOTIF \"Paket berhasil dibeli...\" maka tunggu juga selama 10 menit.\r\n\r\nSETELAH ITU, BELI LAGI HINGGA MUNCUL NOTIF \"422 -> Failed call ipaas purchase, with status code:422 : null,\" (DIPAKSA) DI MASING MASING ADDON. INI TUJUAN YANG KITA CARI-CARI, HARUS ADA RESPON 422 SEPERTI INI.\r\n\r\nDI PALING AKHIR JIKA SEMUA ADDON SUDAH DIRITUALKAN DAN DAPAT RESPON \"422 -> Failed call ipaas purchase, with status code:422 : null,\", TINGGAL GANDENG DENGAN XTRA COMBO SPECIAL 8GB, DAN TUNGGU SELAMA MAX 1-2 JAM NANTI SEMUA ADDON AKAN MASUK. SELAMA MASA TUNGGU 1-2 JAM, SEBAIKNYA JANGAN BELI PAKET APA-APA LAGI DULU.\r\n\r\nSEBELUM MELAKUKAN RITUAL DI ATAS, PASTIKAN TIDAK ADA XTRA COMBO APAPUN, WAJIB CEK VIAL DIAL *808# > INFO > INFO Kartu XL-Ku > Stop Langganan. Unreg sampai bersih dari berbau Xtra Combo 100GB, Xtra Combo Plus, Xtra Combo VIP Plus, Xtra Hotrod, dan sejenisnya (kalau Xtra Combo Flex biarkan saja).\r\n\r\n- Jika untuk trik XUT Addon, sebaiknya gunakan pulsa di bawah 10K (tetep nanti harus digandeng Xtra Combo Special 8GB).",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 2000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BONUS_XCP_2GB_FB_1000_P",
            "package_name": "[Method Pulsa] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_description": "- Bonus/Addon Kuota 2GB akses Facebook untuk Paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Paket akan memotong Pulsa Rp1.000.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 2GB lagi.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "BONUS_XCP_2GB_FB_1000_D",
            "package_name": "[Method E-Wallet] Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_name_alias_short": "Bonus Kuota Xtra Combo Plus/VIP Plus Facebook 2GB",
            "package_description": "- Bonus/Addon Kuota 2GB akses Facebook untuk Paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Sediakan Saldo E-wallet Dana kamu minimal Rp1.000 (akan terpotong segini), klik tombol BAYAR yang muncul setelah beli paket untuk memproses pembelian paket ke laman payment gateway Offical Dana.\r\n- Harus ada paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Masa aktif paket mengikuti sisa masa aktif paket Xtra Combo Plus/Xtra Combo VIP Plus.\r\n- Tidak dapat diakumulasi, jika kamu membelinya ulang maka kuota akan tereset ke 2GB lagi.\r\n- Jika tidak dapat dibeli, silakan coba secara berkala.",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 1000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "23:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_5GB_2C5K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_name_alias_short": "Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_description": "Xtra Edukasi 5GB Rp2.500 3 Hari\r\nButuh Pulsa Rp2.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_15GB_7C5K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_name_alias_short": "Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_description": "Xtra Edukasi 15GB Rp7.500 7 Hari\r\nButuh Pulsa Rp7.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_2GB_1K",
            "package_name": "[Metode Pulsa] Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_name_alias_short": "Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_description": "Xtra Edukasi 2GB Rp1.000 1 Hari\r\nButuh Pulsa Rp1.000",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": true,
            "can_scheduled_trx": true,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "BALANCE",
                    "payment_method_display_name": "Metode Pulsa (BALANCE)",
                    "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa (baca dan pahami deskripsi paketnya)"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_2GB_1K_DANA",
            "package_name": "[Metode E-Wallet] Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_name_alias_short": "Xtra Edukasi 2GB Rp1.000 1 Hari",
            "package_description": "Xtra Edukasi 2GB Rp1.000 1 Hari\r\nButuh Saldo E-Wallet DANA Rp1.000",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_5GB_2C5K_DANA",
            "package_name": "[Metode E-Wallet] Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_name_alias_short": "Xtra Edukasi 5GB Rp2.500 3 Hari",
            "package_description": "Xtra Edukasi 5GB Rp2.500 3 Hari\r\nButuh Saldo E-Wallet DANA Rp2.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        },
        {
            "package_code": "XL_EDU_15GB_7C5K_DANA",
            "package_name": "[Metode E-Wallet] Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_name_alias_short": "Xtra Edukasi 15GB Rp7.500 7 Hari",
            "package_description": "Xtra Edukasi 15GB Rp7.500 7 Hari\r\nButuh Saldo E-Wallet DANA Rp7.500",
            "have_daily_limit": true,
            "daily_limit_details": {
                "max_daily_transaction_limit": 5000,
                "current_daily_transaction_count": 0
            },
            "no_need_login": false,
            "can_multi_trx": false,
            "can_scheduled_trx": false,
            "have_cut_off_time": false,
            "cut_off_time": {
                "prohibited_hour_starttime": "22:00",
                "prohibited_hour_endtime": "02:00"
            },
            "need_check_stock": false,
            "is_show_payment_method": true,
            "available_payment_methods": [
                {
                    "order": 1,
                    "payment_method": "DANA",
                    "payment_method_display_name": "Metode DANA",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet DANA"
                },
                {
                    "order": 2,
                    "payment_method": "QRIS",
                    "payment_method_display_name": "Metode QRIS",
                    "desc": "Bisa dibayar di semua e-wallet / mbanking yang telah mendukung fitur pembayaran QRIS"
                },
                {
                    "order": 3,
                    "payment_method": "GOPAY",
                    "payment_method_display_name": "Metode GOPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet GOPAY"
                },
                {
                    "order": 4,
                    "payment_method": "SHOPEEPAY",
                    "payment_method_display_name": "Metode SHOPEEPAY",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet SHOPEEPAY"
                },
                {
                    "order": 5,
                    "payment_method": "OVO",
                    "payment_method_display_name": "Metode OVO",
                    "desc": "Bisa dibayar dengan menggunakan E-Wallet OVO"
                }
            ],
            "is_akrab": false,
            "is_circle": false,
            "sedang_gangguan": false
        }
    ]
}

   
