package day08

func BuildGreeting(name string) string {
	return "Hello, " + name
}

func Welcome(name string) string {
	return MakeGreeting(name) + "!"
}
