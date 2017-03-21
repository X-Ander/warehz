package names

type Name interface {
	FullCanonicalName() string
	ShortCanonicalName() string
	FullFormalName() string
	ShortFormalName() string
}
