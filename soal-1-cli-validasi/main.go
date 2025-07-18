package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	ErrValidation = errors.New("wajib diisi")
	ErrMinAge     = errors.New("umur tidak valid (minimal 18 tahun)")
	ErrorInt      = errors.New("masukkan hanya angka")
)

func isValid(name, age string) (string, error) {
	a, err := strconv.Atoi(age)

	if name == "" || age == "" {
		return "", fmt.Errorf("nama dan umur %w", ErrValidation)
	}
	if err != nil {
		return "", fmt.Errorf("%w", ErrorInt)
	}
	if a <= 18 {
		return "", fmt.Errorf("%w", ErrMinAge)
	}
	return name, nil
}

func reader(r *bufio.Reader) string {
	str, _ := r.ReadString('\n')
	return strings.TrimSpace(str)
}

func main() {
	nReader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama:")
	n := reader(nReader)
	fmt.Print("Umur:")
	a := reader(nReader)

	result, err := isValid(n, a)

	if err != nil {
		log.WithError(err).Error(err)
	} else {
		log.Info("Selamat datang ", result)
	}

}
