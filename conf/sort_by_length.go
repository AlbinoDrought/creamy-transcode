package conf

type variablesByLength []string

func (v variablesByLength) Len() int {
	return len(v)
}

func (v variablesByLength) Swap(i int, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v variablesByLength) Less(i int, j int) bool {
	return len(v[i]) > len(v[j])
}
