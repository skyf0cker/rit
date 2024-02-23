package main

import "rit/ui"

type BrowseCmd struct {
	Name string
}

func (b *BrowseCmd) Run() error {
	return ui.Run()
}
