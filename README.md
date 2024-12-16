# Fp_Go_Web

Link Demo FP Kelompok 32:
[Youtube](https://www.youtube.com/live/IvIPQFeuOug?si=3cPXopVnB9peyNhB&t=75)


## Cara Run
### **Clone Repository Github**
Mulai dengan cloning repository github dengan perintah
```R
git clone https://github.com/SamuelBerkatHulu/FP-PBKK
cd FP-PBKK
 ```
### **Instalasi Dependensi**
Kemudian install dependensi GO, dengan perintah 
```R
go mod init FP-BPKK
```
### **Pastikan Mysql aktif**
Anda dapat menggunakan tools lain untuk mengaktifkan `Mysql`. Pada progam ini saya menggunakan tools `XAMPP` untuk mengaktifkan `Mysql`.
Pada tampilan `XAMPP` untuk memulai klik `Start Apache` dan klik `Start Mysql` pastikan tombol start dalam keadaan `Stop` yang menandahkan bahwa `Apache` dan `Mysql` telah aktif. 

### **Mulai Menjalankan Program**
Setelah instalasi dependensi dan mysql aktif, selanjutkan jalankan program dengan perintah
```R
go run .
```
### **Akses Program dengan Browser**
setelah `go run .` maka pada terminal VsCode akan muncul tampilan `http://localhost:8080`. Anda tidak hanya mengakses pada port `8080` saja tapi menyesuaikan berdasarkan port yang anda tulis pada progam `main.go`
karena pada `main.go` adalah port `8080`, maka anda dapat menggunakan browser dan mengetikan `http://localhost:8080` untuk mengakses aplikasi.

### **Memberhentikan Program Aplikasi**
tekan `Ctrl+C` pada terminal VsCode untuk memberhentikan program aplikasi.
