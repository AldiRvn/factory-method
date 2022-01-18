package main

import (
	"fmt"
	"strings"
)

func ParseLibrary(language string) {
	switch language {
	case "java":
		fmt.Println("mvn dependency:resolve")
	case "golang":
		fmt.Println("go mod tidy")
	}
}

func CodeRunner(fileName string) {
	switch {
	case strings.HasSuffix(fileName, ".jar"):
		ParseLibrary("java")
		fmt.Println("java -jar", fileName)
	case strings.HasSuffix(fileName, ".go"):
		ParseLibrary("golang")
		fmt.Println("go run", fileName)
	}
}

func main() {
	CodeRunner("app.jar")
	fmt.Println()
	CodeRunner("app.go")

	after()
}
