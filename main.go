package main

import (
	"fmt"
	"main/commands"
)

func main() {
    retorno := commands.Cd()
    fmt.Println(retorno)
    // commands.LsTree()
}
