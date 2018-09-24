package gremlinstring

import "strconv"

// Limit will apply a limit starting from 0. (analogous to range)
// scope will also apply a scope to the filter.
func (g *GremlinString) Limit(params ...interface{}) *GremlinString {
	if len(params) < 1 {
		panic("no parameters given to Limit()")
	}

	g.append(".limit(")

	for _, p := range params {
		switch p.(type) {
		case string:
			g.append(p.(string) + ",")
		case int:
			g.append(strconv.Itoa(p.(int)) + ")")
		}
	}

	return g
}
