package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "127.0.0.1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()

	buf := bufio.NewReader(stdout)                  // buf := bufio.NewReader(stdout)
													// num := 1
	for {											// for {
		line, _, _ := buf.ReadLine()				//      line, _, _ := buf.ReadLine()
		fmt.Println(string(line))					//      if num > 3 {
	}												//         os.Exit(0)
													// 		}
													// 		num += 1
													// 		fmt.Println(string(line))
													// }

}

