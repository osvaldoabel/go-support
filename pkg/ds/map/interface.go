package map


type Map interface {
	// allocate() — Allocates the required amount of  memory for a required capacity
	Allocate()
	//  capacity() — Returns the current capacity of the queue
	GetCapacity()
	// clear() — Removes all the elements from the queue
	Clear()
	// copy() — Returns a copy of the queue
	Copy()
	// count() — Returns the number of elements in the queue
	Count()
	// isEmpty() — Returns true if the queue is empty, and false otherwise
	IsEmpty()
	
	// apply() — Updates the values by applying a callback function to each value
	Apply(callback func(key, value interface{}) interface{})

	// diff() — Returns a new map containing key-value pairs that are not present in another map
	Diff(other Map) Map
	Capacity()

	// filter() — Creates a new map using a callback function to determine which pairs to include
	Filter(callback func(key, value interface{}) bool) Map

	// get() — Returns the value for a given key
	Get(key interface{}) interface{}

	// hasKey() — Determines whether the map contains a given key
	HasKey(key interface{}) bool

	// hasValue() — Determines whether the map contains a given value
	HasValue(value interface{}) bool

	// intersect() — Creates a new map by intersecting keys with another map
	Intersect(other Map) Map

	// keys() — Returns a set of the map's keys
	Keys() Set

	// map() — Returns the result of applying a callback to each value
	Map(callback func(key, value interface{}) interface{}) Map
	

	// merge() — Returns the result of adding all given associations
	Merge(other Map) Map

	// Creates a new map by intersecting keys with another map
	Intersection(other Map) Map

	Keys() 

	// put() — Associates a key with a value
	Put(key, value interface{})

	// putAll() — Associates all key-value pairs of a traversable object or array
	PutAll(other Map)

	// reduce() — Reduces the map to a single value using a callback function
	Reduce(callback func(key, value, carry interface{}) interface{}, initial interface{}) interface{}

	// remove() — Removes and returns a value by key
	Remove(key interface{}) interface{}

	// reverse() — Reverses the map in-place
	Reverse()

	ksort() 

	// skip() — Returns the pair at a given positional index
	Skip(position int) Pair


	Ksorted()

	Last() Pair

	Sort()

	// Union
	Union(other Map) Map

	// Values
	Values() Sequence

	// Xor
	Xor(other Map) Map

	// Sum
	Sum() float64

}