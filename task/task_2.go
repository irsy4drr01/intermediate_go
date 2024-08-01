package task

import (
	"fmt"
	"sync"
)

func Task_2() {
	var wg sync.WaitGroup
	fibCh := make(chan int)

	wg.Add(1)
	go fibonacci(fibCh, &wg)
	wg.Add(1)
	go ganjilGenap(fibCh, &wg)
	wg.Wait()
}

// Fungsi yang menghasilkan deret Fibonacci dan mengirimkan hasilnya ke channel
func fibonacci(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

// Fungsi yang memproses angka dari channel Fibonacci dan memisahkan ganjil/genap
func ganjilGenap(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ganjilCh := make(chan int)
	genapCh := make(chan int)

	// Fungsi untuk memproses angka ganjil
	go func() {
		for num := range ch {
			if num%2 == 0 {
				genapCh <- num
			} else {
				ganjilCh <- num
			}
		}
		close(ganjilCh)
		close(genapCh)
	}()

	var outputWg sync.WaitGroup
	outputWg.Add(2)

	// Menampilkan angka ganjil
	go func() {
		defer outputWg.Done()
		for num := range ganjilCh {
			fmt.Println("Ganjil:", num)
		}
	}()

	// Menampilkan angka genap
	go func() {
		defer outputWg.Done()
		for num := range genapCh {
			fmt.Println("Genap:", num)
		}
	}()

	// Tunggu hingga semua goroutine selesai
	outputWg.Wait()
}
