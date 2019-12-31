package adapter

/*
	- ScreencastCatalogAdapter is implement CatalogAdapter
	- Author : Okky Muhamad Budiman
	- Email : budimanokky93@gmail.com
*/

type ScreencastCatalogAdapter struct {
	Screencast Screencast
}

func NewScreencastCatalogAdapter(screencast Screencast) CatalogAdapter {
	return &ScreencastCatalogAdapter{
		Screencast: screencast,
	}
}

func (screencastCatalogAdapter *ScreencastCatalogAdapter) GetCatalogTitle() string {
	return screencastCatalogAdapter.Screencast.Title + " by " + screencastCatalogAdapter.Screencast.Author
}
