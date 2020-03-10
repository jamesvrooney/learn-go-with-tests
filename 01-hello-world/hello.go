package main

import "fmt"

func main() {
    fmt.Println(Hello("James"))
}

func Hello(name string) string {
    return "Hello, " + name + "!"
}
