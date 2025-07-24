package mahasiswa

const maxNilai = 100
const Version = "v1.0.0"

func hitungRataRata(nilai ...int) float64 {
	var sum = 0.0
	for _, v := range nilai {
		sum += float64(v)
	}
	avg := sum / float64(len(nilai))
	return avg
}

func BuatMahasiswa(nama string, umur int, nilai ...int) *Mahasiswa {
	return &Mahasiswa{
		Nama:  nama,
		umur:  umur,
		Nilai: nilai,
	}
}

func GetMaxNilai() int {
	return maxNilai
}
