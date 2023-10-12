package main 

// Implementing search autocomplete 

type Trie struct{
	Root TrieNode 
}

type TrieNode struct{
	Children map[rune]*TrieNode 
	IsTerminal bool 
}

func (t *Trie) Insert(word string){
	curr := t.Root 
	for _, w := range word{
		if _, ok := curr.Children[w]; !ok{
			curr.Children[w] = &TrieNode{
				Children: map[rune]*TrieNode{}, 
				IsTerminal: false,
			}
		}

		curr = *curr.Children[w]
	}
	curr.IsTerminal = true 
}

func (t *Trie) Search(word string) bool{
	curr := t.Root 
	for _, w := range word{
		if _, ok := curr.Children[w]; !ok{
			return false 
		}
		curr = *curr.Children[w]
	}

	return curr.IsTerminal
}

func (t *Trie) StartsWith(prefix string) bool{
	curr := t.Root 
	for _, w := range prefix{
		if _, ok := curr.Children[w]; !ok{
			return false 
		}
		curr = *curr.Children[w]
	}
	return true 
}

type AutoCompleteSystem struct{
	Trie *Trie 
}


func (a *AutoCompleteSystem) Insert(word string, times int) []string{
	return []string{}
}

func Autocomplete(word string) []string{
	return []string{} 
}