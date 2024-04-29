package serializeAndDeserializeBinaryTree

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/AnkushinDaniil/leetcode-go/trees"
)

type TreeNode = trees.TreeNode

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	var buffer bytes.Buffer

	var rec func(*TreeNode)
	rec = func(tn *TreeNode) {
		if tn == nil {
			buffer.WriteString("nil,")
		} else {
			buffer.WriteString(strconv.Itoa(tn.Val))
			buffer.WriteByte(',')

			rec(tn.Left)
			rec(tn.Right)
		}
	}
	rec(root)

	return buffer.String()
}

func (c *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")

	i := 0

	var rec func() *TreeNode
	rec = func() *TreeNode {
		if arr[i] == "nil" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(arr[i])
		i++
		return &TreeNode{
			Val:   val,
			Left:  rec(),
			Right: rec(),
		}
	}

	return rec()
}
