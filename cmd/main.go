package main

import (
	"fmt"
	"time"

	pkg "github.com/vinihiga/go-jwt/pkg"
	"github.com/vinihiga/go-jwt/pkg/models"
)

func main() {
	var payload models.ClaimsModel = models.ClaimsModel{
		Exp: time.Now().UnixMilli() + (10 * 1000), // Time.Now() in miliseconds + 10 seconds
	}

	fmt.Println("<<< Running Mocked JWT Generator >>>")
	fmt.Println("Edit cmd/main.go to test and customize JWT generation")

	var encoder pkg.JwtEncoder = pkg.JwtEncoder{
		SecretKey: "test123",
		Algorithm: "HS256",
	}

	var decoder pkg.JwtDecoder = pkg.JwtDecoder{
		SecretKey: "test123",
		Algorithm: "HS256",
	}

	var encoded string = encoder.NewJwt(payload)

	fmt.Println("\n<<< Result >>>")
	fmt.Println(encoded)
	fmt.Println()

	fmt.Println("<<< Making sure we can is valid >>>")
	isEqual, _ := decoder.Validate(encoded)

	fmt.Printf("Is equal and valid: %t", isEqual)
}
