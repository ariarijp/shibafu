package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bmuller/arrow/lib"
	"github.com/fatih/color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var cal [7][]string

	colors := []*color.Color{
		color.New(color.BgWhite),
		color.New(color.BgYellow),
		color.New(color.BgGreen),
		color.New(color.BgRed),
	}

	wd := arrow.Yesterday().Weekday()

	for i := 0; i < int(wd); i++ {
		cal[i] = append(cal[i], " ")
	}

	for i := 0; i <= 365; i++ {
		if i < 365-90 {
			cal[wd] = append(cal[wd], "?")
		} else {
			col := colors[rand.Int()%len(colors)]
			cal[wd] = append(cal[wd], col.SprintFunc()(" "))
		}

		if wd == 6 {
			wd = 0
			continue
		}

		wd++
	}

	for _, week := range cal {
		fmt.Println(strings.Join(week, ""))
	}
}
