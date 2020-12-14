

# x = 2 (mod 3)
# x = 3 (mod 4)
# x = 1 (mod 5)

m = [3, 4, 5]
x = [2, 3, 1]
N = prod(m)
n = div.(N, m)
i = invmod.(n, m)
mod(sum(x .* n .* i), N)
