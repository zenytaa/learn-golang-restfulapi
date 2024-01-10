package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	BarRepository *BarRepository
}

func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{BarRepository: barRepository}
}
