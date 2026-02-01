package main

import "fmt"

// ============ Visitor Interface ============

type Visitor interface {
	VisitFile(file *File)
	VisitFolder(folder *Folder)
}

// ============ Element Interface ============

type Element interface {
	Accept(visitor Visitor)
}

// ============ Concrete Elements ============

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Accept(visitor Visitor) {
	visitor.VisitFile(f)
}

type Folder struct {
	name     string
	elements []Element
}

func NewFolder(name string) *Folder {
	return &Folder{name: name, elements: make([]Element, 0)}
}

func (f *Folder) AddElement(element Element) {
	f.elements = append(f.elements, element)
}

func (f *Folder) Accept(visitor Visitor) {
	visitor.VisitFolder(f)
	for _, element := range f.elements {
		element.Accept(visitor)
	}
}

// ============ Concrete Visitors ============

type SizeCalculator struct {
	totalSize int
}

func NewSizeCalculator() *SizeCalculator {
	return &SizeCalculator{totalSize: 0}
}

func (s *SizeCalculator) VisitFile(file *File) {
	s.totalSize += file.size
}

func (s *SizeCalculator) VisitFolder(folder *Folder) {
	// Do nothing for folder size
}

func (s *SizeCalculator) GetTotalSize() int {
	return s.totalSize
}

type NamePrinter struct{}

func NewNamePrinter() *NamePrinter {
	return &NamePrinter{}
}

func (n *NamePrinter) VisitFile(file *File) {
	fmt.Printf("File: %s\n", file.name)
}

func (n *NamePrinter) VisitFolder(folder *Folder) {
	fmt.Printf("Folder: %s\n", folder.name)
}

// ============ Main Function ============

func main() {
	fmt.Println("===== Visitor Pattern Example: File System =====\n")

	// Create file system elements
	file1 := NewFile("file1.txt", 100)
	file2 := NewFile("file2.txt", 200)
	folder1 := NewFolder("Folder1")
	folder2 := NewFolder("Folder2")

	folder1.AddElement(file1)
	folder1.AddElement(file2)
	folder2.AddElement(folder1)

	// Calculate total size
	sizeCalculator := NewSizeCalculator()
	folder2.Accept(sizeCalculator)
	fmt.Printf("Total size of all files: %d bytes\n\n", sizeCalculator.GetTotalSize())

	// Print names of all elements
	namePrinter := NewNamePrinter()
	folder2.Accept(namePrinter)
}
