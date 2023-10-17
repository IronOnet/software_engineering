package tries 

/**
	PROBLEM : Word Search 

	Given an mxn board of characters and a list of strings words, return all words 
	on the board. 

	Each word must be constructed from letters of sequentially adjacent cells, where adjacent 
	cells are horizontally or vertically neighborring. The same letter cell may not be used 
	more than once in a word. 

**/

type TrieNode struct{
	Children map[rune]*TrieNode 
	IsTerminal bool 
	Refs int 
}


func (t *TrieNode) AddWord(word string){
	curr := t 
	for _, w := range word{
		if _, ok := curr.Children[w]; !ok{
			curr.Children[w] = &TrieNode{
				Children: make(map[rune]*TrieNode), 
				IsTerminal: false, 
				Refs: 0,
			}
		}
		curr = curr.Children[w]
		curr.Refs += 1 
	}

	curr.IsTerminal = true 
}

func (t *TrieNode) RemoveWord(word string){
	curr := t 
	curr.Refs += 1 
	for i := 0; i < len(word); i++{
		c := rune(word[i]) 
		if _, ok := curr.Children[c]; ok{
			curr = curr.Children[c] 
			curr.Refs -= 1 
		}
	}
}

func findWords(board [][]byte, words []string) []string{
	root := &TrieNode{
		Children: make(map[rune]*TrieNode),
	}

	for _, w := range words{
		root.AddWord(w)
	}

	ROWS, COLS := len(board), len(board[0]) 

	result := make(map[string]bool) 
	visit := make(map[int]bool) 

	var dfs func(r, c int, root *TrieNode, word string)  

	dfs = func(r, c int, node *TrieNode, word string){
		if 
			r < 0 || r >= ROWS || 
			c < 0 || c >= COLS || 
			node.Children[rune(board[r][c])] == nil || 
			node.Children[rune(board[r][c])].Refs < 1 || 
			visit[r*COLS + c]{
				return 
			}

		visit[r*COLS+c] = true 
		word += string(board[r][c]) 
		if node.IsTerminal{
			node.IsTerminal = false 
			result[word] = true 
			root.RemoveWord(word)
		}

		dfs(r+1, c, node, word) 
		dfs(r-1, c, node, word) 
		dfs(r, c+1, node, word) 
		dfs(r, c-1, node, word) 
		visit[r*COLS+c] = false 
	}

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			dfs(r, c, root, "")
		}
	}

	return list(result)
}

func list(mapping map[string]bool) []string{
	var stringList []string 
	for s, _ := range mapping{
		stringList = append(stringList, s) 
	}
	return stringList
}

