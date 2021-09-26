package main

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist Artist  `json:"artist"`
	Price  float64 `json:"price"`
}

type Artist struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
}

type AlbumSale struct {
	ID       string `json:"id"`
	AlbumId  string `json:"albumId"`
	ArtistId string `json:"artistId"`
	QtySold  int64  `json:"QtySold"`
}
