# Trie data structure

```go

type Trie struct{
  Root Trie
  Children map[rune]*Trie  
  IsTerminal bool

}

func (t *Trie) Insert(word string){
  curr := t.Root  
  for _, w := range word{
    if _, ok := curr.Children[w]; !ok{
      curr.Children[w] = &Trie{}
    }
    curr = curr.Children[w]
  }
  cur.IsTerminal = true

}

func (t *Trie) Search(word string) bool{
  curr := t.Root
  for _, w := range word{
    if _, ok := curr.Children[w]; !ok{
      return false
    }
    curr = curr.Children[c]
  }
  return curr.IsTerminal
}

func (t *Trie) StartsWith(prefix string) bool{
  curr := t.Root
  for _, w := range prefix{
    if _, ok: = curr.Children[w]; !ok{
      return false
    }
    curr = curr.Children[w]
  }

  return true
}
