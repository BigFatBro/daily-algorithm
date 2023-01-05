package main

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	isEnd bool
	next  map[string]Node
}

func main() {
	floder := []string{
		"/a", "/a/b", "/c/d", "/c/d/e", "/c/f",
	}
	fmt.Println("Output:", removeSubfolders(floder))

}

func removeSubfolders(folder []string) []string {
	root := Node{
		isEnd: false,
		next:  map[string]Node{},
	}
	result := []string{}
	sort.Strings(folder)
	for _, v := range folder {
		dirs := strings.Split(v, "/")
		node := root
		for i := 0; i < len(dirs); i++ {
			curLevelNode, ok := node.next[dirs[i]]
			if ok && curLevelNode.isEnd == true {
				//存在父目录
				break
			} else if ok && curLevelNode.isEnd == false {
				//目录已被创建，但这一级目录不存在于folder中
				if i != len(dirs)-1 {
					node = node.next[dirs[i]]
				} else {
					curLevelNode.isEnd = true
					result = append(result, v)
					break
				}

			} else {
				//创建目录
				if i != len(dirs)-1 {
					node.next[dirs[i]] = Node{
						isEnd: false,
						next:  map[string]Node{},
					}
					node = node.next[dirs[i]]
				} else {
					node.next[dirs[i]] = Node{
						isEnd: true,
						next:  map[string]Node{},
					}
					result = append(result, v)
					break
				}

			}
		}

	}
	return result
}
