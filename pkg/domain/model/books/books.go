package books

type Books struct {
	id   int32
	name string
}

func (b Books) Id() int32 {
	return b.id
}
func (b Books) Name() string {
	return b.name
}
