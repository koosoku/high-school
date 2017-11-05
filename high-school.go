package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

type Node struct {
	index int
	mailBox chan string
}

func (node *Node) Listening() string{
	return <- node.mailBox
}

func (node *Node) SpreadRumor(message string, mailBoxes []chan string) {
	for i:= range mailBoxes {
		if i != node.index && rand.Intn(100) < 1 {
			mailBoxes[i] <- "Message From: " + strconv.Itoa(node.index) + " Message: " + message;
		}
	}
}

func main(){

	channels  := make([]chan string, 1000000)
	for index := range channels {
		channels[index] = make(chan string)
	}

	for index := range channels {
		go node(channels, index)
	}

	channels[0]<-"Hiyo"

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func node (mailBoxes []chan string, index int) {
	node := Node {
		index: index,
		mailBox: mailBoxes[index],
	}
	message := node.Listening()
	fmt.Print("Message Received: " + message + "\n")
	node.SpreadRumor(message, mailBoxes)
}
