package main

import (
	"fmt"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang-11 -type e -type barfoo example ../testdata/minimal.c

func main() {
	var objs exampleObjects
	if err := loadExampleObjects(&objs, nil); err != nil {
		panic("Can't load objects: " + err.Error())
	}
	defer objs.Close()

	// Do something useful with the program.
	fmt.Println(objs.Filter.String())
}
