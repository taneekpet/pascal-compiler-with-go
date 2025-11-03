package models

type InputSlide struct {
	underlyingSlide []rune
	seeker          int
}

var (
	SPACE rune = []rune(" ")[0]
	EOF   rune = 0
)

func NewSlide(sourceCode string) *InputSlide {
	return &InputSlide{
		underlyingSlide: []rune(sourceCode),
		seeker:          0,
	}
}

func (s *InputSlide) GetIndex() int {
	return s.seeker
}

func (s *InputSlide) IsEnd() bool {
	return len(s.underlyingSlide) <= s.seeker
}

func (s *InputSlide) NextWithoutSpace() rune {
	nextChar := s.Next()
	for ; nextChar == SPACE && !s.IsEnd(); nextChar = s.Next() {
	}
	if s.IsEnd() {
		return rune(0)
	}
	return nextChar
}

func (s *InputSlide) PeekWithoutSpace() rune {
	defer func() {
		s.seeker--
	}()
	return s.NextWithoutSpace()
}

func (s *InputSlide) Next() rune {
	defer func() {
		s.seeker++
	}()
	return s.Peek()
}

func (s *InputSlide) Peek() rune {
	if s.IsEnd() {
		return EOF
	}
	return s.underlyingSlide[s.seeker]
}
