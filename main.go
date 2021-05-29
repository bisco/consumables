package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	testFlag := flag.Bool("test", false, "generate and insert test data to db")
	serverPort := flag.Uint("port", 12000, "the number of port on which server runs")
	flag.Parse()

	dbInit()
	if *testFlag {
		fmt.Println("there is no test mode...")
		return
	}

	r := setupRouter()
	if *serverPort > 65535 {
		log.Fatalf("%v is larger than max port number", *serverPort)
	}
	r.Run(":" + strconv.FormatUint(uint64(*serverPort), 10))
}
