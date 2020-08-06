#!/usr/bin/python

def minCoins(coins, n):
    """
    >>> minCoins([1], 37)
    37
    >>> minCoins([1,5, 25, 30], 30)
    1
    >>> minCoins([1,5, 25], 30)
    2
    """
    M = [0]
    for i in range(1, n + 1):
      candidates = [M[i-coin] + 1 for coin in coins if (i-coin >= 0)]
      M.append(min(candidates))
    return (M[n], C[n])

