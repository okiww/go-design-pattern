package factory

type Tiger struct{}

type Dog struct{}

func NewTiger() AnimalInterface {
	return &Tiger{}
}

func NewDog() AnimalInterface {
	return &Dog{}
}

func (tiger *Tiger) Speak() string {
	return "Roar..."
}

func (dog *Dog) Speak() string {
	return "Wof..."
}
