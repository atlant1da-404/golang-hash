package hasher

type Hash struct {
	salt    byte
	planets byte
}

func NewHash(salt byte) *Hash {
	return &Hash{
		salt:    salt,
		planets: 8,
	}
}

func (h *Hash) GenerateHash(text string, key string) string {

	bText := []byte(text)
	var planets byte = 8

	for index, _ := range bText {
		bText[index] = bText[index] + (h.salt+planets+byte(len(key)/4))/2
	}

	return string(bText)
}

func (h *Hash) UnHash(text string, key string) string {

	bText := []byte(text)
	var planets byte = 8

	for index, _ := range bText {
		bText[index] = bText[index] - (h.salt+planets+byte(len(key)/4))/2
	}

	return string(bText)
}
