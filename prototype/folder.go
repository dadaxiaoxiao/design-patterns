package prototype

import "fmt"

// Folder 文件夹
type Folder struct {
	// 子文件+文件夹
	children []Inode
	name     string
}

func (f *Folder) print(msg string) {
	fmt.Println(msg + f.name)
	for _, i := range f.children {
		i.print("  " + msg)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy_res := i.clone()
		tempChildren = append(tempChildren, copy_res)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func NewFolder(name string, children []Inode) Inode {
	return &Folder{
		name:     name,
		children: children,
	}
}
