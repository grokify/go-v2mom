# Structured Goals

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

A Go library and CLI tool for managing V2MOM (Vision, Values, Methods, Obstacles, Measures) and OKR (Objects & Key Results) strategic planning documents with JSON Schema validation and Marp slide generation.

## Overview

V2MOM is a strategic planning framework created by Marc Benioff at Salesforce. This tool provides:

- **JSON Schema** for validating V2MOM documents
- **Marp Markdown Generation** for executive presentation slides
- **OKR Alignment** - supports nesting Measures (Key Results) under Methods (Objectives)
- **Flexible Structure** - supports flat (traditional), nested (OKR-aligned), or hybrid structures
- **Multiple Terminology** - render using V2MOM, OKR, or hybrid terminology

## Installation

```bash
go install github.com/grokify/go-v2mom/cmd/v2mom@latest
```

Or build from source:

```bash
git clone https://github.com/grokify/go-v2mom.git
cd go-v2mom
go build ./cmd/v2mom
```

## Quick Start

```bash
# Create a new V2MOM template
v2mom init --name "FY2025 Product Strategy" -o product-v2mom.json

# Edit product-v2mom.json with your content...

# Validate your V2MOM
v2mom validate product-v2mom.json

# Generate Marp slides
v2mom generate marp product-v2mom.json -o slides.md

# Convert to PDF (requires Marp CLI)
marp slides.md -o slides.pdf
```

## V2MOM Framework

| Component | Question | OKR Equivalent |
|-----------|----------|----------------|
| **Vision** | What do you want to achieve? | Mission |
| **Values** | What's important to you? | - |
| **Methods** | How do you get it? | Objectives |
| **Obstacles** | What's preventing success? | - |
| **Measures** | How do you know you have it? | Key Results |

## Structure Modes

### Flat (Traditional V2MOM)

Measures and obstacles at the V2MOM level only:

```json
{
  "vision": "...",
  "values": [...],
  "methods": [{"name": "Method 1"}, {"name": "Method 2"}],
  "obstacles": [{"name": "Obstacle 1"}],
  "measures": [{"name": "Measure 1"}, {"name": "Measure 2"}]
}
```

### Nested (OKR-Aligned)

Measures nested under methods (like Key Results under Objectives):

```json
{
  "vision": "...",
  "values": [...],
  "methods": [
    {
      "name": "Method 1",
      "measures": [
        {"name": "Key Result 1", "target": "100"},
        {"name": "Key Result 2", "target": "50%"}
      ]
    }
  ],
  "obstacles": [{"name": "Global obstacle"}]
}
```

### Hybrid

Both global and nested measures/obstacles:

```json
{
  "vision": "...",
  "methods": [
    {
      "name": "Method 1",
      "measures": [{"name": "Method-specific KR"}],
      "obstacles": [{"name": "Method-specific blocker"}]
    }
  ],
  "obstacles": [{"name": "Global obstacle"}],
  "measures": [{"name": "North star metric"}]
}
```

## CLI Commands

### Initialize

Create a new V2MOM template:

```bash
v2mom init [flags]

Flags:
  --name string       Name for the V2MOM (default "My V2MOM")
  -o, --output string Output file path (default "v2mom.json")
  --structure string  Structure mode: flat, nested, hybrid (default "nested")
```

### Validate

Validate a V2MOM JSON file:

```bash
v2mom validate FILE [flags]

Flags:
  --structure string  Structure mode to validate against (flat, nested, hybrid)
```

### Generate Marp

Generate Marp markdown slides:

```bash
v2mom generate marp FILE [flags]

Flags:
  -o, --output string      Output file path (default: stdout)
  --theme string           Slide theme: default, corporate, minimal (default "default")
  --terminology string     Display terminology: v2mom, okr, hybrid
```

## Library Usage

```go
package main

import (
    "fmt"
    "github.com/grokify/go-v2mom/v2mom"
    "github.com/grokify/go-v2mom/render"
    "github.com/grokify/go-v2mom/render/marp"
)

func main() {
    // Read V2MOM from file
    v, err := v2mom.ReadFile("my-v2mom.json")
    if err != nil {
        panic(err)
    }

    // Validate
    errs := v.Validate(v2mom.DefaultValidationOptions())
    if v2mom.HasErrors(errs) {
        for _, e := range errs {
            fmt.Println(e)
        }
        return
    }

    // Generate Marp slides
    renderer := marp.New()
    opts := render.DefaultOptions()
    opts.Theme = "corporate"
    opts.Terminology = v2mom.TerminologyOKR

    output, err := renderer.Render(v, opts)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(output))
}
```

## JSON Schema

The JSON Schema is available at `schema/v2mom.schema.json` and can be referenced in your V2MOM files:

```json
{
  "$schema": "https://github.com/grokify/go-v2mom/schema/v2mom.schema.json",
  "vision": "..."
}
```

## Terminology Options

| Mode | Methods | Measures | Obstacles |
|------|---------|----------|-----------|
| `v2mom` | Methods | Measures | Obstacles |
| `okr` | Objectives | Key Results | Risks |
| `hybrid` | Methods (Objectives) | Measures (Key Results) | Obstacles |

## Themes

### Default

Clean gradient theme with purple/blue accents.

### Corporate

Professional blue theme suitable for business presentations.

### Minimal

Simple grayscale theme for high contrast readability.

## Roadmap

See [ROADMAP.md](ROADMAP.md) for planned features:

- Phase 2: Enhanced Marp features (method detail slides, roadmap visualization)
- Phase 3: Additional renderers (Pandoc, Confluence, HTML)
- Phase 4: External integrations (Aha!, ProductBoard, Jira)

## Related Documents

- [PRD.md](PRD.md) - Product Requirements Document
- [TRD.md](TRD.md) - Technical Requirements Document
- [ROADMAP.md](ROADMAP.md) - Implementation Roadmap

## References

- [Salesforce V2MOM Blog](https://www.salesforce.com/blog/how-to-create-alignment-within-your-company/)
- [SalesforceLabs V2MOM](https://github.com/SalesforceLabs/V2MOM)
- [SalesforceLabs MyV2MOM](https://github.com/SalesforceLabs/MyV2MOM)
- [Marp Documentation](https://marp.app/)

## License

MIT License - see LICENSE file for details.

 [build-status-svg]: https://github.com/grokify/structured-goals/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/structured-goals/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/structured-goals/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/structured-goals/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/structured-goals
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/structured-goals
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/structured-goals
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/structured-goals
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fstructured-goals
 [loc-svg]: https://tokei.rs/b1/github/grokify/structured-goals
 [repo-url]: https://github.com/grokify/structured-goals
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/structured-goals/blob/master/LICENSE
