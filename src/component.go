package main

import "syscall/js"

func main() {
    var doc = js.Global().Get("document")
    var body = doc.Get("body")
    body.Set("innerHTML", "Hello World")
}
