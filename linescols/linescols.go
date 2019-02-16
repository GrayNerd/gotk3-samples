package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}

	win.SetTitle("Text view")
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create vertical box: ", err)
	}

	toolbar, err := gtk.ToolbarNew()
	if err != nil {
		log.Fatal("Unable to create toolbar: ", err)
	}
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

	img, err := gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("application-exit", gtk.ICON_SIZE_BUTTON)
	exit, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create newTb:", err)
	}
	exit.SetName("exit")
	toolbar.Insert(exit, -1)
	vBox.PackStart(toolbar, false, false, 5)

	view, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create text view: ", err)
	}
	view.SetWrapMode(gtk.WRAP_WORD)
	vBox.PackStart(view, true, true, 0)
	view.GrabFocus()
	
	buffer, err := view.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer: ", err)
	}
	
	statusbar, err := gtk.StatusbarNew()
	if err != nil {
		log.Fatal("Unable to create statusbar: ", err)
	}
	vBox.PackStart(statusbar, false, false, 0)
	win.Add(vBox)
	
	exit.Connect("clicked", gtk.MainQuit)
	
	// buffer.Connect("changed", func(buffer *gtk.TextBuffer) {
	// 	iter := buffer.GetIterAtMark(buffer.CreateMark("mark", ???, false))
	// 	updateStatusbar(iter, statusbar)
	// })
	
	buffer.Connect("mark-set", func(buffer *gtk.TextBuffer, iter *gtk.TextIter, mark *gtk.TextMark) {
		updateStatusbar(iter, statusbar)
	})
	
	win.Connect("destroy", gtk.MainQuit)
	
	updateStatusbar(buffer.GetStartIter(), statusbar)
	win.ShowAll()
	gtk.Main()
}

func updateStatusbar(iter *gtk.TextIter, statusbar *gtk.Statusbar) {
	statusbar.Pop(0)
	row := iter.GetLine()
	col := iter.GetLineOffset()
	statusbar.Push(0, fmt.Sprintf("Row: %d, Col: %d", row, col))

}

func markSetCallback(buffer *gtk.TextBuffer, statusbar *gtk.Statusbar) {
	

}
