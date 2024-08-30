package main

import (
	"fmt"
	"strings"
)

const NMAX int = 10

type car struct {
	pabrikan, nama, id string
	tahun, terjual     int
}
type factory struct {
	pabrikan string
	terjual  int
}
type tabCar [NMAX]car
type tabFactory [NMAX]factory

func main() {
	var data tabCar
	var nData int
	var opt string

	for opt != "5" {
		chooseOption(&opt)
		if nData > 0 {
			if opt == "1" {
				carList(&data, &nData)
			} else if opt == "2" {
				addCar(&data, &nData)
			} else if opt == "3" {
				optEdit(&data, nData)
			} else if opt == "4" {
				optDel(&data, &nData)
			} else if opt == "5" {

			} else {
				fmt.Println("Opsi tidak valid.")
			}
		} else if opt == "1" || opt == "2" {
			if opt == "1" {
				carList(&data, &nData)
			} else {
				addCar(&data, &nData)
			}
		} else if opt == "3" || opt == "4" {
			fmt.Println("!!! Data kosong. Silahkan mengisi data terlebih dahulu. !!!")
			fmt.Println()
		} else if opt != "5" {
			fmt.Println("Opsi tidak valid.")
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("Closing app...")
	fmt.Println("See you later.")
	fmt.Println()
}

// coba ntar main pagenya dijadiin ky bentuk prosedur filter
func chooseOption(o *string) {
	fmt.Println()
	fmt.Println("      === APLIKASI DEALER MOBIL ===     ")
	fmt.Println()
	fmt.Println("+--------------------------------------+")
	fmt.Println("|                 MENU                 |")
	fmt.Println("+--------------------------------------+")
	fmt.Println()
	fmt.Print("1. Daftar data mobil\n2. Tambah data mobil\n3. Edit data mobil\n4. Hapus data mobil\n5. Keluar\nPilih opsi: ")
	fmt.Scan(o)
	fmt.Println()
}

func addCar(A *tabCar, n *int) {
	var t car
	if *n != NMAX {
		fmt.Println("=--------------------------------------=")
		fmt.Println()
		fmt.Println("Ketik 'stop' atau 'x' untuk berhenti menambahkan data.")
		fmt.Println("(Ganti spasi dengan tanda underscore('_') apabila nama pabrikan/mobil terdapat lebih dari satu kata)")
		fmt.Println()
		count := *n + 1
		fmt.Printf("DATA KE-%d\n", count)
		fmt.Print("Isi data baru [<id_mobil> <nama_pabrikan> <nama_mobil> <tahun_rilis> <unit_yang_terjual>]: ")
		fmt.Scan(&t.id)
		for t.id != "stop" && t.id != "x" && *n < NMAX {
			fmt.Scan(&t.pabrikan, &t.nama, &t.tahun, &t.terjual)
			t.pabrikan = strings.ToUpper(t.pabrikan)
			t.nama = strings.ToUpper(t.nama)
			if redSearchAdd(*A, t, *n) {
				A[*n].id = t.id
				A[*n].pabrikan = t.pabrikan
				A[*n].nama = t.nama
				A[*n].tahun = t.tahun
				A[*n].terjual = t.terjual
				fmt.Println("Data berhasil ditambah.")
				fmt.Println()
				*n++
				if *n != NMAX {
					count++
					fmt.Printf("DATA KE-%d\n", count)
					fmt.Print("Isi data baru [<id_mobil> <nama_pabrikan> <nama_mobil> <tahun_rilis> <unit_yang_terjual>]: ")
					fmt.Scan(&t.id)
				} else {
					fmt.Println()
					fmt.Println("Tempat penyimpanan sudah penuh.")
				}

			} else {
				fmt.Println()
				fmt.Println("Data mobil atau id mobil tersebut sudah ada.")
				fmt.Println()
				fmt.Print("Isi data baru [<id_mobil> <nama_pabrikan> <nama_mobil> <tahun_rilis> <unit_yang_terjual>]: ")
				fmt.Scan(&t.id)
			}

		}
		fmt.Println("Total data mobil:", *n)
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("!!! Data penuh. Silahkan hapus beberapa data terlebih dahulu. !!!")
	}

}

// SEQUENTIAL SEARCH (redSearch() -> to prevent redundancy)
func redSearchAdd(A tabCar, t car, n int) bool {
	for i := 0; i < n; i++ {
		if A[i].id == t.id {
			return false
		} else if A[i].pabrikan == t.pabrikan && A[i].nama == t.nama && A[i].tahun == t.tahun {
			return false
		}
	}
	return true
}

func tampilan_CarList(A tabCar, n int) {
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("Daftar data mobil:")
	fmt.Println()
	fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
	for i := 0; i < n; i++ {
		fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
	}
	fmt.Println()
	fmt.Println("Total data mobil:", n)
	fmt.Println()
	fmt.Printf("Filter[f] -- Add Car[a] -- Edit[e] -- Delete[d] -- Back[x]")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
}

func carList(A *tabCar, n *int) {
	var opt string
	tampilan_CarList(*A, *n)
	fmt.Scan(&opt)
	if *n > 0 {
		for opt != "x" {
			if opt == "f" {
				filter(A, *n)
			} else if opt == "a" {
				addCar(A, n)
			} else if opt == "e" {
				optEdit(A, *n)
			} else if opt == "d" {
				optDel(A, n)
			} else {
				fmt.Println("Opsi tidak valid.")
			}
			tampilan_CarList(*A, *n)
			fmt.Scan(&opt)
		}
	} else if opt == "a" {
		addCar(A, n)
	} else {
		fmt.Println("!!! Data kosong. Silahkan mengisi data terlebih dahulu. !!!")
		fmt.Println()
	}
}

func filter(A *tabCar, n int) {
	var opt int
	fmt.Println()
	fmt.Println("1. Tampilkan NAMA PABRIKAN tertentu")
	fmt.Println("2. Tampilkan NAMA MOBIL tertentu")
	fmt.Println("3. Tampilkan TAHUN MOBIL tertentu")
	fmt.Println("4. Tampilkan MOBIL dengan JUMLAH PENJUALAN tertentu")
	fmt.Println("5. Tampilkan PABRIKAN dengan JUMLAH PENJUALAN tertentu")
	fmt.Println("6. Urutkan Data")
	fmt.Println("7. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	if opt == 1 {
		showFacName(*A, n)
	} else if opt == 2 {
		showCarName(*A, n)
	} else if opt == 3 {
		showCarYear(*A, n)
	} else if opt == 4 {
		filterCarSales(*A, n)
	} else if opt == 5 {
		filterFacSales(*A, n)
	} else if opt == 6 {
		sortType(A, n)
	} else if opt == 7 {

	} else {
		fmt.Println()
		fmt.Println("Opsi tidak valid.")
		filter(A, n)
	}
}

func showFacName(A tabCar, n int) {
	var x string
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Cari pabrikan: ")
	fmt.Scan(&x)
	x = strings.ToUpper(x)
	fmt.Println()
	count := 0
	fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
	for i := 0; i < n; i++ {
		if A[i].pabrikan == x {
			fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
			count++
		}
	}
	fmt.Println()
	fmt.Println(count, "data ditampilkan.")
	fmt.Println()
}
func showCarName(A tabCar, n int) {
	var x string
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Cari mobil: ")
	fmt.Scan(&x)
	x = strings.ToUpper(x)
	fmt.Println()
	count := 0
	fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
	for i := 0; i < n; i++ {
		if A[i].nama == x {
			fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
			count++
		}
	}
	fmt.Println()
	fmt.Println(count, "data ditampilkan.")
	fmt.Println()
}
func showCarYear(A tabCar, n int) {
	var y int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Cari tahun: ")
	fmt.Scan(&y)
	fmt.Println()
	count := 0
	fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
	for i := 0; i < n; i++ {
		if A[i].tahun == y {
			fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
			count++
		}
	}
	fmt.Println()
	fmt.Println(count, "data ditampilkan.")
	fmt.Println()
}

func filterCarSales(A tabCar, n int) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("1. Jumlah Penjualan TERTINGGI")
	fmt.Println("2. Jumlah Penjualan TERENDAH")
	fmt.Println("3. Mobil dengan jumlah penjualan RENTANG TERTENTU")
	fmt.Println("4. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	fmt.Println()
	switch opt {
	case 1:
		carTopXSales(A, n)
	case 2:
		carBottomXSales(A, n)
	case 3:
		rangeXSales(A, n)
	case 4:

	default:
		fmt.Println("Opsi tidak valid.")
		filterCarSales(A, n)
	}
}

func rangeXSales(A tabCar, n int) {
	var start, end, count int
	fmt.Println()
	fmt.Print("Masukkan range [<nilai_x> <nilai_y>]: ")
	fmt.Scan(&start, &end)
	if start <= end { //asc
		fmt.Println()
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		minSales(&A, n)
		for i := 0; i < n; i++ {
			if A[i].terjual >= start && A[i].terjual <= end {
				fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
				count++
			}
		}
	} else { //desc
		fmt.Println()
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		maxSales(&A, n)
		for i := 0; i < n; i++ {
			if A[i].terjual >= end && A[i].terjual <= start {
				fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
				count++
			}
		}
	}
}

// penjualan tertinggi
func carTopXSales(A tabCar, n int) {
	var x, max int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan berapa data mobil dengan jumlah penjualan tertinggi teratas yang ingin ditampilkan: ")
	fmt.Scan(&x)
	if x > n {
		x = n
	}
	fmt.Println()
	if x == 1 {
		max = findMaxSales(A, 0, n)
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[max].id, A[max].pabrikan, A[max].nama, A[max].tahun, A[max].terjual)
	} else if x >= 2 {
		maxSales(&A, n)
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		for i := 0; i < x; i++ {
			fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
		}
	} else {
		fmt.Println("Terjadi kesalahan input. Pastikan Anda memasukkan nilai bilangan bulat positif.")
		carTopXSales(A, n)
	}
}

// penjualan terendah
func carBottomXSales(A tabCar, n int) {
	var x, min int
	fmt.Println()
	fmt.Print("Masukkan berapa data mobil dengan jumlah penjualan terendah teratas yang ingin ditampilkan: ")
	fmt.Scan(&x)
	if x > n {
		x = n
	}
	fmt.Println()
	if x == 1 {
		min = findMinSales(A, 0, n)
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[min].id, A[min].pabrikan, A[min].nama, A[min].tahun, A[min].terjual)
	} else if x >= 2 {
		minSales(&A, n)
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		for i := 0; i < x; i++ {
			fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[i].id, A[i].pabrikan, A[i].nama, A[i].tahun, A[i].terjual)
		}
	} else {
		fmt.Println("Terjadi kesalahan input. Pastikan Anda memasukkan nilai bilangan bulat positif.")
		carTopXSales(A, n)
	}
}

