package main

import (
	"github.com/bitclout/mongodb-dumper/cmd"
)

func main() {
	// There is a lot of indirection here introduced by the fact that we are using Viper to
	// manage our command-line flags. When this binary is run, a command is passed, such
	// as "run," which triggers a RunE() function defined in the cmd package. For example,
	// calling:
	// $ ./main run
	// would trigger the RunE() function defined in cmd/run.go. The flags that are available
	// are also all defined in this file.
	cmd.Execute()
}
