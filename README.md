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
    id_spp INT PRIMARY KEY,
    tahun INT,
    nominal int
);

CREATE TABLE IF NOT EXISTS tb_kelas(
    id_kelas INT PRIMARY KEY,
    nama_kelas VARCHAR(10),
    kompetensi_keahlian VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS tb_petugas(
    id_petugas INT PRIMARY KEY,
    username VARCHAR(25),
    password VARCHAR(255),
    nama_petugas VARCHAR(35),
    level level
);

CREATE TABLE IF NOT EXISTS tb_siswa(
    nisn VARCHAR(10) PRIMARY KEY,
    nis VARCHAR(8) UNIQUE NOT NULL,
    nama VARCHAR(35),
    id_kelas INT REFERENCES tb_kelas(id_kelas),
    alamat TEXT,
    no_telp VARCHAR(13),
    id_spp INT REFERENCES tb_spp(id_spp)
);

CREATE TABLE IF NOT EXISTS tb_pembayaran(
    id_pembayaran INT PRIMARY KEY,
    id_petugas INT REFERENCES tb_petugas(id_petugas),
    nisn VARCHAR(10) REFERENCES tb_siswa(nisn),
    tgl_bayar TIMESTAMP,
    bulan_dibayar VARCHAR(8),
    tahun_dibayar VARCHAR(4),
    id_spp INT REFERENCES tb_spp(id_spp),
    jumlah_bayar INT
);
```

## Berikut DML untuk Contoh Data Pembayaran Spp:

```sql
-- Insert sample data into tb_spp
INSERT INTO tb_spp (id_spp, tahun, nominal) VALUES
(1, 2022, 2000000),
(2, 2023, 2200000),
(3, 2024, 2500000);

-- Insert sample data into tb_kelas
INSERT INTO tb_kelas (id_kelas, nama_kelas, kompetensi_keahlian) VALUES
(1, 'X AK 1', 'Akuntansi'),
(2, 'XI AP 1', 'Administrasi Perkantoran'),
(3, 'XII RPL 1', 'Rekayasa Perangkat Lunak');

-- Insert sample data into tb_petugas
INSERT INTO tb_petugas (id_petugas, username, password, nama_petugas, level) VALUES
(1, 'admin01', '$2a$12$OopSsL7cjRtdXo8l3V1ZwePFHmjF/ZImC0Jgu6DDC6uyU2P2sX1jW', 'Admin One', 'ADMIN'),
(2, 'petugas01', '$2a$12$OopSsL7cjRtdXo8l3V1ZwePFHmjF/ZImC0Jgu6DDC6uyU2P2sX1jW', 'Petugas One', 'PETUGAS'),
(3, 'petugas02', '$2a$12$OopSsL7cjRtdXo8l3V1ZwePFHmjF/ZImC0Jgu6DDC6uyU2P2sX1jW', 'Petugas Two', 'PETUGAS');

-- Insert sample data into tb_siswa
INSERT INTO tb_siswa (nisn, nis, nama, id_kelas, alamat, no_telp, id_spp) VALUES
('1234567890', '11111111', 'Student One', 1, '123 Street A', '081234567890', 1),
('1234567891', '11111112', 'Student Two', 2, '456 Street B', '081234567891', 2),
('1234567892', '11111113', 'Student Three', 3, '789 Street C', '081234567892', 3);

-- Insert sample data into tb_pembayaran
INSERT INTO tb_pembayaran (id_pembayaran, id_petugas, nisn, tgl_bayar, bulan_dibayar, tahun_dibayar, id_spp, jumlah_bayar) VALUES
(1, 1, '1234567890', '2022-01-10 10:00:00', 'January', '2022', 1, 2000000),
(2, 2, '1234567891', '2023-02-15 11:30:00', 'February', '2023', 2, 2200000),
(3, 3, '1234567892', '2024-03-20 09:45:00', 'March', '2024', 3, 2500000);
```

## JSON API Collections

This repository contains a collection of API endpoints in JSON format, organized by functionality. Below are instructions on how to use each endpoint along with example requests and responses.

## Base Path

All endpoints in this collection share a common base path: `/api/v1`. Ensure that you prepend this base path to all endpoint URLs when making requests.

## Authentication

All endpoints in this API require authentication using JWT (JSON Web Tokens). To access any endpoint, you must include a valid JWT token in the Authorization header of your requests.

## Endpoints

### Authentication

#### Endpoint: "/login/siswa"

- **Description** : Authenticates a student using their NISN and NIS and returns a JWT token upon successful login.
- **Method** : `POST`
- **Header** : `Authorization` : `Basic Auth admin:admin`

- **Request** :

```json
{
  "nisn": 1234567890,
  "nis": 12345678
}
```

- **Response** :

```json
{
  "responseCode": "2000000",
  "responseMessage": "success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkwMTQ3OTcsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOiJBRE1JTiJ9.-qQTObos7rnAjJFoG2FJhYRckwCLjOR6C2PDR2wPu2g"
}
```

#### Endpoint: "/login/petugas"

- **Description** : Authenticates a staff member using their username and password, and returns a JWT token upon successful login.
- **Method** : `POST`
- **Header** : `Authorization` : `Basic Auth admin:admin`

- **Request** :

```json
{
  "username": "admin",
  "password": "admin"
}
```

- **Response** :

```json
{
  "responseCode": "2000000",
  "responseMessage": "success",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkwMTQ3OTcsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOiJBRE1JTiJ9.-qQTObos7rnAjJFoG2FJhYRckwCLjOR6C2PDR2wPu2g"
}
```

### Data Siswa

#### Endpoint: `/data_siswa`

- **Description** : Adds a new student record to the system using the provided student information.
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

- **Description** : Retrieves a list of student records from the system.
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

- **Description** : Retrieves the details of a specific student using their ID.
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

- **Description** : Updates the details of a specific student using their ID.
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

- **Description** : Deletes a specific student record using their ID.
- **Method** : `DELETE`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000105",
  "responseMessage": "success"
}
```

### Data Petugas

#### Endpoint: `/data_petugas`

- **Description** : Adds a new staff member to the system using the provided staff information.
- **Method** : `POST`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "id_petugas": 1,
  "username": "admin01",
  "nama_petugas": "Admin One",
  "level": "ADMIN"
}
```

- **Response** :

```json
{
  "responseCode": "2000201",
  "responseMessage": "success"
}
```

#### Endpoint: `/data_petugas`

- **Description** : Retrieves a list of staff member records from the system.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000202",
  "responseMessage": "success",
  "data": [
    {
      "id_petugas": 1,
      "username": "admin01",
      "nama_petugas": "Admin One",
      "level": "ADMIN"
    },
    {
      "id_petugas": 2,
      "username": "petugas01",
      "nama_petugas": "Petugas One",
      "level": "PETUGAS"
    }
  ],
  "paging": {
    "page": 1,
    "totalPages": 1,
    "totalData": 2
  }
}
```

