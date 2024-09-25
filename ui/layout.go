package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

// MakeWindow creates the main window with custom border and tabs
func MakeWindow(a fyne.App) fyne.Window {
    w := a.NewWindow("Toontown Mod Loader")

    // Tabs for different sections
    tabs := container.NewAppTabs(
        container.NewTabItem("Inject", MakeInjectPage()),  // Updated Inject Tab
        container.NewTabItem("Mods", MakeModsPage()),
        container.NewTabItem("Content Packs", widget.NewLabel("Content Packs")),
        container.NewTabItem("Sound Packs", widget.NewLabel("Sound Packs")),
    )

    // Set window content with tabs
    w.SetContent(container.NewBorder(nil, nil, nil, nil, tabs))
    w.Resize(fyne.NewSize(800, 600))

    return w
}
