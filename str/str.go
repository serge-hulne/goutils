package str

import "strings"

type Str struct {
	Str string
}

func New(s string) Str {
	return Str{s}
}

func (s Str) String() string {
	return s.Str
}

func (s *Str) Capitalize() *Str {
	s.Str = strings.ToUpper(s.Str)
	return s
}
func (s *Str) ReplaceAll(c1 string, c2 string) *Str {
	s.Str = strings.Replace(s.Str, c1, c2, -1)
	return s
}
