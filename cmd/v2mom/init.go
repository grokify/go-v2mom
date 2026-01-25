package main

import (
	"fmt"
	"os"
	"time"

	"github.com/grokify/structured-goals/v2mom"
	"github.com/spf13/cobra"
)

var (
	initName      string
	initOutput    string
	initStructure string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new V2MOM template",
	Long: `Create a new V2MOM JSON template file with example content.

Structure modes:
  flat   - Traditional V2MOM with measures/obstacles at top level
  nested - OKR-aligned with measures under methods (default)
  hybrid - Both levels with examples

Examples:
  v2mom init
  v2mom init --name "FY2025 Product Strategy"
  v2mom init --name "Engineering Goals" -o engineering-v2mom.json --structure=nested`,
	RunE: runInit,
}

func init() {
	initCmd.Flags().StringVar(&initName, "name", "My V2MOM", "Name for the V2MOM")
	initCmd.Flags().StringVarP(&initOutput, "output", "o", "v2mom.json", "Output file path")
	initCmd.Flags().StringVar(&initStructure, "structure", "nested", "Structure mode (flat, nested, hybrid)")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Check if file already exists
	if _, err := os.Stat(initOutput); err == nil {
		return fmt.Errorf("file already exists: %s (use -o to specify a different output path)", initOutput)
	}

	// Create template based on structure
	var template *v2mom.V2MOM

	switch initStructure {
	case "flat":
		template = createFlatTemplate(initName)
	case "hybrid":
		template = createHybridTemplate(initName)
	default: // "nested"
		template = createNestedTemplate(initName)
	}

	// Write to file
	if err := template.WriteFile(initOutput); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	fmt.Printf("Created: %s\n", initOutput)
	fmt.Printf("  Structure: %s\n", initStructure)
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Edit the file to add your vision, values, methods, and measures")
	fmt.Println("  2. Run 'v2mom validate " + initOutput + "' to check your V2MOM")
	fmt.Println("  3. Run 'v2mom generate marp " + initOutput + " -o slides.md' to create slides")

	return nil
}

