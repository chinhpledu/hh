package main

import (
    "fmt"
    s "strings"
    //"os"
)
type point struct {
    x, y int
}
func main() {
    fmt.Println("Contains: ", s.Contains("test", "es"))
    fmt.Println("Count: ", s.Count("test", "t"))
    fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
    fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
    fmt.Println("Index: ", s.Index("test", "e"))
    fmt.Println("Join: ", s.Join([]string{"a", "b"}, "-"))
    fmt.Println("Repeat: ", s.Repeat("a", 5))
    fmt.Println("Replace: ", s.Replace("foo", "o", "0", -1))

    fmt.Println("Replace: ", s.Replace("foo", "o", "0", 1))
    fmt.Println("Split: ", s.Split("a-b-c-d-e", "-"))
    fmt.Println("ToLower: ", s.ToLower("TEST"))
    fmt.Println("ToUpper: ", s.ToUpper("test"))

    fmt.Println("Len: ", len("hello"))
    fmt.Println("Char: ", "hello"[1])

    p := point{1, 2}
    fmt.Println("====================================")
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%T\n", p)

    fmt.Printf("%t\n", true)

    fmt.Printf("%d\n", 123)
    fmt.Printf("%b\n", 14)
    fmt.Printf("%c\n", 33)
    fmt.Printf("%x\n", 456)

    fmt.Printf("%f\n", 78.9)
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    fmt.Printf("%s\n", "\"string\"")
    fmt.Printf("%q\n", "\"string\"")

    fmt.Printf("%x\n", "hex this")
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    fmt.Printf("|%-6s|%-6s|", "foo", "b")

    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
}
