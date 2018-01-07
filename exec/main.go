package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

func main() {
	c := `svn export https://github.com/googlei18n/libphonenumber/trunk/resources/geocoding`
	cmd := exec.Command("/bin/bash", "-c", c)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe()", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("cmd.StderrPipe()", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("cmd.Start()", err)
		return
	}
	bytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputBuf := bufio.NewReader(stdout)

	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}
		fmt.Println(string(output))
	}

	if err = cmd.Wait(); err != nil {
		fmt.Println("cmd.Wait()", err, " stderr: ", string(bytes))
		return
	}
}
