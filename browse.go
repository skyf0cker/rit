package main

import "github.com/skyf0cker/rit/ui"

type BrowseCmd struct {
	Name string
}

func (b *BrowseCmd) Run() error {
	return ui.Run()
}
