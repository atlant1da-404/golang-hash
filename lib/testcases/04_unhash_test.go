package testcases

import (
	"os"
	"simple_hash/lib/hasher"
	"testing"
)

func UnHash(t *testing.T) {

	testMessage := "mike"

	crypt := hasher.NewHash(5).GenerateHash(testMessage, "23nasmn1")
	if os.ErrExist != nil {
		t.Errorf("something wrong when run")
	}

	value := hasher.NewHash(5).UnHash(crypt, "23nasmn1")

	if value != testMessage {
		t.Errorf("hash not readeble")
	}
}
