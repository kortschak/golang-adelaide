package main

import (
	"fmt"
	"log"

	"github.com/kortschak/yaegiconf"
)

func main() {
	type Config string
	var c Config
	err := yaegiconf.EvalTo(&c, `for {}; config.Value("Configured")`)
	if err != nil {
		log.Fatalf("failed to parse configuration: %v", err)
	}
	fmt.Printf("%#v\n", c)
}

-- go.mod --
module github.com/kortschak/golang-adelaide/jokes/code/old

go 1.13

require github.com/kortschak/yaegiconf v0.0.0-20190924035841-837558797c65
