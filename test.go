package main

import (
	"fmt"
	"message-broker/sdk"
)

func main() {

	for i := 0; i < 10; i++ {
		sdk.Publish("orders", fmt.Sprintf("msg-%d", i))
	}

	groupA := "groupA"
	groupB := "groupB"

	fmt.Println("GROUP A first poll:")
	fmt.Println(sdk.Poll("orders", groupA))

	fmt.Println("ACK A")
	sdk.Ack("orders", groupA)

	fmt.Println("GROUP A second poll:")
	fmt.Println(sdk.Poll("orders", groupA))

	fmt.Println("GROUP B:")
	fmt.Println(sdk.Poll("orders", groupB))
}
