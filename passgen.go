package passgen

import (
	charstore "github.com/mrumyantsev/passgen/internal/pkg/char-store"
	"github.com/mrumyantsev/passgen/internal/pkg/randomizer"
)

type PassGen struct {
	charStore *charstore.CharStore
}

func New() *PassGen {
	var (
		passGen    = &PassGen{}
		randomizer = randomizer.New()
		store      = charstore.New(randomizer)
	)

	store.Init()

	passGen.charStore = store

	return passGen
}

func (p *PassGen) Generate(passwordLength int) []byte {
	var (
		password  = make([]byte, passwordLength)
		character byte
	)

	p.charStore.Reset()
	p.charStore.Shuffle()

	for i := 0; i < passwordLength; i++ {
		character = p.charStore.GetCharacter()
		password[i] = character
	}

	return password
}