func findMaxSales(A tabCar, start, end int) int {
	var i, max int
	if end > NMAX {
		end = NMAX
	}
	max = start
	i = start + 1
	for i < end {
		if A[i].terjual > A[max].terjual {
			max = i
		}
		i++
	}
	return max
}

func findMinSales(A tabCar, start, end int) int {
	var i, min int
	if end > NMAX {
		end = NMAX
	}
	min = start
	i = start + 1
	for i < end {
		if A[i].terjual < A[min].terjual {
			min = i
		}
		i++
	}
	return min
}

func maxSales(A *tabCar, n int) {
	var max int
	var t car
	for i := 0; i < n; i++ {
		max = findMaxSales(*A, i, n)
		t = A[max]
		A[max] = A[i]
		A[i] = t
	}
}

func minSales(A *tabCar, n int) {
	var min int
	var t car
	for i := 0; i < n; i++ {
		min = findMinSales(*A, i, n)
		t = A[min]
		A[min] = A[i]
		A[i] = t
	}
}

func filterFacSales(A tabCar, n int) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("1. Jumlah Penjualan TERTINGGI")
	fmt.Println("2. Jumlah Penjualan TERENDAH")
	fmt.Println("3. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	fmt.Println()
	switch opt {
	case 1:
		facTopXSales(A, n)
	case 2:
		facBottomXSales(A, n)
	case 3:

	default:
		fmt.Println("Opsi tidak valid.")
		filterCarSales(A, n)
	}
}

