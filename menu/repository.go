package menu

type Beer struct {
	Id   int32
	Name string
}

type Repository struct {
}

type MenuRepository interface {
	ListBeers() ([]Beer, error)
}

func NewRepository() MenuRepository {
	return &Repository{}
}

func (r *Repository) ListBeers() ([]Beer, error) {
	beers := []Beer{{
		Id:   1,
		Name: "Change",
	}, {
		Id:   2,
		Name: "Singha",
	}}
	return beers, nil
}
