package maps

import "errors"

//ErrNotFound todo
var ErrNotFound error = errors.New("not found")

//InvalidKey todo
var InvalidKey error = errors.New("invalid key")

//KeyAlreadyExists todo
var KeyAlreadyExists error = errors.New("key already exists")

// Dictionary bla
type Dictionary map[string]string

// Search bla
func (d Dictionary) Search(searchword string) (string, error) {
	result, ok := d[searchword]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

// Add bla
func (d Dictionary) Add(key, value string) error {
	if key == "" {
		return InvalidKey
	}
	if _, err := d.Search(key); err != ErrNotFound {
		return KeyAlreadyExists
	}
	d[key] = value
	return nil
}
