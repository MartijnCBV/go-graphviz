package gvc

import (
	_ "unsafe"

	"github.com/MartijnCBV/go-graphviz/cdt"
	"github.com/MartijnCBV/go-graphviz/cgraph"
	"github.com/MartijnCBV/go-graphviz/internal/wasm"
)

//go:linkname toGraph github.com/MartijnCBV/go-graphviz/cgraph.toGraph
func toGraph(*wasm.Graph) *cgraph.Graph

//go:linkname toGraphWasm github.com/MartijnCBV/go-graphviz/cgraph.toGraphWasm
func toGraphWasm(*cgraph.Graph) *wasm.Graph

//go:linkname toNode github.com/MartijnCBV/go-graphviz/cgraph.toNode
func toNode(*wasm.Node) *cgraph.Node

//go:linkname toEdge github.com/MartijnCBV/go-graphviz/cgraph.toEdge
func toEdge(*wasm.Edge) *cgraph.Edge

//go:linkname toDictLink github.com/MartijnCBV/go-graphviz/cdt.toLink
func toDictLink(*wasm.DictLink) *cdt.Link

//go:linkname toDictLinkWasm github.com/MartijnCBV/go-graphviz/cdt.toLinkWasm
func toDictLinkWasm(*cdt.Link) *wasm.DictLink
