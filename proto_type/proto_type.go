package protoType

import "fmt"

type Inode interface {
	Clone() Inode
	Print(string2 string)
}

type File struct {
	Name string
}

func (f *File) Clone() Inode {
	return &File{
		Name: f.Name + "_clone",
	}
}

func (f *File) Print(s string) {
	fmt.Println(s + f.Name)
}

type Folder struct {
	Name string
	Sub  []Inode
}

func (f *Folder) Clone() Inode {
	var sub Folder
	sub.Name = f.Name + "_clone"
	var child []Inode
	for _, v := range f.Sub {
		child = append(child, v.Clone())
	}
	sub.Sub = child
	return &sub
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, v := range f.Sub {
		v.Print(s + s)
	}
}
