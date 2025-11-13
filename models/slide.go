package models

type InputSlide struct {
	underlyingSlide []rune
	seeker          int
}

var (
	SPACE   rune = []rune(" ")[0]
	NEWLINE rune = []rune("\n")[0]
	TAB     rune = []rune("\t")[0]
	EOF     rune = 0
)

var ISEMPTY = map[rune]bool{
	SPACE:   true,
	NEWLINE: true,
	TAB:     true,
}

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
	for ; ISEMPTY[nextChar] && !s.IsEnd(); nextChar = s.Next() {
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

func (s *InputSlide) SetBack(i int) {
	s.seeker -= i
}
