package ch3

import (
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestAnimalShelter(t *testing.T) {

	as := new(AnimalShelter)

	as.enqueue("Brucy", dog)
	as.enqueue("Ginger", cat)
	as.enqueue("Pepper", cat)
	as.enqueue("Berry", cat)
	as.enqueue("Tommy", dog)
	as.enqueue("Phil", dog)

	if as.doglist != nil && as.catlist != nil { 
		assert.Equal(t, as.dequeueAny().name, "Brucy")
		assert.Equal(t, as.dequeueCat().name, "Ginger")
		assert.Equal(t, as.dequeueDog().name, "Tommy")
		assert.Equal(t, as.dequeueAny().name, "Pepper")
	}

}

