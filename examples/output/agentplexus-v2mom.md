---
marp: true
theme: default
paginate: true
header: "AgentPlexus | FY2025 Annual"
footer: "V2MOM | AgentPlexus FY2025 Strategy"
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
    background: linear-gradient(135deg, #1a365d 0%, #2b6cb0 100%);
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
    border-left: 4px solid #2b6cb0;
    padding-left: 1em;
  }
---

<!-- _class: title -->

# AgentPlexus FY2025 Strategy
**Author:** AgentPlexus Team
**Team:** AgentPlexus
**Period:** FY2025 Annual
**Status:** Active

---

<!-- _class: vision -->

## Vision

> Become the definitive open-source Go ecosystem for building production-grade AI agents, providing composable infrastructure for LLMs, secrets management, search, retrieval, voice, and observability that enables developers to ship secure, observable, multi-provider AI applications in days instead of months.

---

## Values

**1. Modularity** - Each component solves one problem well - independent, focused modules that compose elegantly

**2. Security by Default** - Posture checks, credential gating, and encryption built-in from the start

**3. Provider Agnostic** - Interface-based design enabling pluggable backends without vendor lock-in

**4. Observability First** - Hooks and tracing built into every component for production visibility

**5. Zero External Dependencies** - Core libraries maintain minimal dependencies for reliability and security


---

## Methods

| # | Method | Priority | Status |
|---|--------|----------|--------|
| 1 | Complete OmniLLM as the unified LLM abstraction layer | P0 (Critical) | In Progress |
| 2 | Establish OmniObserve as the LLM observability standard | P0 (Critical) | In Progress |
| 3 | Launch OmniRetrieve for production RAG | P1 (High) | In Progress |
| 4 | Build OmniVoice for voice AI applications | P1 (High) | Planning |
| 5 | Mature OmniVault ecosystem | P0 (Critical) | In Progress |
| 6 | Establish AgentKit as the agent development standard | P0 (Critical) | In Progress |
| 7 | Build developer community and adoption | P1 (High) | In Progress |

---

## Method 1: Complete OmniLLM as the unified LLM abstraction layer

Finalize OmniLLM as the production-ready unified SDK for OpenAI, Anthropic, Google Gemini, X.AI Grok, Ollama, and AWS Bedrock with streaming, memory, and observability hooks.

**Owner:** Core Team
**Priority:** P0 (Critical)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Provider coverage | 7 providers (add Mistral, Cohere) | On Track | 85% |
| Streaming support completeness | 100% all providers | On Track | 95% |
| Token counting accuracy | Provider-native counting | On Track | 70% |


### Obstacles

- **Provider API inconsistencies** (Medium): Adapter pattern with comprehensive normalization layer


---

## Method 2: Establish OmniObserve as the LLM observability standard

Complete the vendor-agnostic LLM observability framework supporting Opik, Langfuse, and Phoenix with unified tracing, evaluation, and dataset management.

**Owner:** Observability Team
**Priority:** P0 (Critical)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Backend coverage | 5 backends (add Weights & Biases, MLflow) | On Track | 60% |
| OmniLLM integration depth | Full automatic instrumentation | On Track | 75% |
| OpenTelemetry support | Full OTLP export | At Risk | 10% |


### Obstacles

- **Backend API differences** (Medium): Common semantic conventions with backend-specific adapters


---

## Method 3: Launch OmniRetrieve for production RAG

Complete the unified retrieval library supporting Vector RAG, Graph RAG, and Hybrid strategies with pgvector, Pinecone, and Neo4j providers.

**Owner:** RAG Team
**Priority:** P1 (High)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Vector database providers | 4 (add Pinecone, Weaviate, Qdrant) | On Track | 25% |
| Graph RAG support | Neo4j + Amazon Neptune | Behind | 20% |
| Reranking strategies | 5 strategies (add Cohere, cross-encoder) | On Track | 60% |


### Obstacles

- **Graph RAG complexity** (High): Start with simple traversal, iterate to more complex graph algorithms


---

## Method 4: Build OmniVoice for voice AI applications

Complete the voice AI abstraction layer supporting TTS, STT, and Voice Agents with ElevenLabs, Twilio, and WebRTC transports.

**Owner:** Voice Team
**Priority:** P1 (High)
**Status:** Planning

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| TTS provider coverage | 3 providers (ElevenLabs, AWS Polly, Google) | Not Started | - |
| STT provider coverage | 3 providers (Whisper, Deepgram, Google) | Not Started | - |
| Transport support | WebRTC + WebSocket + PSTN (Twilio) | Not Started | - |


### Obstacles

- **Real-time audio complexity** (High): Start with WebSocket, add WebRTC, partner with Twilio for PSTN


---

## Method 5: Mature OmniVault ecosystem

Complete the unified secret management ecosystem with AWS, 1Password, and native OS keyring support, plus Swift UI for macOS.

**Owner:** Security Team
**Priority:** P0 (Critical)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Provider coverage | 8 providers (add HashiCorp Vault, Azure Key Vault, GCP Secret Manager) | On Track | 50% |
| macOS Swift UI completion | Full native app | On Track | 30% |
| VaultGuard adoption in AgentPlexus | All projects using VaultGuard | On Track | 70% |


