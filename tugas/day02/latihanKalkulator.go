package main

import (
	"fmt"
	"math"
)

func main() {
	pilihan := 0
	for pilihan >= 0 && pilihan < 3 {
		fmt.Println("\nPilih Menu")
		fmt.Println("1. Kabataku")
		fmt.Println("2. Hitung Luas")
		fmt.Println("3. Hitung Volume")
		fmt.Println("4. Exit")
		pilihan = ScanMenu()
		fmt.Println("")
		Pilihan(pilihan)
	}


}

func Pilihan(menu int) {
	switch menu {
	case 1:
		fmt.Println("\n==== Kabataku ====")
		fmt.Println("Masukan angka pertama")
		angka1 := ScanMenu()
		fmt.Println("Masukan angka kedua")
		angka2 := ScanMenu()

		kabataku := Kabataku{float64(angka1), float64(angka2)}
		kabataku.Pertambahan()
		kabataku.Perkalian()
		kabataku.Pembagian()
		kabataku.Pengurangan()
		kabataku.Akar()
		kabataku.Kuadrat()
	case 2:
		fmt.Println("\n====  Hitung Luas ====")
		fmt.Println("1. Persegi")
		fmt.Println("2. Lingkaran")
		fmt.Println("3. Segitiga")

		pilihanLuas := ScanMenu()

		if pilihanLuas == 1 {
			fmt.Println("\nLuas Persegi")
			fmt.Println("Masukan sisi 1")
			sisi1 := ScanMenu()
			fmt.Println("Masukan sisi 2")
			sisi2 := ScanMenu()

			luas := Luas{float64(sisi1), float64(sisi2)}
			luas.Persegi()
		} else if pilihanLuas == 2 {
			fmt.Println("\nLuas Lingkaran")
			fmt.Println("Masukan Jari-Jari")
			jariJari := ScanMenu()
			phi := 3.14

			luas := Luas{float64(phi), float64(jariJari)}
			luas.Persegi()
		} else {
			fmt.Println("Anda Salah Memasukan Pilihan !")
		}
	case 3:
		fmt.Println("\n==== Hitung Volume ====")
		fmt.Println("1. tabung")
		fmt.Println("2. balok")
		fmt.Println("3. prisma")

		pilihanLuas := ScanMenu()

		if pilihanLuas == 1 {
			fmt.Println("\nvolume tabung")
			fmt.Println("Jari-Jari")
			sisi1 := ScanMenu()
			sisi2 := 3.14
			fmt.Println("tinggi")
			sisi3 := ScanMenu()

			volume := Volume{float64(sisi1), sisi2, float64(sisi3)}
			volume.Tabung()
		} else if pilihanLuas == 2 {
			fmt.Println("\nvolume balok")
			fmt.Println("Masukan Sisi 1")
			sisi1 := ScanMenu()
			fmt.Println("Masukan Sisi 2")
			sisi2 := ScanMenu()
			fmt.Println("Masukan Sisi 3")
			sisi3 := ScanMenu()

			volume := Volume{float64(sisi1), float64(sisi2), float64(sisi3)}
			volume.Balok()
		} else if pilihanLuas == 2 {
			fmt.Println("\nvolume balok")
			fmt.Println("Masukan Sisi 1")
			sisi1 := ScanMenu()
			fmt.Println("Masukan Sisi 2")
			sisi2 := ScanMenu()
			fmt.Println("Masukan Sisi 3")
			sisi3 := ScanMenu()

			volume := Volume{float64(sisi1), float64(sisi2), float64(sisi3)}
			volume.Balok()
		} else if pilihanLuas == 3 {
			fmt.Println("\nvolume prisma")
			fmt.Println("Masukan Sisi Alas 1")
			sisi1 := ScanMenu()
			fmt.Println("Masukan Sisi Alas 2")
			sisi2 := ScanMenu()
			fmt.Println("Masukan Sisi Tinggi")
			sisi3 := ScanMenu()

			volume := Volume{float64(sisi1), float64(sisi2), float64(sisi3)}
			volume.Prisma()
		} else {
			fmt.Println("Anda Salah Memasukan Pilihan !")

		}
	default:
		fmt.Print("Terimakasih")

	}
}

func ScanMenu() int {
	var pilihan int
	fmt.Print("Pilihan Anda : ")
	_, err := fmt.Scanf("%d", &pilihan)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return pilihan
}

type Kabataku struct {
	Angka1 float64
	Angka2 float64
}

func (h Kabataku) Pertambahan() {
	fmt.Println("Hasil Pertambahan:", h.Angka1+h.Angka2)
}

func (h Kabataku) Perkalian() {

	fmt.Println("Hasil Perkalian:", h.Angka1*h.Angka2)
}

func (h Kabataku) Pembagian() {
	fmt.Println("Hasil Pembagian:", h.Angka1/h.Angka2)
}

func (h Kabataku) Pengurangan() {
	fmt.Println("Hasil Pengurangan:", h.Angka1-h.Angka2)
}

func (h Kabataku) Akar() {
	fmt.Println("Akar angka 1:", math.Sqrt(h.Angka1))
	fmt.Println("Akar angka 2:", math.Sqrt(h.Angka2))
}

func (h Kabataku) Kuadrat() {
	fmt.Println("Kuadrat angka 1:", h.Angka1* h.Angka1)
	fmt.Println("Kuadrat angka 2:", h.Angka2* h.Angka2)
}

type Luas struct {
	Sisi1 float64
	Sisi2 float64
}

func (l Luas) Lingkaran() float64 {
	// phi * r * r
	hasil := l.Sisi1 * l.Sisi2 * l.Sisi2
	fmt.Println("luas lingkaran : ", hasil)
	return hasil
}

func (l Luas) Persegi() float64 {
	// sisi * sisi
	hasil := l.Sisi1 * l.Sisi2
	fmt.Println("luas persegi : ", hasil)
	return hasil
}

func (l Luas) Segitiga() float64 {
	// 1/2 sisi * sisi
	hasil := 0.5 * l.Sisi1 * l.Sisi2
	fmt.Println("luas persegi : ", hasil)
	return hasil
}

type Volume struct {
	Alas1  float64
	Alas2  float64
	Tinggi float64
}

func (v Volume) Prisma() {
	luasAlas := Luas{v.Alas1, v.Alas2}.Segitiga()
	hasil := luasAlas * v.Tinggi
	fmt.Println("volume prisma", hasil)
}

func (v Volume) Tabung() {
	luasAlas := Luas{v.Alas1, v.Alas2}.Lingkaran()
	hasil := luasAlas * v.Tinggi
	fmt.Println("volume tabung", hasil)
}

func (v Volume) Balok() {
	luasAlas := Luas{v.Alas1, v.Alas2}.Persegi()
	hasil := luasAlas * v.Tinggi
	fmt.Println("volume balok", hasil)

}

//kali bagi tambah kurang
// 2 params

// akar , pangkat
// 1 params

//luas persegi luas lingkaran

//volume tabung π × r² × t,
// 2param
//volume balok  V = p x l x t,  volume prisma 1/2 x panjang x lebar x tinggi.
// 3param

// cari luas alas -> volume tinggal * tinggi
// prisma -> luas segitiga(1/2*panjang*lebar) * t prisma
// balok -> luas persegi (p*l)    * t balok
// tabung -> luas lingkaran (π × r²) & t tabung

//kayaknya bisa pakai yang varadic ...int
