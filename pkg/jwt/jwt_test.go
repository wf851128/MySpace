package jwt

import (
	"fmt"
	"testing"
)

func TestGoSnowflake(t *testing.T) {
	aToken, rToken, err := GenToken(1122, "1")
	if err != nil {
		t.Error("err>>>：", err)
	}
	fmt.Printf("aToken :%s \n rToken: %s ", aToken, rToken)
	fmt.Println(ParseToken(aToken))
}

func BenchmarkJwt(b *testing.B) {
	aToken, rToken, err := GenToken(1122, "1")
	if err != nil {
		b.Error("err>>>：", err)
	}
	for i := 0; i < b.N; i++ {
		fmt.Printf("aToken :%s \n rToken: %s ", aToken, rToken)
		fmt.Println(ParseToken(aToken))
	}
}
