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

Berikut DDL untuk Pembayaran SPP :

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

Berikut endpoint, request dan response:

1. `CREATE`
