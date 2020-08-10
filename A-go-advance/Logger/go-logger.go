package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[demo-test]", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Print("Hello, log file!")
	logger.Panic("panic err.")
	//logger.Fatal(errors.New("err df "))
}
