package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var validRegex, _ = regexp.Compile("^" + strings.Join([]string{"(\\d|[1-9]\\d+)", "(\\d|[1-9]\\d+)", "(\\d|[1-9]\\d+)"}, ".") + "$")

type SemVer struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func NewSemVer(major uint64, minor uint64, patch uint64) *SemVer {
	return &SemVer{Major: major, Minor: minor, Patch: patch}
}

func Parse(sv string) (*SemVer, error) {
	if !IsValid(sv) {
		return nil, errors.New("invalid semver-go string passed")
	}

	pv := strings.Split(".", sv)
	pints := make([]uint64, 3)
	for i, v := range pv {
		pi, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}

		pints[i] = pi
	}
	return NewSemVer(pints[0], pints[1], pints[2]), nil
}

func IsValid(sv string) bool {
	return validRegex.MatchString(sv)
}

func (s *SemVer) BumpMajor() *SemVer {
	ns := s.Clone()
	ns.Major++
	ns.Minor = 0
	ns.Patch = 0

	return ns
}

func (s *SemVer) BumpMinor() *SemVer {
	ns := s.Clone()
	ns.Minor++
	ns.Patch = 0

	return ns
}

func (s *SemVer) BumpPatch() *SemVer {
	ns := s.Clone()
	ns.Patch++

	return ns
}

func (s *SemVer) Clone() *SemVer {
	return NewSemVer(s.Major, s.Minor, s.Patch)
}

func (s *SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}
