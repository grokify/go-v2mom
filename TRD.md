# Technical Requirements Document: go-v2mom

## V2MOM and OKR: Framework Relationship

### Industry Adoption Patterns

Research indicates that V2MOM and OKR are typically adopted as **alternatives rather than together**:

- **Salesforce** uses V2MOM (created by Marc Benioff in 1999)
- **Google** uses OKRs (adopted via John Doerr from Intel)
- **Fermyon** explicitly combines both frameworks, noting they "overlap in clear ways"

The general pattern is:

| Approach | When to Use |
|----------|-------------|
| V2MOM only | Organizations valuing top-down alignment, values-driven culture |
| OKR only | Organizations prioritizing collaborative goal-setting, agility |
| Combined | Organizations wanting V2MOM's strategic depth with OKR's measurability |

### Structural Alignment

This library supports using OKR concepts within V2MOM's Methods section:

| V2MOM Term | OKR Equivalent | Description |
|------------|----------------|-------------|
| Method | Objective | What you will accomplish |
| Measure | Key Result | Quantifiable success criteria |
| Obstacle | - | V2MOM-specific (risk tracking) |
| Values | - | V2MOM-specific (guiding principles) |

### Terminology Options

The rendering layer supports multiple terminology modes:

| Mode | Methods → | Measures → | Use Case |
|------|-----------|------------|----------|
| `v2mom` | Methods | Measures | Traditional V2MOM users |
| `okr` | Objectives | Key Results | OKR-familiar audiences |
| `hybrid` | Methods (Objectives) | Measures (Key Results) | Combined approach |

**Sources:**

