package main

import (
	"fmt"
)

func main() {
	const maksimal = 1000000
	var data [maksimal]int
	var n int
	var bilangan int

	for {
		fmt.Scan(&bilangan)
		if bilangan < 0 {
			break
		}
		data[n] = bilangan
		n++
	}

	// Insertion sort
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}

	// Cetak data yang sudah terurut
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	//periksa jarak
	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := data[1] - data[0]
	sama := true
	for i := 1; i < n-1; i++ {
		if data[i+1]-data[i] != jarak {
			sama = false
			break
		}
	}

	if sama {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
