package queue

type Queue interface {
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
	// peek() — Returns the element at the front of the queue without removing it
	Peek()
	// pop() — Removes and returns the element at the front of the queue
	Pop()
	// push() — Adds an element to the back of the queue
	Push()
	// toSlice() — Returns a slice containing all the elements of the queue
	ToSlice()
}
