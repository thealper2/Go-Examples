package main
import "fmt"
import "../utils"
func main() {
	n1,l1 := utils.ad("Alper","Karaca")
	fmt.Printf("Tam ad: %v, karakter sayisi: %v\n\n",n1,l1)
	n2,l2 := utils.tamad("Arthur","Morgan")
	fmt.Printf("Tam ad: %v, karakter sayisi: %v\n\n",n1,l1)

}