- [Cascade: V2MOM Overview](https://www.cascade.app/blog/the-v2mom-framework)
- [Salesforce: Create Strategic Alignment](https://www.salesforce.com/blog/how-to-create-alignment-within-your-company/)
- [Betterworks: Goal Setting at Salesforce](https://www.betterworks.com/magazine/goal-setting-salesforce-saas-giant/)

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                         go-v2mom                                │
├─────────────────────────────────────────────────────────────────┤
│  CLI (cmd/v2mom)                                                │
│    ├── validate    - Schema validation                          │
│    ├── generate    - Output generation                          │
│    ├── init        - Template creation                          │
│    └── convert     - Format conversion                          │
├─────────────────────────────────────────────────────────────────┤
│  Core Library                                                   │
│    ├── v2mom/          - V2MOM types and JSON handling          │
│    ├── schema/         - JSON Schema definitions                │
│    └── render/         - Output renderers                       │
│          ├── marp/     - Marp markdown renderer                 │
│          ├── pandoc/   - Pandoc markdown renderer (future)      │
│          └── confluence/ - Confluence format (future)           │
├─────────────────────────────────────────────────────────────────┤
│  Adapters (future)                                              │
│    ├── aha/            - Aha! API integration                   │
│    ├── productboard/   - ProductBoard API integration           │
│    └── jira/           - Jira API integration                   │
└─────────────────────────────────────────────────────────────────┘
```

## Project Structure

```
go-v2mom/
├── cmd/
│   └── v2mom/
│       └── main.go              # CLI entry point
├── v2mom/
│   ├── v2mom.go                 # Core V2MOM types
│   ├── v2mom_test.go            # Unit tests
│   ├── validation.go            # JSON validation
│   └── examples.go              # Example data generators
├── schema/
│   ├── v2mom.schema.json        # JSON Schema definition
│   └── embed.go                 # Embedded schema for validation
├── render/
│   ├── renderer.go              # Renderer interface
│   ├── marp/
│   │   ├── marp.go              # Marp renderer implementation
│   │   ├── marp_test.go         # Marp tests
│   │   └── themes/              # Marp theme templates
│   │       ├── default.go
│   │       ├── corporate.go
│   │       └── minimal.go
│   └── pandoc/
│       └── pandoc.go            # Future: Pandoc renderer
├── examples/
│   ├── product-v2mom.json       # Example product V2MOM
│   ├── engineering-v2mom.json   # Example engineering V2MOM
│   └── output/
│       └── product-v2mom.md     # Generated Marp output
├── go.mod
├── go.sum
├── PRD.md
├── TRD.md
├── ROADMAP.md
└── README.md
```

## JSON Schema Definition

### Schema Version

JSON Schema Draft-07 (`http://json-schema.org/draft-07/schema#`)

### Complete Schema

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "https://github.com/grokify/go-v2mom/schema/v2mom.schema.json",
  "title": "V2MOM",
  "description": "V2MOM strategic planning document with OKR alignment",
  "type": "object",
  "required": ["vision", "values", "methods"],
  "properties": {
    "$schema": {
      "type": "string",
      "description": "JSON Schema reference"
    },
    "metadata": {
      "$ref": "#/definitions/Metadata"
    },
    "vision": {
      "type": "string",
      "description": "The desired end state - what you want to achieve",
      "minLength": 1
    },
    "values": {
      "type": "array",
      "description": "Principles guiding decisions, ordered by priority",
      "items": {
        "$ref": "#/definitions/Value"
      },
      "minItems": 1
    },
    "methods": {
      "type": "array",
      "description": "Actions to achieve the vision (aligned with OKR Objectives)",
      "items": {
        "$ref": "#/definitions/Method"
      },
      "minItems": 1
    },
    "obstacles": {
      "type": "array",
      "description": "Global challenges to overcome",
      "items": {
        "$ref": "#/definitions/Obstacle"
      }
    },
    "measures": {
      "type": "array",
      "description": "Global success criteria (for flat V2MOM structure)",
      "items": {
        "$ref": "#/definitions/Measure"
      }
    },
    "projects": {
      "type": "array",
      "description": "Roadmap projects linked to methods",
      "items": {
        "$ref": "#/definitions/Project"
      }
    }
  },
  "definitions": {
    "Metadata": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique identifier for this V2MOM"
        },
        "name": {
          "type": "string",
          "description": "Display name for this V2MOM"
        },
        "author": {
          "type": "string",
          "description": "Author name and title"
        },
        "team": {
          "type": "string",
          "description": "Team or organization"
        },
        "fiscalYear": {
          "type": "string",
          "pattern": "^FY[0-9]{4}$",
          "description": "Fiscal year (e.g., FY2025)"
        },
        "quarter": {
          "type": "string",
          "enum": ["Q1", "Q2", "Q3", "Q4", "H1", "H2", "Annual"],
          "description": "Planning period"
        },
        "version": {
          "type": "string",
          "description": "Document version"
        },
        "status": {
          "type": "string",
          "enum": ["Draft", "In Review", "Approved", "Active", "Completed", "Archived"],
          "description": "Document status"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "parentId": {
          "type": "string",
          "description": "ID of parent V2MOM for cascading alignment"
        },
        "structure": {
          "type": "string",
          "enum": ["flat", "nested", "hybrid"],
          "default": "hybrid",
          "description": "V2MOM organizational structure: flat (traditional), nested (OKR-aligned), or hybrid (both levels)"
        },
        "terminology": {
          "type": "string",
          "enum": ["v2mom", "okr", "hybrid"],
          "default": "v2mom",
          "description": "Display terminology for rendering: v2mom (Methods/Measures), okr (Objectives/Key Results), or hybrid (both)"
        }
      }
    },
    "Value": {
      "type": "object",
      "required": ["name"],
      "properties": {
        "name": {
          "type": "string",
          "description": "Value name"
        },
        "description": {
          "type": "string",
          "description": "Value description"
        },
        "priority": {
          "type": "integer",
          "minimum": 1,
          "description": "Priority rank (1 = highest)"
        }
      }
    },
    "Method": {
      "type": "object",
      "required": ["name"],
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique identifier"
        },
        "name": {
          "type": "string",
          "description": "Method/Objective name"
        },
        "description": {
          "type": "string",
          "description": "Detailed description"
        },
        "priority": {
          "type": "string",
          "enum": ["P0", "P1", "P2", "P3"],
          "description": "Priority level"
        },
        "status": {
          "type": "string",
          "enum": ["Not Started", "Planning", "In Progress", "At Risk", "Completed", "Cancelled"],
          "description": "Current status"
        },
        "owner": {
          "type": "string",
          "description": "Responsible person or team"
        },
        "startDate": {
          "type": "string",
          "format": "date"
        },
        "endDate": {
          "type": "string",
          "format": "date"
        },
        "measures": {
          "type": "array",
          "description": "Key Results for this Method (OKR alignment)",
          "items": {
            "$ref": "#/definitions/Measure"
          }
        },
        "obstacles": {
          "type": "array",
          "description": "Obstacles specific to this Method",
          "items": {
            "$ref": "#/definitions/Obstacle"
          }
        },
        "projects": {
          "type": "array",
          "description": "Project IDs linked to this Method",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Obstacle": {
      "type": "object",
      "required": ["name"],
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "description": "Obstacle name"
        },
        "description": {
          "type": "string",
          "description": "Detailed description"
        },
        "severity": {
          "type": "string",
          "enum": ["Low", "Medium", "High", "Critical"],
          "description": "Impact severity"
        },
        "likelihood": {
          "type": "string",
          "enum": ["Low", "Medium", "High"],
          "description": "Probability of occurrence"
        },
        "mitigation": {
          "type": "string",
          "description": "Mitigation strategy"
        },
        "status": {
          "type": "string",
          "enum": ["Identified", "Mitigating", "Resolved", "Accepted"],
          "description": "Current status"
        }
      }
    },
    "Measure": {
      "type": "object",
      "required": ["name"],
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "description": "Measure/Key Result name"
        },
        "description": {
          "type": "string",
          "description": "Detailed description"
        },
        "baseline": {
          "type": "string",
          "description": "Starting value"
        },
        "target": {
          "type": "string",
          "description": "Target value"
        },
        "current": {
          "type": "string",
          "description": "Current value"
        },
        "unit": {
          "type": "string",
          "description": "Unit of measurement"
        },
        "progress": {
          "type": "number",
          "minimum": 0,
          "maximum": 1,
          "description": "Progress (0.0-1.0, OKR scoring)"
        },
        "timeline": {
          "type": "string",
          "description": "Target timeline"
        },
        "status": {
          "type": "string",
          "enum": ["On Track", "At Risk", "Behind", "Achieved", "Missed"],
          "description": "Current status"
        }
      }
    },
    "Project": {
      "type": "object",
      "required": ["id", "name"],
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique project identifier"
        },
        "name": {
          "type": "string",
          "description": "Project name"
        },
        "description": {
          "type": "string"
        },
        "category": {
          "type": "string",
          "description": "Project category"
        },
        "methodId": {
          "type": "string",
          "description": "Linked Method ID"
        },
        "priority": {
          "type": "string",
          "enum": ["P0", "P1", "P2", "P3"]
        },
        "status": {
          "type": "string",
          "enum": ["Proposed", "Approved", "In Progress", "Completed", "Cancelled"]
        },
        "startDate": {
          "type": "string",
          "format": "date"
        },
        "endDate": {
          "type": "string",
          "format": "date"
        },
        "quarter": {
          "type": "string"
        },
        "dependencies": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "IDs of dependent projects"
        },
        "externalLinks": {
          "type": "object",
          "properties": {
            "jira": {"type": "string"},
            "aha": {"type": "string"},
            "productboard": {"type": "string"},
            "confluence": {"type": "string"}
          },
          "description": "Links to external systems"
        }
      }
    }
  }
}
```

## Dual-Level Structure Support

### Overview

The schema supports Obstacles and Measures at **both** the V2MOM level (global) and the Method level (nested). This provides flexibility for different organizational approaches.

```
V2MOM
├── Vision
├── Values[]
├── Methods[]
│   ├── Method 1
│   │   ├── Measures[] (nested - OKR Key Results)
│   │   └── Obstacles[] (nested - method-specific)
│   └── Method 2
│       ├── Measures[]
│       └── Obstacles[]
├── Obstacles[] (global - cross-cutting challenges)
└── Measures[] (global - north star metrics)
```

### When to Use Each Level

| Level | Obstacles | Measures |
|-------|-----------|----------|
| **Global (V2MOM)** | Cross-cutting challenges affecting multiple methods: "Hiring freeze," "Budget constraints," "Market uncertainty" | North star metrics spanning all methods: "NPS > 50," "Revenue $10M ARR," "Customer retention 95%" |
| **Nested (Method)** | Method-specific blockers: "Legacy API limitation," "Vendor dependency," "Skill gap in team" | OKR-style Key Results: "Ship feature X by Q2," "Reduce latency to <100ms," "Onboard 50 customers" |

### Structure Modes

The `metadata.structure` field controls validation behavior:

| Mode | Global Level | Method Level | Validation Behavior |
|------|--------------|--------------|---------------------|
| `flat` | Required | Forbidden | Traditional V2MOM - measures/obstacles at V2MOM level only |
| `nested` | Optional* | Required | OKR-aligned - measures under Methods, global obstacles allowed |
| `hybrid` | Allowed | Allowed | Maximum flexibility (default) |

*In `nested` mode, global obstacles are allowed (for cross-cutting risks) but global measures are forbidden.

### Validation Rules

```go
type ValidationOptions struct {
    // Structure enforcement mode
    Structure string // "flat", "nested", "hybrid", or "" (no enforcement)

    // Additional validation rules
    RequireMethodMeasures  bool // Each method must have ≥1 measure (recommended for nested/hybrid)
    MaxMeasuresPerMethod   int  // OKR best practice: 3-5 (0 = unlimited)
    RequireGlobalObstacles bool // At least one top-level obstacle required
    WarnEmptyMethods       bool // Warn if methods have no measures (doesn't fail)
}
```

### Validation Implementation

```go
func (v *V2MOM) Validate(opts *ValidationOptions) []ValidationError {
    var errs []ValidationError

    structure := opts.Structure
    if structure == "" && v.Metadata != nil && v.Metadata.Structure != "" {
        structure = v.Metadata.Structure
    }

    switch structure {
    case "flat":
        // Measures must be at V2MOM level only
        for i, m := range v.Methods {
            if len(m.Measures) > 0 {
                errs = append(errs, ValidationError{
                    Path:    fmt.Sprintf("methods[%d].measures", i),
                    Message: "flat mode: measures must be at V2MOM level, not under methods",
                })
            }
            if len(m.Obstacles) > 0 {
                errs = append(errs, ValidationError{
                    Path:    fmt.Sprintf("methods[%d].obstacles", i),
                    Message: "flat mode: obstacles must be at V2MOM level, not under methods",
                })
            }
        }
        if len(v.Measures) == 0 {
            errs = append(errs, ValidationError{
                Path:    "measures",
                Message: "flat mode: V2MOM-level measures required",
            })
        }

    case "nested":
        // Measures must be under methods (global obstacles allowed)
        if len(v.Measures) > 0 {
            errs = append(errs, ValidationError{
                Path:    "measures",
                Message: "nested mode: measures must be under methods, not at V2MOM level",
            })
        }
        for i, m := range v.Methods {
            if len(m.Measures) == 0 {
                errs = append(errs, ValidationError{
                    Path:    fmt.Sprintf("methods[%d].measures", i),
                    Message: fmt.Sprintf("nested mode: method %q must have at least one measure", m.Name),
                })
            }
        }

    case "hybrid", "":
        // No structural enforcement - validate what's present
        if opts.RequireMethodMeasures {
            for i, m := range v.Methods {
                if len(m.Measures) == 0 {
                    errs = append(errs, ValidationError{
                        Path:    fmt.Sprintf("methods[%d].measures", i),
                        Message: fmt.Sprintf("method %q has no measures", m.Name),
                    })
                }
            }
        }
    }

    // OKR best practice: 3-5 key results per objective
    if opts.MaxMeasuresPerMethod > 0 {
        for i, m := range v.Methods {
            if len(m.Measures) > opts.MaxMeasuresPerMethod {
                errs = append(errs, ValidationError{
                    Path:    fmt.Sprintf("methods[%d].measures", i),
                    Message: fmt.Sprintf("method %q has %d measures (max %d recommended)",
                        m.Name, len(m.Measures), opts.MaxMeasuresPerMethod),
                    Severity: "warning",
                })
            }
        }
    }

    return errs
}
```

### Helper Methods

```go
// AllMeasures returns all measures (global + nested), flattened
func (v *V2MOM) AllMeasures() []Measure {
    all := make([]Measure, 0, len(v.Measures))
    all = append(all, v.Measures...)
    for _, m := range v.Methods {
        all = append(all, m.Measures...)
    }
    return all
}

