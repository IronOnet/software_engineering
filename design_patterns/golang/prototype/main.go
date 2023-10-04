package main 

import "fmt" 

func main(){
	file1 := &File{name: "File1"} 
	file2 := &File{name: "File2"} 
	file3 := &File{name: "File3"} 

	folder1 := &Folder{
		children: []INode{file1}, 
		name: "Folder1", 
	}

	folder2 := &Folder{
		children: []INode{folder1, file2, file3}, 
		name: "Folder2",
	}

	fmt.Println("\nPrinting hiarchy for folder2") 
	folder2.print(" ") 

	cloneFolder := folder2.clone() 
	fmt.Println("\nPrinting hieararchhy for clone folder") 
	cloneFolder.print("  ") 
}