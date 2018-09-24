package gremlinstring

import "strconv"

// New will return a new GremlinString to work with.
func New() *GremlinString {
	var str GremlinString = "g"
	return &str
}

// String converts the GremlinString to a string
func (g *GremlinString) String() string {
	return string(*g)
}

// append is used for construction of commands
func (g *GremlinString) append(str string) {
	*g += GremlinString(str)
}

// gatherParam will act like an optional
// parameter.
func gatherParam(params ...int) string {
	switch len(params) {
	case 1:
		return strconv.Itoa(params[0])
	default:
		return ""
	}
}
