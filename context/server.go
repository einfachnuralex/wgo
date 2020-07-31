package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan string)
		go func() {
			data <- store.Fetch(r.Context())
		}()

		ctx := r.Context()

		select {
		case <-ctx.Done():
			store.Cancel()
		case val := <-data:
			fmt.Fprint(w, val)
		}

	}

}

type Store interface {
	Fetch(ctx context.Context) string
	Cancel()
}
