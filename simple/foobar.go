package simple

type FooBarService struct {
	FooService *FooService
	BarService *BarService
}

func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}
