// Package model contains CodeAtlas's language-independent representation of a codebase.
//
// Parsers should translate language-specific syntax into these small primitives.
// Keeping this layer independent from Tree-sitter (or any future parser) means the
// graph, search and context layers don't need to care which language produced a node.
package model

import "strings"

type NodeKind string

const (
	Repository NodeKind = "repository"
	Package    NodeKind = "package"
	File       NodeKind = "file"
	Service    NodeKind = "service"
	Class      NodeKind = "class"
	Function   NodeKind = "function"
	Method     NodeKind = "method"
	Interface  NodeKind = "interface"
	API        NodeKind = "api"
)

type EdgeKind string

const (
	Contains    EdgeKind = "contains"
	Calls       EdgeKind = "calls"
	Imports     EdgeKind = "imports"
	Implements  EdgeKind = "implements"
	DependsOn   EdgeKind = "depends_on"
	Exposes     EdgeKind = "exposes"
)

type Position struct {
	StartLine int `json:"start_line"`
	EndLine   int `json:"end_line"`
}

type Node struct {
	ID        string            `json:"id"`
	Kind      NodeKind          `json:"kind"`
	Name      string            `json:"name"`
	Path      string            `json:"path,omitempty"`
	Language  string            `json:"language,omitempty"`
	Position  Position          `json:"position"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type Edge struct {
	From string   `json:"from"`
	To   string   `json:"to"`
	Kind EdgeKind `json:"kind"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func NormalizeName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}

func (g *Graph) AddNode(node Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(edge Edge) {
	g.Edges = append(g.Edges, edge)
}
