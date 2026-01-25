package main

import (
	"fmt"
	"os"

	"github.com/grokify/structured-goals/v2mom"
	"github.com/spf13/cobra"
)

var (
	validateStructure string
)

var validateCmd = &cobra.Command{
	Use:   "validate FILE",
	Short: "Validate a V2MOM JSON file",
	Long: `Validate a V2MOM JSON file against the schema and structural rules.

Structure modes:
  flat    - Traditional V2MOM (measures/obstacles at V2MOM level only)
  nested  - OKR-aligned (measures under Methods, global obstacles allowed)
  hybrid  - Both levels allowed (default)

Examples:
  v2mom validate my-v2mom.json
  v2mom validate my-v2mom.json --structure=nested`,
	Args: cobra.ExactArgs(1),
	RunE: runValidate,
}

func init() {
	validateCmd.Flags().StringVar(&validateStructure, "structure", "", "Structure mode to validate against (flat, nested, hybrid)")
}

func runValidate(cmd *cobra.Command, args []string) error {
	filepath := args[0]

	// Check file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filepath)
	}

	// Read and parse V2MOM
	v, err := v2mom.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}

	// Set up validation options
	opts := v2mom.DefaultValidationOptions()
	if validateStructure != "" {
		opts.Structure = validateStructure
	}

	// Validate
	errs := v.Validate(opts)

	// Report results
	errors := v2mom.Errors(errs)
	warnings := v2mom.Warnings(errs)

	if len(warnings) > 0 {
		fmt.Println("Warnings:")
		for _, w := range warnings {
			fmt.Printf("  - %s\n", w)
		}
		fmt.Println()
	}

	if len(errors) > 0 {
		fmt.Println("Errors:")
		for _, e := range errors {
			fmt.Printf("  - %s\n", e)
		}
		return fmt.Errorf("validation failed with %d error(s)", len(errors))
	}

	// Print success info
	fmt.Printf("Valid V2MOM: %s\n", filepath)
	fmt.Printf("  Structure: %s\n", v.GetStructure())
	fmt.Printf("  Methods: %d\n", len(v.Methods))
	fmt.Printf("  Total Measures: %d\n", len(v.AllMeasures()))
	fmt.Printf("  Total Obstacles: %d\n", len(v.AllObstacles()))

	if v.Metadata != nil && v.Metadata.Name != "" {
		fmt.Printf("  Name: %s\n", v.Metadata.Name)
	}

	return nil
}
