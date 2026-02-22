package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(10)
}

var letterRunes = getEnglishAlphabet()

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getEnglishAlphabet() []rune {
	var alphabet []rune
	// Добавляем заглавные буквы A-Z (коды Unicode: 65–90)
	for i := 65; i <= 90; i++ {
		alphabet = append(alphabet, rune(i))
	}
	// Добавляем строчные буквы a-z (коды Unicode: 97–122)
	for i := 97; i <= 122; i++ {
		alphabet = append(alphabet, rune(i))
	}
	return alphabet
}
