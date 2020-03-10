package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
    fmt.Println(Hello("James", ""))
}

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    if language == "Spanish" {
        return "Hola, " + name
    }

    return englishHelloPrefix + name
}
