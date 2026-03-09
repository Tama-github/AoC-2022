package fileh

import (
	"fmt"
	"strconv"
	"strings"
)

//"strconv"
// "log"
// "math"

type IHirarchyElem interface {
	getWeight(hierarchy *Hierarchy) int
	Print()
}

type File struct {
	name   string
	parent *Folder
	weight int
}

func (f File) getWeight(hierarchy *Hierarchy) int {
	return f.weight
}

func (f *File) Print() {
	fmt.Printf(" ")
	tmp := f.parent
	for tmp.parent != nil {
		fmt.Printf(" ")
		tmp = tmp.parent
	}
	fmt.Printf("- %s (file, size=%v)\n", f.name, f.weight)
}

type Folder struct {
	name    string
	parent  *Folder
	weight  int
	content []IHirarchyElem
}

func (f *Folder) getWeight(hierarchy *Hierarchy) int {
	totalWeight := 0
	if f.weight > 0 {
		return f.weight
	}
	for _, h := range f.content {
		var ih IHirarchyElem
		ih = f
		if h == ih {
			break
		}
		weight := h.getWeight(hierarchy)
		totalWeight += weight
	}

	// ex1
	if totalWeight <= hierarchy.cap {
		hierarchy.sum += totalWeight
	}
	f.weight = totalWeight

	return totalWeight
}

func (f *Folder) makeCandidateList(hie *Hierarchy, minW int) {
	if f.getWeight(hie) >= minW {
		hie.candidates = append(hie.candidates, f.weight)
	}
	for _, h := range f.content {
		var ih IHirarchyElem
		ih = f
		if h == ih {
			break
		}
		switch v := h.(type) {
		case *Folder:
			v.makeCandidateList(hie, minW)
		}
	}
}

func (f *Folder) Print() {
	//fmt.Println(f.content)
	tmp := f
	for tmp.parent != nil {
		fmt.Printf(" ")
		tmp = tmp.parent
	}

	fmt.Printf("- %s (dir)\n", f.name)
	for _, h := range f.content {
		//fmt.Printf("TEST\n")
		h.Print()
	}
}

func (f *Folder) Find(s string) IHirarchyElem {
	var res IHirarchyElem
	//fmt.Println(f.content)
	for _, f := range f.content {
		switch v := f.(type) {
		case *Folder:
			if v.name == s {
				res = v
			}
		case *File:
			if v.name == s {
				res = v
			}
		}
	}
	return res
}

type Hierarchy struct {
	root    *Folder
	current *Folder
	// To find solution ex1
	cap int
	sum int

	// To find solution ex2
	Capacity    int
	SpaceNeeded int
	candidates  []int
}

func (h Hierarchy) Print() {
	h.root.Print()
}

func (h *Hierarchy) FindFolderToDelete() int {
	sum := h.root.getWeight(h)

	needToDelete := h.SpaceNeeded - (h.Capacity - sum)
	h.root.makeCandidateList(h, needToDelete)

	// search in list
	fmt.Printf("\nThe total size of the sys is %vo, we need to delete %vo\n", sum, needToDelete)
	minS := h.candidates[0]
	for _, size := range h.candidates {
		opres := size - needToDelete
		fmt.Printf("Prev candidate : %vo(%v), New candidate : %vo(%v)\n", minS, minS-needToDelete, size, size-needToDelete)
		if minS-needToDelete < 0 || (opres >= 0 && opres < minS-needToDelete) {
			minS = size
		}
	}
	return minS
}

func (h *Hierarchy) GetCurrentName() string {
	return h.current.name
}

func (h *Hierarchy) mkdir(name string) *Folder {
	//var ih IHirarchyElem
	//fmt.Printf("creating %s with parent %s\n", name, h.current.name)
	f := Folder{name: name, parent: h.current, weight: -1, content: []IHirarchyElem{}}
	//ih = f
	//h.Print()
	h.current.content = append(h.current.content, &f)
	//append(h.current.content, f)
	//h.Print()
	//fmt.Printf("Just made %s its parent is %s\n", f.name, f.parent.name)
	//fmt.Println(f.parent.content)

	return &f
}

func (h *Hierarchy) mkfile(name string, weight int) {
	//var ih IHirarchyElem
	f := File{name: name, weight: weight, parent: h.current}
	//ih = &f
	//fmt.Printf("add %s to %s content\n", f.name, h.current.name)
	h.current.content = append(h.current.content, &f)
	//fmt.Println(h.current.content)
}

func (h *Hierarchy) ComputeSumFolderSizeWithSizeBellow(w int) int {
	h.sum = 0
	h.cap = w
	h.root.getWeight(h)
	return h.sum
}

func (h *Hierarchy) Cd(arg string) {
	fmt.Printf("try to go in folder %s\n", arg)
	switch arg {
	case "/":
		fmt.Printf("go in %s\n", "/")
		h.current = h.root
	case "..":
		//fmt.Printf("current folder : %s\n", h.current.name)
		fmt.Printf("go to parent %s\n", h.current.parent.name)
		h.current = h.current.parent
	default:
		founded := h.current.Find(arg)
		tmp, _ := founded.(*Folder)
		h.current = tmp
		if founded == nil {
			fmt.Printf("Folder %s not found\n", arg)
		} else {
			fmt.Printf("go to %s\n", h.current.name)

		}

	}
}

func (h *Hierarchy) Ls(args []string) {
	for _, todo := range args {
		fmt.Printf("treating : %s\n", todo)
		line := strings.Split(todo, " ")
		switch line[0] {
		case "dir":
			h.mkdir(line[1])
		default:
			weight, _ := strconv.Atoi(line[0])
			h.mkfile(line[1], weight)
		}
	}
}

func CreateHierarchy() *Hierarchy {
	h := &Hierarchy{}
	h.root = &Folder{name: "/", parent: nil, content: []IHirarchyElem{}}
	h.current = h.root
	return h
}
