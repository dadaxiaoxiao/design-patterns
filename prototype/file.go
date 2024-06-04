package prototype

import "fmt"

// File 文件
type File struct {
	name string
}

func (f *File) print(msg string) {
	fmt.Println(msg + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

func NewFile(name string) Inode {
	return &File{
		name: name,
	}
}
