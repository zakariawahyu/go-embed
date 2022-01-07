package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

/**
Hasil Embed di Compile
- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file Go-Lang
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file externalnya dan bahkan jika diubah file externalnya,
isi variabelnya tidak akan berubah lagi

menjalankan compile
go build
./go-embed
*/

//go:embed version.txt
var version string

//go:embed logo.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println("Version : ", version)

	err := ioutil.WriteFile("logo-next.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

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
