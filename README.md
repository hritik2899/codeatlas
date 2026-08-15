# CodeAtlas

> **Google Maps for your codebase.**
>
> Turn a large repository into a queryable architecture graph for developers and AI agents.

CodeAtlas is an open-source developer intelligence platform that parses source code into a navigable graph of repositories, services, APIs, symbols, calls, dependencies, and data flows. It is designed to give humans and coding agents the *right* context instead of dumping an entire codebase into a prompt.

## Why CodeAtlas?

Modern repositories are too large to understand by reading files one at a time. Search finds text, but architecture is about relationships:

```text
Service → API → Handler → Function → Dependency → Database
             ↓
          Kafka Topic
```

CodeAtlas makes those relationships queryable.

## Architecture

```text
                    ┌──────────────────────┐
                    │      codeatlas       │
                    │        CLI           │
                    └──────────┬───────────┘
                               │
                 ┌─────────────▼─────────────┐
                 │       Ingestion Layer     │
                 │  Tree-sitter / Parsers    │
                 └─────────────┬─────────────┘
                               │
                 ┌─────────────▼─────────────┐
                 │   Canonical Code Model    │
                 │ repos/services/symbols    │
                 │ APIs/calls/dependencies   │
                 └──────┬─────────┬──────────┘
                        │         │
                ┌───────▼───┐ ┌──▼──────────┐
                │ Graph     │ │ Semantic    │
                │ Store     │ │ Index       │
                └───────┬───┘ └──────┬──────┘
                        │             │
                        └──────┬──────┘
                               ▼
                     ┌───────────────────┐
                     │ Context Ranker    │
                     │ relevance + graph │
                     │ proximity + cost  │
                     └─────────┬─────────┘
                               │
                     ┌─────────▼─────────┐
                     │   MCP / REST API  │
                     └─────────┬─────────┘
                               ▼
                       Developer / Agent
```

## Design goals

- **Architecture-aware retrieval**, not keyword search alone.
- **Deterministic graph queries** for dependency and call relationships.
- **Context budgets** so agents receive useful information without wasting tokens.
- **Incremental indexing** so a small commit does not require rebuilding everything.
- **Language extensibility** through a canonical intermediate representation.
- **Local-first development** with reproducible Docker-based infrastructure.

## Planned milestones

- [x] Project architecture and product definition
- [ ] Go core and CLI
- [ ] Tree-sitter ingestion
- [ ] Canonical code graph
- [ ] Dependency and call graph queries
- [ ] Incremental indexing
- [ ] Semantic retrieval
- [ ] Context ranking and token budgets
- [ ] MCP server
- [ ] Benchmark suite
- [ ] Web visualization

## Example target workflow

```bash
codeatlas analyze ./payments
codeatlas graph service payment-service
codeatlas find "settlement"
codeatlas context "Where is payment settlement implemented?" --budget 8000
```

## Philosophy

CodeAtlas should make a repository feel less like a pile of files and more like a system you can reason about.

The project is intentionally being built in public: architecture decisions, benchmarks, limitations, and trade-offs will be documented rather than hidden behind marketing language.

## License

Apache-2.0 (planned).
