# FARM Stack Fresher Plan – 7 Months
**Start**: May 18, 2026 | **End**: Dec 17, 2026

---

## Daily Non-Negotiables
| Block | Time | Task |
|---|---|---|
| Morning | 2h | DSA — Striver Sheet / LeetCode |
| Afternoon | 1.5–2.5h | FARM Building + Learning |
| Evening | 30 min | README / Docs / LinkedIn |
| End of day | — | ≥1 Git commit, no exceptions |

---

## LinkedIn Headline
`Python Backend Developer | FastAPI + React + MongoDB | RAG · Webhooks · SaaS Tooling`

---

## Certifications
| # | Cert | Cost | When |
|---|------|------|------|
| 1 | AWS Cloud Practitioner | ₹7.5k–9.5k | Month 1–2 |
| 2 | HashiCorp Terraform Associate | ₹6k–7k | Month 4–5 |
| 3 | AWS SAA *(optional)* | ₹13k+ | Month 6–7 |

---

## Core Skills to Build
- FastAPI — async, Pydantic v2, JWT, background tasks, WebSockets, SSE
- React — Hooks, TanStack Query, Tailwind CSS
- MongoDB — aggregation pipelines, indexing, Atlas Vector Search
- ARQ + Redis — background task queues, retry logic
- Docker, GitHub Actions CI/CD, Render/Railway deployment
- LLM APIs (OpenAI/Groq), embeddings, RAG fundamentals
- Webhook patterns, HMAC signing, event-driven design

---

## 7-Month Timeline
| Month | Dates | Focus |
|---|---|---|
| 1 | May 18 – Jun 17 | FastAPI + ARQ/Redis basics → **Project 1** + Start AWS CCP |
| 2 | Jun 18 – Jul 17 | RAG + embeddings + vector search → **Project 2** + Pass AWS CCP |
| 3 | Jul 18 – Aug 17 | Multi-tenancy + financial logic + PDF → **Project 3** |
| 4 | Aug 18 – Sep 17 | Encryption + CLI tooling (Typer) → **Project 4** + Start Terraform |
| 5 | Sep 18 – Oct 17 | Polish P1–P4 + Docker + CI/CD + architecture write-ups |
| 6 | Oct 18 – Nov 17 | **Project 5** + Pass Terraform Cert |
| 7 | Nov 18 – Dec 17 | Final polish + mock interviews + 150+ applications |

---

## The 5 Projects

---

### Project 1 — Webhook Delivery & Event Routing Platform
**Deadline**: Jun 17 | **Stack**: FastAPI + ARQ + Redis + MongoDB + React

**Features to build**:
- `POST /events` — accepts JSON payload, queues immediately, returns fast
- Endpoint registration — users register URLs + optional signing secret
- HMAC-SHA256 signature on every outgoing request (Stripe-style)
- ARQ workers — deliver event, retry with exponential backoff (1s → 2s → 4s → 8s, max 5 attempts)
- Delivery log per event — timestamp, HTTP status, response body, attempt #
- Dead letter queue — exhausted events stored for manual replay
- Replay button — re-trigger any past event
- React dashboard — live event stream (WebSocket), endpoint health, delivery drill-down

**Key tech**: ARQ + Redis, HMAC-SHA256, WebSockets, MongoDB event/log persistence, decoupled worker process

**AI Prompt**:
> "Guide me to build a Webhook Delivery and Event Routing Platform with FastAPI + ARQ (Redis task queue) + React + MongoDB. Include: event ingestion API, async delivery workers with exponential backoff retry, HMAC-SHA256 signature on outgoing requests, delivery log with attempt history, WebSocket live event stream in React dashboard, dead letter queue with manual replay. Walk me through full architecture and implementation."

---

### Project 2 — RAG-Powered Personal Knowledge Base
**Deadline**: Jul 17 | **Stack**: FastAPI + MongoDB Atlas Vector Search (or FAISS) + OpenAI/Groq + React

