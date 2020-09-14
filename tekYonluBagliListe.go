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

func main() {
	cins1 := cinsOlustur("homo")
	cins1.turEkle("sapiens")
	cins1.turEkle("habilis")
	cins1.turEkle("erectus")
	cins1.tumTurler()
	cins1.turSil("sapiens")
	fmt.Println("-------------")
	cins1.tumTurler()
}
