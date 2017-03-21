package names

import (
	"testing"
)

var _ Name = (*Russian)(nil)

func TestRussian(t *testing.T) {
	name, err := RussianFromString("Иван\tИванович\tИванов")
	if err != nil {
		t.Errorf("Can't parse a correct Russian name source")
	}
	var tests = []struct{
		object string
		expected string
		got string
	}{
		{"name", "Иван", name.nm},
		{"patronymic", "Иванович", name.pt},
		{"surname", "Иванов", name.sn},
		{"full canonical name", "Иван Иванович Иванов",
			name.FullCanonicalName()},
		{"short canonical name", "И. И. Иванов",
			name.ShortCanonicalName()},
		{"full formal name", "Иванов Иван Иванович",
			name.FullFormalName()},
		{"short formal name", "Иванов И. И.",
			name.ShortFormalName()},
	}
	for _, test := range tests {
		if test.expected != test.got {
			t.Errorf("Wrong " + test.object + ": expected '" + test.expected +
				"', got '" + test.got + "'")
		}
	}
}
