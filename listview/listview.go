package main

import (
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	iistItem = iota
	nColumns
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(270, 250)
	win.SetTitle("List view")

	list, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal(err)
	}
	list.SetHeadersVisible(false)

	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal(err)
	}
	vbox.PackStart(list, true, true, 5)

	label, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal(err)
	}
	vbox.PackStart(label, false, false, 5)

	win.Add(vbox)

	store := initList(list)
	addToList(store, "Aliens")
	addToList(store, "Leon")
	addToList(store, "The Verdict")
	addToList(store, "North Face")
	addToList(store, "Der Untergang")

	selection, err := list.GetSelection()
	if err != nil {
		log.Fatal(err)
	}

	selection.Connect("changed", func() {
		onChanged(selection, label)
	})

	win.Connect("destroy", gtk.MainQuit)
	
	win.ShowAll()
	gtk.Main()
}

func initList(list *gtk.TreeView) *gtk.ListStore {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal(err)
	}

	column, err := gtk.TreeViewColumnNewWithAttribute("List Items", cellRenderer, "text", 0)
	if err != nil {
		log.Fatal(err)
	}
	list.AppendColumn(column)

	store, err := gtk.ListStoreNew(glib.TYPE_STRING)
	if err != nil {
		log.Fatal(err)
	}
	list.SetModel(store)
	
	return store
}

func addToList(store *gtk.ListStore, text string) {
	iter := store.Append()
	store.SetValue(iter, 0, text)
}

func onChanged(selection *gtk.TreeSelection, label *gtk.Label) {
	var iter *gtk.TreeIter
	var model gtk.ITreeModel
	var ok bool
	model, iter, ok = selection.GetSelected()

	if ok {
		value, err := model.(*gtk.TreeModel).GetValue(iter, 0)
		if err != nil {
			log.Fatal(err)
		}
		
		text, err := value.GetString()
		if err != nil {
			log.Fatal(err)
		}
			
		label.SetText(text)
	}
}
	

