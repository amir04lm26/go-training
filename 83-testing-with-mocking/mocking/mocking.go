package mocking

// NOTE: Example: Mocking in Tests
/*
	Using interfaces, you can create mock implementations for testing.
*/

type Service interface {
	DoSomething() string
}

type RealService struct{}

func (realService RealService) DoSomething() string {
	return "Doing something"
}

func UseService(service Service) string {
	return service.DoSomething()
}
