package main

import (
	"fmt"
	"log"
	"time"

	device "github.com/sconklin/go-dockerpi-relay"
	"github.com/sconklin/go-i2c"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.NewI2C(0x10, 1)
	checkError(err)
	defer i2c.Close()
	rly, err := device.NewRelay(i2c)
	checkError(err)

	time.Sleep(1 * time.Second)
	fmt.Println("One ON")
	err = rly.On(1)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("One OFF")
	err = rly.Off(1)

	time.Sleep(1 * time.Second)
	fmt.Println("Two ON")
	err = rly.On(2)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Two OFF")
	err = rly.Off(2)

	time.Sleep(1 * time.Second)
	fmt.Println("Three ON")
	err = rly.On(3)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Three OFF")
	err = rly.Off(3)

	time.Sleep(1 * time.Second)
	fmt.Println("Four ON")
	err = rly.On(4)
	checkError(err)
	time.Sleep(1 * time.Second)
	fmt.Println("Four OFF")
	err = rly.Off(4)
}
