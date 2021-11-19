package snowflake

import (
	"fmt"
	"testing"
)

func TestGoSnowflake(t *testing.T) {
	err := Init("2021-07-21", 1)
	if err != nil {
		t.Error("err>>>：", err)
	}
	fmt.Println("GenID\tfmt.Println(\"GenID：\",GenID())\n：", GenID())
}

func BenchmarkSnowflake(b *testing.B) {
	err := Init("2021-07-21", 1)
	if err != nil {
		b.Error("err>>>：", err)
	}
	for i := 0; i < b.N; i++ {
		b.Log("GenID())：", GenID())
	}
}
