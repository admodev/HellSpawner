package hsconsole

import (
	g "github.com/ianling/giu"
	"github.com/ianling/imgui-go"

	"github.com/OpenDiablo2/HellSpawner/hscommon/hsstate"
	"github.com/OpenDiablo2/HellSpawner/hswindow/hstoolwindow"
)

type Console struct {
	*hstoolwindow.ToolWindow
	outputText string
	fontFixed  imgui.Font
}

func (c *Console) Write(p []byte) (n int, err error) {
	c.outputText = string(p) + c.outputText

	return len(p), nil
}

func Create(fontFixed imgui.Font, x, y float32) *Console {
	result := &Console{
		fontFixed:  fontFixed,
		ToolWindow: hstoolwindow.New("Console", hsstate.ToolWindowTypeConsole, x, y),
	}

	return result
}

func (c *Console) Build() {
	c.IsOpen(&c.Visible).
		Size(600, 200).
		Layout(g.Layout{
			g.Custom(func() {
				g.PushFont(c.fontFixed)
			}),
			g.InputTextMultiline("", &c.outputText).
				Size(-1, -1).
				Flags(g.InputTextFlagsReadOnly | g.InputTextFlagsNoUndoRedo),
			g.Custom(func() {
				g.PopFont()
			}),
		})
}
