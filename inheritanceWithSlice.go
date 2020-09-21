package main

import (
	"fmt"
	"os"
	"os/exec"
)

type cins struct {
	ad  string
	tur []string
}
type familya struct {
	ad      string
	cinsler []cins
}

func ekraniTemizle() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func turEkle(c cins, gelen string) []string {
	c.tur = append(c.tur, gelen)
	return c.tur
}

func turEkle3() cins {
	cins1 := cins{}
	ekraniTemizle()
	secim := ""
	fmt.Print("Yeni eklenecek cinsin adini girin : ")
	fmt.Scanln(&secim)
	cins1.ad = secim
	fmt.Println("Cikmak icin (x)")
	for {
		fmt.Print("Bu cinse eklenecek turleri giriniz : ")
		fmt.Scanln(&secim)
		if secim == "x" {
			break
		}
		cins1.tur = turEkle(cins1, secim)
	}

	return cins1
}
func main() {
	familya1 := familya{}
	for {
		print("Cins ve tur eklemek icin (E) Cinsleri yazdirmak için (C) Turleri yazdirmak icin (T) Cikmak icin (x)")
		print("\n Seciminiz ")
		secim := ""
		fmt.Scanln(&secim)
		if secim == "E" || secim == "e" {
			ekraniTemizle()
			familya1.cinsler = append(familya1.cinsler, turEkle3())
		} else if secim == "C" || secim == "c" {
			ekraniTemizle()
			for i := range familya1.cinsler {
				println(familya1.cinsler[i].ad)
			}
		} else if secim == "T" || secim == "t" {
			ekraniTemizle()
			for i := range familya1.cinsler {
				println(i, familya1.cinsler[i].ad)
			}
			print("Hangi cinsin turlerini yazdirmak istiyorsunuz ")
			secim2 := 0
			fmt.Scanln(&secim2)
			for i := range familya1.cinsler[secim2].tur {
				println(familya1.cinsler[secim2].tur[i])
			}
		} else if secim == "x" || secim == "X" {
			break
		}
	}
}
