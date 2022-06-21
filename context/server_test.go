package context

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const mockData = "hello, world"

type MockStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func newSpyStore(t *testing.T) *MockStore {
	t.Helper()
	return &MockStore{response: mockData, t: t}
}

func (s *MockStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, char := range s.response {
			select {
			case <-ctx.Done():
				log.Println("mock store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-data:
		return result, nil
	}
}

func (s *MockStore) Cancel() {
	s.cancelled = true
}

func (s *MockStore) assertCancelled() bool {
	s.t.Helper()
	return assert.True(s.t, s.cancelled)
}

func (s *MockStore) assertNotCancelled() bool {
	s.t.Helper()
	return assert.False(s.t, s.cancelled)
}

func (s *MockStore) assertReceivedResponse(response *httptest.ResponseRecorder) {
	s.t.Helper()
	assert.Equal(s.t, s.response, response.Body.String())
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write(bytes []byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		store := newSpyStore(t)
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertReceivedResponse(response)
		store.assertNotCancelled()
	})
	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		store := newSpyStore(t)
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		assert.False(t, response.written, "response should not have been written")
	})
}
