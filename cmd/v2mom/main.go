// Command v2mom is a CLI tool for managing V2MOM strategic planning documents.
// It validates V2MOM JSON files and generates Marp presentation slides.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version information (set by build flags)
var (
	version = "0.1.0"
	commit  = "dev"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "v2mom",
	Short: "V2MOM strategic planning tool",
	Long: `v2mom is a CLI tool for managing V2MOM (Vision, Values, Methods, Obstacles, Measures)
strategic planning documents.

It supports:
- JSON validation against the V2MOM schema
- Marp markdown slide generation
- Both traditional flat V2MOM and OKR-aligned nested structures
- Multiple terminology modes (V2MOM, OKR, hybrid)

Examples:
  v2mom validate my-v2mom.json
  v2mom generate marp my-v2mom.json -o slides.md
  v2mom init --name "FY2025 Product Strategy" -o product-v2mom.json`,
	Version: fmt.Sprintf("%s (commit: %s)", version, commit),
}

func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(initCmd)
}
