package testcases

import (
	"os"
	"simple_hash/lib/hasher"
	"testing"
)

func RunHash(t *testing.T) {

	hasher.NewHash(5).GenerateHash("mike", "23nasmn1")
	if os.ErrExist != nil {
		t.Errorf("something wrong when run")
	}
}
