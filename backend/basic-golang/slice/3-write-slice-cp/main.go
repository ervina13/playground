package main

import "fmt"

// Lanjutan nomor 2
// Sehabis menambahkan "Olleh" pada slice tersebut coba ubah nilai "World" menjadi "Marcus"
// dan "Olleh" menjadi "Aurelius"
func main() {
	slice := []string{"Hello", "World"}

	// Dibawah ini adalah jawaban nomor 2: silahkan kalian copy paste dari jawaban nomor 2
	// TODO: answer here
	slice = append(slice, "Olleh") //kalau g pakai ini kan g perlu pake for trus tinggal langsung print bisa kan ya........
	// Dibawah ini adalah jawaban nomor 3 silahkan kalian isi
	// TODO: answer here

	slice[0] = "marcus"
	slice[1] = "Audrelius"

	for _, value := range slice {
		fmt.Println(value)
	}
}