// AllObstacles returns all obstacles (global + nested), flattened
func (v *V2MOM) AllObstacles() []Obstacle {
    all := make([]Obstacle, 0, len(v.Obstacles))
    all = append(all, v.Obstacles...)
    for _, m := range v.Methods {
        all = append(all, m.Obstacles...)
    }
    return all
}

// InferStructure detects the structure based on data placement
func (v *V2MOM) InferStructure() string {
    hasGlobalMeasures := len(v.Measures) > 0
    hasNestedMeasures := false
    for _, m := range v.Methods {
        if len(m.Measures) > 0 {
            hasNestedMeasures = true
            break
        }
    }

    switch {
    case hasGlobalMeasures && !hasNestedMeasures:
        return "flat"
    case !hasGlobalMeasures && hasNestedMeasures:
        return "nested"
    case hasGlobalMeasures && hasNestedMeasures:
        return "hybrid"
    default:
        return "empty" // No measures anywhere
    }
}

// HasNestedStructure returns true if any method has nested measures or obstacles
func (v *V2MOM) HasNestedStructure() bool {
    for _, m := range v.Methods {
        if len(m.Measures) > 0 || len(m.Obstacles) > 0 {
            return true
        }
    }
    return false
}
```

## Terminology and Rendering

### Terminology Modes

The `metadata.terminology` field controls display labels in rendered output:

| Mode | Methods | Measures | Obstacles | Use Case |
|------|---------|----------|-----------|----------|
| `v2mom` | "Methods" | "Measures" | "Obstacles" | Salesforce-aligned, traditional V2MOM |
| `okr` | "Objectives" | "Key Results" | "Risks" | Google-aligned, OKR-familiar audiences |
| `hybrid` | "Methods (Objectives)" | "Measures (Key Results)" | "Obstacles" | Educational, showing both terminologies |

### Terminology Implementation

```go
// Terminology defines display labels for V2MOM components
type Terminology struct {
    Methods        string // "Methods", "Objectives", or "Methods (Objectives)"
    MethodSingular string // "Method", "Objective", or "Method (Objective)"
    Measures       string // "Measures", "Key Results", or "Measures (Key Results)"
    MeasureSingular string
    Obstacles      string // "Obstacles" or "Risks"
    ObstacleSingular string
}

