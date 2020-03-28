package maps

import (
	"reflect"
	"testing"
)

func ErrCheck(t *testing.T, err DictionaryErr, want DictionaryErr) {
	t.Helper()
	if err != want {
		t.Errorf("Got unexpected error %v, want %v", err, want)
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name       string
		dictionary Dictionary
		term       string
		wantString string
		wantErr    DictionaryErr
	}{
		{
			"Single known word",
			Dictionary{"test": "this is just a test"},
			"test",
			"this is just a test",
			"",
		}, {
			"Unknown word",
			Dictionary{},
			"test",
			"",
			ErrWordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.dictionary.Search(tt.term); err != "" {
				ErrCheck(t, err, tt.wantErr)
			} else if !reflect.DeepEqual(got, tt.wantString) {
				t.Errorf("got %s, want %s", got, tt.wantString)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		term       string
		definition string
	}
	tests := []struct {
		name    string
		d       Dictionary
		args    args
		wantErr DictionaryErr
	}{
		{
			"Add successful",
			Dictionary{},
			args{
				term:       "foo",
				definition: "foo bar baz",
			},
			"",
		}, {
			"Word exists",
			Dictionary{"foo": ""},
			args{
				term:       "foo",
				definition: "",
			},
			ErrWordExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.add(tt.args.term, tt.args.definition); err != tt.wantErr {
				ErrCheck(t, err, tt.wantErr)
			}
			if def, ok := tt.d[tt.args.term]; !ok {
				if !reflect.DeepEqual(def, tt.args.definition) {
					t.Errorf("Wanted def: %s\nGot def: %s", tt.args.definition, def)
				}
			}
		})
	}
}
