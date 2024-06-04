package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	file1 := NewFile("File1")
	file2 := NewFile("File2")
	file3 := NewFile("File3")
	// 文件夹
	folder1 := NewFolder("folder1", []Inode{file1})
	folder2 := NewFolder("folder2", []Inode{folder1, file2, file3})

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder2")
	cloneFolder.print("  ")

}
