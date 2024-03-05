package main

import (
	"fmt"

	pkg "github.com/vinihiga/go-jwt/pkg"
)

func main() {
	var config string = "{ \"typ\": \"JWT\", \"alg\": \"HS256\" }"
	var mock string = "{ \"foo\": \"bar\" }"

	fmt.Println("<<< Running Mocked JWT Generator >>>")
	fmt.Println("Edit cmd/main.go to test and customize JWT generation")

	var encoder pkg.JwtEncoder = pkg.JwtEncoder{
		SecretKey: "test123",
	}

	var encoded string = encoder.NewJwt(config, mock)

	fmt.Println("<<< Result >>>")
	fmt.Println(encoded)
}
