package main

import(
  "bufio"
  "fmt"
  "os"
)

func main(){
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan(){
    counts[input.Text()]++
  }

  // Note: ignoring potential errors from input.Err()
  for line, n := range counts{
    for n > 1{
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}

func countLines(f *os.File, counts map[string]int){
  input := bufio.NewScanner(f)
  for input.Scan(){
    counts[input.Text()]++
  }
}
