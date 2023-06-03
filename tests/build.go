package main

import (
	"log"

	golbs "github.com/Dominic-Wassef/golbs"
)

// buildNormalProgram
func buildNormalProgram() error {
	sh := &golbs.Sh{}
	const command string = "go build -o nprog ./programs/normal/main.go"
	return sh.Init(command).Run().Error()
}

// runNormalProgram
func runNormalProgram() error {
	const command string = "./nprog"
	return new(golbs.Sh).Init(command).Run().Error()
}

// cleanNormalProgram
func cleanNormalProgram() error {
	const command string = "rm nprog"
	return new(golbs.Sh).Init(command).Run().Error()
}

// buildInputProgram
func buildInputProgram() error {
	const command string = "go build -o inprog ./programs/input/main.go"
	return new(golbs.Sh).Init(command).Run().Error()
}

// runInputProgram
func runInputProgram() error {
	const command string = "./inprog"
	return new(golbs.Sh).Init(command).In("hello, world").Error()
}

// cleanInputProgram
func cleanInputProgram() error {
	const command string = "rm inprog"
	return new(golbs.Sh).Init(command).Run().Error()
}

func main() {

	// build the script and runs itself. if the current file has changed.
	ok, err := golbs.GoBuildYourSelf("build.go")
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return
	}

	opts := []golbs.BuildFuncOpt{
		{
			FuncName: "buildNormalProgram",
			Func:     buildNormalProgram,
		},
		{
			FuncName: "runNormalProgram",
			Func:     runNormalProgram,
		},
		{
			FuncName: "cleanNormalProgram",
			Func:     cleanNormalProgram,
		},
	}

	if err := golbs.Build(opts...); err != nil {
		log.Fatal(err)
	}

	opts = []golbs.BuildFuncOpt{
		{FuncName: "buildInputProgram", Func: buildInputProgram},
		{FuncName: "runInputProgram", Func: runInputProgram},
		{FuncName: "cleanInputProgram", Func: cleanInputProgram},
	}

	if err := golbs.Build(opts...); err != nil {
		log.Fatal(err)
	}
}
