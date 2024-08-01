package task

import (
	"fmt"
	"sync"
	"time"
)

func Task_3() {
	var wg sync.WaitGroup
	bufferSize := 5
	bufferedCh := make(chan int, bufferSize) // Buffer channel dengan kapasitas 5

	wg.Add(2) // Tambah jumlah goroutine yang akan ditunggu

	go generateNumbers(bufferedCh, &wg)
	go printNumbers(bufferedCh, &wg)
	go monitorChannel(bufferedCh, bufferSize, &wg)

	wg.Wait()
}

// Fungsi yang menghasilkan angka urut dan mengirimkannya ke channel
func generateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)       // Menampilkan angka yang dikirim
		time.Sleep(50 * time.Millisecond) // Delay untuk memperlihatkan buffering
	}
	close(ch) // Menutup channel setelah selesai mengirim semua angka
}

// Fungsi yang memonitor kapasitas channel dan menampilkan statusnya
func monitorChannel(ch <-chan int, bufferSize int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(50 * time.Millisecond) // Memeriksa kapasitas setiap 50ms
		bufferedChLen := len(ch)
		if bufferedChLen == bufferSize {
			fmt.Println("Channel is full")
		}
		if bufferedChLen == 0 && len(ch) == 0 {
			break // Hentikan monitoring jika channel telah ditutup dan kosong
		}
	}
}

// Fungsi yang menerima angka dari channel dan menampilkannya
func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Received: %d\n", num)   // Menampilkan angka yang diterima
		time.Sleep(1000 * time.Millisecond) // Delay lebih lama untuk membuat channel bisa penuh
	}
}
