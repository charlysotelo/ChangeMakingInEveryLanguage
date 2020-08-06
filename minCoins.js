
function minCoins(coins, n) {
  const M = [0];
  for (let i = 1; i <= n; i++) {
    const candidates = coins.filter(coin => (i - coin) >= 0).map(coin => (M[i - coin] + 1));
    M.push(Math.min(...candidates));
  }
  return M[n]; 
}


const coins = [1,5, 25, 30];
console.log(minCoins(coins, 37));
