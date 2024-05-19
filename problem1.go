package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	var scores []int

	fmt.Println("Masukkan skor siswa :")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error membaca input:", err)
		return
	}

	if input == "" {
		fmt.Println("Input tidak boleh kosong.")
		return
	}

	parts := strings.Split(input, " ")
	for _, part := range parts {
		score, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid:", part)
			continue
		}
		scores = append(scores, score)
	}

	for _, score := range scores {
		grade := tentukanNilai(score)
		fmt.Printf("Skor: %d, Nilai: %s\n", score, grade)
	}
}

func tentukanNilai(score int) string {
	switch {
	case score >= 80 && score <= 100:
		return "A"
	case score >= 65 && score <= 79:
		return "B+"
	case score >= 50 && score <= 64:
		return "B"
	case score >= 35 && score <= 49:
		return "C"
	case score >= 0 && score <= 34:
		return "E"
	default:
		return "Skor Tidak Valid"
	}
}
