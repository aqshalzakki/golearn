package main

func main() {
	// Slice
	// days := [...]string{"senin, ", "selasa, ", "rabu, ", "kamis, ", "jumat, ", "sabtu, ", "minggu, "}

	// daysSliced := days[5:]

	// // Akan buat array baru, tidak akan append array lama
	// newDaySliced := append(daysSliced, "Minggu Kedua")

	// fmt.Println(newDaySliced, cap(newDaySliced))

	// ------------------

	// var slice []string = make([]string, 2, 5)

	// slice[0] = "Aqshal"
	// slice[1] = "Jajang"
	// // slice[2] = "Nurjaman" error harus menggunakan append

	// newSlice := append(slice, "Nurjaman")
	// newSlice[0] = "Budi"

	// fmt.Println(newSlice, len(newSlice), cap(newSlice))
	// fmt.Println(slice, len(slice), cap(slice))

	// ------------------
	// days := []string{"senin, ", "selasa, ", "rabu, ", "kamis, ", "jumat, ", "sabtu, ", "minggu, "}

	// fromSlice := days[:]
	// toSlice := make([]string, len(days), cap(days))

	// copy(toSlice, fromSlice)

	// println()

	// ---------------
	// * Di golang rata2 menggunakan slice dibandingkan array
	// iniArray := [...]int{1, 2, 3, 4, 5}
	// iniSlice := []int{1, 2, 3, 4, 5}

	// fmt.Println(iniArray, len(iniArray), cap(iniArray))
	// fmt.Println(iniSlice, len(iniSlice), cap(iniSlice))
}
