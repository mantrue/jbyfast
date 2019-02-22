package go_testing_examples

import (
	"fmt"
	"testing"
)

type cate struct {
	id   int
	name string
	pid  int
	//son  *cate
}

func Run(t *testing.T) {
	allCate := []cate{
		cate{1, "a", 0},
		cate{2, "b", 0},
		cate{3, "c", 0},
		cate{4, "aa", 1},
		cate{5, "bb", 2},
		cate{6, "cc", 3},
		cate{7, "aaa", 4},
		cate{8, "bbb", 5},
		cate{9, "ccc", 6},
	}

	arr := superCategoryTree(allCate, 0)
	fmt.Println(arr)
}

type cateTree struct {
	id   int
	name string
	pid  int
	son  []cateTree
}

//递归实现
func superCategoryTree(allCate []cate, pid int) []cateTree {
	var arr []cateTree
	for _, v := range allCate {
		if pid == v.pid {
			ctree := cateTree{}
			ctree.id = v.id
			ctree.pid = v.pid
			ctree.name = v.name

			sonCate := superCategoryTree(allCate, v.id)

			ctree.son = sonCate

			arr = append(arr, ctree)
		}
	}
	return arr
}
