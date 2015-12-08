package multifilter

type FilterType func(param ...interface{})

type MultiFilter struct {
	Filters []FilterType
}

func NewMultiFilter() *MultiFilter {
	return &MultiFilter{
		Filters: make([]FilterType, 0, 10),
	}
}

func (multiFilter *MultiFilter) AddFilter(filter FilterType) {
	multiFilter.Filters = append(multiFilter.Filters, filter)
}

func (multiFilter *MultiFilter) Call(param ...interface{}) {
	for _, filter := range multiFilter.Filters {
		filter(param...)
	}
}
