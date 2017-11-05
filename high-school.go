package main

import (
	"fmt"
	"strconv"
	//memberlist "github.com/hashicorp/memberlist"
)

func main(){
	//list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	//if err != nil {
	//	panic("Failed to create memberlist: " + err.Error())
	//}
	//// Ask for members of the cluster
	//for _, member := range list.Members() {
	//	fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	//}

	channels := []chan string{
		make(chan string),
		make(chan string),
	}

	go node(channels, 0)
	go node(channels, 1)


	channels[0]<-"Hiyo"
	//channel := make(chan string)
	//go sendNode(channel)
	//go listenNode(channel)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}


func node (mailBoxes []chan string, index int) {
	mailBox := mailBoxes[index]
	for i := 0;i<100; i++ {
		message:= <-mailBox
		fmt.Print("Message Received: "+ message + "\n")
		for i:= range mailBoxes {
			if i != index {
				mailBoxes[i] <- "Message From: " + strconv.Itoa(index) + " Message: " + message;
			}
		}
	}
}

func sendNode (sending chan string) {
	sending<-"Hello"
}

func listenNode(listening chan string) {
	fmt.Print(<-listening)
}

func printNumbers(from string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(from, ":", i)
	}
}