func facTopXSales(A tabCar, n int) {
	var x, nF int
	var F tabFactory
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan berapa data pabrikan dengan jumlah penjualan tertinggi teratas yang ingin ditampilkan: ")
	fmt.Scan(&x)
	fmt.Println()
	descFactory(&A, n)
	scanFac(A, n, &F, &nF)
	descFacSales(&F, nF)
	if x > nF {
		x = nF
	}
	if x >= 1 {
		fmt.Printf("%-10s | %-10s\n", "Pabrikan", "Terjual")
		for i := 0; i < nF; i++ {
			fmt.Printf("%-10s   %-10d\n", F[i].pabrikan, F[i].terjual)
		}
	} else {
		fmt.Println("Terjadi kesalahan input. Pastikan Anda memasukkan nilai bilangan bulat positif.")
		fmt.Println()
		facBottomXSales(A, n)
	}
}

func facBottomXSales(A tabCar, n int) {
	var x, nF int
	var F tabFactory
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan berapa data pabrikan dengan jumlah penjualan terendah teratas yang ingin ditampilkan: ")
	fmt.Scan(&x)
	fmt.Println()
	ascFactory(&A, n)
	scanFac(A, n, &F, &nF)
	ascFacSales(&F, nF)
	if x > nF {
		x = nF
	}
	if x >= 1 {
		fmt.Printf("%-10s | %-10s\n", "Pabrikan", "Terjual")
		for i := 0; i < nF; i++ {
			fmt.Printf("%-10s   %-10d\n", F[i].pabrikan, F[i].terjual)
		}
	} else {
		fmt.Println("Terjadi kesalahan input. Pastikan Anda memasukkan nilai bilangan bulat positif.")
		fmt.Println()
		facBottomXSales(A, n)
	}

}

