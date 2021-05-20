package main

import (
	"fmt"
	"io/ioutil"
	"opencodes/cmd"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	//defer profile.Start().Stop()
	savePid()
	cmd.Execute()
}

func savePid() {
	pidFileName := "pid/" + filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()
	fmt.Println(pid)
	ioutil.WriteFile(pidFileName, []byte(strconv.Itoa(pid)), 075)
}
