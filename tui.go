package main

import (
	"github.com/rivo/tview"
)

func GetHomePage() *tview.Grid {
	testText := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("zfetch")

	return tview.NewGrid().
		SetRows(-1).
		SetColumns(5).
		SetBorders(true).
		AddItem(testText, 0, 0, 1, 3, 0, 0, true)
}

func StartTui() {
	app := tview.NewApplication()

	//box := tview.NewBox().SetBorder(true).SetBorderColor(tcell.Color45)
	layout := tview.NewFlex().SetDirection(tview.FlexRow).SetFullScreen(true)
	//layout.AddItem(box, 0, 1, false)
	layout.AddItem(GetHomePage(), 0, 1, true)

	app.SetRoot(layout, true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}