func scanFac(A tabCar, nC int, F *tabFactory, nF *int) {
	var t string
	t = A[0].pabrikan
	F[0].pabrikan = t
	F[0].terjual = A[0].terjual
	*nF = 1
	for i := 1; i < nC; i++ {
		if A[i].pabrikan == t {
			F[*nF-1].terjual += A[i].terjual
		} else {
			t = A[i].pabrikan
			F[*nF].pabrikan = t
			F[*nF].terjual = A[i].terjual
			*nF++
		}
	}
}

func descFacSales(F *tabFactory, n int) { //max->min
	var i, j, max int
	var t factory
	i = 1
	for i < n {
		j = i
		max = i - 1
		for j < n {
			if F[j].terjual > F[max].terjual {
				max = j
			}
			j++
		}
		t = F[max]
		F[max] = F[i-1]
		F[i-1] = t
		i++
	}
}

func ascFacSales(F *tabFactory, n int) { //min->max
	var i, j, min int
	var t factory
	i = 1
	for i < n {
		j = i
		min = i - 1
		for j < n {
			if F[j].terjual < F[min].terjual {
				min = j
			}
			j++
		}
		t = F[min]
		F[min] = F[i-1]
		F[i-1] = t
		i++
	}
}

func optDel(A *tabCar, n *int) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("Hapus data")
	fmt.Println("1. Berdasarkan ID MOBIL")
	fmt.Println("2. Berdasarkan NAMA PABRIKAN")
	fmt.Println("3. Berdasarkan NAMA MOBIL")
	fmt.Println("4. Berdasarkan TAHUN")
	fmt.Println("5. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	fmt.Println()
	for opt != 5 {
		if opt == 1 {
			delCar_Id(A, n)
		} else if opt == 2 {
			delCar_Factory(A, n)
		} else if opt == 3 {
			delCar_Name(A, n)
		} else if opt == 4 {
			delCar_Year(A, n)
		}
		fmt.Println("Hapus data")
		fmt.Println("1. Berdasarkan ID MOBIL")
		fmt.Println("2. Berdasarkan NAMA PABRIKAN")
		fmt.Println("3. Berdasarkan NAMA MOBIL")
		fmt.Println("4. Berdasarkan TAHUN")
		fmt.Println("5. Kembali")
		fmt.Println()
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&opt)
	}
}

