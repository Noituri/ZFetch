package main

import (
	"encoding/base64"
	"fmt"
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

var click = -1

func GetMainGrid(data OsInfo) *ui.Grid {
	grid := ui.NewGrid()
	termWidth, _ := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, int(float64(termWidth) / 3.5))

	imgLogo, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(TEST_LOGO)))

	if click == -1 {
		imgLogo = imaging.Resize(imgLogo, termWidth / 5, termWidth / 10, imaging.Lanczos)
	} else if click % 2 == 0 {
		imgLogo = imaging.Resize(imgLogo, termWidth / 5, termWidth / 10, imaging.Lanczos)
	} else {
		imgLogo = imaging.Resize(imgLogo, termWidth / 4, termWidth / 9, imaging.Lanczos)
	}

	if err != nil {
		log.Fatalf("Could not load the image " + err.Error())
	}

	imgWidget := widgets.NewImage(imgLogo)
	imgWidget.SetRect(0, 0, 5, 5)
	imgWidget.PaddingLeft = termWidth / 2 - (termWidth / 5) / 2
	imgWidget.Border = false

	p := widgets.NewParagraph()

	p.Text = fmt.Sprintf("%s@%s\n" +
		"OS: %s\n" +
		"Kernel: %s\n" +
		"Uptime: %s\n" +
		"CPU: %s\n" +
		"Cores: %s\n" +
		"GPU: %s\n" +
		"Terminal: %s\n" +
		"Shell: %s\n" +
		"RAM: %s/%s\n",
		data.Username,
		data.Hostname,
		data.OS,
		data.Kernel,
		data.Uptime,
		data.CPU,
		data.Cores,
		data.GPU,
		data.Terminal,
		data.Shell,
		data.UsedRAM,
		data.MaxRam,
	)

	p.SetRect(0, 0, 25, 5)
	p.PaddingLeft = termWidth / 2 - 5
	p.Border = false

	if click == -1 {
		grid.Set(
			ui.NewRow(1.0 / 2.5,
				ui.NewCol(1.0, imgWidget),
			),
			ui.NewRow(1.0 / 2.5,
				ui.NewCol(1.0, p),
			),
		)
	} else {

		clicker := widgets.NewParagraph()
		clicker.Text = fmt.Sprintf("%d CLICKED!!!", click)

		size := 2.5

		if click % 2 == 0 {
			clicker.TextStyle.Fg = ui.ColorMagenta
			size = 2.5
		} else {
			clicker.TextStyle.Fg = ui.ColorCyan
			size = 3
		}

		clicker.SetRect(0, 0, 25, 5)
		clicker.PaddingLeft = termWidth / 2 - 5
		clicker.PaddingTop = 1
		clicker.Border = false

		grid.Set(
			ui.NewRow(1.0 / size,
				ui.NewCol(1.0, imgWidget),
			),
			ui.NewRow(1.0 / 2.5,
				ui.NewCol(1.0, p),
			),
			ui.NewRow(1.0 / 2.5,
				ui.NewCol(1.0, clicker),
			),
		)
	}

	return grid
}

func StartTui() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	data := OsInfo{}
	data.GetInfo()

	ui.Render(GetMainGrid(data))

	for {
		select {
		case e := <-ui.PollEvents():
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<MouseRelease>":
				click++
				ui.Clear()
				ui.Render(GetMainGrid(data))
			case "<Resize>":
				ui.Clear()
				ui.Render(GetMainGrid(data))
			}
		}
	}
}