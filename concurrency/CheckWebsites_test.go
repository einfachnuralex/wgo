package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	myCoolWebsites := []string{"google.com", "web.de", "yahoo.com"}

	got := CheckWebsites(subbaSlowChecker, myCoolWebsites)
	want := map[string]bool{
		"google.com": false,
		"web.de":     false,
		"yahoo.com":  false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expceted %v got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for index := range urls {
		urls[index] = "lalallalalala"
	}

	for j := 0; j < b.N; j++ {
		CheckWebsites(subbaSlowChecker, urls)
	}
}

func mockChecker(_ string) bool {
	return false
}

func subbaSlowChecker(uri string) bool {
	time.Sleep(80 * time.Millisecond)
	return false
}
