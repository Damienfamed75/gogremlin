package gremlinstring

import "strconv"

// Range will apply a filter from minimum to maximum
func (g *GremlinString) Range(min, max int) *GremlinString {
	g.append(".range(" + strconv.Itoa(min) + "," + strconv.Itoa(max) + ")")
	return g
}
