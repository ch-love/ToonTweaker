package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
    "image/color"
)

// ToontownTheme implements fyne.Theme to create a custom look
type ToontownTheme struct{}

// LoadCustomTheme loads the custom theme for the app
func LoadCustomTheme() fyne.Theme {
    return &ToontownTheme{}
}

// Font overrides the default font with a custom Toontown font
func (t *ToontownTheme) Font(s fyne.TextStyle) fyne.Resource {
    // Add Toontown font path in assets
    return theme.DefaultTheme().Font(s)
}

// PrimaryColor changes the primary color to a vibrant Toontown-like color
func (t *ToontownTheme) PrimaryColor() color.Color {
    return &color.RGBA{R: 255, G: 183, B: 64, A: 255} // Orange/yellow tone
}

// BackgroundColor overrides the background color
func (t *ToontownTheme) BackgroundColor() color.Color {
    return color.RGBA{R: 224, G: 240, B: 255, A: 255} // Soft blue
}

// BorderColor overrides the border color for custom window
func (t *ToontownTheme) BorderColor() color.Color {
    return color.RGBA{R: 50, G: 50, B: 150, A: 255} // Deep blue
}

// ...Other overrides for colors, hover states, etc.
// Load the custom logos for dropdown options
var resourceCorporateClashLogo = theme.NewThemedResource(fyne.NewStaticResource("CorporateClashLogo.png", nil))
var resourceToontownRewrittenLogo = theme.NewThemedResource(fyne.NewStaticResource("ToontownRewrittenLogo.png", nil))
