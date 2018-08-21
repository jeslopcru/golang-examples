package hello

const HELLO = "Hello, "

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return HELLO + name
}
