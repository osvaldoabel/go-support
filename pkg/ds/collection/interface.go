package collection

type Collection interface {
	Get(index string) interface{}
	Clear()
	Copy() Collection
	IsEmpty() bool
	ToSlice() []interface{}
	Count() int

	Filter(func(interface{}) bool) Collection
	Map(func(interface{}) ItemCollection) Collection
	Add(ItemCollection)
}
