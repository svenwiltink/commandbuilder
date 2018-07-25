package main

import (
	"github.com/svenwiltink/commandbuilder"
	"fmt"
)

func main() {
	builder := commandbuilder.New("echo")
	builder.AddArgument("banaan")

	cmd := builder.ToCmd()
	output, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
