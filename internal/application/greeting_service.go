package application

type GreetingService struct {
}

func (g *GreetingService) Greet(name string) string {
	return "hello" + name
}
