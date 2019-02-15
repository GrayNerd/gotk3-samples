package main

import (
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}

	win.SetTitle("Tree view")
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(350, 300)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	if err != nil {
		log.Fatal(err)
	}

	view, model := createViewAndModel()
	_ = model
	selection, err := view.GetSelection()
	if err != nil {
		log.Fatal(err)
	}
	vBox.PackStart(view, true, true, 1)

	statusbar, err := gtk.StatusbarNew()
	if err != nil {
		log.Fatal(err)
	}
	vBox.PackStart(statusbar, false, false, 0)

	selection.Connect("changed", func() {
		onChanged(selection, statusbar)
	})
	win.Connect("destroy", gtk.MainQuit)

	win.Add(vBox)
	win.ShowAll()
	gtk.Main()
}

func createViewAndModel() (*gtk.TreeView, *gtk.TreeStore) {
	view, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal(err)
	}
	col, err := gtk.TreeViewColumnNew()
	if err != nil {
		log.Fatal(err)
	}
	col.SetTitle("Programming languages")
	view.AppendColumn(col)

	renderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal(err)
	}
	col.PackStart(renderer, true)
	col.AddAttribute(renderer, "text", 0)

	model := createAndFillModel()
	view.SetModel(model)

	return view, model
}

func createAndFillModel() *gtk.TreeStore {
	var topLevel *gtk.TreeIter

	store, err := gtk.TreeStoreNew(glib.TYPE_STRING)
	if err != nil {
		log.Fatal(err)
	}
	topLevel = store.Insert(nil, -1)
	err = store.SetValue(topLevel, 0, "Scripting Languages")
	if err != nil {
		log.Fatal("Unable to set topLevel value: ", err)
	}

	child := store.Insert(topLevel, -1)
	store.SetValue(child, 0, "Python")
	child = store.Insert(topLevel, -1)
	store.SetValue(child, 0, "Perl")
	child = store.Insert(topLevel, -1)
	store.SetValue(child, 0, "PHP")

	topLevel = store.Insert(nil, -1)
	err = store.SetValue(topLevel, 0, "Compiled Languages")
	if err != nil {
		log.Fatal("Unable to set topLevel value: ", err)
	}

	child = store.Insert(topLevel, -1)
	store.SetValue(child, 0, "C")
	child = store.Insert(topLevel, -1)
	store.SetValue(child, 0, "Go")
	child = store.Insert(topLevel, -1)
	store.SetValue(child, 0, "Rust")

	return store
}

func onChanged(selection *gtk.TreeSelection, sb *gtk.Statusbar) {
	model, iter, ok := selection.GetSelected()
	if ok {
		model, iter, ok = selection.GetSelected()
		if ok {
			tpath, err := model.(*gtk.TreeModel).GetValue(iter, 0)
			if err != nil {
				log.Printf("treeSelectionChangedCB: Could not get path from model: %s\n", err)
				return
			}
			str, err := tpath.GetString()
			if err != nil {
				log.Fatal("Unable to get string: ", err)
			}
			contextID := sb.GetContextId("text")
			sb.Push(contextID, str)
		}
	}

}
