package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strconv"
    "flag"
    "sort"
)

func parse_input(path string) ([]int, error) {

    elves := make([]int, 1)

    f, err := os.Open(path)
    if err != nil {
        return elves, err
    }
    defer f.Close()


    scanner := bufio.NewScanner(f)

    curElf := 0

    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            curNum, err := strconv.Atoi(line)
            if err != nil {
                return elves, err
            }
            curElf += curNum
        } else {
            elves = append(elves, curElf)
            curElf = 0
        }
    }

    if err := scanner.Err(); err != nil {
        return elves, err
	}

    return elves, nil
}

func part1(calories []int) {

    most := 0
    richest := 0

    for i, c := range calories {

        if c > most {
            most = c
            richest = i + 1
        }
    }

    fmt.Println("Richest Elf:", richest)
    fmt.Println("Calories:", most)

}

func part2(calories *[]int) {

    sort.Ints(*calories)

    top3 := (*calories)[len(*calories)-3:]
    sum := 0

    for _, cal := range top3 {
        sum += cal
    }

    fmt.Println(top3)
    fmt.Println("Calories:", sum)

}

func main() {

    flag_part := flag.Int("part", 0, "Which part of the task to execute (if more than one)")
    flag.Parse()

    switch *flag_part {
    case 1:
        calories, err := parse_input("input.dat")
        if err != nil {
            log.Fatal(err)
        }
        part1(calories)
    case 2:
        calories, err := parse_input("input.dat")
        if err != nil {
            log.Fatal(err)
        }
        part2(&calories)
    default:
        flag.Usage()
    }

}
