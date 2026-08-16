# Tree-sitter parser adapter

This package is the planned integration boundary for Tree-sitter-backed source parsing.

The important architectural rule is that Tree-sitter stays an implementation detail:

```text
source bytes
    ↓
Tree-sitter
    ↓
language AST
    ↓
parser adapter
    ↓
parser.Result
    ↓
CodeAtlas graph model
```

Keeping this boundary small means the ingestion pipeline can support additional parsers later without changing graph storage, retrieval, or MCP layers.

The concrete grammar integration will be introduced in the next increment after the adapter contract is validated.
