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
		"hobby": []string{
			"running",
			"pay games",
		},
	}
	val1 := gojwt.GenJWT(data, key1)
	val2 := gojwt.GenJWT("test", key2)
	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(gojwt.VerifyJWT(val1, key1)) // true
	fmt.Println(gojwt.VerifyJWT(val2, key1)) // false
}