### Obstacles

- **Cross-platform keyring differences** (Medium): oscompat package abstractions with graceful fallbacks


---

## Method 6: Establish AgentKit as the agent development standard

Complete AgentKit as the production-ready framework for building A2A Protocol agents with multi-cloud deployment (Kubernetes, AWS AgentCore).

**Owner:** Platform Team
**Priority:** P0 (Critical)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| Deployment target coverage | 4 (add AWS AgentCore, GCP Cloud Run) | On Track | 75% |
| IaC template coverage | Helm, CDK, Pulumi, Terraform | On Track | 75% |
| Boilerplate reduction | 1500+ lines | Achieved | 100% |



---

## Method 7: Build developer community and adoption

Grow the AgentPlexus community through documentation, examples, blog posts, and reference implementations.

**Owner:** DevRel Team
**Priority:** P1 (High)
**Status:** In Progress

### Measures

| Measure | Target | Status | Progress |
|---------|--------|--------|----------|
| GitHub stars across all repos | 1000+ | Behind | 5% |
| Reference implementations | 5 complete examples | On Track | 40% |
| Documentation coverage | Full API docs + tutorials for all components | On Track | 50% |


### Obstacles

- **Discoverability in crowded AI space** (High): Focus on Go ecosystem (underserved), emphasize security and observability differentiators


---

## Obstacles

| Obstacle | Severity | Likelihood | Status |
|----------|----------|------------|--------|
| Go ecosystem maturity for AI | High | High | Accepted |
| Rapid LLM provider API changes | Medium | High | Mitigating |
| Resource constraints (open source project) | Medium | Medium | Accepted |


### Mitigation Strategies

- **Go ecosystem maturity for AI:** Focus on areas where Go excels (performance, concurrency, deployment) and provide bridges to Python tools
- **Rapid LLM provider API changes:** Abstraction layer insulates users; version providers independently; automated compatibility testing
- **Resource constraints (open source project):** Prioritize ruthlessly; automated testing; community contributions


---

## Measures Dashboard

| Measure | Target | Progress | Status |
|---------|--------|----------|--------|
| Provider coverage | 7 providers (add Mistral, Cohere) | [########--] 85% | On Track |
| Streaming support completeness | 100% all providers | [#########-] 95% | On Track |
| Token counting accuracy | Provider-native counting | [#######---] 70% | On Track |
| Backend coverage | 5 backends (add Weights & Biases, MLflow) | [######----] 60% | On Track |
| OmniLLM integration depth | Full automatic instrumentation | [#######---] 75% | On Track |
| OpenTelemetry support | Full OTLP export | [#---------] 10% | At Risk |
| Vector database providers | 4 (add Pinecone, Weaviate, Qdrant) | [##--------] 25% | On Track |
| Graph RAG support | Neo4j + Amazon Neptune | [##--------] 20% | Behind |
| Reranking strategies | 5 strategies (add Cohere, cross-encoder) | [######----] 60% | On Track |
| TTS provider coverage | 3 providers (ElevenLabs, AWS Polly, Google) | - | Not Started |
| STT provider coverage | 3 providers (Whisper, Deepgram, Google) | - | Not Started |
| Transport support | WebRTC + WebSocket + PSTN (Twilio) | - | Not Started |
| Provider coverage | 8 providers (add HashiCorp Vault, Azure Key Vault, GCP Secret Manager) | [#####-----] 50% | On Track |
| macOS Swift UI completion | Full native app | [###-------] 30% | On Track |
| VaultGuard adoption in AgentPlexus | All projects using VaultGuard | [#######---] 70% | On Track |
| Deployment target coverage | 4 (add AWS AgentCore, GCP Cloud Run) | [#######---] 75% | On Track |
| IaC template coverage | Helm, CDK, Pulumi, Terraform | [#######---] 75% | On Track |
| Boilerplate reduction | 1500+ lines | [##########] 100% | Achieved |
| GitHub stars across all repos | 1000+ | [----------] 5% | Behind |
| Reference implementations | 5 complete examples | [####------] 40% | On Track |
| Documentation coverage | Full API docs + tutorials for all components | [#####-----] 50% | On Track |

---

## Roadmap Projects

| Project | Priority | Quarter | Status |
|---------|----------|---------|--------|
| OmniLLM Mistral Provider | P1 | Q1 | Proposed |
| OmniLLM Cohere Provider | P2 | Q2 | Proposed |
| OmniLLM Token Counting | P1 | Q1 | In Progress |
| OmniObserve OpenTelemetry Export | P1 | Q2 | Proposed |
| OmniRetrieve Pinecone Provider | P1 | Q2 | Proposed |
| OmniRetrieve Neo4j Provider | P1 | Q2 | Proposed |
| OmniVoice ElevenLabs Integration | P1 | Q2 | Proposed |
| OmniVault macOS Swift UI | P2 | Q2 | In Progress |
| AgentKit Terraform Templates | P2 | Q3 | Proposed |
| Stats Agent Team v2.0 | P1 | Q1 | In Progress |
| AgentPlexus Documentation Site | P0 | Q1 | In Progress |

---

