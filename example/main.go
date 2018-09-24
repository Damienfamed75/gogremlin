package main

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	grem "github.com/damienfamed75/gogremlin"
)

var localhost = "ws://127.0.0.1:8182/gremlin"

func main() {
	// Setup ---------------------
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	c := grem.NewClient(localhost)

	command := c.Graph.V().Limit(1).Label()

	// Output the command's string
	log.Info().Str("cmd", command.String()).Msg("executing command")

	// retrieve the result without error checking
	res, _ := c.ExecGremlinQuery(command)

	// Write result to json file
	ioutil.WriteFile("output.json", res, 0644)

	// Output the resulting marshalled object
	log.Debug().RawJSON("result", res).Msg("command finished")
}
