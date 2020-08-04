package main

import "fmt"

func main() {
	for i := 0; i < 12; i++ {
		namaBulan, atributBulan := bulan(i + 1)
		awalBulan := atributBulan[0]
		maksHari := atributBulan[1]

		fmt.Printf("%-10s %-10s %10s\n", "===========", namaBulan, "===========")
		fmt.Printf("%-5s%-5s%-5s%-5s%-5s%-5s%-5s\n", "S", "S", "R", "K", "J", "S", "M")

		loopTanggal(maksHari, awalBulan)

	}
}

/*
 Fungsi ini berfungsi untuk mencocokan bulan yang sesuai

 Parameter : Urutan Bulan
 return	String : Nama Bulan
 return	Array : Awal Bulan & Total Hari
*/
func bulan(i int) (string, [2]int) {
	switch i {
	case 1:
		return "Januari", [2]int{3, 31}
	case 2:
		return "Februari", [2]int{6, 29}
	case 3:
		return "Maret", [2]int{7, 31}
	case 4:
		return "April", [2]int{3, 31}
	case 5:
		return "Mei", [2]int{5, 31}
	case 6:
		return "Juni", [2]int{1, 30}
	case 7:
		return "Juli", [2]int{3, 31}
	case 8:
		return "Agustus", [2]int{6, 31}
	case 9:
		return "September", [2]int{2, 30}
	case 10:
		return "Oktober", [2]int{4, 31}
	case 11:
		return "November", [2]int{7, 30}
	case 12:
		return "Desember", [2]int{2, 31}
	default:
		return "bulan error", [2]int{0, 0}
	}
}

/*
	Fungsi ini untuk melakukan looping tanggal sampai akhir bulan

	parameter maksHari : menentukan total hari dalam sebulan
	parameter awalBulan : menentukan hari yang menjadi awal bulan
 */

func loopTanggal(maksHari int, awalBulan int) {
	var hariKe7 int8 = 0 // membuat baris baru

	for hari := 1; hari <= maksHari; hari++ {
		//mencetak pada hari sesuai dengan awal bulan
		for awalBulan > 1 {
			fmt.Printf("%-5s", "")
			awalBulan--
			hariKe7++
		}

		//cetak tanggal
		fmt.Printf("%-5d", hari)

		hariKe7++

		//kembali ke baris baru
		if hariKe7 == 7 {
			fmt.Println("")
			hariKe7 = 0
		}

	}
	fmt.Println("\n")
}


/*
Output:

=========== Januari    ===========
S    S    R    K    J    S    M
          1    2    3    4    5
6    7    8    9    10   11   12
13   14   15   16   17   18   19
20   21   22   23   24   25   26
27   28   29   30   31

=========== Februari   ===========
S    S    R    K    J    S    M
                         1    2
3    4    5    6    7    8    9
10   11   12   13   14   15   16
17   18   19   20   21   22   23
24   25   26   27   28   29

=========== Maret      ===========
S    S    R    K    J    S    M
                              1
2    3    4    5    6    7    8
9    10   11   12   13   14   15
16   17   18   19   20   21   22
23   24   25   26   27   28   29
30   31

=========== April      ===========
S    S    R    K    J    S    M
          1    2    3    4    5
6    7    8    9    10   11   12
13   14   15   16   17   18   19
20   21   22   23   24   25   26
27   28   29   30   31

=========== Mei        ===========
S    S    R    K    J    S    M
                    1    2    3
4    5    6    7    8    9    10
11   12   13   14   15   16   17
18   19   20   21   22   23   24
25   26   27   28   29   30   31


=========== Juni       ===========
S    S    R    K    J    S    M
1    2    3    4    5    6    7
8    9    10   11   12   13   14
15   16   17   18   19   20   21
22   23   24   25   26   27   28
29   30

=========== Juli       ===========
S    S    R    K    J    S    M
          1    2    3    4    5
6    7    8    9    10   11   12
13   14   15   16   17   18   19
20   21   22   23   24   25   26
27   28   29   30   31

=========== Agustus    ===========
S    S    R    K    J    S    M
                         1    2
3    4    5    6    7    8    9
10   11   12   13   14   15   16
17   18   19   20   21   22   23
24   25   26   27   28   29   30
31

=========== September  ===========
S    S    R    K    J    S    M
     1    2    3    4    5    6
7    8    9    10   11   12   13
14   15   16   17   18   19   20
21   22   23   24   25   26   27
28   29   30

=========== Oktober    ===========
S    S    R    K    J    S    M
               1    2    3    4
5    6    7    8    9    10   11
12   13   14   15   16   17   18
19   20   21   22   23   24   25
26   27   28   29   30   31

=========== November   ===========
S    S    R    K    J    S    M
                              1
2    3    4    5    6    7    8
9    10   11   12   13   14   15
16   17   18   19   20   21   22
23   24   25   26   27   28   29
30

=========== Desember   ===========
S    S    R    K    J    S    M
     1    2    3    4    5    6
7    8    9    10   11   12   13
14   15   16   17   18   19   20
21   22   23   24   25   26   27
28   29   30   31


 */