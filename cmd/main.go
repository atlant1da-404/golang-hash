package main

import (
	"fmt"
	"simple_hash/lib/hasher"
)

func main() {
	hash := hasher.NewHash(5)

	message := `This world need new heroes! Spider Man, Iron Man and etc.
PostreSQL is cool. MongoDB not bad. Rust cool.`

	value := hash.GenerateHash(message, "password")

	fmt.Println(value)

	fmt.Println(hash.UnHash(value, "password"))
}
