package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)
func main() {
  style := lipgloss.NewStyle().Bold(true).SetString("Hello,").Background(lipgloss.Color("#7D56F4"))
  
  fmt.Println(style.Render("kitty.")) // Hello, kitty.
  fmt.Println(style.Render("puppy.")) // Hello, puppy.
  fmt.Println("Hello, puppy.")
  fmt.Println(style.Width(24).Align(lipgloss.Center).Render("puppy.")) // Hello, puppy.
}
