package collection

import "fmt"

type Callback func(interface{}) bool
type CallbackMap func(interface{}) interface{}, error

type ItemsCollection map[string]interface{}

type collection struct {
	Total int
	Items ItemsCollection
}

func NewCollection() Collection {
	return &collection{
		Items: make(ItemsCollection),
		Total: 0,
	}

}
func initilize() map[string]interface{} {
	return make(map[string]interface{})

}

func (c *collection) Copy() Collection {
	return &collection{
		Items: c.Items,
		Total: c.Total,
	}
}

func (c *collection) IsEmpty() bool {
	return len(c.Items) == 0
}

func (c *collection) ToSlice() []interface{} {
	return c.Items
}

func (c *collection) Count() int {
	return len(c.Items)
}

func (c *collection) Filter(f Callback) Collection {
	newCollection := NewCollection()
	for _, item := range c.Items {
		if !f(item) {
			continue
		}

		newCollection.Add(item)
	}
	return newCollection
}

func (c *collection) Add(item interface{}) {
	nextPos := fmt.Sprintf("%s", c.Items)
	c.Items[nextPos] = item
}

func (c *collection) Get(index string) interface{} {
	return c.Items[index]
}

func (c *collection) Clear() {
	c.Items = initilize()
}

func (c *collection) Map(f CallbackMap) Collection {
	newCollection := NewCollection()
	for _, item := range c.Items {
		result, err := f(item)
		if err != nil {
			continue
		}

		if result == nil {
			continue
		}

		newCollection.Add(result)
	}

	return newCollection
}
