package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	analyze := flag.String("analyze", "", "analyze a repository path")
	find := flag.String("find", "", "find symbols or APIs matching a query")
	flag.Parse()

	switch {
	case *analyze != "":
		fmt.Printf("CodeAtlas analyzer is ready for %s\n", *analyze)
		fmt.Println("Next: parse source files into the canonical code graph.")
	case *find != "":
		fmt.Printf("Searching CodeAtlas index for %q\n", *find)
		fmt.Println("Next: graph + semantic retrieval will rank matching context.")
	default:
		fmt.Fprintln(os.Stderr, "CodeAtlas — architecture intelligence for codebases")
		fmt.Fprintln(os.Stderr, "usage: codeatlas -analyze ./repo | -find 'payment settlement'")
		os.Exit(2)
	}
}
