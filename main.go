package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/eyslce/clashx/contants"
	"github.com/eyslce/clashx/data"
	"github.com/eyslce/clashx/theme"
)

func main() {
	a := app.New()
	a.SetIcon(data.Logo)
	a.Settings().SetTheme(new(theme.DarkTheme))

	w := a.NewWindow(contants.Title)
	w.Resize(fyne.NewSize(800, 600))

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(contants.Title,
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	w.Show()

	a.Run()
}
