package main

import (
	"fmt"

	"github.com/damienfamed75/gogremlin"
)

var (
	localhost = "ws://127.0.0.1:8182/gremlin"

	commands = []string{
		"graph = JanusGraphFactory.open('conf/janusgraph-cassandra-es.properties')",
		"GraphOfTheGodsFactory.load(graph)",
		"g = graph.traversal()",
	}
)

func main() {
	c := gogremlin.NewClient(localhost)

	for _, v := range c.Vertices {
		fmt.Println(v.Value.Properties["name"][0].Value.Value)
	}

	res, _ := c.ExecuteQuery(c.Graph.V().String())

	fmt.Println(string(res))
}
