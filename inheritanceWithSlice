package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	raider := bufio.NewReader(os.Stdin)
	ekraniTemizle()
	fmt.Print("Yeni eklenecek cinsin adini girin : ")
	secim, _ := raider.ReadString('\n')
	secim2 := strings.TrimSpace(secim)
	cins1.ad = secim2
	fmt.Println("Cikmak icin (x)")
	for {

		fmt.Print("Bu cinse eklenecek turleri giriniz : ")
		secim, _ := raider.ReadString('\n')
		secim2 := strings.TrimSpace(secim)
		if secim2 == "x" {
			break
		}
		cins1.tur = turEkle(cins1, secim2)
	}

	return cins1
}
func main() {
	familya1 := familya{}

	okuyucu := bufio.NewReader(os.Stdin)
	for {
		print("Cins ve tur eklemek icin (E) Cinsleri yazdirmak için (C) Turleri yazdirmak icin (T) Cikmak icin (x)")
		print("\n Seciminiz ")
		secim2, _ := okuyucu.ReadString('\n')
		secim := strings.TrimSpace(secim2)
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
			secim2, _ := okuyucu.ReadString('\n')
			secim, _ := strconv.Atoi(secim2)
			for i := range familya1.cinsler[secim].tur {
				println(familya1.cinsler[secim].tur[i])
			}
		} else if secim == "x" || secim == "X" {
			break
		}
	}
}
