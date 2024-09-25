package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/theme"
    "os"
    "runtime"
    "path/filepath"
)

// MakeInjectPage creates the "Inject" tab content
func MakeInjectPage() fyne.CanvasObject {
    injectLabel := widget.NewLabel("Select Toontown Client for Mod Injection")

    // Create dropdown with two options and logos
    toontownClients := []string{"Toontown: Corporate Clash", "Toontown Rewritten"}
    clientDropDown := widget.NewSelect(toontownClients, func(selected string) {
        autoSelectPath(selected)
    })

    // Add logos to the dropdown options
    clashLogo := canvas.NewImageFromResource(theme.NewThemedResource(resourceCorporateClashLogo, nil))
    rewrittenLogo := canvas.NewImageFromResource(theme.NewThemedResource(resourceToontownRewrittenLogo, nil))

    // Combine logos with dropdown options (in this example, simplified)
    clientDropDown.PlaceHolder = "Select Client"
    logoContainer := container.NewVBox(clashLogo, rewrittenLogo)

    // Create layout with dropdown and logo
    layout := container.NewVBox(
        injectLabel,
        logoContainer,
        clientDropDown,
    )

    return layout
}

// autoSelectPath auto-detects the installation path based on the OS and client choice
func autoSelectPath(selectedClient string) {
    var installPath string

    switch selectedClient {
    case "Toontown: Corporate Clash":
        if runtime.GOOS == "windows" {
            // Find path for Windows
            currentUser := os.Getenv("USERNAME")
            installPath = filepath.Join("C:", "Users", currentUser, "AppData", "Local", "Corporate Clash")
        } else if runtime.GOOS == "darwin" {
            // Find path for macOS
            installPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Corporate Clash")
        }
        // For Linux, Corporate Clash does not exist (can add specific cases if needed)
    case "Toontown Rewritten":
        if runtime.GOOS == "windows" {
            installPath = "C:\\Program Files (x86)\\Toontown Rewritten"
        } else if runtime.GOOS == "linux" {
            installPath = filepath.Join(os.Getenv("HOME"), ".local", "share", "toontown-rewritten")
        } else if runtime.GOOS == "darwin" {
            installPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Toontown Rewritten")
        }
    }

    // Display detected path (can be used to populate a form or other UI element)
    println("Auto-detected install path: ", installPath)
}
