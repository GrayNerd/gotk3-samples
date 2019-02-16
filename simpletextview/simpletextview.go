package main

import (
	"log"

	"github.com/gotk3/gotk3/pango"

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

	table, err := gtk.TextTagTableNew()
	if err != nil {
		log.Fatal("Unable to create tag table: ", err)
	}

	tag, _ := gtk.TextTagNew("lmarg")
	tag.SetProperty("left-margin", 5)
	table.Add(tag)

	tag, _ = gtk.TextTagNew("blue-fg")
	tag.SetProperty("foreground", "blue")
	table.Add(tag)

	tag, _ = gtk.TextTagNew("gray-bg")
	tag.SetProperty("background", "gray")
	table.Add(tag)

	tag, _ = gtk.TextTagNew("ital")
	tag.SetProperty("style", pango.STYLE_ITALIC)
	table.Add(tag)

	tag, _ = gtk.TextTagNew("bold")
	tag.SetProperty("weight", pango.WEIGHT_BOLD)
	table.Add(tag)

	buffer, err := gtk.TextBufferNew(table)
	if err != nil {
		log.Fatal("Unable to get buffer: ", err)
	}

	view, err := gtk.TextViewNewWithBuffer(buffer)
	if err != nil {
		log.Fatal("Unable to create text view: ", err)
	}
	vBox.PackStart(view, true, true, 0)

	buffer.Insert(buffer.GetStartIter(), "Plain text\n")

	// weirdness ... if this mark isn't named, then indenting may not be consistent
	// depending on how the other marks are named
	indent := buffer.CreateMark("ident", buffer.GetEndIter() , true)

	start := buffer.CreateMark("", buffer.GetEndIter(), true )
	buffer.InsertAtCursor("Colored Text\n")
	iter := buffer.GetIterAtMark(start)
	buffer.ApplyTagByName("blue-fg", iter, buffer.GetEndIter())
	
	start = buffer.CreateMark("", buffer.GetEndIter(), true )
	buffer.InsertAtCursor("Text with colored background\n")
	iter = buffer.GetIterAtMark(start)
	buffer.ApplyTagByName("gray-bg", iter, buffer.GetEndIter())
	
	start = buffer.CreateMark("", buffer.GetEndIter(), true )
	buffer.InsertAtCursor("Text in italics\n")
	iter = buffer.GetIterAtMark(start)
	buffer.ApplyTagByName("ital", iter, buffer.GetEndIter())

	start = buffer.CreateMark("", buffer.GetEndIter(), true )
	buffer.InsertAtCursor("Bold text\n")
	iter = buffer.GetIterAtMark(start)
	buffer.ApplyTagByName("bold", iter, buffer.GetEndIter())

	iter = buffer.GetIterAtMark(indent)
	buffer.ApplyTagByName("lmarg", iter, buffer.GetEndIter())

	win.Connect("destroy", gtk.MainQuit)

	win.Add(vBox)

	win.ShowAll()
	gtk.Main()
}
