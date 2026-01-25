package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grokify/structured-goals/render"
	"github.com/grokify/structured-goals/render/marp"
	"github.com/grokify/structured-goals/v2mom"
	"github.com/spf13/cobra"
)

var (
	generateOutput      string
	generateTheme       string
	generateTerminology string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate output from V2MOM",
	Long:  `Generate various output formats from a V2MOM JSON file.`,
}

var generateMarpCmd = &cobra.Command{
	Use:   "marp FILE",
	Short: "Generate Marp markdown slides",
	Long: `Generate Marp markdown presentation slides from a V2MOM JSON file.

Themes:
  default   - Clean gradient theme (default)
  corporate - Professional blue theme
  minimal   - Simple grayscale theme

Terminology:
  v2mom  - Use V2MOM terms: Methods, Measures, Obstacles (default)
  okr    - Use OKR terms: Objectives, Key Results, Risks
  hybrid - Show both: Methods (Objectives), Measures (Key Results)

Examples:
  v2mom generate marp my-v2mom.json
  v2mom generate marp my-v2mom.json -o slides.md
  v2mom generate marp my-v2mom.json --theme=corporate --terminology=okr`,
	Args: cobra.ExactArgs(1),
	RunE: runGenerateMarp,
}

func init() {
	generateCmd.AddCommand(generateMarpCmd)

	generateMarpCmd.Flags().StringVarP(&generateOutput, "output", "o", "", "Output file path (default: stdout)")
	generateMarpCmd.Flags().StringVar(&generateTheme, "theme", "default", "Slide theme (default, corporate, minimal)")
	generateMarpCmd.Flags().StringVar(&generateTerminology, "terminology", "", "Display terminology (v2mom, okr, hybrid)")
}

func runGenerateMarp(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Check file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", inputPath)
	}

	// Read and parse V2MOM
	v, err := v2mom.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("reading V2MOM: %w", err)
	}

	// Create renderer and options
	renderer := marp.New()
	opts := render.DefaultOptions()
	opts.Theme = generateTheme
	if generateTerminology != "" {
		opts.Terminology = generateTerminology
	}

	// Render
	output, err := renderer.Render(v, opts)
	if err != nil {
		return fmt.Errorf("rendering Marp: %w", err)
	}

	// Write output
	if generateOutput != "" {
		// Ensure output directory exists
		dir := filepath.Dir(generateOutput)
		if dir != "." && dir != "" {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("creating output directory: %w", err)
			}
		}

		if err := os.WriteFile(generateOutput, output, 0600); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Printf("Generated: %s\n", generateOutput)
	} else {
		// Write to stdout
		fmt.Print(string(output))
	}

	return nil
}
