package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response   string
	isCanceled bool
}

func (s *StubStore) Fetch(ctx context.Context) string {
	time.Sleep(10 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.isCanceled = true
}

func TestHandler(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&StubStore{data, false})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("cancel path", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{data, false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(ctx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.isCanceled {
			t.Errorf("not canceled")
		}
	})
}
