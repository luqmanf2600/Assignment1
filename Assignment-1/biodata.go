package main

import "fmt"

type student struct {
	name       string
	address    string
	profession string
	reason     string
}

func data() {
	var s1 = student{}
	s1.name = "Luqman Fauzi"
	s1.address = "Tangerang"
	s1.profession = "Mahasiswa"
	s1.reason = "Ingin mempunyai skill khusus pada bagian back-end"

	var s2 = student{}
	s2.name = "Ahmad Faris"
	s2.address = "Kalimantan"
	s2.profession = "Mahasiswa"
	s2.reason = "Mendalami pemahaman"

	var s3 = student{}
	s3.name = "Rakha Rizqillah Pratama Saputra"
	s3.address = "Jakarta"
	s3.profession = "Mahasiswa"
	s3.reason = "Ingin lebih mempelajari dan memahami lebih banyak bahasa golang"

	var s4 = student{}
	s4.name = "I Gede Diva Dwijayana"
	s4.address = "Bali"
	s4.profession = "Mahasiswa"
	s4.reason = "Tertarik mengikuti pemrograman golang"

	var s5 = student{}
	s5.name = "Muhammad Daffa Haryadi putra"
	s5.address = "Aceh"
	s5.profession = "Mahasiswa"
	s5.reason = "Ingin menjadi back-end programmer"

	var s6 = student{}
	s6.name = "Muhammad Yoga Irvandi"
	s6.address = "Sulawesi Tengah"
	s6.profession = "Mahasiswa"
	s6.reason = "Mencoba hal baru yang belum dipelajari"

	var s7 = student{}
	s7.name = "Wafianda Azhar"
	s7.address = "Solo"
	s7.profession = "Mahasiswa"
	s7.reason = "Mengikuti golang ini membuat saya semangat dalam mempelajari back-end"

	var s8 = student{}
	s8.name = "Muhammad Haiqal Malik"
	s8.address = "Bogor"
	s8.profession = "Mahasiswa"
	s8.reason = "Ingin menambah pengetahuan bahasa golang"

	var s9 = student{}
	s9.name = "Rubenhard Manat Hasigolan Lumbantobing"
	s9.address = "Medan"
	s9.profession = "Mahasiswa"
	s9.reason = "Suka menambah ilmu baru"

	var s10 = student{}
	s10.name = "Denada Ramschie"
	s10.address = "Medan"
	s10.profession = "Mahasiswa"
	s10.reason = "Tertarik mengikuti bahasa pemrograman selain c"

	var i int
	fmt.Print("Pilih data 1-10 yang akan dicetak: ")
	fmt.Scanln(&i)
	switch i {
	case 1:
		fmt.Println("Name : ", s1.name)
		fmt.Println("address : ", s1.address)
		fmt.Println("profession :", s1.profession)
		fmt.Println("Reason : ", s1.reason)
	case 2:
		fmt.Println("Name : ", s2.name)
		fmt.Println("address : ", s2.address)
		fmt.Println("profession :", s2.profession)
		fmt.Println("Reason : ", s2.reason)
	case 3:
		fmt.Println("Name : ", s3.name)
		fmt.Println("address : ", s3.address)
		fmt.Println("profession :", s3.profession)
		fmt.Println("Reason : ", s3.reason)
	case 4:
		fmt.Println("Name : ", s4.name)
		fmt.Println("address : ", s4.address)
		fmt.Println("profession :", s4.profession)
		fmt.Println("Reason : ", s4.reason)
	case 5:
		fmt.Println("Name : ", s5.name)
		fmt.Println("address : ", s5.address)
		fmt.Println("profession :", s5.profession)
		fmt.Println("Reason : ", s5.reason)
	case 6:
		fmt.Println("Name : ", s6.name)
		fmt.Println("address : ", s6.address)
		fmt.Println("profession :", s6.profession)
		fmt.Println("Reason : ", s6.reason)
	case 7:
		fmt.Println("Name : ", s7.name)
		fmt.Println("address : ", s7.address)
		fmt.Println("profession :", s7.profession)
		fmt.Println("Reason : ", s7.reason)
	case 8:
		fmt.Println("Name : ", s8.name)
		fmt.Println("address : ", s8.address)
		fmt.Println("profession :", s8.profession)
		fmt.Println("Reason : ", s8.reason)
	case 9:
		fmt.Println("Name : ", s9.name)
		fmt.Println("address : ", s9.address)
		fmt.Println("profession :", s9.profession)
		fmt.Println("Reason : ", s9.reason)
	case 10:
		fmt.Println("Name : ", s10.name)
		fmt.Println("address : ", s10.address)
		fmt.Println("profession :", s10.profession)
		fmt.Println("Reason : ", s10.reason)
	default:
		fmt.Println("Input salah! Data tidak ada")
	}
}

func main() {
	data()
	data()
	data()
}
