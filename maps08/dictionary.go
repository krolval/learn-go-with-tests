package maps08

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (obj Dictionary) Search(key string) (string, error) {
	find, err := obj[key]
	if !err {
		return "", ErrNotFound
	}
	return find, nil
}

func (obj Dictionary) Add(key, value string) error {
	_, err := obj.Search(value)

	switch err {
	case ErrNotFound:
		obj[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (obj Dictionary) Update(key, value string) {
	obj[key] = value
}

func (obj Dictionary) Delete(key string) {
	delete(obj, key)
}
