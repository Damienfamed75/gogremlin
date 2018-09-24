package gremlinstring

// Label is to access Vertice labels.
func (g *GremlinString) Label() *GremlinString {
	g.append(".label()")
	return g
}
