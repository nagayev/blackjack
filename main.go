package main

import (
	"os/signal"
	"syscall"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"io/ioutil"
	"strconv"
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
func checkError(err error){
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func botChoise(botScore *int) {
	if randInt(1, 3) != 1 && *botScore <= 19 {
		//AI))) Bot is smart enougth
		*botScore += randInt(2, 11)
	}
	if *botScore > 21 {
		fmt.Printf("Бот проиграл)\n")
		input("Нажми Enter для выхода!")
		os.Exit(0)
	}
}
func bot(){
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
				fmt.Printf("Ты проиграл :(\nУ тебя %d очков\n", score)
				input("Нажми Enter для выхода!")
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
		fmt.Printf("Бот выиграл %s\nУ бота %d очков, а у тебя - %d\n", username, botScore, score)
	} else if botScore < score {
		fmt.Printf("Ты выиграл!\nУ тебя %d очков, а у бота - %d\n", score, botScore)
	} else {
		fmt.Printf("Ничья!\n")
	}
}
func checkEnd()bool{
	files,err:=ioutil.ReadDir("./21")
	if err!=nil{
		panic(err)
	}
	for _,file:=range files{
		//check file
		content,_:=ioutil.ReadFile("./21/"+file.Name())
		lines:=strings.Split(string(content),"\n")
		if(lines[1]!="true"){
			return false
		}
	}
	return true
}
func getWinner()(string,int){
	user:=""
	max:=0
	files,_:=ioutil.ReadDir("./21")
	for _,file:=range files{
		content,_:=ioutil.ReadFile("./21/"+file.Name())
		line:=strings.Split(string(content),"\n")[0]
		score,_:=strconv.ParseInt(line,10,64)
		if score>int64(max) && score<22{
			max=int(score)
			user=file.Name()[:len(file.Name())-4]
		}
	}
	return user,max
}
func writeFile(username,text string)error{
	f,err:=os.OpenFile("./21/"+username+".txt",os.O_RDWR|os.O_CREATE,0777)
	fmt.Fprintf(f, text)
	if err!=nil{
		return err
	}
	err=f.Close()
	return err
}
func man(){
	score:=0
	running := true
	username := input("Введи ник: ")
	writeFile(username,"0\nfalse")
	for running{
		fmt.Printf("У тебя %d очков\n", score)
		answer := strings.ToLower(input("Взять еще (д/н) ?"))
		if answer == "д" {
			score += randInt(2, 11)
			if score > 21 {
				fmt.Printf("Ты проиграл :(\nУ тебя %d очков\n", score)
			}
		} else if answer == "н" {
			running = false
		} else { //neither y nor n
			fmt.Println("Неправильный ввод!")
		}	
	}
	writeFile(username,strconv.Itoa(score)+"\ntrue")
	fmt.Println("Ждем остальных участников...")
	seconds:=0
	for{
		if checkEnd() || seconds >=20{
			break
		}
		time.Sleep(2000 * time.Millisecond)
		seconds+=2
	}
	if seconds>=20{
		fmt.Println("Ждали одного из участников 20 секунд")
	}
	user,score:=getWinner()
	err:=os.Remove("./21/"+username+".txt")
	checkError(err)
	fmt.Printf("Победил игрок %s (%d очков)\n",user,score)
}
func main() {
	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Handle Ctrl+c, exiting...")
        os.Exit(1)
    }()
	config,err:=ioutil.ReadFile("config.json")
	if err!=nil{
		fmt.Println("Конфигурация не найдена!!!")
		os.Exit(1)
	}
	fmt.Println("Добро пожаловать в 21!")
	isbot,err:= strconv.ParseBool(string(config))
	if isbot{
		bot()
	} else{
		man()
	}
	input("Нажми Enter для выхода!")
}
