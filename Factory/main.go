package main

import (
	"fmt"
	"okky/go-design-pattern/Factory/factory"
)

func main()  {
	tiger := factory.NewAnimalFactory("tiger")
	soundTiger := tiger.Speak()

	dog := factory.NewAnimalFactory("dog")
	soundDog := dog.Speak()

	fmt.Println(soundTiger)
	fmt.Println(soundDog)
}
