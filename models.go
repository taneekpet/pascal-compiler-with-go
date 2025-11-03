package main

type inputSlide struct {
	underlyingSlide []rune
	seeker          int
}

func NewSlide(sourceCode string) *inputSlide {
	return &inputSlide{
		underlyingSlide: []rune(sourceCode),
		seeker:          0,
	}
}

func (s *inputSlide) IsEnd() bool {
	return len(s.underlyingSlide) <= s.seeker
}

func (s *inputSlide) NextWithoutSpace() string {
	nextChar := s.Next()
	for ; nextChar == " " && !s.IsEnd(); nextChar = s.Next() {
	}
	if s.IsEnd() {
		return ""
	}
	return nextChar
}

func (s *inputSlide) PeekWithoutSpace() string {
	defer func(tmp int) {
		s.seeker = tmp
	}(s.seeker)
	return s.NextWithoutSpace()
}

func (s *inputSlide) Next() string {
	defer func() {
		s.seeker++
	}()
	return s.Peek()
}

func (s *inputSlide) Peek() string {
	if s.IsEnd() {
		return ""
	}
	return string(s.underlyingSlide[s.seeker])
}
