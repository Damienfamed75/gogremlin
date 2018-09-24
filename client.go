package gogremlin

import (
	"encoding/json"

	"github.com/damienfamed75/gogremlin/gremlinstring"
	"github.com/go-gremlin/gremlin"
	"github.com/pkg/errors"
)

// Client is a wrapper struct to manage
// gremlin Querying to simplify the workflow
type Client struct {
	g *gremlin.Client

	Graph    *gremlinstring.GremlinString
	Vertices []Vertex
}

// NewClient will return a ready client for gremlin
func NewClient(host string, cfgs ...gremlin.OptAuth) *Client {
	c := &Client{}
	gremlinClient, err := gremlin.NewClient(host, cfgs...)
	if err != nil {
		panic(err)
	}
	c.g = gremlinClient
	c.Graph = gremlinstring.New()
	c.loadVertices()

	return c
}

func (c *Client) loadVertices() error {
	res, err := c.ExecuteQuery("g.V()")
	if err != nil {
		return err
	}

	json.Unmarshal(res, &c.Vertices)

	return nil
}

// ExecuteQuery will run a simple query to gremlin-server
func (c *Client) ExecuteQuery(query string) ([]byte, error) {
	res, err := c.g.ExecQuery(query)
	if err != nil {
		return []byte{}, errors.Wrap(err, "couldn't execute query")
	}

	return res, nil
}

// ExecGremlinQuery will run a query using a gremlinstring instead of a regular string.
func (c *Client) ExecGremlinQuery(query *gremlinstring.GremlinString) ([]byte, error) {
	res, err := c.ExecuteQuery(query.String())
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}