func createNestedTemplate(name string) *v2mom.V2MOM {
	now := time.Now()
	return &v2mom.V2MOM{
		Schema: "../schema/v2mom.schema.json",
		Metadata: &v2mom.Metadata{
			Name:        name,
			Author:      "Your Name",
			Team:        "Your Team",
			FiscalYear:  fmt.Sprintf("FY%d", now.Year()),
			Quarter:     "Q1",
			Version:     "1.0.0",
			Status:      v2mom.StatusDraft,
			Structure:   v2mom.StructureNested,
			Terminology: v2mom.TerminologyV2MOM,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Vision: "Describe your vision here - what do you want to achieve?",
		Values: []v2mom.Value{
			{
				Name:        "Value 1",
				Description: "What's most important to you?",
				Priority:    1,
			},
			{
				Name:        "Value 2",
				Description: "What's the second most important principle?",
				Priority:    2,
			},
		},
		Methods: []v2mom.Method{
			{
				ID:          "method-1",
				Name:        "First Method/Objective",
				Description: "How will you achieve your vision? What's the first major initiative?",
				Priority:    v2mom.PriorityP0,
				Status:      "Not Started",
				Measures: []v2mom.Measure{
					{
						ID:       "m1-kr1",
						Name:     "Key Result 1",
						Target:   "Define your target",
						Status:   "Not Started",
						Progress: 0,
					},
					{
						ID:       "m1-kr2",
						Name:     "Key Result 2",
						Target:   "Define your target",
						Status:   "Not Started",
						Progress: 0,
					},
				},
				Obstacles: []v2mom.Obstacle{
					{
						Name:       "Method-specific obstacle",
						Severity:   "Medium",
						Mitigation: "How will you address this?",
					},
				},
			},
			{
				ID:          "method-2",
				Name:        "Second Method/Objective",
				Description: "What's the second major initiative?",
				Priority:    v2mom.PriorityP1,
				Status:      "Not Started",
				Measures: []v2mom.Measure{
					{
						ID:       "m2-kr1",
						Name:     "Key Result 1",
						Target:   "Define your target",
						Status:   "Not Started",
						Progress: 0,
					},
				},
			},
		},
		Obstacles: []v2mom.Obstacle{
			{
				ID:          "obs-global-1",
				Name:        "Global obstacle",
				Description: "What's preventing success across multiple methods?",
				Severity:    "High",
				Likelihood:  "Medium",
				Mitigation:  "How will you mitigate this risk?",
				Status:      "Identified",
			},
		},
	}
}

func createFlatTemplate(name string) *v2mom.V2MOM {
	now := time.Now()
	return &v2mom.V2MOM{
		Schema: "../schema/v2mom.schema.json",
		Metadata: &v2mom.Metadata{
			Name:        name,
			Author:      "Your Name",
			Team:        "Your Team",
			FiscalYear:  fmt.Sprintf("FY%d", now.Year()),
			Quarter:     "Q1",
			Version:     "1.0.0",
			Status:      v2mom.StatusDraft,
			Structure:   v2mom.StructureFlat,
			Terminology: v2mom.TerminologyV2MOM,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Vision: "Describe your vision here - what do you want to achieve?",
		Values: []v2mom.Value{
			{
				Name:        "Value 1",
				Description: "What's most important to you?",
				Priority:    1,
			},
			{
				Name:        "Value 2",
				Description: "What's the second most important principle?",
				Priority:    2,
			},
		},
		Methods: []v2mom.Method{
			{
				Name:        "First Method",
				Description: "How will you achieve your vision?",
				Priority:    v2mom.PriorityP0,
			},
			{
				Name:        "Second Method",
				Description: "What's the second major action?",
				Priority:    v2mom.PriorityP1,
			},
		},
		Obstacles: []v2mom.Obstacle{
			{
				Name:        "Obstacle 1",
				Description: "What's preventing success?",
				Severity:    "High",
				Mitigation:  "How will you address this?",
			},
		},
		Measures: []v2mom.Measure{
			{
				Name:   "Measure 1",
				Target: "Define your target",
				Status: "Not Started",
			},
			{
				Name:   "Measure 2",
				Target: "Define your target",
				Status: "Not Started",
			},
		},
	}
}

func createHybridTemplate(name string) *v2mom.V2MOM {
	now := time.Now()
	return &v2mom.V2MOM{
		Schema: "../schema/v2mom.schema.json",
		Metadata: &v2mom.Metadata{
			Name:        name,
			Author:      "Your Name",
			Team:        "Your Team",
			FiscalYear:  fmt.Sprintf("FY%d", now.Year()),
			Quarter:     "Q1",
			Version:     "1.0.0",
			Status:      v2mom.StatusDraft,
			Structure:   v2mom.StructureHybrid,
			Terminology: v2mom.TerminologyV2MOM,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		Vision: "Describe your vision here - what do you want to achieve?",
		Values: []v2mom.Value{
			{
				Name:        "Value 1",
				Description: "What's most important to you?",
				Priority:    1,
			},
		},
		Methods: []v2mom.Method{
			{
				ID:          "method-1",
				Name:        "Method with nested measures",
				Description: "This method has its own key results",
				Priority:    v2mom.PriorityP0,
				Status:      "Not Started",
				Measures: []v2mom.Measure{
					{
						Name:   "Method-specific KR",
						Target: "Define target",
					},
				},
			},
			{
				ID:          "method-2",
				Name:        "Method without nested measures",
				Description: "This method uses global measures",
				Priority:    v2mom.PriorityP1,
			},
		},
		Obstacles: []v2mom.Obstacle{
			{
				Name:        "Global obstacle",
				Description: "Affects multiple methods",
				Severity:    "High",
				Mitigation:  "Mitigation strategy",
			},
		},
		Measures: []v2mom.Measure{
			{
				Name:        "North star metric",
				Description: "Global measure spanning all methods",
				Target:      "Define target",
			},
		},
	}
}
