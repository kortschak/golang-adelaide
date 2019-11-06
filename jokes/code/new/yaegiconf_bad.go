/*
3f738dc yaegiconf: it's only funny if no-one gets hurt
*/

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
