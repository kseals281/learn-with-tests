package maps

const (
	ErrWordNotFound     = DictionaryErr("unable to find word in dictionary")
	ErrWordExists       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exist within the dictionary")
)

type DictionaryErr string

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(term string) (string, error) {
	if m, ok := d[term]; !ok {
		return "", ErrWordNotFound
	} else {
		return m, nil
	}
}

func (d Dictionary) Add(term, definition string) error {
	if _, ok := d[term]; ok {
		return ErrWordExists
	}
	d[term] = definition
	return nil
}

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)
	if err != nil {
		return ErrWordDoesNotExist
	}
	d[term] = definition
	return nil
}

func (d Dictionary) Delete(term string) error {
	_, err := d.Search(term)
	if err != nil {
		return ErrWordDoesNotExist
	}
	delete(d, term)
	return nil
}
