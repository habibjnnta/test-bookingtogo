package main

import (
	"fmt"
	"bytes"
	"strings"
	"strconv"
	"sort"
)

func repeat() int {
	var batas int

	fmt.Printf("Masukkan Jumlah Array : ")
	fmt.Scan(&batas)

	if  batas < 2 || batas > 100 {
		batas = repeat()
	}

	return batas
}

func input(i int) int {
	var a int
	fmt.Printf("Masukkan Jumlah Angka ke - %d : ", i+1)
	fmt.Scan(&a)

	if a < 1 || a > 99 {
		a = input(i)
	}

	return a
}


func main() {
	var batas int

	fmt.Printf("Masukkan Jumlah Array : ")
	fmt.Scan(&batas)

	if  batas < 2 || batas > 100 {
		batas = repeat()
	}

	var a []int

	for i := 0; i < batas; i++ {
		a = append(a, input(i))
	}

	sort.Sort(sort.IntSlice(a))

	var str bytes.Buffer
	total := 1
	angka := a[0]

	for i := 0; i < len(a); i++ {
		str.WriteString(strconv.Itoa(a[i]) + " ")
	}

	for i := 0; i < len(a); i++ {
		cek := strings.Count(str.String(), strconv.Itoa(a[i]))
		if cek > total {
			total = cek
			angka = a[i]
		}
	}

	atas := angka+1
	bawah := angka-1
	cekAtas := strings.Count(str.String(), strconv.Itoa(atas))
	cekBawah := strings.Count(str.String(), strconv.Itoa(bawah))

	if cekAtas == cekBawah {
		total += cekAtas
	}else if cekAtas > cekBawah {
		total += cekAtas
	}else if cekAtas < cekBawah {
		total += cekBawah
	}

	// fmt.Println(str.String())
	// fmt.Println(angka)
	fmt.Printf("Jawabannya adalah : %d\n", total)
	// fmt.Println(a)

}
