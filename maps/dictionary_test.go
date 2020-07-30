package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""
		assertStrings(t, got, want)
		assertErrors(t, err, ErrorNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("add normal entry", func(t *testing.T) {
		dictionary := Dictionary{}
		erradd := dictionary.Add("test", "this is just a test")
		assertErrors(t, erradd, nil)

		want := "this is just a test"
		got, errsearch := dictionary.Search("test")
		assertStrings(t, got, want)
		assertErrors(t, errsearch, nil)
	})
	t.Run("add same entry", func(t *testing.T) {
		dictionary := Dictionary{}
		err1 := dictionary.Add("test", "this is just a test")
		err2 := dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, errsearch := dictionary.Search("test")
		assertStrings(t, got, want)
		assertErrors(t, err1, nil)
		assertErrors(t, err2, ErrorKeyAlreadyExists)
		assertErrors(t, errsearch, nil)
	})
	t.Run("add same entry", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("", "this is not allowed")

		want := ""
		got, errsearch := dictionary.Search("")
		assertStrings(t, got, want)
		assertErrors(t, err, ErrorInvalidKey)
		assertErrors(t, errsearch, ErrorNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertErrors(t *testing.T, errgot, errwant error) {
	t.Helper()
	if errgot != errwant {
		t.Errorf("got %q want %q given, %q", errgot, errwant, "test")
	}
}
