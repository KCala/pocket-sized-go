package main

import (
	"reflect"
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)
			if err != nil && !testCase.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if err == nil && testCase.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}
		})

	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{handmaidsTale: 2, theBellJar: 1, oryxAndCrake: 1, janeEyre: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Joe", Books: []Book{}},
			},
			want: map[Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale}},
			},
			want: map[Book]uint{handmaidsTale: 2},
		},
	}

	for name, testCase := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(testCase.input)
			if !equalBooksCount(t, got, testCase.want) {
				t.Fatalf("lists of books different, got: %v, want: %v", got, testCase.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, handmaidsTale}},
			},
			want: []Book{handmaidsTale},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, handmaidsTale}},
				{Name: "Joe", Books: []Book{janeEyre}},
			},
			want: []Book{janeEyre, handmaidsTale},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// equalBooksCount is a helper to test the equality of two maps of books count.
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}
