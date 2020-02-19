package main

import (
	"fmt"
	"github.com/teamwork/reload"
	"log"
	"os"
)

func main() {

	go func() {
		err := reload.Do(log.Printf,
			reload.Dir("./tmp", func() { log.Printf("/tmp changed") }),
			reload.Dir(".", reload.Exec))
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println(os.Args)
	fmt.Println(os.Environ())
	ch := make(chan bool)
	<-ch
}
