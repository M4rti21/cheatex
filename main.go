package main

import (
	"fmt"
	"os"

	"m4rti.dev/cheatex/parsers"
)

type Options struct {
	parser string
	action string   // compile || render
	files  []string // optional targets
}

func printUsage() {
	fmt.Println("USAGE: cheatex <parser> <compile | render> [optional: file1 file2 file3 ...]")
}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	opt := Options{
		parser: os.Args[1],
		action: os.Args[2],
		files:  os.Args[3:],
	}

	switch opt.action {
	case "compile":
		compileOption(opt.parser, opt.files)
		break
	case "render":
		renderOption(opt.parser, opt.files)
		break
	default:
		fmt.Println("This action doesnt exist")
		printUsage()
		break
	}

	fmt.Print(opt)
}

func compileOption(parser string, files []string) {
	p := parsers.Parsers[parser]
	if len(files) > 0 {
	}
	binds := p.Compile()
	fmt.Print(binds)
}

func renderOption(parser string, files []string) {
}