// GetTerminology returns terminology based on mode
func GetTerminology(mode string) Terminology {
    switch mode {
    case "okr":
        return Terminology{
            Methods:          "Objectives",
            MethodSingular:   "Objective",
            Measures:         "Key Results",
            MeasureSingular:  "Key Result",
            Obstacles:        "Risks",
            ObstacleSingular: "Risk",
        }
    case "hybrid":
        return Terminology{
            Methods:          "Methods (Objectives)",
            MethodSingular:   "Method (Objective)",
            Measures:         "Measures (Key Results)",
            MeasureSingular:  "Measure (Key Result)",
            Obstacles:        "Obstacles",
            ObstacleSingular: "Obstacle",
        }
    default: // "v2mom"
        return Terminology{
            Methods:          "Methods",
            MethodSingular:   "Method",
            Measures:         "Measures",
            MeasureSingular:  "Measure",
            Obstacles:        "Obstacles",
            ObstacleSingular: "Obstacle",
        }
    }
}
```

### Rendering Options Update

```go
// Options contains rendering options
type Options struct {
    // Visual options
    Theme           string // Theme name (default, corporate, minimal)
    IncludeProjects bool   // Include roadmap/projects slides
    IncludeStatus   bool   // Include status indicators
    CustomCSS       string // Custom CSS for Marp

    // Terminology options
    Terminology     string // "v2mom", "okr", or "hybrid"

    // Structure handling
    Structure       string // Override structure detection: "flat", "nested", "hybrid"
    FlattenMeasures bool   // Combine global + nested measures into single slide

    // Additional metadata
    Metadata        map[string]string
}
```

### Marp Slide Generation by Structure

| Structure | Slides Generated |
|-----------|------------------|
| `flat` | Single "Measures" slide, single "Obstacles" slide (global only) |
| `nested` | Per-method detail slides with measures/obstacles, optional global obstacles slide |
| `hybrid` | Global measures/obstacles slides + per-method detail slides |

## Go Types

### Core Types

```go
package v2mom

