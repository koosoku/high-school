package main

import (
	"fmt"
	memberlist "github.com/hashicorp/memberlist"
)

func main(){
	list, err := memberlist.Create(memberlist.DefaultLocalConfig())
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}
}
