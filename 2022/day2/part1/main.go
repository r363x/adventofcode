package main

import (
	"fmt"
    "bufio"
    "os"
    "strings"
)

type Move string
type Outcome string

const (
	ThemRock     Move    = "A"
	ThemPaper    Move    = "B"
	ThemScissors Move    = "C"
	UsRock       Move    = "X"
	UsPaper      Move    = "Y"
	UsScissors   Move    = "Z"
    Win          Outcome = "W"
    Loss         Outcome = "L"
    Draw         Outcome = "D"
)

func (m Move) i() int64 {
    switch m {
    case UsRock:
        return 1
    case UsPaper:
        return 2
    case UsScissors:
        return 3
    }
    return -1 
}

func (o Outcome) i() int64 {
    switch o {
    case Win:
        return 6
    case Loss:
        return 0
    case Draw:
        return 3
    }
    return -1
}

type Round struct {
	MoveThem Move
	MoveUs   Move
}

func (r *Round) play() (Outcome, int64){
    switch r.MoveThem {
    case ThemRock:

        switch r.MoveUs {
        case UsPaper:
            return Win, Win.i() + UsPaper.i()
        case UsScissors:
            return Loss, Loss.i() + UsScissors.i()
        case UsRock:
            return Draw, Draw.i() + UsRock.i()
        }

    case ThemPaper:

        switch r.MoveUs {
        case UsScissors:
            return Win, Win.i() + UsScissors.i()
        case UsRock:
            return Loss, Loss.i() + UsRock.i()
        case UsPaper:
            return Draw, Draw.i() + UsPaper.i()
        }

    case ThemScissors:

        switch r.MoveUs {
        case UsRock:
            return Win, Win.i() + UsRock.i()
        case UsPaper:
            return Loss, Loss.i() + UsPaper.i()
        case UsScissors:
            return Draw, Draw.i() + UsScissors.i()
        }
    }

    return Outcome("INVALID"), -1
}

type Tournament struct {
    Rounds []Round
}

type TournamentResults struct {
    Score int64
    Wins int64
    Losses int64
    Draws int64
}

func (t *TournamentResults) add(outcome Outcome, points int64) {
    t.Score += points
    switch outcome {
    case Win:
        t.Wins++
    case Loss:
        t.Losses++
    case Draw:
        t.Draws++
    }
}

func (t TournamentResults) String() string {
    return fmt.Sprintf(
        "Wins: %d\nLosses: %d\nDraws: %d\nScore: %d",
        t.Wins,
        t.Losses,
        t.Draws,
        t.Score,
    )
}

func (t *Tournament) results() TournamentResults {

    var results TournamentResults

    for _, round := range t.Rounds {
        results.add(round.play())
    }

    return results
}

func main() {

    f, err := os.Open("input.dat")
    if err != nil {
        fmt.Println("Error")
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    var tournament Tournament

    for scanner.Scan() {
        line := scanner.Text()

        tokens := strings.Split(line, " ")
        if len(tokens) == 2 {
            tournament.Rounds = append(tournament.Rounds, Round{
                MoveThem: Move(tokens[0]),
                MoveUs: Move(tokens[1]),
            })
        }
    }

    results := tournament.results()
    fmt.Println("Tournament Results")
    fmt.Println("------------------")
    fmt.Println(results)

}

