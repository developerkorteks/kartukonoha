# 📦 Packages API Guide - GET vs POST

Penjelasan lengkap perbedaan antara endpoint GET dan POST untuk packages dalam Nadia API v2.0.

## 🔍 Overview

Ada **2 endpoint** berbeda untuk mengakses data packages:

1. **GET /api/packages** - Mengambil semua paket tanpa filter
2. **POST /api/packages/search** - Mencari paket dengan filter kompleks

## 📋 Perbandingan Detail

| Aspek | GET /api/packages | POST /api/packages/search |
|-------|-------------------|---------------------------|
| **Method** | GET | POST |
| **Body** | Tidak ada | JSON dengan filter |
| **Query Params** | `?limit=100` | Tidak ada |
| **Fungsi** | Ambil semua paket | Cari paket dengan filter |
| **Performance** | Lebih cepat | Sedikit lebih lambat (filtering) |
| **Use Case** | List semua paket | Pencarian spesifik |

## 🚀 1. GET /api/packages

### Deskripsi
Mengambil **semua paket** yang tersedia tanpa filtering. Cocok untuk:
- Menampilkan daftar lengkap paket
- Dropdown/select options
- Cache data paket
- Overview semua paket

### Request
```bash
GET /api/packages?limit=50
```

### Query Parameters
- `limit` (optional): Batasi jumlah paket yang dikembalikan (default: 100)

### Response
```json
{
  "statusCode": 200,
  "message": "Retrieved 50 packages",
  "success": true,
  "data": [
    {
      "package_code": "XL_MASTIF_30D_P_V1",
      "package_name": "XL Masa Aktif 30 Hari",
      "package_price": 2000,
      "available_payment_methods": [
        {
          "payment_method": "BALANCE",
          "payment_method_display_name": "Saldo"
        }
      ]
    }
    // ... more packages
  ]
}
```

### Contoh Penggunaan

#### cURL
```bash
curl -X GET "http://localhost:8080/api/packages?limit=20"
```

#### JavaScript
```javascript
const response = await fetch('http://localhost:8080/api/packages?limit=20');
const data = await response.json();
console.log(`Found ${data.data.length} packages`);
```

#### Python
```python
import requests

response = requests.get('http://localhost:8080/api/packages', params={'limit': 20})
data = response.json()
print(f"Found {len(data['data'])} packages")
```

## 🔍 2. POST /api/packages/search

### Deskripsi
Mencari paket dengan **filter kompleks**. Cocok untuk:
- Pencarian berdasarkan nama/deskripsi
- Filter berdasarkan harga
- Filter berdasarkan metode pembayaran
- Kombinasi multiple filter

### Request
```bash
POST /api/packages/search
Content-Type: application/json

{
  "query": "masa aktif",
  "payment_method": "BALANCE",
  "max_price": 10000,
  "min_price": 1000
}
```

### Request Body (PackageSearchRequest)
```json
{
  "query": "string (optional)",         // Cari dalam nama/deskripsi
  "payment_method": "string (optional)", // BALANCE, DANA, QRIS
  "max_price": "int (optional)",        // Harga maksimal
  "min_price": "int (optional)"         // Harga minimal
}
```

### Response
```json
{
  "statusCode": 200,
  "message": "Found 15 packages matching your criteria",
  "success": true,
  "data": [
    {
      "package_code": "XL_MASTIF_30D_P_V1",
      "package_name": "XL Masa Aktif 30 Hari",
      "package_price": 2000,
      "available_payment_methods": [
        {
          "payment_method": "BALANCE",
          "payment_method_display_name": "Saldo"
        }
      ]
    }
    // ... filtered packages only
  ]
}
```

### Contoh Penggunaan

#### cURL - Cari paket murah
```bash
curl -X POST http://localhost:8080/api/packages/search \
  -H "Content-Type: application/json" \
  -d '{
    "max_price": 5000,
    "payment_method": "BALANCE"
  }'
```

#### cURL - Cari paket masa aktif
```bash
curl -X POST http://localhost:8080/api/packages/search \
  -H "Content-Type: application/json" \
  -d '{
    "query": "masa aktif",
    "min_price": 1000,
    "max_price": 10000
  }'
```

#### JavaScript - Pencarian kompleks
```javascript
const searchData = {
  query: "combo",
  payment_method: "DANA",
  max_price: 15000,
  min_price: 5000
};

const response = await fetch('http://localhost:8080/api/packages/search', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify(searchData)
});

const data = await response.json();
console.log(`Found ${data.data.length} packages matching criteria`);
```

#### Python - Filter berdasarkan metode bayar
```python
import requests

search_data = {
    "payment_method": "QRIS",
    "max_price": 20000
}

response = requests.post(
    'http://localhost:8080/api/packages/search',
    json=search_data
)

data = response.json()
print(f"Found {len(data['data'])} packages with QRIS payment")
```

## 🎯 Kapan Menggunakan Yang Mana?

### Gunakan GET /api/packages ketika:
- ✅ Butuh semua paket tanpa filter
- ✅ Membuat dropdown/select options
- ✅ Caching data paket
- ✅ Performance adalah prioritas
- ✅ Implementasi sederhana

