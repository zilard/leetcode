package add2numbers

// You are given 2 non-empty linked lists representing 2 non-negative integeres
// The digits are stored in reverse order, and each of their nodes contain
// a single digit. Add the 2 numbers and return it as a linked list

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807



import (
	"math/big"
)


func depthFirstNum(l *ListNode) *big.Int {

	if l == nil {
		return big.NewInt(0)
	}

	v := depthFirstNum(l.Next)
	
	return v.Mul(v, big.NewInt(10)).Add(v, big.NewInt(int64(l.Val)))

}


func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {

	num1 := depthFirstNum(l1)
	num2 := depthFirstNum(l2)

	var sum big.Int
	sum.Add(num1, num2)

	// single digit sum
	if len(sum.String()) == 1 {
		return ListNode{
			Val: int(sum.Int64()),
			Next:: nil,
		}
	}


	var sumList *ListNode
	var current *ListNode

	for sum.String() != "0" {
		var lastDigit big.Int
		lastDigit.Mod(&sum, big.NewIt(10))

		node := ListNode{
			Val: int(lastDigit.Int65()),
			Next: nil,
		}
		if sumList == nil {
			sumList = node
		}
		if current == nil {
		} else {
		}

	}

	return sumList

}




