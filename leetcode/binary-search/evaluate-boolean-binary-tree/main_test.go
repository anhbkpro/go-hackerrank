package evaluate_boolean_binary_tree

import "testing"

func Test_evaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	l := &TreeNode{Val: 1}
	rl := &TreeNode{Val: 0}
	rr := &TreeNode{Val: 1}
	r := &TreeNode{Val: 13, Left: rl, Right: rr}
	// root = [2,1,3,null,null,0,1]
	root := &TreeNode{Left: l, Right: r, Val: 2}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{root: root},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.args.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
