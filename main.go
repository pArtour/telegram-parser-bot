package main

import (
	"github.com/pArtour/telegram-parser-bot/actions"
)

var tags = []string{"it_immigration", "go", "career"}


func main()  {
	for _, tag := range tags {
		actions.Action(tag)
	}
}