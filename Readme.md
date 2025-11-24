# File Finder

File Finder adalah aplikasi CLI dalam Golang untuk mencari file berdasarkan direktori, ekstensi, dan kata kunci.

## 📦 Instalasi

Pastikan Go sudah terpasang di sistem Anda.

Untuk menginstal File Finder, jalankan perintah berikut:

```bash
go install github.com/MuhammadHakim33/file-finder@latest
```

## 🚀 Cara Menggunakan

```
-d string
    Path to the directory to scan
-e string
    Filter files by extension, e.g., txt, png, pdf
-s string
    Keyword to search for within files

```

contoh 
```
# Cari semua file dengan ekstensi .png di folder /home/user/Downloads
file-finder -d /home/user/Downloads -e png

# Cari file-file .txt yang mengandung kata "report"
file-finder -d /path/to/folder -e txt -s report

# Cari file apapun (semua ekstensi) yang mengandung kata “invoice”
file-finder -d /another/folder -s invoice

```
