package main


import (
	"strings"
	"testing"

	"os"
	"fmt"
)
var a = []string{"2,2,1,23,2,2"}
func main(){
	var s, sep string
	for i := 1; i <len(a); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}


func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(a, " "))

	}
}
