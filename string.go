package jstypes

import (
	"strings"
)

type String string

func StringFrom(str string) *String {
	s := String(str)
	return &s
}

func (str *String) At(index int) String {
	arr := Array[string](strings.Split(str.String(), ""))
	return String(*arr.At(index))
}

func (str *String) Concat(strings ...String) String {
	s := *str
	for _, i := range strings {
		s += i
	}
	return s
}

func (str *String) Includes(s String) bool {
	return strings.Contains(str.String(), s.String())
}

func (str *String) ReplaceAll(old String, new String) String {
	return String(strings.ReplaceAll(str.String(), old.String(), new.String()))
}

func (str *String) Replace(old String, new String) String {
	return String(strings.Replace(str.String(), old.String(), new.String(), 1))
}

func (str *String) Reverse() String {
	sp := Array[string](strings.Split(str.String(), ""))
	sp.Reverse()
	return String(strings.Join(sp, ""))
}

func (str *String) ToLowerCase() String {
	return String(strings.ToLower(str.String()))
}

func (str *String) Length() int {
	arr := Array[string](strings.Split(str.String(), ""))
	return arr.Length()
}

func (str *String) ToUpperCase() String {
	return String(strings.ToUpper(str.String()))
}

func (str *String) Trim() String {
	return String(strings.TrimSpace(str.String()))
}

func (str *String) Substring(start int, e ...int) String {
	end := str.Length() - 1
	if len(e) == 1 {
		end = int(e[0])
	}
	sp := strings.Split(str.String(), "")
	return String(strings.Join(sp[start:end], ""))
}

func (str *String) StartsWith(s String) bool {
	return strings.HasPrefix(str.String(), s.String())
}

func (str *String) EndsWith(s String) bool {
	return strings.HasSuffix(str.String(), s.String())
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

func (str *String) Split(splitter String) (s Array[String]) {
	sp := strings.Split(str.String(), splitter.String())
	for _, i := range sp {
		s.Push(String(i))
	}
	return
}

func (str *String) String() string {
	return string(*str)
}
