// does not run
// exit status 127

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "package_cloud yank vericity/rpm/el/7 nbx-1.0.0-1487.x86_64.rpm")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the output
	fmt.Println(string(stdout))
}
