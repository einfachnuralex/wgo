package maps

import "errors"

//ErrNotFound todo
var ErrNotFound error = errors.New("not found")

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
