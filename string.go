package jstypes

import (
	"slices"
	"strings"
)

type String string

func (str *String) At(index int) String {
	return String(strings.Split(string(*str), "")[index])
}

func (str *String) Concat(strings ...String) String {
	s := *str
	for _, i := range strings {
		s += i
	}
	return s
}

func (str *String) Includes(s String) bool {
	return strings.Contains(string(*str), string(s))
}

func (str *String) ReplaceAll(old String, new String) String {
	return String(strings.ReplaceAll(string(*str), string(old), string(new)))
}

func (str *String) Replace(old String, new String) String {
	return String(strings.Replace(string(*str), string(old), string(new), 1))
}

func (str *String) Reverse() String {
	sp := strings.Split(string(*str), "")
	slices.Reverse(sp)
	return String(strings.Join(sp, ""))
}

func (str *String) ToLowerCase() String {
	return String(strings.ToLower(string(*str)))
}

func (str *String) Length() int {
	return len(strings.Split(string(*str), ""))
}

func (str *String) ToUpperCase() String {
	return String(strings.ToUpper(string(*str)))
}

func (str *String) Trim() String {
	return String(strings.TrimSpace(string(*str)))
}

func (str *String) Substring(start int, e ...int) String {
	end := str.Length() - 1
	if len(e) == 1 {
		end = int(e[0])
	}
	sp := strings.Split(string(*str), "")
	return String(string(strings.Join(sp[start:end], "")))
}

func (str *String) StartsWith(s String) bool {
	return strings.HasPrefix(string(*str), string(s))
}

func (str *String) EndsWith(s String) bool {
	return strings.HasSuffix(string(*str), string(s))
}

func (str *String) PadEnd(length int, pad String) String {
	if str.Length() >= length {
		return *str
	}
	s := *str
	for ; s.Length() < length; s += pad {
	}
	return s
}

func (str *String) PadStart(length int, pad String) String {
	if str.Length() >= length {
		return *str
	}
	s := *str
	for ; s.Length() < length; s = pad + s {
	}
	return s
}

func (str *String) Repeat(times int) (st String) {
	s := *str
	for i := 0; i < times; i++ {
		st += s
	}
	return
}

func (str *String) Split(splitter String) (s []String) {
	sp := strings.Split(string(*str), string(splitter))
	for _, i := range sp {
		s = append(s, String(i))
	}
	return
}

func (str *String) String() string {
	return string(*str)
}
