package main

import "charm.land/lipgloss/v2"

var (
	StylePrompt     = lipgloss.NewStyle().Foreground(lipgloss.Color("#bc95e9")).Bold(true)
	StyleError      = lipgloss.NewStyle().Foreground(lipgloss.Color("#d14c4c")).Bold(true)
	StyleEntity     = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ff3cc")).Italic(true)
	StyleHealth     = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff808d"))
	StyleInitiative = lipgloss.NewStyle().Foreground(lipgloss.Color("#8adaf3"))
)
