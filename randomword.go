package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "math/rand"
    "time"
    "os"
    "strconv"
)

func main() {
    cla := os.Args
    count := 1
    cmdlen := len(cla)
    if cmdlen > 2 {
        fmt.Println("Expecting 0 or 1 argument, not more than that!")
        return
    }

    if cmdlen == 2 {
        arg, err := strconv.Atoi(cla[1])
        if err != nil {
            fmt.Println(err)
            return
        } else {
            count = arg
        }
    }

    filebytes, err := ioutil.ReadFile("words_alpha.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    str := string(filebytes)

    words := strings.Fields(str)

    wordslen := len(words)

    if count >= wordslen {
        for _, elem := range words {
            fmt.Println(elem)
        }
        return
    }

    rand.Seed(time.Now().Unix())

    for i:=0; i<count; i++ {
        fmt.Println(words[rand.Intn(wordslen)])
    }
}