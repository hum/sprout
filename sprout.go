package sprout

type Sprout struct {
	reddit *Reddit
}

type Post struct {
	Name   string
	Author string
	Link   string
}

func New() *Sprout {
	return &Sprout{}
}

func (s *Sprout) Reddit() *Reddit {
	if s.reddit == nil {
		return &Reddit{}
	}
	return s.reddit
}
