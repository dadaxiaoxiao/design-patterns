package composite

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:       name,
		components: make([]Component, 0),
	}
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
