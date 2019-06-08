package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

var click = -1

func GetMainGrid(data OsInfo) *ui.Grid {
	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	logo := widgets.NewParagraph()
	logo.Text = GetASCII(data.OS)
	logo.SetRect(0, 0, 25, 5)
	logo.PaddingLeft = termWidth / 2 - 20
	logo.Border = false

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
	p.PaddingLeft = termWidth / 2 - 10
	p.Border = false

	if click == -1 {
		grid.Set(
			ui.NewRow(1.0 / 2,
				ui.NewCol(1.0, logo),
			),
			ui.NewRow(1.0/2,
				ui.NewCol(1.0, p),
			),
		)
	} else {

		clicker := widgets.NewParagraph()
		clicker.Text = fmt.Sprintf("%d CLICKED!!!", click)

		size := 2.5

		if click % 2 == 0 {
			clicker.TextStyle.Fg = ui.ColorMagenta
			size = 2
		} else {
			clicker.TextStyle.Fg = ui.ColorCyan
			size = 2.1
		}

		clicker.SetRect(0, 0, 25, 5)
		clicker.PaddingLeft = termWidth / 2 - 5
		clicker.PaddingTop = 1
		clicker.Border = false

		grid.Set(
			ui.NewRow(1.0 / size,
				ui.NewCol(1.0, logo),
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