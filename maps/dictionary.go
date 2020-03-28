package maps

const (
	ErrWordNotFound = DictionaryErr("unable to find word in dictionary")
	ErrWordExists   = DictionaryErr("word already exists in dictionary")
)

type DictionaryErr string

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, DictionaryErr) {
	if m, ok := d[term]; !ok {
		return "", ErrWordNotFound
	} else {
		return m, ""
	}
}

func (d Dictionary) add(term, definition string) DictionaryErr {
	if _, ok := d[term]; ok {
		return ErrWordExists
	}
	d[term] = definition
	return ""
}
