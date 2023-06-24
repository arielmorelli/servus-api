package main

import (
	// "os"

	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/arielmorelli/servus-api/api"
	"github.com/arielmorelli/servus-api/filehandler"
	"github.com/gin-gonic/gin"
)

func main() {
	var debug, version bool
	var port, inputFile string

	// set cli arguments
	flag.StringVarP(&port, "port", "p", "8080", "Port to run.")
	flag.StringVarP(&inputFile, "file", "f", "", "Input file")
	flag.BoolVarP(&debug, "debug", "d", false, "Debug mode")
	flag.BoolVarP(&version, "version", "v", false, "Version")
	flag.Parse()

	if version {
		fmt.Println("VERSION_TAG")
		os.Exit(0)
	}

	if inputFile != "" {
		if err := filehandler.LoadRoutesFromFile(inputFile); err != nil {
			fmt.Println("Unable to parse file inputFile", err)
			os.Exit(1)
		}
	}

	// API setup
	if !debug {
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Listening and serving HTTP on localhost:%s\n", port)
	}
	router := gin.Default()
	router.Any("/*route", api.DynamicRouting)

	router.Run(fmt.Sprintf("localhost:%s", port))
}
