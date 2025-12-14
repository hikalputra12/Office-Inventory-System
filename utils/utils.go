package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Inisialisasi bufio.Reader SEKALI di scope package
var reader = bufio.NewReader(os.Stdin)

// ReadLine adalah fungsi utilitas yang aman untuk membaca input baris penuh
func ReadLine(prompt string) string {
	fmt.Print(prompt)

	// Membaca seluruh baris hingga karakter newline (\n)
	input, _ := reader.ReadString('\n')

	// Membersihkan spasi di awal/akhir dan karakter newline
	return strings.TrimSpace(input)
}
