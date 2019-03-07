package xmld

type stack struct {
	data []interface{}
}

func (s *stack)empty() bool  {
	return len(s.data) == 0
}

func (s *stack)push(value interface{})  {
	s.data = append(s.data, value)
}

func (s *stack)pop() interface{}  {
	value := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = nil
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *stack)peek() interface{}  {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}
