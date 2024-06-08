package ch3

// an animal shelter, which holds only dogs and cats, operates on a strictly
// "first in, first out" basis. People must adopt either the "oldest" (based on arrival time)
// of all animals at the shelter, or they can selected whether they would prefer a dog or a cat
// (and will receive the oldest animal of that type). They cannot select which specific animal they would
// like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny,
// dequeueDog, and dequeueCat. You may use the build-int LinkeList data structure


import (
	"time"
)


const (
	dog = "dog"
	cat  = "cat"
)


type Node struct {
	name	string
	timestamp	int
	next	*Node

}


type AnimalShelter struct {
	catlist *Node
	doglist *Node
}


func (as *AnimalShelter) enqueue(name, category string) {

	newnode := &Node{name: name, timestamp: time.Now().Nanosecond()}

	if category == cat {

		if as.catlist == nil {
			as.catlist = newnode
		} else {
			n := as.catlist
			for n.next != nil {
				n = n.next
			}
			n.next = newnode
		}

	} else {

		if as.doglist == nil {
			as.doglist = newnode
		} else {
			n := as.doglist
			for n.next != nil {
				n = n.next
			}
			n.next = newnode
		}
	}
}


func (as *AnimalShelter) dequeueDog() *Node {
	if as.doglist == nil {
		return nil
	}

	dequeued := as.doglist
	as.doglist = as.doglist.next

	return dequeued

}


func (as *AnimalShelter) dequeueCat() *Node {

	if as.catlist == nil {
		return nil
	}

	dequeued := as.catlist
	as.catlist = as.catlist.next

	return dequeued

}




func (as *AnimalShelter) dequeueAny() *Node {

	firstcat := as.catlist
	firstdog := as.doglist


	if firstcat == nil && firstdog == nil {
		return nil
	}

	if firstcat == nil && firstdog != nil {
		return firstdog
	}

	if firstdog == nil && firstcat != nil {
		return firstcat
	}


	if firstcat.timestamp < firstdog.timestamp {
		return as.dequeueCat()
	}

	return as.dequeueDog()

}
















