package main

import (
	"fmt"

	pkg "github.com/vinihiga/go-jwt/pkg"
)

func main() {
	var header string = "{ \"typ\": \"JWT\", \"alg\": \"HS256\" }"
	var payload string = "{ \"foo\": \"bar\" }"

	fmt.Println("<<< Running Mocked JWT Generator >>>")
	fmt.Println("Edit cmd/main.go to test and customize JWT generation")

	var encoder pkg.JwtEncoder = pkg.JwtEncoder{
		SecretKey: "test123",
	}

	var encoded string = encoder.NewJwt(header, payload)

	fmt.Println("<<< Result >>>")
	fmt.Println(encoded)
}
