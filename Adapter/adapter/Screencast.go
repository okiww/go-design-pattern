package adapter

/*
	- Class Screencast
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

type Screencast struct {
	Title  string
	Author string
}

func NewScreencast(title string, author string) Screencast {
	return Screencast{
		Title:  title,
		Author: author,
	}
}

func (screen *Screencast) GetTitle() string {
	return screen.Title
}

func (screen *Screencast) GetAuthor() string {
	return screen.Author
}
