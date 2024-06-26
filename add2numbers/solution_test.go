package add2numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "add two numbers",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "add two numbers",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
								Next: nil,
							},
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 1,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(tt.args.l1, tt.args.l2))
			assert.Equal(t, tt.want, addTwoNumbers2(tt.args.l1, tt.args.l2))
		})
	
	}

}


