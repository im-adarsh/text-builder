package main

import (
	"fmt"
	"github.com/im-adarsh/text-builder/txtBuilder"
)

func main() {
	str := txtBuilder.NewTextBuilder().Template("hello %v").Params([]string{"world"}).GetText()
	fmt.Println(str)
}