### Gunakan POST /api/packages/search ketika:
- ✅ Butuh pencarian berdasarkan kata kunci
- ✅ Filter berdasarkan harga
- ✅ Filter berdasarkan metode pembayaran
- ✅ Kombinasi multiple filter
- ✅ User experience yang lebih baik

## 🔄 Filter Options untuk Search

### 1. Text Search (`query`)
Mencari dalam `package_name` dan `package_description`:
```json
{
  "query": "masa aktif"     // Cari paket yang mengandung "masa aktif"
}
```

### 2. Payment Method Filter (`payment_method`)
Filter berdasarkan metode pembayaran yang tersedia:
```json
{
  "payment_method": "BALANCE"  // Hanya paket yang bisa dibayar dengan saldo
}
```

**Available Payment Methods:**
- `BALANCE` - Saldo akun
- `DANA` - E-wallet DANA
- `QRIS` - QR Code payment

### 3. Price Range Filter (`min_price`, `max_price`)
Filter berdasarkan rentang harga:
```json
{
  "min_price": 1000,    // Minimal Rp 1.000
  "max_price": 10000    // Maksimal Rp 10.000
}
```

### 4. Kombinasi Filter
Semua filter bisa dikombinasikan:
```json
{
  "query": "combo",
  "payment_method": "DANA",
  "min_price": 5000,
  "max_price": 15000
}
```

## 📊 Performance Comparison

| Endpoint | Response Time | Data Size | Best For |
|----------|---------------|-----------|----------|
| GET /api/packages | ~200ms | Full dataset | Initial load, caching |
| POST /api/packages/search | ~250ms | Filtered dataset | User searches |

## 🛠️ Implementation Examples

### React Component
```jsx
import React, { useState, useEffect } from 'react';

function PackageList() {
  const [packages, setPackages] = useState([]);
  const [searchFilters, setSearchFilters] = useState({});
  
  // Load all packages on mount
  useEffect(() => {
    fetch('/api/packages?limit=50')
      .then(res => res.json())
      .then(data => setPackages(data.data));
  }, []);
  
  // Search with filters
  const searchPackages = async (filters) => {
    const response = await fetch('/api/packages/search', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(filters)
    });
    const data = await response.json();
    setPackages(data.data);
  };
  
  return (
    <div>
      <SearchForm onSearch={searchPackages} />
      <PackageGrid packages={packages} />
    </div>
  );
}
```

### Vue.js Component
```vue
<template>
  <div>
    <search-form @search="searchPackages" />
    <package-list :packages="packages" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      packages: []
    };
  },
  
  async mounted() {
    // Load all packages
    const response = await fetch('/api/packages?limit=50');
    const data = await response.json();
    this.packages = data.data;
  },
  
  methods: {
    async searchPackages(filters) {
      const response = await fetch('/api/packages/search', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(filters)
      });
      const data = await response.json();
      this.packages = data.data;
    }
  }
};
</script>
```

## 🚨 Error Handling

### GET /api/packages
```json
{
  "statusCode": 500,
  "message": "Failed to fetch packages: connection timeout",
  "success": false
}
```

### POST /api/packages/search
```json
{
  "statusCode": 400,
  "message": "Invalid search request: max_price must be greater than min_price",
  "success": false
}
```

## 📝 Best Practices

### 1. Caching Strategy
```javascript
// Cache all packages for dropdown
const allPackages = await fetch('/api/packages').then(r => r.json());
localStorage.setItem('packages_cache', JSON.stringify(allPackages));

// Use search for user queries
const searchResults = await fetch('/api/packages/search', {
  method: 'POST',
  body: JSON.stringify({ query: userInput })
}).then(r => r.json());
```

### 2. Progressive Loading
```javascript
// Load basic list first
const basicPackages = await fetch('/api/packages?limit=20').then(r => r.json());
displayPackages(basicPackages.data);

// Then allow search
document.getElementById('search').addEventListener('input', debounce(async (e) => {
  if (e.target.value.length > 2) {
    const results = await fetch('/api/packages/search', {
      method: 'POST',
      body: JSON.stringify({ query: e.target.value })
    }).then(r => r.json());
    displayPackages(results.data);
  }
}, 300));
```

### 3. Error Handling
```javascript
async function getPackages(useSearch = false, filters = {}) {
  try {
    let response;
    
    if (useSearch) {
      response = await fetch('/api/packages/search', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(filters)
      });
    } else {
      response = await fetch('/api/packages');
    }
    
    const data = await response.json();
    
    if (!data.success) {
      throw new Error(data.message);
    }
    
    return data.data;
    
  } catch (error) {
    console.error('Failed to fetch packages:', error);
    return [];
  }
}
```

---

**Summary:**
- **GET /api/packages** = Ambil semua paket (simple, fast)
- **POST /api/packages/search** = Cari paket dengan filter (advanced, flexible)

Pilih endpoint yang sesuai dengan kebutuhan aplikasi Anda!