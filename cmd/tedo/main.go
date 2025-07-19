package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/WasathTheekshana/tedo/internal/ui"
	"github.com/WasathTheekshana/tedo/internal/version"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Command line flags
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help information")
	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Println(version.GetFullVersion())
		os.Exit(0)
	}

	// Handle help flag
	if *showHelp {
		fmt.Printf("%s\n\n", version.GetFullVersion())
		fmt.Println("A beautiful, interactive command-line todo application.")
		fmt.Println("\nUsage:")
		fmt.Println("  tedo            Start the application")
		fmt.Println("  tedo -version   Show version information")
		fmt.Println("  tedo -help      Show this help message")
		fmt.Println("\nFor more information, visit: https://github.com/WasathTheekshana/Tedo")
		os.Exit(0)
	}

	// Create the application model
	model := ui.NewModel()

	// Create the Bubble Tea program
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
