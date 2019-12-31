package main

/*
	- Implement Adapter Pattern in Go
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

import (
	"fmt"
	"okky/go-design-pattern/Adapter/adapter"
)

/*
	- Assume we have an online shop
	- Assume we have products Book and Screencast
	- Assume we want show all product in Catalog
*/
func main() {
	// create list of array type CatalogAdapter
	var lists []adapter.CatalogAdapter

	// new instance of BookCatalogAdapter
	bookCatalogAdapter := adapter.NewBookCatalogAdapter(adapter.NewBook("book", "Okky"))

	// new instance of ScreencastCatalogAdapter
	screencastCatalogAdapter := adapter.NewScreencastCatalogAdapter(adapter.NewScreencast("screncast", "Okky"))

	// append data to lists
	lists = append(lists, bookCatalogAdapter)
	lists = append(lists, screencastCatalogAdapter)

	// print catalog title
	for i, _ := range lists {
		fmt.Println(lists[i].GetCatalogTitle())
	}
}
