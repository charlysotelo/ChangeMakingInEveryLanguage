#!/bin/lua

function minCoins(coins, n)
  M = {0}
  for i = 1, n do
    candidates = {}
    for j = 1, #coins do
      if (i - coins[j]) >= 0 then
        table.insert(candidates, M[i - coins[j] + 1] + 1) -- Remember lua is 1-indexed
      end
    end
    table.insert(M, math.min(table.unpack(candidates))) --Unpack is basically the spread operator from python/javascript
  end
  return M[n + 1]
end

coins = {1, 5, 25, 30}
n = 37

print(minCoins(coins, n))
