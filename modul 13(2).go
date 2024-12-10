package main

import (
	"fmt"
)

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 1; i <= *n; i++ {
		var id, judul, penulis, penerbit string
		var eksemplar, tahun, rating int
		fmt.Scan(&id, &judul, &penulis, &penerbit, &eksemplar, &tahun, &rating)
		pustaka[i] = Buku{id, judul, penulis, penerbit, eksemplar, tahun, rating}
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	terfavorit := pustaka[1]
	for i := 2; i <= n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

// (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		kunci := pustaka[i]
		j := i - 1
		for j > 0 && pustaka[j].rating < kunci.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = kunci
	}
}

// Procedure Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("Lima Buku dengan Rating Tertinggi:")
	for i := 1; i <= n && i <= 5; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

// (Binary Search)
func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri, kanan := 1, n
	ketemu := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			buku := pustaka[tengah]
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
			ketemu = true
			// rating yang sama
			for i := tengah - 1; i >= 1 && pustaka[i].rating == r; i-- {
				fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
					pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
			}
			for i := tengah + 1; i <= n && pustaka[i].rating == r; i++ {
				fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
					pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
			}
			break
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !ketemu {
		fmt.Println("Buku Tidak Ditemukan")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var rating int
	DaftarkanBuku(&pustaka, &n)
	fmt.Scan(&rating)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, rating)
}