func delCar_Id(A *tabCar, n *int) {
	var id string
	var found bool
	var left, right, mid int
	ascId(A, *n)
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan nama mobil yang akan dihapus: ")
	fmt.Scan(&id)
	left = 0
	right = *n - 1
	found = false
	for left <= right && !found {
		mid = (left + right) / 2
		if id < A[mid].id {
			right = mid - 1
		} else if A[mid].id < id {
			left = mid + 1
		} else {
			found = true
			for mid < *n {
				A[mid] = A[mid+1]
				mid++
			}
			*n--
		}
	}
	if found {
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func delCar_Name(A *tabCar, n *int) {
	var name string
	var found bool
	var count int
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan nama mobil yang akan dihapus: ")
	fmt.Scan(&name)
	name = strings.ToUpper(name)
	for i := 0; i < *n; i++ {
		if A[i].nama == name {
			found = true
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			count++
		}
	}
	if found {
		fmt.Printf("Mobil dengan nama %s berhasil dihapus (sebanyak %d)\n", name, count)
	} else {
		fmt.Printf("Mobil dengan nama %s tidak ditemukan\n", name)
	}
	fmt.Println()
}

func delCar_Factory(A *tabCar, n *int) {
	var fac string
	var found bool
	var count int
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan nama pabrikan yang akan dihapus: ")
	fmt.Scan(&fac)
	fac = strings.ToUpper(fac)
	for i := 0; i < *n; i++ {
		if A[i].pabrikan == fac {
			found = true
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			count++
		}
	}
	if found {
		fmt.Printf("Pabrikan dengan nama %s berhasil dihapus (sebanyak %d)\n", fac, count)
	} else {
		fmt.Printf("Pabrikan dengan nama %s tidak ditemukan\n", fac)
	}
	fmt.Println()
}

func delCar_Year(A *tabCar, n *int) {
	var year int
	var found bool
	var count int
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan nama pabrikan yang akan dihapus: ")
	fmt.Scan(&year)
	for i := 0; i < *n; i++ {
		if A[i].tahun == year {
			found = true
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			count++
		}
	}
	if found {
		fmt.Printf("Pabrikan dengan nama %d berhasil dihapus (sebanyak %d)\n", year, count)
	} else {
		fmt.Printf("Pabrikan dengan nama %d tidak ditemukan\n", year)
	}
	fmt.Println()
}

func optEdit(A *tabCar, n int) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Printf("\nOpsi pencarian data untuk menemukan data yang akan diedit\n")
	fmt.Println("1. Cari berdasarkan ID mobil")
	fmt.Println("2. Cari berdasarkan Nama Pabrikan, Nama Mobil, dan tahun rilis")
	fmt.Println("3. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	switch opt {
	case 1:
		editCarById(A, n)
	case 2:
		editExactCar(A, n)
	}

}

func redSearchEdit(A tabCar, t car, n int, idEdit int) bool {
	for i := 0; i < n; i++ {
		if i != idEdit {
			if A[i].id == t.id {
				return false
			} else if A[i].pabrikan == t.pabrikan && A[i].nama == t.nama && A[i].tahun == t.tahun {
				return false
			}
		}
	}
	return true
}

func editCarById(A *tabCar, n int) {
	var i, found, opt int
	var id string
	var tnew car
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Print("Masukkan ID mobil: ")
	fmt.Scan(&id)
	found = -1
	for i < n && found == -1 {
		if A[i].id == id {
			found = i
		}
		i++
	}
	if found != -1 {
		tnew = A[found]
		fmt.Println()
		fmt.Println("Data berhasil ditemukan.")
		fmt.Println()
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[found].id, A[found].pabrikan, A[found].nama, A[found].tahun, A[found].terjual)
		fmt.Println()
		fmt.Println("Opsi edit")
		fmt.Println("1. Edit ID")
		fmt.Println("2. Edit Pabrikan")
		fmt.Println("3. Edit Nama")
		fmt.Println("4. Edit Tahun")
		fmt.Println("5. Edit Jumlah Terjual")
		fmt.Println("6. Kembali")
		fmt.Println()
		fmt.Print("Pilih bagian yang akan diubah: ")
		fmt.Scan(&opt)
		fmt.Println()
		for opt != 6 {
			if opt == 1 {
				fmt.Print("Masukkan ID yang baru: ")
				fmt.Scan(&tnew.id)
			} else if opt == 2 {
				fmt.Print("Masukkan nama pabrikan yang baru: ")
				fmt.Scan(&tnew.pabrikan)
				tnew.pabrikan = strings.ToUpper(tnew.pabrikan)
			} else if opt == 3 {
				fmt.Print("Masukkan nama mobil yang baru: ")
				fmt.Scan(&tnew.nama)
				tnew.nama = strings.ToUpper(tnew.nama)
			} else if opt == 4 {
				fmt.Print("Masukkan tahun yang baru: ")
				fmt.Scan(&tnew.tahun)
			} else if opt == 5 {
				fmt.Printf("[positif = bertambah (ex: 5) || negatif = berkurang (ex: -5)]\nTambah nilai unit terjual: ")
				//pengurangan unit terjual mungkin dibutuhkan apabila salah mengetik angka saat ingin menambahkan data unit terjual
				fmt.Scan(&tnew.terjual)
				tnew.terjual += A[found].terjual
			}
			if redSearchEdit(*A, tnew, n, found) && tnew != A[found] {
				A[found] = tnew
				fmt.Println("Data berhasil diubah!")
				fmt.Println()
				fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
				fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[found].id, A[found].pabrikan, A[found].nama, A[found].tahun, A[found].terjual)
			} else if tnew != A[found] {
				fmt.Println("Data gagal diubah. (data identik/duplikat terdeteksi)")
			} else {
				fmt.Println("Opsi tidak valid.")
			}
			fmt.Println()
			fmt.Println("Opsi edit")
			fmt.Println("1. Edit ID")
			fmt.Println("2. Edit Pabrikan")
			fmt.Println("3. Edit Nama")
			fmt.Println("4. Edit Tahun")
			fmt.Println("5. Edit Jumlah Terjual")
			fmt.Println("6. Kembali")
			fmt.Println()
			fmt.Print("Pilih bagian yang akan diubah: ")
			fmt.Scan(&opt)
		}
	} else {
		fmt.Println()
		fmt.Println("Data tidak ditemukan.")
	}
}

