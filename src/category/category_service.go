package category

type CategoryService struct {}

func NewService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) Create() (bool, error) {
	return true, nil
}

func (s *CategoryService) Read() ([]Category, error) {
	var c []Category

	return c, nil
}

func (s *CategoryService) Update() (bool, error) {
	return true, nil
}

func (s *CategoryService) Delete() (bool, error) {
	return true, nil
}