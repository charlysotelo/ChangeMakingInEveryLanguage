fn min_coin(coins: &[u32], n: u32) -> u32 {
    let mut m: Vec<u32> = vec![0; (n + 1) as usize];
    m[0] = 0;
    for i in 1u32..(n + 1) {
        let mut min_val = std::u32::MAX;
        for coin in coins.iter() {
            if i < *coin {
                continue;
            }

            let local_min = 1 + m[(i - *coin) as usize];
            if min_val > local_min {
                min_val = local_min;
            }
        }
        m[i as usize] = min_val;
    }
    m[n as usize]
}

fn main() {
    let coins: [u32; 4] = [1, 5, 25, 30];
    let n: u32 = 37;
    println!("{}", min_coin(&coins, n));
}
