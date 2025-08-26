// Nadia API v2.0 - JavaScript Example
// Usage: node javascript_example.js

// Configure API base URL - change this to match your server
// For local development: 'http://localhost:8080/api' or 'http://127.0.0.1:8080/api'
// For production: 'https://your-domain.com/api'
const API_BASE = process.env.NADIA_API_BASE || 'http://localhost:8080/api';

// Helper function for API calls
async function apiCall(method, endpoint, data = null) {
    const options = {
        method: method,
        headers: {
            'Content-Type': 'application/json',
        },
    };
    
    if (data) {
        options.body = JSON.stringify(data);
    }
    
    try {
        const response = await fetch(`${API_BASE}${endpoint}`, options);
        const result = await response.json();
        return result;
    } catch (error) {
        console.error('API call failed:', error);
        throw error;
    }
}

// Example 1: Search for packages
async function searchPackages() {
    console.log('🔍 Searching for packages...');
    
    const searchData = {
        query: 'masa aktif',
        payment_method: 'BALANCE',
        max_price: 10000
    };
    
    const result = await apiCall('POST', '/packages', searchData);
    console.log(`Found ${result.data?.length || 0} packages`);
    
    if (result.data && result.data.length > 0) {
        console.log('First package:', {
            name: result.data[0].package_name,
            price: result.data[0].package_price,
            code: result.data[0].package_code
        });
    }
}

// Example 2: Complete purchase flow
async function purchasePackage(phoneNumber, packageCode, paymentMethod) {
    console.log('🛒 Starting purchase flow...');
    
    try {
        // Step 1: Request OTP
        console.log('Step 1: Requesting OTP...');
        const otpResult = await apiCall('POST', '/otp/request', {
            phone_number: phoneNumber
        });
        
        if (!otpResult.success) {
            throw new Error(`OTP request failed: ${otpResult.message}`);
        }
        
        console.log('✅ OTP sent successfully!');
        console.log('📱 Please check your SMS and enter the OTP code');
        
        // In real application, you would get OTP from user input
        // For demo purposes, we'll simulate it
        const otpCode = '123456'; // This should come from user input
        
        // Step 2: Purchase with OTP
        console.log('Step 2: Purchasing package...');
        const purchaseResult = await apiCall('POST', '/purchase', {
            phone_number: phoneNumber,
            package_code: packageCode,
            payment_method: paymentMethod,
            otp_code: otpCode
        });
        
        if (purchaseResult.success) {
            console.log('🎉 Purchase successful!');
            console.log('Transaction ID:', purchaseResult.data?.trx_id);
            console.log('Package:', purchaseResult.data?.package_name);
        } else {
            console.log('❌ Purchase failed:', purchaseResult.message);
        }
        
        return purchaseResult;
        
    } catch (error) {
        console.error('Purchase flow failed:', error.message);
        throw error;
    }
}

// Example 3: Check card status
async function checkCardStatus(phoneNumber) {
    console.log('🏥 Checking card status...');
    
    try {
        // Step 1: Request OTP
        const otpResult = await apiCall('POST', '/otp/request', {
            phone_number: phoneNumber
        });
        
        if (!otpResult.success) {
            throw new Error(`OTP request failed: ${otpResult.message}`);
        }
        
        console.log('✅ OTP sent for card status check');
        
        // In real application, get OTP from user
        const otpCode = '123456';
        
        // Step 2: Check status
        const statusResult = await apiCall('POST', '/card/status', {
            phone_number: phoneNumber,
            otp_code: otpCode
        });
        
        if (statusResult.success) {
            console.log('📊 Card Status:', {
                msisdn: statusResult.data?.msisdn,
                status: statusResult.data?.subscription_status,
                balance: statusResult.data?.pulsa_real,
                active_until: statusResult.data?.active_until,
                location: statusResult.data?.location
            });
        } else {
            console.log('❌ Status check failed:', statusResult.message);
        }
        
        return statusResult;
        
    } catch (error) {
        console.error('Card status check failed:', error.message);
        throw error;
    }
}

// Example 4: Get account balance
async function getBalance() {
    console.log('💰 Getting account balance...');
    
    const result = await apiCall('GET', '/balance');
    
    if (result.success) {
        console.log('💳 Balance:', result.data?.balance || 'N/A');
    } else {
        console.log('❌ Failed to get balance:', result.message);
    }
    
    return result;
}

// Example 5: Check transaction status
async function checkTransaction(transactionId) {
    console.log('📊 Checking transaction status...');
    
    const result = await apiCall('POST', '/transaction/check', {
        transaction_id: transactionId
    });
    
    if (result.success) {
        console.log('📋 Transaction Status:', result.data);
    } else {
        console.log('❌ Transaction check failed:', result.message);
    }
    
    return result;
}

// Main demo function
async function runDemo() {
    console.log('🚀 Nadia API v2.0 - JavaScript Demo');
    console.log('===================================');
    
    try {
        // Demo 1: Search packages
        await searchPackages();
        console.log('');
        
        // Demo 2: Get balance
        await getBalance();
        console.log('');
        
        // Demo 3: Purchase flow (commented out to avoid actual purchase)
        // await purchasePackage('087786388052', 'XL_MASTIF_30D_P_V1', 'BALANCE');
        console.log('💡 Purchase demo skipped (uncomment to test with real data)');
        console.log('');
        
        // Demo 4: Card status check (commented out to avoid OTP request)
        // await checkCardStatus('087786388052');
        console.log('💡 Card status demo skipped (uncomment to test with real data)');
        console.log('');
        
        console.log('✅ Demo completed successfully!');
        
    } catch (error) {
        console.error('❌ Demo failed:', error.message);
    }
}

// Run the demo
if (require.main === module) {
    runDemo();
}

// Export functions for use in other modules
module.exports = {
    apiCall,
    searchPackages,
    purchasePackage,
    checkCardStatus,
    getBalance,
    checkTransaction
};