package main

import (
	"fmt"
	"test/api/komaArrangements"
	"test/api/komas"
	"test/api/boards"
)

func main() {
	var ok = 0
	var ng = 0
	komaArrangements.TestGet(&ok, &ng)
	komas.TestGet(&ok, &ng)
	boards.TestGet(&ok, &ng)

	fmt.Printf("TEST RESULTS\n OK: %d\n NG: %d\n", ok, ng)

}
