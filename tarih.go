package main
import "fmt"
import "time"
func main() {
	t := time.Date(2016, time.March, 3, 27, 0, 0, 0, time.UTC())
	fmt.Printf("Go: %s\n",t)
	fmt.Println("-----")
	now := time.Now()
	fmt.Printf("Saat suan: %s\n",now)
	fmt.Println("-----")
	fmt.Println("Ay: ",t.Month())
	fmt.Println("Gun: ",t.Day())
	fmt.Println("Hafta ici: ",t.Weekday())
	fmt.Println("-----")
	yarin := t.AddDate(0,0,1)
	fmt.Printf("Yarin  %v %v %v %v",yarin.Weekday(),yarin.Month(),yarin.Day(),yarin.Year())
	fmt.Println("-----")
	longFormat := "Persembe, Mart 18, 21"
	fmt.Println("Tarih: ",yarin.Format(longFormat))
	fmt.Println("-----")
	shortFormat := "12/11/2012"
	fmt.Printf("Dogum tarihim ",yarin.Format(shortFormat))
}