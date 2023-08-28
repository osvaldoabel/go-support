package map


type Map struct {
	Total int
	Items map[string]interface{}
}

func New() *Map {
	return &Map{}
}


