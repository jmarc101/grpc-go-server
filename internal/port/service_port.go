package port

type GreetingService interface {
	Greet(name string) string
}
