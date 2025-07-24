package mahasiswa

import "strconv"

type Mahasiswa struct {
	Nama      string
	Nilai     []int
	umur      int
	nilainAvg float64
}
type Deskripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

func (m *Mahasiswa) GetUmur() int {
	return m.umur
}

func (m *Mahasiswa) RataRata() float64 {
	m.nilainAvg = hitungRataRata(m.Nilai...)
	return m.nilainAvg
}

func (m *Mahasiswa) Info() string {
	strN := "Nama: " + m.Nama + ", Umur: " + strconv.Itoa(m.umur)
	return strN
}
