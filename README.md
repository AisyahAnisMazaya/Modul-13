*NOMOR 1*
Deskripsi Program
Program ini menerima input berupa bilangan bulat non-negatif yang disimpan dalam array. Program kemudian:
Mengurutkan bilangan menggunakan algoritma Insertion Sort.
Mencetak bilangan yang sudah terurut.
Memeriksa apakah selisih antar bilangan (jarak) tetap atau tidak, kemudian memberikan output sesuai hasil pemeriksaan.
Penjelasan Variabel
Konstanta
maksimal: Batas maksimum elemen array, yaitu 1.000.000.
Variabel di main
data: Array dengan ukuran tetap untuk menyimpan bilangan yang diinputkan.
Tipe: [maksimal]int.
Contoh: Jika pengguna memasukkan 5, 10, 15, maka data = [5, 10, 15, ...].
n: Jumlah bilangan yang telah dimasukkan ke dalam array.
Tipe: int.
bilangan: Bilangan yang sedang dibaca dari input pengguna.
Tipe: int.
Variabel lokal lainnya:
key: Elemen yang sedang diatur posisinya dalam Insertion Sort.
j: Indeks untuk iterasi dalam Insertion Sort.
jarak: Selisih tetap antar elemen array (jika ada).
sama: Penanda apakah semua selisih antar elemen sama.
Alur Program
1. Input Bilangan
Program membaca bilangan secara berulang hingga bilangan negatif dimasukkan.
Bilangan negatif akan menghentikan proses input.
2. Urutkan Bilangan
Insertion Sort digunakan untuk mengurutkan elemen dalam array data secara menaik.
3. Cetak Bilangan Terurut
Setelah pengurutan, elemen array dicetak dalam satu baris, dipisahkan oleh spasi.
4. Periksa Selisih Tetap
Program menghitung selisih antara elemen berturut-turut.
Jika semua selisih sama, program mencetak jarak tersebut.
Jika ada selisih yang berbeda, program mencetak bahwa data berjarak tidak tetap.

*NOMOR 2*
Deskripsi Program
Program ini adalah sistem manajemen perpustakaan sederhana yang memungkinkan pengguna untuk:
Mendaftarkan sejumlah buku ke dalam daftar pustaka.
Menampilkan buku dengan rating tertinggi.
Mengurutkan buku berdasarkan rating secara menurun.
Menampilkan lima buku dengan rating tertinggi.
Mencari buku berdasarkan rating menggunakan algoritma Binary Search.
Penjelasan Variabel dan Struktur
Konstanta
nMax: Jumlah maksimum buku yang dapat didaftarkan (7919 buku).
Tipe Data
Buku: Struktur untuk menyimpan informasi tentang sebuah buku, dengan atribut:
id: String yang menunjukkan ID buku.
judul: Judul buku.
penulis: Nama penulis buku.
penerbit: Nama penerbit buku.
eksemplar: Jumlah eksemplar buku yang tersedia (integer).
tahun: Tahun penerbitan buku (integer).
rating: Rating buku (integer).
DaftarBuku: Array berukuran tetap ([nMax]Buku) untuk menyimpan semua buku yang didaftarkan.
Variabel Global
pustaka: Array dari tipe DaftarBuku untuk menyimpan data buku.
n: Jumlah buku yang telah didaftarkan.
rating: Rating yang akan dicari di daftar buku.
Penjelasan Fungsi
1. DaftarkanBuku
Fungsi untuk mengisi daftar buku berdasarkan input pengguna.
Parameter:
pustaka *DaftarBuku: Referensi ke daftar buku.
n *int: Referensi ke jumlah buku yang akan didaftarkan.
Input:
ID, judul, penulis, penerbit, eksemplar, tahun, dan rating untuk setiap buku.
2. CetakTerfavorit
Menampilkan buku dengan rating tertinggi dari daftar.
Parameter:
pustaka DaftarBuku: Daftar buku.
n int: Jumlah buku dalam daftar.
3. UrutBuku
Mengurutkan daftar buku berdasarkan rating secara menurun menggunakan Insertion Sort.
Parameter:
pustaka *DaftarBuku: Referensi ke daftar buku.
n int: Jumlah buku dalam daftar.
4. Cetak5Terbaru
Menampilkan lima buku dengan rating tertinggi setelah daftar diurutkan.
Parameter:
pustaka DaftarBuku: Daftar buku.
n int: Jumlah buku dalam daftar.
5. CariBuku
Mencari buku berdasarkan rating menggunakan Binary Search.
Menampilkan semua buku yang memiliki rating sama.
Parameter:
pustaka DaftarBuku: Daftar buku.
n int: Jumlah buku dalam daftar.
r int: Rating yang dicari.
6. main
Fungsi utama yang mengatur alur program:
Mendaftarkan buku ke dalam daftar pustaka.
Menampilkan buku dengan rating tertinggi.
Mengurutkan buku berdasarkan rating.
Menampilkan lima buku dengan rating tertinggi.
Mencari buku berdasarkan rating tertentu.
Alur Program
Daftarkan Buku:
Masukkan jumlah buku yang akan didaftarkan (n).
Masukkan data untuk setiap buku.
Cetak Buku dengan Rating Tertinggi:
Temukan buku dengan rating tertinggi menggunakan perulangan.
Urutkan Buku Berdasarkan Rating:
Urutkan array pustaka menggunakan Insertion Sort.
Cetak 5 Buku dengan Rating Tertinggi:
Tampilkan hingga lima buku pertama dari daftar yang telah diurutkan.
Cari Buku Berdasarkan Rating:
Cari buku dengan rating tertentu menggunakan Binary Search.
Tampilkan semua buku dengan rating yang sama.
