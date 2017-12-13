package stream

type Stream struct {
	ignoreNext   bool
	inGarbage    bool
	inGroup      bool
	groupNest    int
	GroupCount   int
	GarbageCount int
}

func (s *Stream) ProcessStream(stream string) {
	for _, v := range stream {
		if s.ignoreNext {
			s.ignoreNext = false
			continue
		}
		switch v {
		case '<':
			s.lessThan()
		case '>':
			s.greaterThan()
		case '!':
			s.bang()
		case '{':
			s.openBrace()
		case '}':
			s.closeBrace()
		default:
			s.other()
		}
	}
}

func (s *Stream) lessThan() {
	if s.inGarbage == true {
		s.GarbageCount++
	}
	s.inGarbage = true
}

func (s *Stream) greaterThan() {
	if s.inGarbage == true {
		s.inGarbage = false
	}
}

func (s *Stream) bang() {
	s.ignoreNext = true
}

func (s *Stream) openBrace() {
	if s.inGarbage {
		s.GarbageCount++
	} else {
		s.inGroup = true
		s.groupNest++
	}
}

func (s *Stream) closeBrace() {
	if s.inGarbage {
		s.GarbageCount++
	} else {
		s.GroupCount += s.groupNest
		s.groupNest--
		if s.groupNest == 0 {
			s.inGroup = false
		}
	}
}

func (s *Stream) other() {
	if s.inGarbage {
		s.GarbageCount++
	}
}
