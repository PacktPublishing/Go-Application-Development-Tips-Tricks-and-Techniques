package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

type some_struct struct {
	called   int
	filename string
}

func (this *some_struct) WriteToFile(what string, ctx context.Context) (error, int) {
	this.called += 1
	finished := make(chan error)

	go func() {
		defer close(finished)

		f, err := os.Create(this.filename)
		if err != nil {
			finished <- err
			return
		}

		_, err = f.Write([]byte(what))
		finished <- err
	}()

	select {
	case <-ctx.Done():
		return errors.New(fmt.Sprintf("cancelled writing %s", what)), 0
	case <-finished:
		return nil, this.called
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	a := some_struct{filename: "foo.dat"}
	a.WriteToFile("hello", ctx)
}
