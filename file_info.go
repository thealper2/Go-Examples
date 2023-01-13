package main
import "fmt"
import "os"
import "log"
var {
	fileInfo *os.fileInfo
	err error
}
func main() {
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dosya adi: ",fileInfo.Name())
	fmt.Println("Byte: ",fileInfo.Size())
	fmt.Println("Yetkileri: ",fileInfo.Mode())
	fmt.Println("Son islem: ",fileInfo.ModTime())
	fmt.Println("Dizin: ",fileInfo.IsDir())
	fmt.Println("Sistem Arayuz: ",fileInfo.Sys())
	fmt.Println("Sistem Bilgisi: %+v\n\n",fileInfo.Sys())
}