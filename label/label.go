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
	win.SetDefaultSize(300, 250)
	win.SetTitle("No sleep")
	win.SetBorderWidth(15)

	label, err := gtk.LabelNew(`I've always been too lame
To see what's before me
And I know nothing sweeter than
Champaign from last New Years
Sweet music in my ears
And a night full of no fears
	
But if I had one wish fulfilled tonight
I'd ask for the sun to never rise
If God passed a mic to me to speak
I'd say "Stay in bed, world,
Sleep in peace"`)

	label.SetJustify(gtk.JUSTIFY_LEFT)
	win.Add(label)

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}