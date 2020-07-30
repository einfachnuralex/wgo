package maps

import "errors"

//ErrorNotFound todo
var ErrorNotFound error = errors.New("not found")

//ErrorInvalidKey todo
var ErrorInvalidKey error = errors.New("invalid key")

//ErrorKeyAlreadyExists todo
var ErrorKeyAlreadyExists error = errors.New("key already exists")

// Dictionary bla
type Dictionary map[string]string

// Search bla
func (d Dictionary) Search(searchword string) (string, error) {
	result, ok := d[searchword]
	if !ok {
		return "", ErrorNotFound
	}
	return result, nil
}

// Add bla
func (d Dictionary) Add(key, value string) error {
	if key == "" {
		return ErrorInvalidKey
	}
	if _, err := d.Search(key); err != ErrorNotFound {
		return ErrorKeyAlreadyExists
	}
	d[key] = value
	return nil
}
