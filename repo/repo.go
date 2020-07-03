package repo

type Repository interface {
	Insert(string) string
	Get() string
}

func NewRepository() Repository {
	return Repo{}
}

type Repo struct{}

func (Repo) Insert(data string) string {
	return data
}

func (Repo) Get() string {
	return "Get Repo"
}
