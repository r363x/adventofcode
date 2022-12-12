package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strconv"
)

func main() {

    f, err := os.Open("input.dat")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    elves := make([]int, 10)

    scanner := bufio.NewScanner(f)

    curElf := 0

    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            curNum, err := strconv.Atoi(line)
            if err != nil {
                log.Fatal(err)
            }
            curElf += curNum
        } else {
            elves = append(elves, curElf)
            curElf = 0
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatal("reading line:", err)
	}

    most := 0
    richest := 0

    for i, c := range elves {

        if c > most {
            most = c
            richest = i + 1
        }
    }

    fmt.Println("Richest Elf:", richest)
    fmt.Println("Calories:", most)

}
