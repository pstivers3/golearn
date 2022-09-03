// does not run
// error code 126

package main

import (
	"fmt"
	//"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "/usr/local/bin/aws", "s3 --profile cis-dev ls cisint-nb-staging --recursive --summarize | tail -1")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the output
	fmt.Println(string(stdout))
}
