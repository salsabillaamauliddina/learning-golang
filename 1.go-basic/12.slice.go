package main

import "fmt"

func main() {
	// ! potongan dari data array
	// ? slice mirip seperti array yg beda adalah ukuran slica bisa berubah
	// * slice dan array selalu terkoneksi, slice data yg menakses sebagian atau seluruh data di array
	// ? slice selalu terkoneksi , slice adalah data yg  mengakses sebagian / seluruh array

	// tipe dataslice 3 -> pointer,length dan capacity
	// pointer-> penunjuk data pertama di array pada slice
	// length -> panjang dari slice
	// capacity -> kapasitas dari slice
	// dimana legth !> capacity

	// ! NOTES
	//? array[low:high] => membuat slice dr array dimuai index low s/d index sblm high
	//? array[low:] => slice dr array dimulai index low s/d index akhir di array
	//? array[:] => slice dari aaray dimulai index 0 samoai index akhir di array

	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// ! index 0 di slice 1 yaitu mei
	// slice1[0] = "Mei"
	// fmt.Println(months)

	//? append(slice,data) -> buat slice baru dg menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice1 := days[5:]
	daySlice1[0] = "new Saturday"
	daySlice1[1] = "new Sunday"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "New Holiday")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// ? make([]TypeData,length,capacity) => membuat slice baru

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Rama"
	newSlice[1] = "Aca"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// ? copy(destination,source) => menyalin slice dari source ke destination

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// [5]
	iniArray := [...]int{
		1,
		2,
		3,
		4,
		5,
	}

	iniSlices := []int{
		1,
		2,
		3,
	}

	fmt.Println(iniArray)
	fmt.Println(iniSlices)
}
