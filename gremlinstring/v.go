package gremlinstring

// V is to acces all vertices within a graph
func (g *GremlinString) V(params ...int) *GremlinString {
	strParam := gatherParam(params...)
	g.append(".V(" + strParam + ")")
	return g
}
