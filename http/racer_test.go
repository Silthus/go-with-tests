package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fastest url wins", func(t *testing.T) {
		slowServer := createDelayedServer(10 * time.Millisecond)
		fastServer := createDelayedServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		winner, err := Racer(slowUrl, fastUrl)
		assert.Equal(t, fastUrl, winner)
		assert.NoError(t, err)
	})
	t.Run("error if both urls timeout (>10s)", func(t *testing.T) {
		server1 := createDelayedServer(21 * time.Millisecond)
		server2 := createDelayedServer(22 * time.Millisecond)
		defer server1.Close()
		defer server2.Close()

		winner, err := ConfigurableRacer(server1.URL, server2.URL, 20*time.Millisecond)
		assert.Error(t, err)
		assert.Empty(t, winner)
	})
}

func BenchmarkRacer(b *testing.B) {
	slowServer := createDelayedServer(20 * time.Millisecond)
	fastServer := createDelayedServer(0)
	defer slowServer.Close()
	defer fastServer.Close()

	for i := 0; i < b.N; i++ {
		Racer(slowServer.URL, fastServer.URL)
	}
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