func editExactCar(A *tabCar, n int) {
	var i, found, opt int
	var t, tnew car
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("Masukkan nilai <pabrikan> <mobil> <tahun> yang akan diedit: ")
	fmt.Scan(&t.pabrikan, &t.nama, &t.tahun)
	t.pabrikan = strings.ToUpper(t.pabrikan)
	t.nama = strings.ToUpper(t.nama)
	found = -1
	for i < n && found == -1 {
		if (A[i].pabrikan == t.pabrikan) && (A[i].nama == t.nama) && (A[i].tahun == t.tahun) {
			found = i
		}
		i++
	}
	if found != -1 {
		tnew = A[found]
		fmt.Println()
		fmt.Println("Data berhasil ditemukan.")
		fmt.Println()
		fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
		fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[found].id, A[found].pabrikan, A[found].nama, A[found].tahun, A[found].terjual)
		fmt.Println()
		fmt.Println("Opsi edit")
		fmt.Println("1. Edit ID")
		fmt.Println("2. Edit Pabrikan")
		fmt.Println("3. Edit Nama")
		fmt.Println("4. Edit Tahun")
		fmt.Println("5. Edit Jumlah Terjual")
		fmt.Println("6. Kembali")
		fmt.Println()
		fmt.Print("Pilih bagian yang akan diubah: ")
		fmt.Scan(&opt)
		fmt.Println()
		for opt != 6 {
			if opt == 1 {
				fmt.Print("Masukkan ID yang baru: ")
				fmt.Scan(&tnew.id)
			} else if opt == 2 {
				fmt.Print("Masukkan nama pabrikan yang baru: ")
				fmt.Scan(&tnew.pabrikan)
				tnew.pabrikan = strings.ToUpper(tnew.pabrikan)
			} else if opt == 3 {
				fmt.Print("Masukkan nama mobil yang baru: ")
				fmt.Scan(&tnew.nama)
				tnew.nama = strings.ToUpper(tnew.nama)
			} else if opt == 4 {
				fmt.Print("Masukkan tahun yang baru: ")
				fmt.Scan(&tnew.tahun)
			} else if opt == 5 {
				fmt.Printf("[positif = bertambah (ex: 5) || negatif = berkurang (ex: -5)]\nTambah nilai unit terjual: ")
				//pengurangan unit terjual mungkin dibutuhkan apabila salah mengetik angka saat ingin menambahkan data unit terjual
				fmt.Scan(&tnew.terjual)
				tnew.terjual += A[found].terjual
			}
			if redSearchEdit(*A, tnew, n, found) && tnew != A[found] {
				A[found] = tnew
				fmt.Println("Data berhasil diubah!")
				fmt.Println()
				fmt.Printf("%-10s | %-10s | %-15s | %-5s | %-6s\n", "ID", "Pabrikan", "Nama Mobil", "Tahun", "Terjual")
				fmt.Printf("%-10s   %-10s   %-15s   %-5d   %-6d\n", A[found].id, A[found].pabrikan, A[found].nama, A[found].tahun, A[found].terjual)
			} else if tnew != A[found] {
				fmt.Println("Data gagal diubah. (data identik/duplikat terdeteksi)")
			} else {
				fmt.Println("Opsi tidak valid.")
			}
			fmt.Println()
			fmt.Println("Opsi edit")
			fmt.Println("1. Edit ID")
			fmt.Println("2. Edit Pabrikan")
			fmt.Println("3. Edit Nama")
			fmt.Println("4. Edit Tahun")
			fmt.Println("5. Edit Jumlah Terjual")
			fmt.Println("6. Kembali")
			fmt.Println()
			fmt.Print("Pilih bagian yang akan diubah: ")
			fmt.Scan(&opt)
		}
	} else {
		fmt.Println()
		fmt.Println("Data tidak ditemukan.")
	}
}

