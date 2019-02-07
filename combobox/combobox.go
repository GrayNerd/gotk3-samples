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
	win.SetTitle("GtkComboBox")
	win.SetBorderWidth(15)

	hBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		log.Fatal("Unable to create hBox:", err)
	}
	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create vBox", err)
	}

	combo, err := gtk.ComboBoxTextNew()
	combo.AppendText("Ubuntu")
	combo.AppendText("Arch")
	combo.AppendText("Fedora")
	combo.AppendText("Mint")
	combo.AppendText("Gentoo")
	combo.AppendText("Debian")

	vBox.PackStart(combo, false, false, 0)

	label, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	vBox.PackStart(label, false, false, 0)
	
	hBox.PackStart(vBox, false, false, 0)
	win.Add(hBox)

	win.Connect("destroy", gtk.MainQuit)
	combo.Connect("changed", func() {
		comboSelected(combo, label)
	})

	win.ShowAll()
	gtk.Main()
}

func comboSelected(c *gtk.ComboBoxText, l *gtk.Label) {
	t := c.GetActiveText()
	l.SetLabel(t)
}