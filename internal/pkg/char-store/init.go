package charstore

import (
	"github.com/mrumyantsev/passgen/internal/pkg/consts"
	"github.com/mrumyantsev/passgen/internal/pkg/node"
)

// Initializes specialized char sets.
// E.g. A-Z, a-z, 0-9, etc.
func (c *CharStore) Init() {
	c.collection = node.Make(consts.DEFAULT_CHAR_SETS_COUNT)

	collectionSettings := []struct {
		pattern   string
		charSetId byte
	}{
		{"0-9", consts.ID_NUMBERS},
		{"@#$%&", consts.ID_SPEC_CHARS},
		{"a-z", consts.ID_LOW_LETTERS},
		{"A-Z", consts.ID_HIGH_LETTERS},
	}

	var node *node.Node

	for _, setting := range collectionSettings {
		node = makeCharsWithPattern(setting.pattern)
		node.SetValue(setting.charSetId)
		c.collection.Push(node)
	}

	c.initAvailableItems()
}

func (c *CharStore) initAvailableItems() {
	c.availableItems = []int{
		consts.ID_NUMBERS,
		consts.ID_SPEC_CHARS,
		consts.ID_LOW_LETTERS,
		consts.ID_HIGH_LETTERS,
	}

	c.count = consts.DEFAULT_CHAR_SETS_COUNT
}

func makeCharsWithPattern(pattern string) *node.Node {
	if len(pattern) == 3 {
		if pattern[1] == '-' {
			return makeCharsWithRange(pattern[0], pattern[2])
		}
	}

	return makeCharsWithSpecifiedChars(pattern)
}

func makeCharsWithRange(startChar, endChar byte) *node.Node {
	var (
		setting = node.Make(int(endChar-startChar) + 1)
	)

	for ch := startChar; ch <= endChar; ch++ {
		charItem := node.New()
		charItem.SetValue(ch)
		setting.Push(charItem)
	}

	return setting
}

func makeCharsWithSpecifiedChars(specifiedChars string) *node.Node {
	var (
		setting = node.Make(len(specifiedChars))
	)

	for _, ch := range specifiedChars {
		charItem := node.New()
		charItem.SetValue(byte(ch))
		setting.Push(charItem)
	}

	return setting
}
