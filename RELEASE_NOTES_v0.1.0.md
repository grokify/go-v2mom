# Release Notes: v0.1.0

**Release Date:** 2026-01-18

This is the initial release of go-v2mom, a Go library and CLI tool for managing V2MOM (Vision, Values, Methods, Obstacles, Measures) strategic planning documents with integrated OKR (Objectives and Key Results) alignment.

## Overview

go-v2mom bridges Salesforce's V2MOM framework with modern OKR practices, providing:

- A canonical JSON schema for V2MOM documents
- Flexible structure modes (flat, nested, hybrid) to support various planning styles
- Professional Marp presentation generation for executive reviews
- A CLI tool for validation and document generation

## Installation

```bash
go install github.com/grokify/go-v2mom/cmd/v2mom@v0.1.0
```

## Features

### Core Library

- **V2MOM Types**: Complete type system including `V2MOM`, `Metadata`, `Value`, `Method`, `Measure`, `Obstacle`, and `Project`
- **Structure Modes**:
  - `flat`: Traditional V2MOM with measures/obstacles at top level only
  - `nested`: OKR-aligned with measures under Methods (Objectives)
  - `hybrid`: Maximum flexibility with both levels allowed
- **Terminology Support**: Display labels for v2mom, okr, or hybrid terminology
- **Validation**: Structure enforcement with path-aware error reporting and configurable options

### JSON Schema

- Draft-07 compliant schema for V2MOM documents
- Embedded in binary via `//go:embed` for zero-dependency validation
- Comprehensive field definitions with enums for status, priority, and quarters

### Marp Renderer

- Professional slide deck generation from V2MOM JSON
- Theme support: default, corporate, minimal
- Slide types: title, vision, values, methods overview, obstacles, measures, progress dashboard
- Terminology-aware labels based on configuration

### CLI Commands

| Command | Description |
|---------|-------------|
| `v2mom init` | Generate a template V2MOM JSON file |
| `v2mom validate` | Validate V2MOM JSON with structure detection |
| `v2mom generate marp` | Generate Marp markdown slides |

### CLI Flags

**init**:

- `--name` - V2MOM name (default: "My V2MOM")
- `--output`, `-o` - Output file path (default: v2mom.json)
- `--structure` - Template structure: flat, nested, hybrid (default: nested)

**validate**:

- `--structure` - Enforce specific structure mode

**generate marp**:

- `--output`, `-o` - Output file path (default: stdout)
- `--theme` - Slide theme: default, corporate, minimal
- `--terminology` - Display terminology: v2mom, okr, hybrid

## Quick Start

```bash
# Create a new V2MOM
v2mom init --name "Q1 2026 Product Strategy" -o strategy.json

# Edit the JSON file with your vision, values, methods, etc.

# Validate the document
v2mom validate strategy.json

# Generate presentation slides
v2mom generate marp strategy.json -o strategy.md

# Convert to PDF/HTML using Marp CLI
marp strategy.md -o strategy.pdf
```

## Library Usage

```go
package main

import (
    "fmt"
    "github.com/grokify/go-v2mom/v2mom"
    "github.com/grokify/go-v2mom/render/marp"
)

func main() {
    // Load V2MOM from file
    doc, err := v2mom.ReadFile("my-v2mom.json")
    if err != nil {
        panic(err)
    }

    // Validate
    errs := doc.Validate(v2mom.DefaultValidationOptions())
    if errs.HasErrors() {
        for _, e := range errs.Errors() {
            fmt.Printf("Error at %s: %s\n", e.Path, e.Message)
        }
    }

    // Generate Marp slides
    renderer := marp.NewRenderer()
    output, err := renderer.Render(doc, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(output))
}
```

## Examples

The `examples/` directory contains sample V2MOM documents:

- `product-v2mom.json` - Full-featured product strategy example
- `minimal-v2mom.json` - Minimum valid V2MOM
- `agentplexus-v2mom.json` - OKR-aligned nested structure example

## Known Limitations

- JSON Schema validation uses manual validation logic; schema.Validate() not yet implemented
- Theme CSS is referenced but not fully customized (uses Marp defaults)
- Test coverage is minimal (tests to be added in v0.2.0)

## Roadmap

See [ROADMAP.md](ROADMAP.md) for the full implementation plan. Upcoming features:

- **v0.2.0**: Comprehensive test suite, JSON Schema validation
- **v0.3.0**: Method detail slides, progress dashboard enhancements
- **v0.4.0**: Additional renderers (Pandoc, Confluence, HTML)

## Contributing

Contributions welcome! Priority areas:

- Test coverage
- Additional output format renderers
- Theme designs
- Documentation improvements

## License

MIT License - see [LICENSE](LICENSE) for details.

## Links

- Repository: https://github.com/grokify/go-v2mom
- Documentation: [README.md](README.md)
- Technical Spec: [TRD.md](TRD.md)
