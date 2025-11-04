package tool


type ToolService struct {
	// DB
}

func NewService() *ToolService {
	return &ToolService{}
}

func (s *ToolService) Create(t Tool) (bool, error) {
	return true, nil
}

func (s *ToolService) Read() ([]Tool, error){
	var tool []Tool

	return tool, nil
}

func (s *ToolService) Update(t Tool) (bool, error) {
	return true, nil
}

func (s *ToolService) Delete() (bool, error) {
	return true, nil
}

