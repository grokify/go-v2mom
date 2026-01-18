---
marp: true
theme: default
paginate: true
header: "Product | FY2025 Annual"
footer: "V2MOM | Product Strategy FY2025"
style: |
  section {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  }
  section.title {
    text-align: center;
  }
  section.title h1 {
    font-size: 2.5em;
  }
  section.vision {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: #ffffff;
  }
  section.vision h2 {
    color: #ffffff;
  }
  table {
    font-size: 0.85em;
    width: 100%;
  }
  th {
    background: #f7fafc;
  }
  .status-done { color: #38a169; }
  .status-progress { color: #3182ce; }
  .status-risk { color: #e53e3e; }
  .status-planning { color: #805ad5; }
  blockquote {
    font-size: 1.3em;
    font-style: italic;
    border-left: 4px solid #764ba2;
    padding-left: 1em;
  }
---

<!-- _class: title -->

# Product Strategy FY2025
**Author:** Jane Smith, VP Product
**Team:** Product
**Period:** FY2025 Annual
**Status:** Active

---

<!-- _class: vision -->

## Vision

> Become the leading platform for enterprise workflow automation, enabling 10,000+ organizations to reduce operational costs by 40% through intelligent process optimization.

---

## Values

**1. Customer Obsession** - Every decision starts with customer impact and value delivery

**2. Simplicity** - Complexity is the enemy of adoption - make it easy

**3. Speed** - Ship fast, learn faster, iterate continuously

**4. Transparency** - Open communication at all levels, share context widely


---

## Methods

| # | Method | Priority | Status |
|---|--------|----------|--------|
| 1 | Launch self-service onboarding | P0 (Critical) | In Progress |
| 2 | Expand enterprise integrations | P1 (High) | Planning |
| 3 | Build partner ecosystem | P2 (Medium) | Not Started |

---

## Method 1: Launch self-service onboarding

Enable customers to sign up, configure, and start using the platform without sales or support intervention.

**Owner:** Onboarding Team
**Priority:** P0 (Critical)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Self-service signup conversion rate | 25% | On Track | 48% |
| Time to first value | < 1 day | On Track | 78% |
| Support tickets from new users | < 10 per week | At Risk | 56% |


### Obstacles

- **Legacy authentication system** (High): Phased migration to Auth0 in Q1


---

## Method 2: Expand enterprise integrations

Build native integrations with top 10 enterprise systems to reduce implementation time.

**Owner:** Integrations Team
**Priority:** P1 (High)
**Status:** Planning

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Native integrations shipped | 15 | On Track | 20% |
| Average integration setup time | < 2 hours | On Track | 50% |


### Obstacles

- **API rate limits** (Medium): Implement intelligent batching and caching


---

## Method 3: Build partner ecosystem

Create a partner program enabling system integrators and consultants to implement and extend our platform.

**Owner:** Partnerships Team
**Priority:** P2 (Medium)
**Status:** Not Started

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Certified partners | 25 | Not Started | - |
| Partner-sourced revenue | $2M ARR | Not Started | - |



---

## Obstacles

| Obstacle | Severity | Likelihood | Status |
|----------|----------|------------|--------|
| Hiring freeze | High | High | Accepted |
| Market uncertainty | Medium | Medium | Mitigating |


### Mitigation Strategies

- **Hiring freeze:** Focus on automation and contractor support for non-core work
- **Market uncertainty:** Emphasize ROI and cost-reduction messaging


---

## Measures Dashboard

| Measure | Target | Progress | Status |
|---------|--------|----------|--------|
| Self-service signup conversion rate | 25% | [####------] 48% | On Track |
| Time to first value | < 1 day | [#######---] 78% | On Track |
| Support tickets from new users | < 10 per week | [#####-----] 56% | At Risk |
| Native integrations shipped | 15 | [##--------] 20% | On Track |
| Average integration setup time | < 2 hours | [#####-----] 50% | On Track |
| Certified partners | 25 | - | Not Started |
| Partner-sourced revenue | $2M ARR | - | Not Started |

---

## Roadmap Projects

| Project | Priority | Quarter | Status |
|---------|----------|---------|--------|
| Self-Service Onboarding Flow | P0 | Q1 | In Progress |
| Auth0 Migration | P0 | Q1 | In Progress |
| Salesforce Native Integration | P1 | Q2 | Proposed |

---

