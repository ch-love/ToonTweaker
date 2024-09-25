package ui

import (
    "fyne.io/fyne/v2/widget"
    "time"
)

// AddButtonHoverAnimation adds hover effects for buttons
func AddButtonHoverAnimation(button *widget.Button) {
    button.OnMouseIn = func() {
        // Example: Grow button slightly on hover
        fyne.CurrentApp().Driver().CanvasForObject(button).Animate(func(float64) {
            button.Resize(button.Size().Add(fyne.NewSize(5, 5)))
        }, 200*time.Millisecond)
    }

    button.OnMouseOut = func() {
        // Return to original size on mouse out
        fyne.CurrentApp().Driver().CanvasForObject(button).Animate(func(float64) {
            button.Resize(button.Size().Subtract(fyne.NewSize(5, 5)))
        }, 200*time.Millisecond)
    }
}
