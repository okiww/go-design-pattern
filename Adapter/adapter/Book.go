package adapter

/*
	- Class Book
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

type Book struct {
	Title  string
	Author string
}

func NewBook(title string, author string) Book {
	return Book{
		Title:  title,
		Author: author,
	}
}

func (book *Book) GetTitle() string {
	return book.Title
}

func (book *Book) GetAuthor() string {
	return book.Author
}