**Features to build**:
- PDF + text ingestion — pypdf parsing, recursive text chunking (hand-implement, don't use LangChain)
- Embedding generation — OpenAI `text-embedding-3-small` or Sentence Transformers (free)
- Vector storage — MongoDB Atlas Vector Search or local FAISS index
- Query pipeline — embed question → retrieve top-5 chunks → pass as context to LLM → stream response via SSE
- Source attribution — every answer shows which document chunks it used
- Document manager UI — upload, list, delete; each doc shows chunk count + embedding status
- Streaming chat UI — SSE-based, not a basic POST/response
- Per-user namespacing — separate knowledge bases, no cross-contamination

**Key tech**: Text chunking strategy, cosine similarity search, RAG pipeline, SSE streaming, async background ingestion tasks

**AI Prompt**:
> "Help me build a RAG-powered Personal Knowledge Base using FastAPI + MongoDB Atlas Vector Search + React. Include: PDF/text ingestion with recursive chunking (explain the strategy), embeddings via OpenAI API, vector similarity search, RAG query pipeline with context passed to LLM, SSE response streaming, source attribution per answer, React UI with document manager and streaming chat. Guide me through every component."

---

### Project 3 — Multi-Tenant SaaS Billing Engine
**Deadline**: Aug 17 | **Stack**: FastAPI + MongoDB + ReportLab + React

**Features to build**:
- Multi-tenant schema — Organizations → Subscriptions → Invoices → Usage Records
- Plan management — monthly/annual cycles, feature limits, usage quotas
- Subscription lifecycle — activate, upgrade (with proration), downgrade, cancel with grace period
- Proration math — credit for unused days on old plan; use `Decimal`, never `float`
- Monthly invoice generation — PDF via ReportLab, line items + tax + totals
- Dunning flow — failed payment → retry day 3 → retry day 7 → suspend → notify (background tasks)
- Usage metering — track API calls or storage per tenant per billing period
- Admin dashboard — MRR chart, subscriptions by plan, overdue invoices, tenant billing health
- Webhook events — emits `invoice.created`, `subscription.upgraded`, `payment.failed` → routes through Project 1

**Key tech**: Multi-tenant MongoDB schema design, Decimal arithmetic, ReportLab, ARQ background tasks, proration calculation

**AI Prompt**:
> "Build a FARM stack Multi-Tenant SaaS Subscription and Billing Engine. FastAPI + MongoDB + React. Include: multi-tenant data model, plan management, full subscription lifecycle with proration on upgrades (explain the math), PDF invoice generation via ReportLab, dunning management with retry schedule, usage metering per tenant, React admin dashboard with MRR chart. Emit webhook events for billing state changes. Guide me through architecture and every feature."

---

### Project 4 — Developer Config & Secrets Management (Vault-Lite)
**Deadline**: Sep 17 | **Stack**: FastAPI + MongoDB + React + Python CLI (Typer)

**Features to build**:
- Hierarchy — Project → Environment → Config Key (e.g. `myapp → production → DATABASE_URL`)
- Fernet encryption at rest — no plaintext secrets ever stored in MongoDB
- Access tokens — read-only vs read-write scopes, expiry timestamps
- Python CLI (`configctl`) built with Typer:
  - `configctl pull --project myapp --env production` → writes `.env` file
  - Warns if token is expiring within 24h
  - Structured for `pip install`
- Secret rotation — generate new token, mark old as expiring, CLI warns on next pull
- Audit log — who changed/accessed what key, timestamp, IP, action type
- Environment diff — compare dev vs staging vs prod, highlight missing/mismatched keys
- React dashboard — masked secret viewer (reveal on click), token manager, audit trail timeline

**Key tech**: Fernet encryption, Typer CLI (PyPI-publishable structure), token-based auth separate from user JWT, diff algorithm

**AI Prompt**:
> "Help me build a Developer Config and Secrets Management Platform (Doppler-lite) using FastAPI + MongoDB + React. Include: project/environment/key hierarchy, Fernet encryption at rest, read/write access tokens with expiry, Python CLI with Typer that runs `configctl pull` to write a .env file (pip-installable), secret rotation with expiry warnings, audit logs, environment diff view, React dashboard with masked secret viewer. Walk me through every component including CLI packaging."

---

### Project 5 — Internal Developer Platform (Self-Service Portal)
**Deadline**: Nov 17 | **Stack**: FastAPI + MongoDB + Jinja2 + React

**Features to build**:
- Service catalog — register microservices with owner, repo, tech stack, SLOs, health status, on-call; full-text search
- Self-service provisioning — simulate DB creation / API key issuance with async status workflow: Requested → Provisioning → Ready
- Envoy config generator — input upstream host/port/cluster → generate YAML via Jinja2 → preview + download
- Team usage analytics — API call counts, resource consumption, last-active (MongoDB aggregation pipelines)
- Runbook viewer — Markdown stored in MongoDB, rendered with syntax highlighting, searchable
- RBAC — Admin (full), Developer (provision + edit), Viewer (read-only) via JWT claims in middleware
- Golden path templates — downloadable FastAPI + Docker Compose starter with env var conventions (references Project 4)
- Service health cards — uptime %, last incident, SLO compliance indicator per service

**Key tech**: Jinja2 templating, MongoDB full-text search, RBAC via JWT claims, aggregation pipelines, Markdown rendering in React

**AI Prompt**:
> "Help me build a FARM stack Internal Developer Platform and Self-Service Engineering Portal. FastAPI + React + MongoDB. Include: searchable service catalog with SLOs, async resource provisioning with status workflow, Envoy config generator via Jinja2, team usage analytics via MongoDB aggregation, Markdown runbook viewer with full-text search, RBAC via JWT claims, and golden path service templates. Guide me through architecture, data modeling, and full implementation."

---

## Project Integration (Mention in Every Interview)
- **P1 + P3**: Billing engine emits `invoice.created` / `payment.failed` events → delivered by the webhook platform
- **P4 + P5**: IDP provisions API keys and env configs using the secrets manager
- **P2**: Standalone — your AI literacy signal

Say in interviews: *"Two of my projects actually integrate — my billing engine routes events through my webhook platform. I designed it that way intentionally."*

---

## Weekly Checklist (Every Sunday)
- [ ] 12–16h DSA done
- [ ] 12–18h FARM building done
- [ ] 1 major feature shipped
- [ ] README updated with 1 architectural decision explained
- [ ] 1 LinkedIn post — specific technical problem solved, not "made progress on X"
- [ ] Cert material revised
- [ ] 1 engineering blog read (Stripe blog → P1, Doppler blog → P4, Honeycomb → observability)

---

## README Checklist (All 5 Projects)
- [ ] Problem statement — 2 sentences, product-person tone
- [ ] Architecture diagram — Excalidraw PNG embedded
- [ ] Tech stack + one-line justification per choice
- [ ] 2–3 key technical challenges (specific, not vague)
- [ ] Demo GIF or screenshots
- [ ] Local setup — works in <5 min with `docker compose up`
- [ ] "What I'd do differently" section

---

## Interview Pitch Formula
> "I built [X] because [real problem real teams face]. The hard part was [specific challenge]. I solved it with [your approach]. Here's what it enables: [concrete outcome]."

**P1 example**: *"I built a webhook delivery platform because every SaaS product needs to reliably push events to customer servers even when those servers are temporarily down. The hard part was guaranteeing delivery without blocking the main API. I used ARQ with Redis for async delivery, exponential backoff retry, and HMAC-SHA256 signing on every outgoing request — exactly like Stripe does."*

---

## Dec 17 Success Metrics
- [ ] 500+ LeetCode (250 Easy / 200 Medium / 50 Hard)
- [ ] 5 deployed, documented projects on GitHub
- [ ] AWS Cloud Practitioner ✓
- [ ] Terraform Associate ✓
- [ ] Projects 1 + 3 integrated and demo-able live
- [ ] Can explain every architectural decision without notes