package controllers

type ScoreWatcher func(int)

type Score struct {
	value    int
	watchers []ScoreWatcher
}

func NewScore() *Score {
	return &Score{
		value: 0,
	}
}

func (s *Score) Value() int {
	return s.value
}

func (s *Score) Add(points int) {
	s.value += points
	for _, w := range s.watchers {
		w(s.value)
	}
}

func (s *Score) AddWatcher(watcher ScoreWatcher) {
	s.watchers = append(s.watchers, watcher)
	watcher(s.value)
}
