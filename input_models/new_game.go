package input_models

type NewGame struct {
	Title string
	Description string
	ReleaseDate string
	Price int
	TagIds []int
	GenreIds []int
	PublisherId int
	DeveloperId int
	GameHeaderImage string
	GameImages []string
}
