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
	win.SetDefaultSize(300, 200)
	win.SetTitle("GtkHSeparator")
	win.SetResizable(false)
	win.SetBorderWidth(10)

	label1, err := gtk.LabelNew(`Zinc is a moderately reactive, blue gray metal
that tarnishes in moist air and burns in air with a bright bluish-green flame,
giving off fumes of zinc oxide. It reacts with acids, alkalis and other non-metals.
If not completely pure, zinc reacts with dilute acids to release hydrogen.`)
	if err != nil {
		log.Fatal("Unable to create label1:", err)
	}

	label2, err := gtk.LabelNew(`Copper is an essential trace nutrient to all high
plants and animals. In animals, including humans, it is found primarily in
the bloodstream, as a co-factor in various enzymes, and in copper-based pigments.
However, in sufficient amounts, copper can be poisonous and even fatal to organisms.`)
	if err != nil {
		log.Fatal("Unable to create label2:", err)
	}

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create vBox:", err)
	}

	hSeparator, err := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		log.Fatal("Unable to create hSeparator:", err)
	}
	
	vBox.PackStart(label1, false, true, 0)
	vBox.PackStart(hSeparator, false, true, 10)
	vBox.PackStart(label2, false, true, 0)
	
	win.Add(vBox)
	win.Connect("destroy", gtk.MainQuit)
	
	win.ShowAll()
	gtk.Main()
}
