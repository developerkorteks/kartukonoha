#!/usr/bin/env python3
"""
Nadia API v2.0 - Python Example
Usage: python3 python_example.py
"""

import requests
import json
import time
from typing import Dict, Any, Optional

API_BASE = 'http://localhost:8080/api'

class NadiaAPI:
    def __init__(self, base_url: str = API_BASE):
        self.base_url = base_url
        self.session = requests.Session()
        self.session.headers.update({
            'Content-Type': 'application/json',
            'User-Agent': 'Nadia-Python-Client/2.0'
        })
    
    def _make_request(self, method: str, endpoint: str, data: Optional[Dict] = None) -> Dict[str, Any]:
        """Make API request with error handling"""
        url = f"{self.base_url}{endpoint}"
        
        try:
            if method.upper() == 'GET':
                response = self.session.get(url, params=data)
            else:
                response = self.session.request(method, url, json=data)
            
            response.raise_for_status()
            return response.json()
            
        except requests.exceptions.RequestException as e:
            print(f"❌ API request failed: {e}")
            raise
        except json.JSONDecodeError as e:
            print(f"❌ Failed to parse JSON response: {e}")
            raise
    
    def search_packages(self, query: str = None, payment_method: str = None, 
                       max_price: int = None, min_price: int = None) -> Dict[str, Any]:
        """Search for packages with filters"""
        print("🔍 Searching for packages...")
        
        search_data = {}
        if query:
            search_data['query'] = query
        if payment_method:
            search_data['payment_method'] = payment_method
        if max_price:
            search_data['max_price'] = max_price
        if min_price:
            search_data['min_price'] = min_price
        
        result = self._make_request('POST', '/packages', search_data)
        
        if result.get('success'):
            packages = result.get('data', [])
            print(f"✅ Found {len(packages)} packages")
            
            if packages:
                first_package = packages[0]
                print(f"📦 First package: {first_package.get('package_name')}")
                print(f"💰 Price: Rp {first_package.get('package_price'):,}")
                print(f"🔑 Code: {first_package.get('package_code')}")
        else:
            print(f"❌ Search failed: {result.get('message')}")
        
        return result
    
    def request_otp(self, phone_number: str) -> Dict[str, Any]:
        """Request OTP for phone number"""
        print(f"📱 Requesting OTP for {phone_number}...")
        
        result = self._make_request('POST', '/otp/request', {
            'phone_number': phone_number
        })
        
        if result.get('success'):
            print("✅ OTP sent successfully!")
            print("📨 Please check your SMS for the OTP code")
            expires_in = result.get('data', {}).get('expires_in', 300)
            print(f"⏰ OTP expires in {expires_in} seconds")
        else:
            print(f"❌ OTP request failed: {result.get('message')}")
        
        return result
    
    def purchase_package(self, phone_number: str, package_code: str, 
                        payment_method: str, otp_code: str) -> Dict[str, Any]:
        """Purchase package with OTP"""
        print("🛒 Purchasing package...")
        
        purchase_data = {
            'phone_number': phone_number,
            'package_code': package_code,
            'payment_method': payment_method,
            'otp_code': otp_code
        }
        
        result = self._make_request('POST', '/purchase', purchase_data)
        
        if result.get('success'):
            print("🎉 Purchase successful!")
            data = result.get('data', {})
            print(f"📦 Package: {data.get('package_name')}")
            print(f"📱 Phone: {data.get('msisdn')}")
            print(f"🆔 Transaction ID: {data.get('trx_id')}")
            if data.get('package_processing_fee'):
                print(f"💳 Processing fee: Rp {data.get('package_processing_fee'):,}")
        else:
            print(f"❌ Purchase failed: {result.get('message')}")
        
        return result
    
    def check_card_status(self, phone_number: str, otp_code: str) -> Dict[str, Any]:
        """Check XL card status"""
        print("🏥 Checking card status...")
        
        result = self._make_request('POST', '/card/status', {
            'phone_number': phone_number,
            'otp_code': otp_code
        })
        
        if result.get('success'):
            print("✅ Card status retrieved!")
            data = result.get('data', {})
            print(f"📱 Phone: {data.get('msisdn')}")
            print(f"📊 Status: {data.get('subscription_status')}")
            print(f"💰 Balance: {data.get('pulsa_real')}")
            print(f"📅 Active until: {data.get('active_until')}")
            print(f"📍 Location: {data.get('location')}")
        else:
            print(f"❌ Status check failed: {result.get('message')}")
        
        return result
    
    def check_active_packages(self, phone_number: str, otp_code: str) -> Dict[str, Any]:
        """Check active packages on card"""
        print("📦 Checking active packages...")
        
        result = self._make_request('POST', '/card/packages', {
            'phone_number': phone_number,
            'otp_code': otp_code
        })
        
        if result.get('success'):
            print("✅ Active packages retrieved!")
            data = result.get('data', {})
            quotas = data.get('quotas', [])
            print(f"📦 Found {len(quotas)} active packages")
            
            for i, quota in enumerate(quotas, 1):
                print(f"  {i}. {quota.get('name')} - Expires: {quota.get('expired_at')}")
        else:
            print(f"❌ Package check failed: {result.get('message')}")
        
        return result
    
    def get_balance(self) -> Dict[str, Any]:
        """Get account balance"""
        print("💰 Getting account balance...")
        
        result = self._make_request('GET', '/balance')
        
        if result.get('success'):
            balance = result.get('data', {}).get('balance', 'N/A')
            print(f"💳 Current balance: {balance}")
        else:
            print(f"❌ Failed to get balance: {result.get('message')}")
        
        return result
    
    def check_transaction(self, transaction_id: str) -> Dict[str, Any]:
        """Check transaction status"""
        print(f"📊 Checking transaction {transaction_id}...")
        
        result = self._make_request('POST', '/transaction/check', {
            'transaction_id': transaction_id
        })
        
        if result.get('success'):
            print("✅ Transaction status retrieved!")
            data = result.get('data', {})
            print(f"📋 Status: {data}")
        else:
            print(f"❌ Transaction check failed: {result.get('message')}")
        
        return result
    
    def get_package_stock(self) -> Dict[str, Any]:
        """Get package stock information"""
        print("📊 Getting package stock...")
        
        result = self._make_request('GET', '/packages/stock')
        
        if result.get('success'):
            stocks = result.get('data', [])
            print(f"📦 Found stock info for {len(stocks)} packages")
            
            # Show packages with stock > 0
            available = [s for s in stocks if s.get('stok', 0) > 0]
            print(f"✅ {len(available)} packages available in stock")
        else:
            print(f"❌ Failed to get stock: {result.get('message')}")
        
        return result

