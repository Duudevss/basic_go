package main

import "fmt"

func main() {
	{
		msg := "Patinya pimpo" + " " + "Kook!"
		age := 28
		price, check := 100.00, true

		fmt.Println("Name:", msg)
		fmt.Println("Age:", age)
		fmt.Println("Price:", price)
		fmt.Println("Check:", check)
	}
}

func workshop() {
	{
		movieName := "Avengers: Endgame"
		year := 2019
		rating := 8.4
		genre := "Sci-Fi"
		isSuperHero := true

		fmt.Println("เรื่อง:", movieName)
		fmt.Println("ปี:", year)
		fmt.Println("เรตติ้ง:", rating)
		fmt.Println("ประเภท:", genre)
		fmt.Println("ซุปเปอร์ฮีโร่:", isSuperHero)
	}
}