#### Endpoint: `/data_petugas/:id`

- **Description** : Retrieves the details of a specific staff member using their ID.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000203",
  "responseMessage": "success",
  "data": {
    "id_petugas": 1,
    "username": "admin01",
    "nama_petugas": "Admin One",
    "level": "ADMIN"
  }
}
```

#### Endpoint: `/data_petugas/:id`

- **Description** : Updates the details of a specific staff member using their ID.
- **Method** : `PUT`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "nama_petugas": "Admin One-One"
}
```

- **Response** :

```json
{
  "responseCode": "2000204",
  "responseMessage": "success",
  "data": {
    "id_petugas": 1,
    "username": "admin01",
    "nama_petugas": "Admin One-One",
    "level": "ADMIN"
  }
}
```

#### Endpoint: `/data_petugas/:id`

- **Description** : Deletes a specific staff member record using their ID.
- **Method** : `DELETE`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000205",
  "responseMessage": "success"
}
```

### Data Kelas

#### Endpoint: `/data_kelas`

- **Description** : Adds a new class record to the system using the provided class information.
- **Method** : `POST`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "id_kelas": 1,
  "nama_kelas": "X AK 1",
  "kompetensi_keahlian": "Akuntansi"
}
```

- **Response** :

```json
{
  "responseCode": "2000301",
  "responseMessage": "success"
}
```

