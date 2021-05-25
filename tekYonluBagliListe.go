package main

import "fmt"

type tur struct {
	ad      string
	sonraki *tur
}
type cins struct {
	ad     string
	turler *tur
}

func cinsOlustur(gelenAd string) *cins {
	return &cins{
		ad: gelenAd,
	}
}

func (c *cins) turEkle(gelenAd string) error {
	gecici := &tur{
		ad: gelenAd,
	}
	if c.turler == nil {
		c.turler = gecici
	} else {
		gecici2 := c.turler
		for gecici2.sonraki != nil {
			gecici2 = gecici2.sonraki
		}
		gecici2.sonraki = gecici
	}
	return nil
}

func (c *cins) turSil(gelenAd string) error {
	gecici := c.turler
	gecici2 := c.turler
	if gecici == nil {
		fmt.Println("bos deger atamasi.")
	} else {
		if gecici.ad == gelenAd {
			c.turler = c.turler.sonraki
		} else {
			for gecici.ad != gelenAd {
				gecici2 = gecici
				gecici = gecici.sonraki
			}
			if gecici.sonraki == nil {
				gecici2.sonraki = nil
			} else {
				gecici2.sonraki = gecici.sonraki
			}
		}
	}
	return nil
}

func (c *cins) tumTurler() error {
	gecici := c.turler
	if gecici == nil {
		fmt.Println("herhangi bir sey bulunamadi.")
	} else {
		fmt.Printf("%+v\n", *gecici)
		for gecici.sonraki != nil {
			gecici = gecici.sonraki
			fmt.Printf("%+v\n", *gecici)
		}
	}
	return nil
}

func ekraniTemizle() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	cins1 := cinsOlustur("homo")
	for {
		print("Ne yapmak istersiniz?\n Ekleme(E) Silme(S) Tüm Türler(T) Çıkış(C)\n Seçiminiz ")
		secim := ""
		fmt.Scanln(&secim)
		if secim == "E" || secim == "e" {
			ekraniTemizle()
			print("eklemek istediğiniz türün adı ")
			fmt.Scanln(&secim)
			cins1.turEkle(secim)
			cins1.tumTurler()
		} else if secim == "S" || secim == "s" {
			ekraniTemizle()
			print("silmek istediğiniz türün adı ")
			fmt.Scanln(&secim)
			cins1.turSil(secim)
			cins1.tumTurler()
		} else if secim == "T" || secim == "t" {
			ekraniTemizle()
			cins1.tumTurler()
		} else if secim == "C" || secim == "c" {
			break
		}

}

