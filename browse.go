package main

import "fmt"

type BrowseCmd struct {
	Name string
}

func (b *BrowseCmd) Run() error {
	fmt.Println("hi")
	return nil
}
