package protoType

import (
	"fmt"
	"testing"
)

func TestFile_Clone(t *testing.T) {
	file1 := &File{Name: "file1"}
	file2 := &File{Name: "file2"}
	file3 := &File{Name: "file3"}
	folder1 := &Folder{
		Name: "folder1",
		Sub:  []Inode{file1},
	}
	folder2 := &Folder{
		Name: "folder2",
		Sub:  []Inode{folder1, file2, file3},
	}
	fmt.Println("Printing folder2")
	folder2.Print(" ")
	cloneFolder2 := folder2.Clone()
	fmt.Println("Printing clone folder2")
	cloneFolder2.Print(" ")
}
