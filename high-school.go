package main

import (
	"fmt"
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

	channel := make(chan string)
	go sendNode(channel)
	go listenNode(channel)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
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