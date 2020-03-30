package maps

import (
	"reflect"
	"testing"
)

func ErrCheck(t *testing.T, err error, want error) {
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
		wantErr    error
	}{
		{
			"Single known word",
			Dictionary{"test": "this is just a test"},
			"test",
			"this is just a test",
			nil,
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
			if got, err := tt.dictionary.Search(tt.term); err != nil {
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
		wantErr error
	}{
		{
			"New word",
			Dictionary{},
			args{
				term:       "foo",
				definition: "foo bar baz",
			},
			nil,
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
			if err := tt.d.Add(tt.args.term, tt.args.definition); err != tt.wantErr {
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

func TestUpdate(t *testing.T) {
	type args struct {
		term       string
		definition string
	}
	tests := []struct {
		name    string
		d       Dictionary
		args    args
		want    error
		wantDef string
	}{
		{
			"word exists",
			Dictionary{"foo": ""},
			args{
				term:       "foo",
				definition: "bar",
			},
			nil,
			"bar",
		}, {
			"word doesn't exist",
			Dictionary{},
			args{
				term:       "foo",
				definition: "bar",
			},
			ErrWordDoesNotExist,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			term := tt.args.term
			if got := tt.d.Update(term, tt.args.definition); got != tt.want {
				t.Errorf("update() = %v, want %v", got, tt.want)
			}
			if definition := tt.d[term]; !reflect.DeepEqual(tt.wantDef, definition) {
				t.Errorf("expected definition %s, got %s", tt.wantDef, definition)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		d       Dictionary
		term    string
		wantErr error
	}{
		{
			"Word exists",
			Dictionary{"foo": ""},
			"foo",
			nil,
		}, {
			"Word does not exist",
			Dictionary{},
			"foo",
			ErrWordDoesNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Delete(tt.term); err != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, ok := tt.d[tt.term]; ok {
				t.Errorf("term %s exists in dictionary", tt.term)
			}
		})
	}
}