import "time"

// V2MOM represents a complete V2MOM document with optional OKR alignment
type V2MOM struct {
    Schema   string    `json:"$schema,omitempty"`
    Metadata *Metadata `json:"metadata,omitempty"`
    Vision   string    `json:"vision"`
    Values   []Value   `json:"values"`
    Methods  []Method  `json:"methods"`
    // Global obstacles (traditional V2MOM)
    Obstacles []Obstacle `json:"obstacles,omitempty"`
    // Global measures (traditional V2MOM, or use Method.Measures for OKR)
    Measures []Measure  `json:"measures,omitempty"`
    Projects []Project  `json:"projects,omitempty"`
}

// Metadata contains document metadata
type Metadata struct {
    ID          string    `json:"id,omitempty"`
    Name        string    `json:"name,omitempty"`
    Author      string    `json:"author,omitempty"`
    Team        string    `json:"team,omitempty"`
    FiscalYear  string    `json:"fiscalYear,omitempty"`
    Quarter     string    `json:"quarter,omitempty"`
    Version     string    `json:"version,omitempty"`
    Status      string    `json:"status,omitempty"`
    CreatedAt   time.Time `json:"createdAt,omitempty"`
    UpdatedAt   time.Time `json:"updatedAt,omitempty"`
    ParentID    string    `json:"parentId,omitempty"`
    // Structure defines the V2MOM organizational style
    // - "flat": Traditional V2MOM (measures/obstacles at V2MOM level only)
    // - "nested": OKR-aligned (measures under Methods, global obstacles allowed)
    // - "hybrid": Both levels allowed (default)
    Structure   string    `json:"structure,omitempty"`
    // Terminology defines display labels for rendering
    // - "v2mom": Methods/Measures/Obstacles (default)
    // - "okr": Objectives/Key Results/Risks
    // - "hybrid": Methods (Objectives)/Measures (Key Results)/Obstacles
    Terminology string    `json:"terminology,omitempty"`
}