func sortType(A *tabCar, n int) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("Opsi pengurutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("3. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	fmt.Println()
	for opt != 3 {
		if opt == 1 {
			sortBy(A, n, true)
		} else if opt == 2 {
			sortBy(A, n, false)
		}
		fmt.Println()
		fmt.Println("Opsi pengurutan:")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Println("3. Kembali")
		fmt.Println()
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&opt)
	}
}

func sortBy(A *tabCar, n int, asc bool) {
	var opt int
	fmt.Println()
	fmt.Println("=--------------------------------------=")
	fmt.Println()
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Berdasarkan Nama Pabrikan")
	fmt.Println("3. Berdasarkan Nama Mobil")
	fmt.Println("4. Berdasarkan Tahun")
	fmt.Println("5. Berdasarkan Jumlah Terjual")
	fmt.Println("6. Kembali")
	fmt.Println()
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opt)
	fmt.Println()
	for opt != 6 {
		if opt == 1 {
			if asc {
				ascId(A, n)
			} else {
				descId(A, n)
			}
		} else if opt == 2 { //pabrikan
			if asc {
				ascFactory(A, n)
			} else {
				descFactory(A, n)
			}
		} else if opt == 3 { //nama
			if asc {
				ascName(A, n)
			} else {
				descName(A, n)
			}
		} else if opt == 4 { //tahun
			if asc {
				ascYear(A, n)
			} else {
				descYear(A, n)
			}
		} else if opt == 5 { //penjualan
			if asc {
				minSales(A, n)
			} else {
				maxSales(A, n)
			}
		}
		fmt.Println("Data berhasil diurutkan.")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&opt)
	}
}

