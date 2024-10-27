package cell

import (
	"fmt"
	"testing"

	"math/rand"
)

func TestPriorityQueue(t *testing.T) {

	// Create Cells with random weights record the minimum weight, second minimum weight, max weight
	c1 := CreateBaseCell(0, 0)
	c1.SetWeight(5)

	pq := PriorityQueue{}
	pq.Push(c1)

	c2 := CreateBaseCell(0, 1)
	c2.SetWeight(10)

	pq.Push(c2)

	if pq.Peek() != c1 {
		t.Error("Wrong")
	}

	pq.Pop()
	if pq.Peek() != c2 {
		t.Error("Wrong")
	}

	// min
	minWeight := 11
	numMin := 0

	// second lowest
	minWeight2 := 12
	numMin2 := 0

	pq2 := PriorityQueue{}

	for i := 0; i < 20; i++ {

		var weight int
		if i%2 == 0 {

			weight = 1
		} else {
			weight = 2
		}

		if weight == minWeight {
			numMin++
		} else if weight == minWeight2 {
			numMin2++
		} else if weight < minWeight {
			minWeight2 = minWeight
			numMin2 = 1
			minWeight = weight
			numMin = 1
		} else if weight < minWeight2 {
			minWeight2 = weight
			numMin2 = 1
		}

		c1 := CreateBaseCell(0, 0)
		c1.SetWeight(weight)
		pq2.Push(c1)
	}

	if len(pq2) != 20 {
		t.Errorf("pq2 does not have the correct length; got: %d, expected: 10", len(pq2))
	}

	fmt.Println("numMin: ", numMin)
	for i := 0; i < numMin; i++ {
		w := pq2.Pop().Weight()
		if w != minWeight {
			t.Errorf("Incorrect minimum weight; got: %d, expected: %d", w, minWeight)
		}
	}

	for i := 0; i < numMin2; i++ {
		w := pq2.Pop().Weight()
		if w != minWeight2 {
			t.Errorf("Incorrect minimum weight 2; got: %d, expected: %d", w, minWeight2)
		}
	}

	pq3 := PriorityQueue{}

	c3 := CreateBaseCell(0, 0)
	c3.SetWeight(0)
	pq3.Push(c3)

	c4 := CreateBaseCell(0, 0)
	c4.SetWeight(10)
	pq3.Push(c4)

	c5 := CreateBaseCell(0, 0)
	c5.SetWeight(5)
	pq3.Push(c5)

	fmt.Println(pq3.Pop().Weight())
	fmt.Println(pq3.Pop().Weight())
	fmt.Println(pq3.Pop().Weight())

	// min
	minWeight = 11
	numMin = 0

	// second lowest
	minWeight2 = 12
	numMin2 = 0

	for i := 0; i < 20; i++ {

		weight := rand.Intn(4) + 1

		if weight == minWeight {
			numMin++
		} else if weight == minWeight2 {
			numMin2++
		} else if weight < minWeight {
			minWeight2 = minWeight
			numMin2 = 1
			minWeight = weight
			numMin = 1
		} else if weight < minWeight2 {
			minWeight2 = weight
			numMin2 = 1
		}

		c1 := CreateBaseCell(0, 0)
		c1.SetWeight(weight)
		pq3.Push(c1)
	}

	if len(pq3) != 20 {
		t.Errorf("pq3 does not have the correct length; got: %d, expected: 10", len(pq3))
	}

	for i := 0; i < numMin; i++ {
		w := pq3.Pop().Weight()
		if w != minWeight {
			t.Errorf("Incorrect minimum weight; got: %d, expected: %d", w, minWeight)
		}
	}

	for i := 0; i < numMin2; i++ {
		w := pq3.Pop().Weight()
		if w != minWeight2 {
			t.Errorf("Incorrect minimum weight 2; got: %d, expected: %d", w, minWeight2)
		}
	}
}
