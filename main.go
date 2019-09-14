package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text) //delete \r \n
}
func botChoise(botScore *int) {
	if randint(1, 3) != 1 && *botScore <= 19 {
		//AI))) Bot is smart enougth
		*botScore += randInt(2, 11)
	}
	if *botScore > 21 {
		fmt.Printf("Бот проиграл)")
		os.Exit(0)
	}
}
func main() {
	score := 0
	botScore := 0
	running := true
	username := input("Введи ник: ")
	rand.Seed(time.Now().UnixNano()) //init randomiser
	for running {
		fmt.Printf("У тебя %d очков\n", score)
		answer := strings.ToLower(input("Взять еще (д/н) ?"))
		if answer == "д" {
			score += randInt(2, 11)
			if score > 21 {
				fmt.Printf("Ты проиграл :(\nУ тебя %d очков", score)
				os.Exit(0)
			}
			botChoise(&botScore)
		} else if answer == "н" {
			botChoise(&botScore)
			running = false
		} else { //neither y nor n
			fmt.Println("Неправильный ввод!")
		}
	}
	if botScore > score {
		fmt.Printf("Бот выиграл %s\nУ бота %d очков, а у тебя - %d", username, botScore, score)
	} else if botScore < score {
		fmt.Printf("Ты выиграл!\nУ тебя %d очков, а у бота - %d", score, botScore)
	} else {
		fmt.Printf("Ничья!")
	}
}
