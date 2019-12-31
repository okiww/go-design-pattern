package adapter

/*
	- BookCatalogAdapter is implement CatalogAdapter
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

type BookCatalogAdapter struct {
	Book Book
}

func NewBookCatalogAdapter(book Book) CatalogAdapter {
	return &BookCatalogAdapter{
		Book: book,
	}
}

func (bookCatalogAdapter *BookCatalogAdapter) GetCatalogTitle() string {
	return bookCatalogAdapter.Book.Title + " by " + bookCatalogAdapter.Book.Author
}
