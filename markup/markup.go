package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 100)
	win.SetTitle("Markup label")
	win.SetBorderWidth(10)
	
	str := `“The glory which is built upon a lie soon becomes a most unpleasant incumbrance. 
…  How easy it is to make people believe a lie, and how hard it is to undo that work again!” 
																	<b>- Mark Twain</b>`

	label, err := gtk.LabelNew("")																 
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	label.SetMarkup(str)
	win.Add(label)
	
	win.Connect("destroy", gtk.MainQuit)
	
	win.ShowAll()
	gtk.Main()
}