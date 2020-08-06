class minCoins {
  public static void main(String[] args) {
    int coins[] = new int[]{1, 5, 25, 30};
    int change = 37;
    System.out.println(minCoins(coins, change));
  }

  public static int minCoins(int[] coins, int n) {
    int M[] = new int[n + 1];
    M[0] = 0;
    for (int i = 1; i <= n; i++) {
      int minVal = Integer.MAX_VALUE;
      for (int coin: coins) {
        if (i - coin < 0)
          continue;
        minVal = (minVal < 1 + M[i - coin]) ? minVal : (1 + M[i - coin]);
      }
      M[i] = minVal;
    }

    return M[n];
  }
}
