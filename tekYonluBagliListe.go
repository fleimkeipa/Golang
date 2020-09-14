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
		fmt.Println("ben aslinda yoğum")
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
		fmt.Println("yoh yoh ha zorla mı getireyim kardeşim")
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
	okuyucu := bufio.NewReader(os.Stdin)
	for true {
		print("Ne yapmak istersiniz?\n Ekleme(E) Silme(S) Tüm Türler(T) Çıkış(C)\n Seçiminiz ")
		secim, _ := okuyucu.ReadString('\n')
		secim2 := strings.TrimSpace(secim)
		if secim2 == "E" || secim2 == "e" {
			ekraniTemizle()
			print("eklemek istediğiniz türün adı ")
			secim, _ = okuyucu.ReadString('\n')
			cins1.turEkle(secim)
			cins1.tumTurler()
		} else if secim2 == "S" || secim2 == "s" {
			ekraniTemizle()
			print("silmek istediğiniz türün adı ")
			secim, _ = okuyucu.ReadString('\n')
			cins1.turSil(secim)
			cins1.tumTurler()

		} else if secim2 == "T" || secim2 == "t" {
			ekraniTemizle()
			cins1.tumTurler()
		} else if secim2 == "C" || secim2 == "c" {
			break
		}
	}

}

