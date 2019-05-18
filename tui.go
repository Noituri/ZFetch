package main

import (
	"encoding/base64"
	"github.com/disintegration/imaging"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strings"
)

const TEST_LOGO = `iVBORw0KGgoAAAANSUhEUgAAAFAAAABNCAMAAAAGhxPaAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAhFBMVEU1v1wAAAA1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1w1v1wAAACLICWOAAAAKnRSTlMAAOAf/Ozp6evNHOVDKCkqJAUCEhUUGrrW1O3v0B0uLykGEREPAs/Qthmsk4gjAAAAAWJLR0QB/wIt3gAAAAd0SU1FB+MFEQgvL8ufKIoAAACWSURBVFjD7dNZCsJAEEXRMnGeTWIS53nc/wJ1AdWCT9AI9/4WfaCg2u7vF8VmcRQYGiAgICAgICDgx2C90WwFane6AtjrD4aBRuOJAD7Xqr2oGqAlaTZ1ytJEW9nyopw5lUUugqFH0tkA/hKcuy2WKxVcu222OxUUvteXwb3b4XhSVz67Xa63ypwNICAgICAgICDgP4MPIjViHiX1RUQAAAAldEVYdGRhdGU6Y3JlYXRlADIwMTktMDUtMTdUMDg6NDc6NDctMDQ6MDCAF1wrAAAAJXRFWHRkYXRlOm1vZGlmeQAyMDE5LTA1LTE3VDA4OjQ3OjQ3LTA0OjAw8UrklwAAAABJRU5ErkJggg==`
func GetMainGrid() *ui.Grid {
	grid := ui.NewGrid()
	termWidth, _ := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, int(float64(termWidth) / 3.5))

	imgLogo, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(TEST_LOGO)))
	imgLogo = imaging.Resize(imgLogo, termWidth / 5, termWidth / 10, imaging.Lanczos)

	if err != nil {
		log.Fatalf("Could not load the image " + err.Error())
	}

	imgWidget := widgets.NewImage(imgLogo)
	imgWidget.SetRect(0, 0, 5, 5)
	imgWidget.PaddingLeft = termWidth / 2 - (termWidth / 5) / 2
	imgWidget.Border = false

	p := widgets.NewParagraph()

	p.Text = "noituri@me\n" +
		"OS: Manjaro Linux\n" +
		"CPU: intel i3\n" +
		"Cores: 4\n" +
		"GPU: 920M\n" +
		"Terminal: st\n" +
		"Shell: zsh"

	p.SetRect(0, 0, 25, 5)
	p.PaddingLeft = termWidth / 2 - 5
	p.Border = false

	grid.Set(
		ui.NewRow(1.0 / 2.5,
			ui.NewCol(1.0, imgWidget),
		),
		ui.NewRow(1.0 / 2.5,
			ui.NewCol(1.0, p),
		),
	)

	return grid
}

func StartTui() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	ui.Render(GetMainGrid())

	for {
		select {
		case e := <-ui.PollEvents():
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				ui.Clear()
				ui.Render(GetMainGrid())
			}
		}
	}
}