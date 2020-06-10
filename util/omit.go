package util

import "strconv"

type StrOmit struct {
	Value    string
	is_Valid bool
}

func (s *StrOmit) Set(val string) *StrOmit {
	s.Value = val

	return s
}

func (s *StrOmit) ChValid(isValid bool) {
	s.is_Valid = isValid
}

func (s *StrOmit) IsValid() bool {
	if s.is_Valid {
		return true
	}

	return s.Value != ""
}

func (s *StrOmit) Get() string {
	return s.Value
}

func (s *StrOmit) Int() int {

	if !s.IsValid() {
		return 0
	}

	v := s.Get()

	i, _ := strconv.Atoi(v)

	return i
}

func (s *StrOmit) Bool() bool {
	if !s.IsValid() {
		return false
	}

	v := s.Get()

	if v == "" || v == "false" || v == "FALSE" {
		return false
	}

	return true
}
