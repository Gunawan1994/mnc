package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var soal int
	fmt.Print("Pilih soal: ")
	fmt.Scanln(&soal)

	switch soal {
	case 1:
		Soal1()
	case 2:
		Soal2()
	case 3:
		Soal3()
	case 4:
		Soal4()
	default:
		fmt.Println("Soal tidak ada!")
	}
}

func Soal1() {
	var s int
	fmt.Print("Input: ")
	fmt.Scanln(&s)

	strArray := make([]string, s)
	r := bufio.NewReader(os.Stdin)

	for i := 0; i < s; i++ {
		fmt.Printf("Masukkan string %d:", i+1)
		str, _ := r.ReadString('\n')
		strArray[i] = strings.ToLower(str)
	}

	for i := 0; i < s-1; i++ {
		for j := i + 1; j < s; j++ {
			if strArray[i] == strArray[j] {
				fmt.Printf("output %d %d", i+1, j+1)
				return
			}
		}
	}

	fmt.Println(false)
}

func Soal2() {
	var totalBlanja int
	fmt.Print("total belanja: ")
	fmt.Scanln(&totalBlanja)

	var jumlahBayar int
	fmt.Print("jumlah bayar: ")
	fmt.Scanln(&jumlahBayar)

	if jumlahBayar < totalBlanja {
		fmt.Println(false, ", kurang bayar")
		return
	}

	kembalian := jumlahBayar - totalBlanja
	kembalianFix := int(math.Floor(float64(kembalian)/100) * 100)

	fmt.Printf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d\n", kembalian, kembalianFix)
	pecahanUang := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	result := make(map[int]int)

	sisa := kembalianFix
	for _, pecahan := range pecahanUang {
		if sisa >= pecahan {
			jumlahPecahan := sisa / pecahan
			sisa = sisa % pecahan
			result[pecahan] = jumlahPecahan
		}
	}

	fmt.Println("Pecahan uang:")
	for _, val := range pecahanUang {
		if i, v := result[val]; v {
			if val >= 1000 {
				fmt.Printf("%d lembar %d\n", i, val)
			} else {
				fmt.Printf("%d koin %d\n", i, val)
			}
		}
	}
}

func Soal3() {
	fmt.Print("Input simbol: ")
	var s string
	fmt.Scanln(&s)

	pairs := map[rune]rune{
		'>': '<',
		'}': '{',
		']': '[',
	}

	var stack []rune
	for _, char := range s {
		switch char {
		case '<', '{', '[':
			stack = append(stack, char)
		case '>', '}', ']':
			if len(stack) == 0 {
				fmt.Println(false)
				return
			}

			if stack[len(stack)-1] != pairs[char] {
				fmt.Println(false)
				return
			}

			stack = stack[:len(stack)-1]
		default:
			fmt.Println(false)
			return
		}
	}
	fmt.Println(len(stack) == 0)
	return
}

func Soal4() {
	fmt.Println("Contoh tanggal 2024-09-28 foremat (yyyy-mm-dd)")
	cutiKantor := 14
	maxCuti := 3
	cutiKaryawan := 180

	var cutiBersama int
	fmt.Print("Jumlah Cuti Bersama: ")
	fmt.Scanln(&cutiBersama)

	var tanggalJoin string
	fmt.Print("Tanggal join karyawan: ")
	fmt.Scanln(&tanggalJoin)

	var tanggalCuti string
	fmt.Print("Tanggal rencana cuti: ")
	fmt.Scanln(&tanggalCuti)

	var durasiCuti int
	fmt.Print("Durasi cuti: ")
	fmt.Scanln(&durasiCuti)

	layout := "2006-01-02"

	tanggalJoinDate, err := time.Parse(layout, tanggalJoin)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	tanggalCutiDate, err := time.Parse(layout, tanggalCuti)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	validateCuti := tanggalCutiDate.Sub(tanggalJoinDate)
	checkCuti := validateCuti.Hours() / 24

	if int(checkCuti) < cutiKaryawan {
		fmt.Println("Tidak berhak mengambil cuti")
		return
	}

	cutiPribadi := cutiKantor - cutiBersama

	endOf180 := tanggalJoinDate.AddDate(0, 0, 180)
	yearEnd := time.Date(tanggalCutiDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	remainingDays := int(yearEnd.Sub(endOf180).Hours() / 24)

	hakCuti := int(math.Floor(float64(remainingDays) / 365.0 * float64(cutiPribadi)))
	if durasiCuti > hakCuti {
		valStr := strconv.Itoa(hakCuti)
		x := strings.Replace(valStr, "-", "kurang ", 1)
		fmt.Printf("Karyawan memiliki %s hari cuti pribadi di tahun ini. %v", x, false)
		return
	}

	if durasiCuti > maxCuti {
		fmt.Println("Cuti pribadi max 3 hari berturutan.", false)
		return
	}
	fmt.Println("Karyawan boleh mengambil cuti pribadi.", true)
}
