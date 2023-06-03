package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	golbs "github.com/Dominic-Wassef/golbs"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	cancel := make(chan struct{})

	go func() {
		<-quit
		cancel <- struct{}{}
	}()

	if err := golbs.LiveBuild("../", func(...golbs.BuildFuncOpt) error {
		fmt.Println("build ....")
		return nil
	}, cancel); err != nil {
		log.Fatal(err)
	}
}