// Value represents a guiding principle
type Value struct {
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    Priority    int    `json:"priority,omitempty"`
}

// Method represents an action/objective (OKR Objective equivalent)
type Method struct {
    ID          string     `json:"id,omitempty"`
    Name        string     `json:"name"`
    Description string     `json:"description,omitempty"`
    Priority    string     `json:"priority,omitempty"`
    Status      string     `json:"status,omitempty"`
    Owner       string     `json:"owner,omitempty"`
    StartDate   string     `json:"startDate,omitempty"`
    EndDate     string     `json:"endDate,omitempty"`
    // Nested measures (OKR Key Results)
    Measures  []Measure  `json:"measures,omitempty"`
    // Method-specific obstacles
    Obstacles []Obstacle `json:"obstacles,omitempty"`
    // Linked project IDs
    Projects  []string   `json:"projects,omitempty"`
}

// Obstacle represents a challenge or risk
type Obstacle struct {
    ID          string `json:"id,omitempty"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    Severity    string `json:"severity,omitempty"`
    Likelihood  string `json:"likelihood,omitempty"`
    Mitigation  string `json:"mitigation,omitempty"`
    Status      string `json:"status,omitempty"`
}

// Measure represents a success metric (OKR Key Result equivalent)
type Measure struct {
    ID          string  `json:"id,omitempty"`
    Name        string  `json:"name"`
    Description string  `json:"description,omitempty"`
    Baseline    string  `json:"baseline,omitempty"`
    Target      string  `json:"target,omitempty"`
    Current     string  `json:"current,omitempty"`
    Unit        string  `json:"unit,omitempty"`
    Progress    float64 `json:"progress,omitempty"`
    Timeline    string  `json:"timeline,omitempty"`
    Status      string  `json:"status,omitempty"`
}

// Project represents a roadmap project
type Project struct {
    ID            string            `json:"id"`
    Name          string            `json:"name"`
    Description   string            `json:"description,omitempty"`
    Category      string            `json:"category,omitempty"`
    MethodID      string            `json:"methodId,omitempty"`
    Priority      string            `json:"priority,omitempty"`
    Status        string            `json:"status,omitempty"`
    StartDate     string            `json:"startDate,omitempty"`
    EndDate       string            `json:"endDate,omitempty"`
    Quarter       string            `json:"quarter,omitempty"`
    Dependencies  []string          `json:"dependencies,omitempty"`
    ExternalLinks map[string]string `json:"externalLinks,omitempty"`
}
```

### Renderer Interface

```go
package render

import "github.com/grokify/go-v2mom/v2mom"

// Renderer defines the interface for output format renderers
type Renderer interface {
    // Render converts a V2MOM to the target format
    Render(v *v2mom.V2MOM, opts *Options) ([]byte, error)
    // Format returns the output format name
    Format() string
    // FileExtension returns the file extension for this format
    FileExtension() string
}

// Options contains rendering options
type Options struct {
    Theme           string            // Theme name (default, corporate, minimal)
    IncludeProjects bool              // Include roadmap/projects slides
    IncludeStatus   bool              // Include status indicators
    CustomCSS       string            // Custom CSS for Marp
    Metadata        map[string]string // Additional metadata
}
```

## Marp Output Specification

### Front Matter

```yaml
---
marp: true
theme: default
paginate: true
header: "{{.Metadata.Team}} | {{.Metadata.FiscalYear}} {{.Metadata.Quarter}}"
footer: "V2MOM | {{.Metadata.Name}}"
style: |
  section.title {
    text-align: center;
  }
  section.vision {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
  }
---
```

### Slide Templates

1. **Title Slide** (`<!-- _class: title -->`)
2. **Vision Slide** (`<!-- _class: vision -->`)
3. **Values Slide** - Ordered list with descriptions
4. **Methods Overview** - Table with priority and status
5. **Method Detail** (per method) - Description, measures, obstacles
6. **Obstacles Slide** - Risk matrix or list
7. **Measures Dashboard** - Progress indicators
8. **Roadmap Slide** - Timeline/Gantt view (ASCII or Mermaid)

## CLI Specification

### Commands

```bash
# Initialize a new V2MOM template
v2mom init [--name NAME] [--output FILE] [--structure flat|nested|hybrid]

# Validate a V2MOM JSON file against the schema
v2mom validate FILE [--schema SCHEMA_FILE] [--structure flat|nested|hybrid]

# Generate output in various formats
v2mom generate marp FILE [--output FILE] [--theme THEME] [--terminology v2mom|okr|hybrid]
v2mom generate json FILE [--output FILE]  # Pretty-print/normalize

# Convert between formats (future)
v2mom convert confluence FILE [--output FILE]
v2mom convert pandoc FILE [--format FORMAT] [--output FILE]
```

### Flag Reference

| Flag | Commands | Values | Description |
|------|----------|--------|-------------|
| `--structure` | init, validate | flat, nested, hybrid | Structure mode for validation/template |
| `--terminology` | generate | v2mom, okr, hybrid | Display terminology in output |
| `--theme` | generate marp | default, corporate, minimal | Visual theme for slides |
| `--output`, `-o` | all | filepath | Output file path |
| `--quiet`, `-q` | all | - | Suppress non-error output |
| `--verbose`, `-v` | all | - | Enable debug output |

### Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | File not found |
| 4 | Schema validation error |
| 5 | Render error |

## Dependencies

### Required (Core)

| Package | Purpose | Version |
|---------|---------|---------|
| `encoding/json` | JSON marshal/unmarshal | stdlib |
| `text/template` | Marp template rendering | stdlib |
| `embed` | Embed schema and templates | stdlib |

### Optional (CLI)

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/spf13/cobra` | CLI framework | v1.8+ |
| `github.com/santhosh-tekuri/jsonschema/v5` | JSON Schema validation | v5.3+ |

## Testing Strategy

### Unit Tests

- V2MOM type marshaling/unmarshaling
- Schema validation (valid and invalid documents)
- Marp rendering for each slide type
- Template rendering with various data shapes

### Integration Tests

- CLI command execution
- End-to-end JSON to Marp conversion
- Schema validation with embedded schema

### Test Data

- Minimal valid V2MOM
- Full-featured V2MOM with all fields
- OKR-aligned V2MOM (nested measures)
- Invalid V2MOMs for validation testing

## Performance Requirements

| Operation | Target |
|-----------|--------|
| JSON parsing (10KB) | < 10ms |
| Schema validation | < 50ms |
| Marp generation | < 100ms |
| CLI startup | < 200ms |

## Security Considerations

1. **Input Validation**: All JSON input validated against schema
2. **Template Injection**: Use `text/template` with proper escaping
3. **File Paths**: Validate and sanitize file paths in CLI
4. **No Network**: Core library has no network access

## Future Technical Considerations

### Adapter Interface (v2.0)

```go
// Adapter defines integration with external systems
type Adapter interface {
    // Push sends V2MOM data to external system
    Push(v *v2mom.V2MOM, opts *AdapterOptions) error
    // Pull retrieves data from external system as V2MOM
    Pull(id string, opts *AdapterOptions) (*v2mom.V2MOM, error)
    // Sync bidirectionally syncs V2MOM with external system
    Sync(v *v2mom.V2MOM, opts *AdapterOptions) (*v2mom.V2MOM, error)
}
```

### Supported Adapters (Planned)

- `aha.Adapter` - Aha! Roadmaps API
- `productboard.Adapter` - ProductBoard API
- `jira.Adapter` - Jira/Atlassian API
- `confluence.Adapter` - Confluence Storage Format
- `linear.Adapter` - Linear GraphQL API
