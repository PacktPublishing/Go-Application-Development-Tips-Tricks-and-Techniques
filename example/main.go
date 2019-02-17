package main

type Example struct {
	Name string `json:"name,omitempty"`
}

func (e *Example) Read(p []byte) (n int, err error) {
	panic("not implemented")
}

func main() {
}
