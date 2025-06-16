-- Membuat database bernama `company_db` jika belum ada sebelumnya
CREATE DATABASE IF NOT EXISTS company_db;

-- Mengatur penggunaan database `company_db`
USE company_db;

-- Membuat tabel `employees` yang menyimpan data karyawan
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY, -- Kolom id sebagai primary key dan auto increment
    name VARCHAR(100) NOT NULL,        -- Kolom name untuk menyimpan nama karyawan (tidak boleh null)
    email VARCHAR(100) NOT NULL UNIQUE, -- Kolom email yang wajib diisi dan harus unik
    phone VARCHAR(100) NOT NULL,       -- Kolom phone untuk menyimpan nomor telepon karyawan (tidak boleh null)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Kolom timestamp untuk mencatat kapan data dibuat (default: waktu saat insert)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Kolom timestamp untuk mencatat waktu update terakhir
);
