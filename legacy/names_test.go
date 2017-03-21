package main

import (
	"testing"
)

func TestNames (t *testing.T) {
	var russianName NameType
	russianName.fcnPat = NewNamePattern("a b c")
	russianName.scnPat = NewNamePattern("A. B. c")
	russianName.ffnPat = NewNamePattern("c a b")
	russianName.sfnPat = NewNamePattern("c A. B.")
	name1 := NewName("Иван\tИванович\tИванов", &russianName)
	name2 := NewName("Пётр\t\tПетров", &russianName)
	name3 := NewName("Сидор\tСидорович", &russianName)
	name4 := NewName("Николай", &russianName)
	var tests = []struct{
		object string
		expected string
		got string
	}{
		{"full canonical name", "Иван Иванович Иванов",
			name1.FullCanonicalName()},
		{"short canonical name", "И. И. Иванов",
			name1.ShortCanonicalName()},
		{"full formal name", "Иванов Иван Иванович",
			name1.FullFormalName()},
		{"short formal name", "Иванов И. И.",
			name1.ShortFormalName()},
		{"full canonical name", "Пётр Петров",
			name2.FullCanonicalName()},
		{"short canonical name", "П. Петров",
			name2.ShortCanonicalName()},
		{"full formal name", "Петров Пётр",
			name2.FullFormalName()},
		{"short formal name", "Петров П.",
			name2.ShortFormalName()},
		{"full canonical name", "Сидор Сидорович",
			name3.FullCanonicalName()},
		{"short canonical name", "С. С.",
			name3.ShortCanonicalName()},
		{"full formal name", "Сидор Сидорович",
			name3.FullFormalName()},
		{"short formal name", "С. С.",
			name3.ShortFormalName()},
		{"full canonical name", "Николай",
			name4.FullCanonicalName()},
		{"short canonical name", "Н.",
			name4.ShortCanonicalName()},
		{"full formal name", "Николай",
			name4.FullFormalName()},
		{"short formal name", "Н.",
			name4.ShortFormalName()},
	}
	for _, test := range tests {
		if test.expected != test.got {
			t.Errorf("Wrong " + test.object + ": expected '" + test.expected +
				"', got '" + test.got + "'")
		}
	}
}
