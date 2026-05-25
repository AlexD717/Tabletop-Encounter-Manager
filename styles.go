package main

import "charm.land/lipgloss/v2"

var (
	StylePrompt      = lipgloss.NewStyle().Foreground(lipgloss.Color("#fcff96")).Bold(true)
	StyleError       = lipgloss.NewStyle().Foreground(lipgloss.Color("#d14c4c")).Bold(true)
	StyleEntity      = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ff3cc")).Italic(true)
	StyleHealth      = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff808d"))
	StyleInitiative  = lipgloss.NewStyle().Foreground(lipgloss.Color("#8adaf3"))
	StyleCommand     = lipgloss.NewStyle().Foreground(lipgloss.Color("#fcffa3")).Width(10)
	StyleArguments   = lipgloss.NewStyle().Foreground(lipgloss.Color("#fafafa"))
	StyleDescription = lipgloss.NewStyle().Foreground(lipgloss.Alpha(lipgloss.Color("#c8c8c8"), 0.8))
)
