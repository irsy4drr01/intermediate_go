package task

import "fmt"

func Task_1() {
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)

	// Menjalankan goroutine untuk menjumlahkan bagian dari slice
	go sum(a[:len(a)/2], chn)
	go sum(a[len(a)/2:], chn)

	// Menerima hasil dari channel dan mencetaknya
	result1 := <-chn
	result2 := <-chn

	fmt.Printf("Result from first half: %d\n", result1)
	fmt.Printf("Result from second half: %d\n", result2)
	fmt.Printf("Total sum: %d\n", result1+result2)
}

// Fungsi untuk menjumlahkan angka-angka dalam slice dan mengirimkan hasilnya melalui channel
func sum(d []int, ch chan int) {
	total := 0
	for _, v := range d {
		total += v
	}
	ch <- total // kirim hasil ke channel
}
