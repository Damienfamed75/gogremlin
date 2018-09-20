package gremlinstring

// New will return a new GremlinString to work with.
func New() *GremlinString {
	var str GremlinString = "g"
	return &str
}

// V is to acces all vertices within a graph
func (g *GremlinString) V() *GremlinString {
	g.append(".V()")
	return g
}

// String converts the GremlinString to a string
func (g *GremlinString) String() string {
	return string(*g)
}

// append is used for construction of commands
func (g *GremlinString) append(str string) {
	*g += GremlinString(str)
}
