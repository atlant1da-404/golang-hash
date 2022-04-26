package testcases

import (
	"os"
	"simple_hash/lib/hasher"
	"testing"
)

func InitTest(t *testing.T) {

	hasher.NewHash(5)
	if os.ErrExist != nil {
		t.Errorf("something wrong when run")
	}
}
