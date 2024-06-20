# Pembayaran SPP

## Deskripsi Aplikasi

## Fitur dan Role dalam Aplikasi

| Fitur                      | Administrator | Petugas | Siswa |
| -------------------------- | ------------- | ------- | ----- |
| Login                      | √             | √       | √     |
| Logout                     | √             | √       | √     |
| CRUD Data Siswa            | √             |         |       |
| CRUD Data Petugas          | √             |         |       |
| CRUD Data Kelas            | √             |         |       |
| CRUD Data Spp              | √             |         |       |
| Entri Transaksi Pembayaran | √             | √       |       |
| Lihat History Pembayaran   | √             | √       | √     |
| Generate Laporan           | √             |         |       |

## Berikut DDL untuk Pembayaran SPP :

```sql
CREATE DATABASE IF NOT EXISTS pembayaran_spp;

\c pembayaran_spp;

CREATE TYPE level AS ENUM ('ADMIN', 'PETUGAS');

CREATE TABLE IF NOT EXISTS tb_spp(
    id_spp int(11) PRIMARY KEY,
    tahun int(11),
    nominal int(11)
);

CREATE TABLE IF NOT EXISTS tb_kelas(
    id_kelas int(11) PRIMARY KEY,
    nama_kelas varchar(10),
    kompetensi_keahlian varchar(50)
);

CREATE TABLE IF NOT EXISTS tb_petugas(
    id_petugas int(11) PRIMARY KEY,
    username varchar(25),
    password varchar(35),
    nama_petugas varchar(35),
    level level
);

CREATE TABLE IF NOT EXISTS tb_siswa(
    nisn varchar(10) PRIMARY KEY,
    nis varchar(8) UNIQUE NOT NULL,
    nama varchar(35),
    id_kelas int(11) REFERENCES tb_kelas(id_kelas),
    alamat text,
    no_telp varchar(13),
    id_spp int(11) REFERENCES tb_spp(id_spp)
);

CREATE TABLE IF NOT EXISTS tb_pembayaran(
    id_pembayaran int(11) PRIMARY KEY,
    id_petugas int(11) REFERENCES tb_petugas(id_petugas),
    nisn varchar(10) REFERENCES tb_siswa(nisn),
    tgl_bayar date,
    bulan_dibayar varchar(8),
    tahun_dibayar varchar(4),
    id_spp int(11) REFERENCES tb_spp(id_spp),
    jumlah_bayar int(11)
);
```

## JSON API Collections

This repository contains a collection of API endpoints in JSON format, organized by functionality. Below are instructions on how to use each endpoint along with example requests and responses.

## Base Path

All endpoints in this collection share a common base path: `/api/v1`. Ensure that you prepend this base path to all endpoint URLs when making requests.

## Authentication

All endpoints in this API require authentication using JWT (JSON Web Tokens). To access any endpoint, you must include a valid JWT token in the Authorization header of your requests.

## Endpoints

### Data Siswa

#### Endpoint: `/data_siswa`

- **Description** :
- **Method** : `POST`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "nisn": 1234567890,
  "nis": 12345678,
  "nama": "Asep Saepudin",
  "id_kelas": 1,
  "alamat": "Desa Ciuyah",
  "no_telp": "085722456782",
  "id_spp": 1
}
```

- **Response** :

```json
{
  "responseCode": "2000101",
  "responseMessage": "success"
}
```

#### Endpoint: `/data_siswa`

- **Description** :
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000102",
  "responseMessage": "success",
  "data": [
    {
      "nisn": 1234567890,
      "nis": 12345678,
      "nama": "Asep Saepudin",
      "id_kelas": 1,
      "alamat": "Desa Ciuyah",
      "no_telp": "085722456782",
      "id_spp": 1
    },
    {
      "nisn": 1234567899,
      "nis": 12345677,
      "nama": "Wahyudi Hamisi",
      "id_kelas": 1,
      "alamat": "Desa Waledasem",
      "no_telp": "085722456797",
      "id_spp": 1
    }
  ],
  "paging": {
    "page": 1,
    "totalPages": 1,
    "totalData": 2
  }
}
```

#### Endpoint: `/data_siswa/:id`

- **Description** :
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000103",
  "responseMessage": "success",
  "data": {
    "nisn": 1234567890,
    "nis": 12345678,
    "nama": "Asep Saepudin",
    "kelas": {
      "id_kelas": 1,
      "nama_kelas": "XII",
      "kompetensi_keahlian": "Rekayasa Perangkat Lunak"
    },
    "alamat": "Desa Ciuyah",
    "no_telp": "085722456782",
    "id_spp": {
      "id_spp": 1,
      "tahun": 2024,
      "nominal": 150.0
    }
  }
}
```

#### Endpoint: `/data_siswa/:id`

- **Description** :
- **Method** : `PUT`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "nis": 12345678,
  "nama": "Asep Saepudin",
  "id_kelas": 1,
  "alamat": "Desa Ciuyah",
  "no_telp": "085722456782",
  "id_spp": 1
}
```

- **Response** :

```json
{
  "responseCode": "2000104",
  "responseMessage": "success"
}
```

#### Endpoint: `/data_siswa/:id`

- **Description** :
- **Method** : `DELETE`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000105",
  "responseMessage": "success"
}
```
