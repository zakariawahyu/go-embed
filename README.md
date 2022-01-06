**Embed Package**
- Sejak Go-Lang versi 1.16 terdapat package baru dengan nama embed
- Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan isi filenya kedalam variabel
- https://pkg.go.dev/embed

**Cara Embed File**
- Untuk melakukan embed file ke variabel, kita bisa mengimport package embed terlebih dahulu 
- Selanjutnya kita bisa tambahkan komentar //go:embed di ikuti nama filenya, diatas variabel yang kita tuju
- Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yan gkita inginkan secara otomatis ketika kode GO-Lang di compile
- Variable yang di tuju tidak bisa disimpan di dalam function