#### Endpoint: `/data_kelas`

- **Description** : Retrieves a list of class records from the system.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000302",
  "responseMessage": "success",
  "data": [
    {
      "id_kelas": 1,
      "nama_kelas": "X AK 1",
      "kompetensi_keahlian": "Akuntansi"
    },
    {
      "id_kelas": 2,
      "nama_kelas": "XI AP 1",
      "kompetensi_keahlian": "Administrasi Perkantoran"
    }
  ],
  "paging": {
    "page": 1,
    "totalPages": 1,
    "totalData": 2
  }
}
```

#### Endpoint: `/data_kelas/:id`

- **Description** : Retrieves the details of a specific class using its ID.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000303",
  "responseMessage": "success",
  "data": {
    "id_kelas": 1,
    "nama_kelas": "X AK 1",
    "kompetensi_keahlian": "Akuntansi"
  }
}
```

#### Endpoint: `/data_kelas/:id`

- **Description** : Updates the details of a specific class using its ID.
- **Method** : `PUT`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "nama_kelas": "X AK 1",
  "kompetensi_keahlian": "Akuntansi"
}
```

- **Response** :

```json
{
  "responseCode": "2000304",
  "responseMessage": "success",
  "data": {
    "id_kelas": 1,
    "nama_kelas": "X AK 1",
    "kompetensi_keahlian": "Akuntansi"
  }
}
```

#### Endpoint: `/data_kelas/:id`

- **Description** : Deletes a specific class record using its ID.
- **Method** : `DELETE`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000305",
  "responseMessage": "success"
}
```

### Data Spp

#### Endpoint: `/data_spp`

- **Description** : Adds a new school fee (SPP) record to the system using the provided information.
- **Method** : `POST`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "id_spp": 1,
  "tahun": 2022,
  "nominal": 100000
}
```

- **Response** :

```json
{
  "responseCode": "2000401",
  "responseMessage": "success"
}
```

#### Endpoint: `/data_spp`

- **Description** : Retrieves a list of school fee (SPP) records from the system.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000402",
  "responseMessage": "success",
  "data": [
    {
      "id_spp": 1,
      "tahun": 2022,
      "nominal": 100000
    },
    {
      "id_spp": 2,
      "tahun": 2023,
      "nominal": 150000
    }
  ],
  "paging": {
    "page": 1,
    "totalPages": 1,
    "totalData": 2
  }
}
```

#### Endpoint: `/data_spp/:id`

- **Description** : Retrieves the details of a specific school fee (SPP) using its ID.
- **Method** : `GET`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000403",
  "responseMessage": "success",
  "data": {
    "id_spp": 1,
    "tahun": 2022,
    "nominal": 100000
  }
}
```

#### Endpoint: `/data_spp/:id`

- **Description** : Updates the details of a specific school fee (SPP) using its ID.
- **Method** : `PUT`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Request** :

```json
{
  "tahun": 2022,
  "nominal": 100000
}
```

- **Response** :

```json
{
  "responseCode": "2000404",
  "responseMessage": "success",
  "data": {
    "id_spp": 1,
    "tahun": 2022,
    "nominal": 100000
  }
}
```

#### Endpoint: `/data_spp/:id`

- **Description** : Deletes a specific school fee (SPP) record using its ID.
- **Method** : `DELETE`
- **Header** : `Authorization` : `Bearer <your JWT Token>`

- **Response** :

```json
{
  "responseCode": "2000405",
  "responseMessage": "success"
}
```

### Entri Transaksi Pembayaran

#### Endpoint: `/transaksi_pembayaran`
