package mydict

import (
	"errors"
)

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errExist = errors.New("exist")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if ok {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = value
	} else if err == nil {
		return errExist
	}

	return nil
}

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errNotFound
	} else {
		d[word] = value
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
