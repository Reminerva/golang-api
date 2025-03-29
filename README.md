# Golang API

## 📌 Deskripsi
Golang API adalah layanan berbasis REST API yang memungkinkan pengguna untuk melakukan transaksi keuangan, mengelola data merchant, serta melakukan autentikasi pengguna.

## 📥 Instalasi dan Menjalankan API

### 1️⃣ Clone Repository
```sh
git clone https://github.com/username/bank-api.git
cd bank-api
```

### 3️⃣ Menjalankan API
API akan berjalan di `http://localhost:8080`

## 🚀 Endpoint API

### 📌 1. **Autentikasi**
| Method | Endpoint | Deskripsi |
|--------|---------|-----------|
| `POST` | `/auth/register` | Registrasi pengguna baru |
| `POST` | `/auth/login` | Login pengguna |
| `POST` | `/auth/logout` | Logout pengguna |
| `GET` | `/auth` | Mendapatkan semua data auth |

### 📌 2. **Customer**
| Method | Endpoint | Deskripsi |
|--------|---------|-----------|
| `GET` | `/customers` | Mendapatkan semua customer |
| `GET` | `/customers/id/{id}` | Mendapatkan customer berdasarkan ID |
| `GET` | `/customers/email/{email}` | Mendapatkan customer berdasarkan email |

### 📌 3. **Merchant**
| Method | Endpoint | Deskripsi |
|--------|---------|-----------|
| `POST` | `/merchants` | Menambahkan merchant baru |
| `GET` | `/merchants` | Mendapatkan semua merchant |
| `GET` | `/merchants/id/{id}` | Mendapatkan merchant berdasarkan ID |

### 📌 4. **Transaksi**
| Method | Endpoint | Deskripsi |
|--------|---------|-----------|
| `POST` | `/transactions` | Memproses transaksi |
| `GET` | `/transactions` | Mendapatkan semua transaksi |
| `GET` | `/transactions/customer/{customerID}` | Mendapatkan transaksi berdasarkan ID pelanggan |
| `GET` | `/transactions/merchant/{merchantID}` | Mendapatkan transaksi berdasarkan ID merchant |

## 📝 Contoh Request
### 📌 Register User
**Request:**
```json
{
  "name": "Panda",
  "email": "panda@gmail.com",
  "password": "123456"
}
```

### 📌 Login User
**Request:**
```json
{
  "email": "panda@gmail.com",
  "password": "123456"
}
```

### 📌 Create Merchant
**Request:**
```json
{
  "name": "Merchant A",
  "balance": 11000,
  "address": "Jl. mana"
}
```

### 📌 Create Transaction
**Request:**
```json
{
  "customer_id": "c-4806",
  "merchant_id": "m-7486",
  "amount": 1000
}
```