def complete_purchase_flow_demo():
    """Demo of complete purchase flow"""
    print("🚀 Complete Purchase Flow Demo")
    print("==============================")
    
    api = NadiaAPI()
    phone_number = "087786388052"  # Replace with actual XL number
    
    try:
        # Step 1: Search for packages
        print("\n1️⃣ Searching for affordable packages...")
        packages = api.search_packages(
            query="masa aktif",
            payment_method="BALANCE",
            max_price=5000
        )
        
        if not packages.get('success') or not packages.get('data'):
            print("❌ No packages found, stopping demo")
            return
        
        # Get first package
        package = packages['data'][0]
        package_code = package['package_code']
        package_name = package['package_name']
        
        print(f"\n📦 Selected package: {package_name}")
        print(f"🔑 Package code: {package_code}")
        
        # Step 2: Request OTP
        print(f"\n2️⃣ Requesting OTP for {phone_number}...")
        otp_result = api.request_otp(phone_number)
        
        if not otp_result.get('success'):
            print("❌ OTP request failed, stopping demo")
            return
        
        # In real application, user would enter OTP from SMS
        print("\n💡 In real application, user would enter OTP from SMS")
        print("💡 For demo purposes, we'll simulate OTP entry")
        
        # Simulate user input (in real app, use input())
        otp_code = "123456"  # This should come from user input
        print(f"🔢 Simulated OTP: {otp_code}")
        
        # Step 3: Purchase package
        print(f"\n3️⃣ Purchasing package...")
        purchase_result = api.purchase_package(
            phone_number=phone_number,
            package_code=package_code,
            payment_method="BALANCE",
            otp_code=otp_code
        )
        
        if purchase_result.get('success'):
            transaction_id = purchase_result.get('data', {}).get('trx_id')
            
            # Step 4: Check transaction status
            if transaction_id:
                print(f"\n4️⃣ Checking transaction status...")
                time.sleep(2)  # Wait a bit before checking
                api.check_transaction(transaction_id)
        
        print("\n✅ Purchase flow demo completed!")
        
    except Exception as e:
        print(f"❌ Demo failed: {e}")

def card_management_demo():
    """Demo of card management features"""
    print("🏥 Card Management Demo")
    print("=======================")
    
    api = NadiaAPI()
    phone_number = "087786388052"  # Replace with actual XL number
    
    try:
        # Step 1: Request OTP
        print(f"\n1️⃣ Requesting OTP for {phone_number}...")
        otp_result = api.request_otp(phone_number)
        
        if not otp_result.get('success'):
            print("❌ OTP request failed, stopping demo")
            return
        
        # Simulate OTP entry
        otp_code = "123456"
        print(f"🔢 Simulated OTP: {otp_code}")
        
        # Step 2: Check card status
        print(f"\n2️⃣ Checking card status...")
        api.check_card_status(phone_number, otp_code)
        
        # Step 3: Check active packages
        print(f"\n3️⃣ Checking active packages...")
        api.check_active_packages(phone_number, otp_code)
        
        print("\n✅ Card management demo completed!")
        
    except Exception as e:
        print(f"❌ Demo failed: {e}")

def main():
    """Main demo function"""
    print("🚀 Nadia API v2.0 - Python Demo")
    print("================================")
    
    api = NadiaAPI()
    
    try:
        # Demo 1: Basic API calls
        print("\n📊 Basic API Calls Demo")
        print("-----------------------")
        
        # Get balance
        api.get_balance()
        
        # Search packages
        api.search_packages(query="combo", max_price=10000)
        
        # Get package stock
        api.get_package_stock()
        
        # Demo 2: Complete flows (commented to avoid actual API calls)
        print("\n💡 Advanced demos available:")
        print("   - complete_purchase_flow_demo() - Full purchase process")
        print("   - card_management_demo() - Card status and package checking")
        print("   - Uncomment in main() to test with real data")
        
        # Uncomment these to test with real data:
        # complete_purchase_flow_demo()
        # card_management_demo()
        
        print("\n✅ Demo completed successfully!")
        
    except Exception as e:
        print(f"❌ Demo failed: {e}")

if __name__ == "__main__":
    main()