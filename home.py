# PENTING: Skrip ini hanya untuk tujuan edukasi.
# Gunakan dengan bijak dan hormati aturan layanan Telegram serta pemilik bot.
# Risiko pemblokiran akun ditanggung sendiri.

from telethon.sync import TelegramClient
import time
import re

# -----------------------------------------------------------
# KONFIGURASI - Ganti dengan data Anda
# -----------------------------------------------------------
api_id = 22527852
api_hash = '4f595e6aac7dfe58a2cf6051360c3f14' 
phone_number = '+6287742028130' 
target_bot_username = '@CendrawasihStore_bot'
session_file = 'my_telegram_session'
# -----------------------------------------------------------


with TelegramClient(session_file, api_id, api_hash) as client:
    print("Session berhasil dimulai...")
    
    # 1. Mulai percakapan
    print(f"Mengirim perintah /start ke {target_bot_username}...")
    client.send_message(target_bot_username, '/start')
    time.sleep(3)
    
    # 2. Baca balasan pertama
    print("Membaca balasan dari bot...")
    messages = client.get_messages(target_bot_username, limit=1)
    
    if not messages:
        print("Bot tidak memberikan balasan. Skrip berhenti.")
    else:
        last_message = messages[0]
        
        # 3. Ekstrak semua teks tombol dari menu utama
        print("\n--- DAFTAR SEMUA MENU UTAMA ---")
        menu_options = []
        if last_message.buttons:
            # Iterasi melalui setiap baris tombol
            for row in last_message.buttons:
                # Iterasi melalui setiap tombol di dalam baris
                for button in row:
                    menu_options.append(button.text)
            
            # Tampilkan semua menu yang berhasil diekstrak
            for i, option in enumerate(menu_options, 1):
                print(f"{i}. {option}")
        else:
            print("Tidak ada tombol yang ditemukan di pesan ini.")
        print("--------------------------------\n")

        # 4. Contoh: Mengklik tombol spesifik berdasarkan teksnya
        # Ganti teks di bawah ini dengan menu yang ingin Anda klik
        target_button_text = "🌐 Beli Akun VPN" 
        
        print(f"Mencoba mencari dan mengklik tombol: '{target_button_text.strip()}'...")
        
        button_found = False
        if last_message.buttons:
            for row in last_message.buttons:
                for button in row:
                    if button.text == target_button_text:
                        print("Tombol ditemukan! Mengklik...")
                        button_found = True
                        
                        # Klik tombol yang cocok
                        button.click()
                        time.sleep(3) # Tunggu balasan
                        
                        # Baca balasan baru setelah klik
                        new_messages = client.get_messages(target_bot_username, limit=1)
                        if new_messages:
                            print("\n--- BALASAN SETELAH KLIK TOMBOL ---")
                            print(new_messages[0].text)
                            print("-----------------------------------")
                        else:
                            print("Tidak menerima balasan baru setelah klik.")
                        
                        # Hentikan loop setelah tombol ditemukan dan diklik
                        break 
                if button_found:
                    break
        
        if not button_found:
            print(f"Tombol dengan teks '{target_button_text.strip()}' tidak ditemukan.")

print("\nSkrip selesai dijalankan.")
