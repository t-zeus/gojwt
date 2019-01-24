# gojwt

concept：
```text
jwt string  = <header>.<payload>.<signature>
<header>    = Base64URL(header json)
<payload>   = Base64URL(payload json)
<signature> = Base64URL(signature)
signature   = HMAC-SHA256(<header>.<payload>) 
```

## APIs

```text
# generate jwt string
func GenJWT(v interface{}, key string) string

# verify jwt string
func VerifyJWT(jwt, key string) bool
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/t-zeus/gojwt"
)

func main() {
	key1 := "secret key"
	key2 := "secret key2"
	data := map[string]interface{}{
		"user_id": 1,
		"gender":  "male",
		"hobby": []string {
			"running",
			"pay games",
		},
	}
	val1 := gojwt.GenJWT(data, key1)
	val2 := gojwt.GenJWT("test", key2)
	fmt.Println(gojwt.VerifyJWT(val1, key1)) // true
	fmt.Println(gojwt.VerifyJWT(val2, key1)) // false
}
```