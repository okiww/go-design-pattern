package factory

type AnimalFactory struct {
	Type string
}

func NewAnimalFactory(typeAnimal string) AnimalInterface  {
	if typeAnimal == "tiger" {
		return NewTiger()
	} else {
		return NewDog()
	}
}
