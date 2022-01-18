package main

import "fmt"

type Code interface {
	ParseLibrary()
	Run()
}

type Go struct{ Filename string }

func (this Go) ParseLibrary() {
	fmt.Println("go mod tidy")
}
func (this Go) Run() {
	fmt.Println("go run", this.Filename)
}

type Java struct{ Filename string }

func (this Java) ParseLibrary() {
	fmt.Println("mvn dependency:resolve")
}
func (this Java) Run() {
	fmt.Println("java -jar", this.Filename)
}

func After_CodeRunner(code Code) {
	code.ParseLibrary()
	code.Run()
}
func after() {
	fmt.Println("\nafter")

	After_CodeRunner(Java{Filename: "app.jar"})
	fmt.Println()
	After_CodeRunner(Go{Filename: "app.go"})
}
