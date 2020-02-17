package main
import (
	"fmt"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main(){
		fmt.Println(morestrings.ReverseRuns("hello world"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}