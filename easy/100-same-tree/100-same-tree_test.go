package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "both trees nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "p nil, q not nil",
			p:    nil,
			q:    &TreeNode{Val: 1},
			want: false,
		},
		{
			name: "p not nil, q nil",
			p:    &TreeNode{Val: 1},
			q:    nil,
			want: false,
		},
		{
			name: "single node equal",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 5},
			want: true,
		},
		{
			name: "single node different",
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 7},
			want: false,
		},
		{
			name: "simple trees equal",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "simple trees different structure",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: nil,
			},
			q: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "different values same structure",
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 3},
			},
			want: false,
		},
		{
			name: "deep trees equal",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{Val: 5},
			},
			want: true,
		},
		{
			name: "deep trees different",
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
