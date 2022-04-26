package testcases

import (
	"os"
	"simple_hash/lib/hasher"
	"testing"
)

func ReadHash(t *testing.T) {

	value := hasher.NewHash(5).GenerateHash("mike", "23nasmn1")
	if os.ErrExist != nil {
		t.Errorf("something wrong when run")
	}

	if value == "" && len(value) <= 0 {
		t.Errorf("hash not work")
	}
}

