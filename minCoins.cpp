#include <bits/stdc++.h>
#include <iostream>
#include <limits.h>
using namespace std;

int minCoins(int coins[], const int numCoins, int n) {
  int M[numCoins + 1];
  M[0] = 0;
  for (int i = 1; i <= n; i++) {
    int minVal = INT_MAX;
    for (int j = 0; j < numCoins; j++) {
      if (i - coins[j] < 0)
        continue;

      int localMin = 1 + M[i - coins[j]];
      if (minVal > localMin) {
        minVal = localMin;
      }
    }
    M[i] = minVal;
  }
  return M[n];  
}

int main() {
  int coins[] = {1, 5, 25, 30};
  cout << minCoins(coins, 4, 37) << endl;
}
