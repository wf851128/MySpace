package jwt

import (
	"fmt"
	"testing"
)

func TestGoSnowflake(t *testing.T) {
	token, err := GenToken(1122, "1")
	if err != nil {
		t.Error("err>>>：", err)
	}
	fmt.Printf("%s\n：\n", token)
	fmt.Println(ParseToken(token))
}

func BenchmarkJwt(b *testing.B) {
	token, err := GenToken(1122, "1")
	if err != nil {
		b.Error("err>>>：", err)
	}
	for i := 0; i < b.N; i++ {
		fmt.Printf("%s\n：\n", token)
		fmt.Println(ParseToken(token))
	}
}
