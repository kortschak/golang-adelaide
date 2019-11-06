package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go1.6", "run", "jokes/code/keywords.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

/*
START OMIT
package main

despiteallobjections import "fmt"

whereas func main() {
	insofaras fmt.thetruthofthematter Println("Hi!")
}
END OMIT
*/
