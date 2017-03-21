package main

import (
	"strings"
	"unicode/utf8"
)

type NamePatternItem struct {
	pre string
	suf string
	partN int
	abbr bool
}

type NamePattern []NamePatternItem

type NameType struct {
	fcnPat NamePattern  // full canonical name pattern
	scnPat NamePattern  // short canonical name pattern
	ffnPat NamePattern  // full formal name pattern
	sfnPat NamePattern  // short formal name pattern
}

type Name struct {
	parts []string
	tp *NameType
}

func NewName(src string, tp *NameType) *Name {
	var n Name
	n.parts = strings.Split(src, "\t")
	n.tp = tp
	return &n
}

func isLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z'
}

func NewNamePattern(src string) NamePattern {
	var res NamePattern
	cnt := len(src)
	i := 0;
	for i < cnt {
		for i < cnt && src[i] == ' ' { // skip repeating spaces
			i++
		}
		pre := ""
		if i < cnt && src[i] == '^' { // a prefix starts
			i++
			preStart := i
			for i < cnt && !isLetter(src[i]) {
				i++
			}
			pre = src[preStart:i]
		}
		partN := -1
		abbr := false
		if i < cnt && isLetter(src[i]) { // a part number (letter)
			if 'A' <= src[i] && src[i] <= 'Z' {
				abbr = true
				partN = int(src[i]) - 'A'
			} else {
				partN = int(src[i]) - 'a'
			}
			i++
		}
		suf := ""
		if i < cnt { // a suffix starts
			sufStart := i
			for i < cnt && src[i] != '^' && !isLetter(src[i]) {
				i++
			}
			suf = src[sufStart:i]
		}
		if pre != "" || suf != "" || partN >= 0 {
			res = append(res, NamePatternItem{pre, suf, partN, abbr})
		}
	}
	return res
}

// Apply name pattern to name parts
func (p NamePattern) apply(parts []string) string {
	var res []byte
	for _, item := range p {
		n := item.partN
		if n < 0 || n < len(parts) && parts[n] != "" {
			res = append(res, item.pre...)
			if n >= 0 {
				if item.abbr {
					rn, _ := utf8.DecodeRuneInString(parts[n])
					res = append(res, string(rn)...)
				} else {
					res = append(res, parts[n]...)
				}
			}
			res = append(res, item.suf...)
		}
	}
	// trim tail spaces
	for len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[0:len(res)-1]
	}
	return string(res)
}

func (n *Name) FullCanonicalName() string {
	return n.tp.fcnPat.apply(n.parts)
}

func (n *Name) ShortCanonicalName() string {
	return n.tp.scnPat.apply(n.parts)
}

func (n *Name) FullFormalName() string {
	return n.tp.ffnPat.apply(n.parts)
}

func (n *Name) ShortFormalName() string {
	return n.tp.sfnPat.apply(n.parts)
}
