package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		fmt.Println("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.Show()

	gtk.Main()
}
