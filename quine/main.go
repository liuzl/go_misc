package main

import "fmt"

func main() {
	p := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tp := %q\n\tfmt.Printf(p, p)\n}\n"
	fmt.Printf(p, p)
}
