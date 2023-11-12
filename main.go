package main

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	// wg.Done() akan mengurangi 1 pada wg counter
	// atau memberikan informasi kepada wg bahwa 1 goruntine yang dijalankan sudah selesai
	defer wg.Done()
	fmt.Println(message)
}

// dibuat untuk sinkronasi goruntine yang berjalan
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		// wg.Add(1) akan menambahkan 1 ke wg counter
		// atau wg.Add(1) akan memberikan informasi kepada wg bahwa jumplah goruntine yang sedang di proses bertambah satu
		wg.Add(1)
		go doPrint(&wg, data)
	}
	// wg.Wait() akan menunggu  sampai wg counter bernilai 0
	// proses setelahnya tidak akan di eksekusi ke baris selanjutnya sebelum semua goruntine selesai dijalankan.
	wg.Wait()
}
