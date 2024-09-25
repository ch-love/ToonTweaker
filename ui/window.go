package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

// MakeWindow creates the main window for the application
func MakeWindow(a fyne.App) fyne.Window {
    w := a.NewWindow("Toontown Mod Loader")

    // Toontown-inspired layout (can add background or images later)
    menu := MakeMenu()

    // Placeholder for mod list or options panel
    content := widget.NewLabel("Welcome to the Toontown Mod Loader!")

    // Main layout
    layout := container.NewBorder(menu, nil, nil, nil, content)
    w.SetContent(layout)
    
    return w
}
