# Product Requirements Document: go-v2mom

## Overview

go-v2mom is a Go library and CLI tool for managing V2MOM (Vision, Values, Methods, Obstacles, Measures) strategic planning documents using a JSON Intermediate Representation (IR). The tool enables deterministic generation of executive presentation slides (Marp) and other output formats from a single source of truth.

## Problem Statement

Organizations using V2MOM for strategic alignment face several challenges:

1. **Format Fragmentation**: V2MOMs exist in various formats (Google Docs, Confluence, PowerPoint) making them difficult to maintain and synchronize
2. **Manual Slide Creation**: Converting V2MOMs to executive presentations requires manual effort and introduces inconsistencies
3. **No Structured Data**: V2MOMs stored as prose documents cannot be programmatically processed, validated, or integrated with other tools
4. **OKR Alignment Gap**: Organizations using both V2MOM and OKR lack a unified structure that bridges both frameworks

## Goals

1. Establish a canonical JSON schema for V2MOM documents
2. Provide automatic generation of Marp markdown slides from V2MOM data
3. Support OKR alignment by nesting Measures (Key Results) under Methods (Objectives)
4. Enable future integrations with product management tools (Aha, ProductBoard, Jira)

## Target Users

- **Executives**: Review and present V2MOM slides
- **Product Managers**: Author and maintain V2MOM documents
- **Engineering Leaders**: Integrate V2MOM into planning workflows
- **DevOps/Platform Teams**: Automate V2MOM artifact generation in CI/CD

## V2MOM Framework

### Traditional V2MOM Components

| Component | Description | Question Answered |
|-----------|-------------|-------------------|
| **Vision** | The desired end state | What do you want to achieve? |
| **Values** | Principles guiding decisions | What's important to you? |
| **Methods** | Actions to achieve the vision | How do you get it? |
| **Obstacles** | Challenges to overcome | What's preventing success? |
| **Measures** | Success criteria | How do you know you have it? |

### V2MOM/OKR Alignment

This project introduces a structural alignment between V2MOM and OKR (Objectives and Key Results):

| V2MOM Component | OKR Equivalent | Relationship |
|-----------------|----------------|--------------|
| Vision | Mission/North Star | Strategic direction |
| Values | - | Guiding principles (V2MOM-specific) |
| Method | Objective | What you will accomplish |
| Obstacles | - | Challenges (can be global or per-method) |
| Measure | Key Result | Quantifiable success criteria |

### MOM as a Unit

The key structural innovation is treating **MOM (Method-Obstacle-Measure)** as a cohesive unit:

```
Method (Objective)
├── Description
├── Priority
├── Obstacles (challenges specific to this method)
└── Measures (Key Results)
    ├── Measure 1 (target, timeline, status)
    ├── Measure 2 (target, timeline, status)
    └── Measure 3 (target, timeline, status)
```

This structure allows:

- Each Method to have its own measurable outcomes (like OKR)
- Method-specific obstacles alongside global obstacles
- Clear ownership and accountability per method
- Progress tracking at the method level

## Functional Requirements

### FR-1: JSON Intermediate Representation

- **FR-1.1**: Define a JSON schema for V2MOM documents
- **FR-1.2**: Support both flat (traditional) and nested (OKR-aligned) structures
- **FR-1.3**: Validate V2MOM JSON documents against the schema
- **FR-1.4**: Support metadata (author, date, version, fiscal year/quarter)

### FR-2: Marp Markdown Generation

- **FR-2.1**: Generate Marp-compatible markdown from V2MOM JSON
- **FR-2.2**: Support multiple slide themes (default, corporate, minimal)
- **FR-2.3**: Generate title slide with V2MOM metadata
- **FR-2.4**: Generate one slide per V2MOM component
- **FR-2.5**: Support method detail slides with nested measures
- **FR-2.6**: Include roadmap/timeline visualization

### FR-3: CLI Tool

- **FR-3.1**: `v2mom validate <file.json>` - Validate against schema
- **FR-3.2**: `v2mom generate marp <file.json> -o slides.md` - Generate Marp slides
- **FR-3.3**: `v2mom init` - Create template V2MOM JSON file
- **FR-3.4**: `v2mom convert <format> <file.json>` - Convert to other formats

### FR-4: Go Library

- **FR-4.1**: Provide Go types for V2MOM structures
- **FR-4.2**: JSON marshaling/unmarshaling with validation
- **FR-4.3**: Marp renderer as a reusable package
- **FR-4.4**: Extensible adapter interface for output formats

## Non-Functional Requirements

- **NFR-1**: Zero external runtime dependencies for core functionality
- **NFR-2**: JSON Schema compliant with JSON Schema Draft-07
- **NFR-3**: Generated Marp compatible with Marp CLI v3.x
- **NFR-4**: Go 1.21+ compatibility

## Output Format: Marp Slides

### Slide Structure

1. **Title Slide**: V2MOM name, author, date, fiscal period
2. **Vision Slide**: Full vision statement with visual emphasis
3. **Values Slide**: Prioritized list of values with descriptions
4. **Methods Overview**: Summary of all methods with priorities
5. **Method Detail Slides** (one per method):
   - Method description
   - Associated measures/key results
   - Method-specific obstacles
   - Timeline/status
6. **Obstacles Slide**: Global obstacles and mitigation strategies
7. **Measures Summary**: Dashboard view of all measures with status
8. **Roadmap Slide**: Timeline visualization of methods/projects

### Example Slide Output

```markdown
---
marp: true
theme: default
paginate: true
---

# FY2025 Product Strategy V2MOM

**Author:** Jane Smith, VP Product
**Period:** FY2025 Q1-Q4
**Last Updated:** 2025-01-15

---

## Vision

> Become the leading platform for enterprise workflow automation,
> enabling 10,000+ organizations to reduce operational costs by 40%

---

## Values

1. **Customer Obsession** - Every decision starts with customer impact
2. **Simplicity** - Complexity is the enemy of adoption
3. **Speed** - Ship fast, learn faster
4. **Transparency** - Open communication at all levels

---

## Methods

| # | Method | Priority | Status |
|---|--------|----------|--------|
| 1 | Launch self-service onboarding | P0 | In Progress |
| 2 | Expand enterprise integrations | P1 | Planning |
| 3 | Build partner ecosystem | P2 | Not Started |
```

## Future Integrations (Out of Scope for v1.0)

The following adapters are planned for future releases:

- **Confluence/Wiki**: Export to Confluence Storage Format
- **Pandoc**: Generate PDF, DOCX, HTML via Pandoc
- **Aha!**: Sync with Aha! roadmap via API
- **ProductBoard**: Push features/objectives to ProductBoard
- **Jira**: Create epics/initiatives from methods
- **Linear**: Sync with Linear projects

## Success Metrics

1. **Adoption**: 100+ GitHub stars within 6 months
2. **Reliability**: Zero critical bugs in JSON schema validation
3. **Performance**: Generate slides for 50+ method V2MOM in <1 second
4. **Extensibility**: Community contributions for 2+ output adapters

## References

- [Salesforce V2MOM Blog](https://www.salesforce.com/blog/how-to-create-alignment-within-your-company/)
- [V2MOM vs OKRs Comparison](https://www.v2mom.io/blog/v2mom-vs-okrs-discover-which-goal-setting-framework-aligns-best-with-your-business-objectives/)
- [SalesforceLabs V2MOM](https://github.com/SalesforceLabs/V2MOM)
- [SalesforceLabs MyV2MOM](https://github.com/SalesforceLabs/MyV2MOM)
- [Marp Documentation](https://marp.app/)
