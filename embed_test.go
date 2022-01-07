package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

/**
Embed File ke String
- Embed file bisa kita lakukan ke variabel dengan tipe data string
- Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variabel tersebut
- File boleh di embed berkali-kali walaupun file sama atau berbeda-beda
*/

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println("Current version : ", version)
	fmt.Println("Current version : ", version2)
}

/**
Embed File ke []Byte
- Selain ke tipe data String, embed file juga bisa dilakukan ke variabel tipe data []byte
- Ini cocok sekali jika ktia ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain
- File-file yang tidak bisa dibuka text editor
*/

//go:embed logo.jpg
var logo []byte

func TestByte(t *testing.T) {
	// test dengan memindahkan file baru dengan data sama
	// cek apakah data gambarnya sama
	err := ioutil.WriteFile("logo-next.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

/**
Embed Multiple File
- Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
- Hal ini juga bisa dilakukan menggunakan embed package
- Kita bisa menambahkan komentar //go:embed lebih dari satu baris
- Selain itu variabelnya bisa kita gunakan tipe data embed.FS (singaktan File System)
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

	if err != nil {
		panic(err)
	}
}

/**
Path Macther
- Selain manual satu persat, kita bisa menggunakan path macther untuk membaca multiple file yang kita inginkan
- Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
- Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match
https://pkg.go.dev/path#Match
*/

//go:embed files/*.txt
var path embed.FS

func TestPathMatc(t *testing.T) {
	dir, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}