func ascId(A *tabCar, n int) {
	var i, j int
	var t car
	i = 1
	for i < n {
		j = i
		t = A[j]
		for j > 0 && t.id < A[j-1].id {
			A[j] = A[j-1]
			j--
		}
		A[j] = t
		i++
	}
}

func descId(A *tabCar, n int) {
	var i, j, max int
	var t car
	i = 1
	for i < n {
		j = i
		max = i - 1
		for j < n {
			if A[j].id > A[max].id {
				max = j
			}
			j++
		}
		t = A[max]
		A[max] = A[i-1]
		A[i-1] = t
		i++
	}
}

func ascFactory(A *tabCar, n int) {
	var i, j int
	var t car
	i = 1
	for i < n {
		j = i
		t = A[j]
		for j > 0 && t.pabrikan < A[j-1].pabrikan {
			A[j] = A[j-1]
			j--
		}
		A[j] = t
		i++
	}
}

func descFactory(A *tabCar, n int) {
	var i, j, max int
	var t car
	i = 1
	for i < n {
		j = i
		max = i - 1
		for j < n {
			if A[j].pabrikan > A[max].pabrikan {
				max = j
			}
			j++
		}
		t = A[max]
		A[max] = A[i-1]
		A[i-1] = t
		i++
	}
}

func ascName(A *tabCar, n int) {
	var i, j int
	var t car
	i = 1
	for i < n {
		j = i
		t = A[j]
		for j > 0 && t.nama < A[j-1].nama {
			A[j] = A[j-1]
			j--
		}
		A[j] = t
		i++
	}
}

func descName(A *tabCar, n int) {
	var i, j, max int
	var t car
	i = 1
	for i < n {
		j = i
		max = i - 1
		for j < n {
			if A[j].nama > A[max].nama {
				max = j
			}
			j++
		}
		t = A[max]
		A[max] = A[i-1]
		A[i-1] = t
		i++
	}
}

func ascYear(A *tabCar, n int) {
	var i, j int
	var t car
	i = 1
	for i < n {
		j = i
		t = A[j]
		for j > 0 && t.tahun < A[j-1].tahun {
			A[j] = A[j-1]
			j--
		}
		A[j] = t
		i++
	}
}

func descYear(A *tabCar, n int) {
	var i, j, max int
	var t car
	i = 1
	for i < n {
		j = i
		max = i - 1
		for j < n {
			if A[j].tahun > A[max].tahun {
				max = j
			}
			j++
		}
		t = A[max]
		A[max] = A[i-1]
		A[i-1] = t
		i++
	}
}
