package names

import (
	"fmt"
	"strings"
)

type Russian struct {
	nm string  // name
	pt string  // patronymic
	sn string  // surname
}

func RussianFromString(src string) (*Russian, error) {
	parts := strings.Split(src, "\t")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Russian name source must consist of " +
			"three tab-separated parts")
	}
	res := Russian{nm: parts[0], pt: parts[1], sn: parts[2]};
	return &res, nil
}

func (n *Russian) FullCanonicalName() string {
	var res string
	if n.nm != "" {
		res = n.nm
	}
	if n.pt != "" {
		if res != "" { res += " " }
		res += n.pt
	}
	if n.sn != "" {
		if res != "" { res += " " }
		res += n.sn
	}
	return res
}

func (n *Russian) ShortCanonicalName() string {
	return "" // TODO
}

func (n *Russian) FullFormalName() string {
	return "" // TODO
}

func (n *Russian) ShortFormalName() string {
	return "" // TODO
}
