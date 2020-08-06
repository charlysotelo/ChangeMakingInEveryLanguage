package main

import "fmt"

func minCoin(coins []int, n int) (int, []int) {
  M := make([]int, n + 1)      //Stores minimum number of coins needed for value i
  C := make([][]int, n + 1)    //Stores which coins are used
  M[0] = 0
  C[0] = []int{}
  for i := 1; i <= n; i++ {
    minVal := int(int(^uint(0) >> 1))
    var minCoins []int
    for _, coin := range coins {
      if i - coin < 0 {
        continue
      }

      localMin := 1 + M[i - coin]
      if minVal > localMin {
        minVal = localMin
        minCoins = append(C[i - coin], coin)
      }
    }
    M[i] = minVal
    C[i] = minCoins
  }
  return M[n], C[n]
}

func main() {
  fmt.Println(minCoin([]int{1, 5, 25, 30}, 37))
}
