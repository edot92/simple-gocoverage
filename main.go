package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/edot92/simple-gocoverage/lib"
)

// Buat sebuah fungsi untuk mendapatkan nilai terbesar dari value array yang berderet dan
// muncul lagi setelahnya dengan urutan deret kebalikannya.
// contoh :
// - [1, 2, 3, 8, 9, 3, 2, 1] nilai terbesarnya 3 (dari deret [1, 2, 3])
// - [1, 2, 1, 2, 2, 1] → 2
// - [7, 1, 2, 9, 7, 2, 1] → 2
// Buat dengan bahasa Go beserta unit test dengan coverage minimal 90%

func main() {
	type contoh struct {
		strContoh  string
		paramDeret []int
	}
	derets := []contoh{
		{
			"contoh 1", []int{1, 2, 3},
		},
		{
			"contoh 2", []int{1, 2},
		},
		{
			"contoh 3", []int{7, 1, 2},
		},
	}
	for index := 0; index < len(derets); index++ {
		results, err := lib.CariDeretMax(derets[index].paramDeret)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("-------" + derets[index].strContoh + "---------")
		fmt.Print("Val Input Array  : ")
		fmt.Println(derets[index].paramDeret)
		fmt.Print("Val Result Array : ")
		fmt.Println(results.ValComputeArray)
		fmt.Println("Nilai Terbesar   : " + strconv.Itoa(results.ValBigger))
		fmt.Println()
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("wait ctrl + c")
	<-done
	fmt.Println("exit program")
}
