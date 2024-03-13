# go-jwt (Experimental)
A JWT implementation from scratch without external dependencies and just using Go Language.

### Usage

To create a new token, you can use `JwtEncoder`. To use you need to have to instantiate that struct and call the inner method `JwtEncoder.NewJwt`.

To verify if a token is valid, you can use `JwtDecoder`. You need to instantiate that struct and call `JwtDecoder.Validate` method with the Jwt token. It returns if the validation was succesful, otherwise returns some parsing error, invalid format or if couldn't find the expiration claim.

### To Contribute

Feel free to open new `pull requests` as you desired and explaining why the change is needed. Since it's a brand new project, we don't have any patterns, but of course, we can discuss the best strategies to make this code more scalable, cleaners and faster.

### Author

Original maintainer: Vinícius Hiroshi Higa (vinihiga)
