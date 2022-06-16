package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockWebsiteChecker struct {
	Called map[string]int
}

func (m *MockWebsiteChecker) CheckWebsite(website string) bool {
	m.Called[website] += m.Called[website] + 1
	if website == "http://fail.test" {
		return false
	}
	return true
}

func CreateMockWebsiteChecker() *MockWebsiteChecker {
	return &MockWebsiteChecker{Called: make(map[string]int)}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckSingleWebsite(t *testing.T) {
	checker := CreateMockWebsiteChecker()
	assertWebsiteCheck(t, checker, "https://michaelreichenbach.de", true)
	assertWebsiteCheck(t, checker, "http://fail.test", false)
}

func TestCheckWebsites(t *testing.T) {
	checker := CreateMockWebsiteChecker()
	result := CheckWebsites(checker.CheckWebsite, []string{"https://michaelreichenbach.de", "http://fail.test"})
	assert.Equal(t, map[string]bool{
		"https://michaelreichenbach.de": true,
		"http://fail.test":              false,
	}, result)
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func assertWebsiteCheck(t *testing.T, checker *MockWebsiteChecker, url string, expected bool) {
	t.Helper()
	result := CheckWebsite(checker.CheckWebsite, url)
	assert.Equal(t, expected, result)
	assert.Greater(t, checker.Called[url], 0)
}
