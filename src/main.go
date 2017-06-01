package main
import (
    "fmt"
    "os"
)

func convert(data string) string {
    return data + " covfefe"
}

func main() {
    to_convert := os.Args
    length := len(to_convert)
    if length < 2 {
        fmt.Println("Usage: govfefe [\"string_1\", ...]")
    } else {
        for i:=1 ; i<len(to_convert); i++ {
            fmt.Printf("%d. %s\n", i, convert(to_convert[i]))
        }
    }
}
