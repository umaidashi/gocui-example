package main

import (
	"github.com/rivo/tview"
)

func createCommandList() *tview.List {
	commandList := tview.NewList()
	commandList.SetBorder(true).SetTitle("Command")
	return commandList
}

func createLayout(sidebar tview.Primitive, main tview.Primitive) *tview.Flex {
	header := tview.NewTextView()
	header.SetBorder(true)
	header.SetText("Header")
	header.SetTextAlign(tview.AlignCenter)

	bodyLayout := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(sidebar, 20, 1, true).
		AddItem(main, 0, 1, false)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(bodyLayout, 0, 1, true)

	return layout
}

func createApplication() *tview.Application {
	app := tview.NewApplication()
	pages := tview.NewPages()

	sidebar := tview.NewList()
	sidebar.SetBorder(true)
	sidebar.SetTitle("Sidebar")
	sidebar.AddItem("Command", "", 'c', nil)
	sidebar.AddItem("Setting", "", 's', nil)
	sidebar.AddItem("Help", "", 'h', nil)
	sidebar.AddItem("Exit", "", 'q', func() {
		app.Stop()
	})
	main := tview.NewTextView()
	main.SetBorder(true)
	main.SetText("Hello, world!")
	main.SetTextAlign(tview.AlignCenter)

	layout := createLayout(sidebar, main)
	pages.AddPage("main", layout, true, true)

	app.SetRoot(pages, true)
	return app
}

func main() {
	app := createApplication()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
