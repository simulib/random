# Some code to calculate the jump constants for pcg54dxsm

mod = 2**128
a = (2549297995355413924 << 64) | 4865540595714422341
c = (6364136223846793005 << 64) | 1442695040888963407

def lcg_jump(steps):
    """
    Compute jump-ahead multiplier and increment for an LCG modulo 2^128.

    Returns (jump_mul, jump_inc) such that:
        x_{n+steps} = jump_mul * x_n + jump_inc  (mod 2^128)
    """
    jump_mul = 1
    jump_inc = 0

    cur_mul = a
    cur_inc = c

    while steps > 0:
        if steps & 1:
            jump_inc = (jump_inc * cur_mul + cur_inc) % mod
            jump_mul = (jump_mul * cur_mul) % mod

        # advance (cur_mul, cur_inc) by doubling the step
        cur_inc = (cur_inc * (cur_mul + 1)) % mod
        cur_mul = (cur_mul * cur_mul) % mod

        steps >>= 1

    return jump_mul, jump_inc


A, C = lcg_jump(2**64)
print((A>>64),(A%(2**64)))
print((C>>64),(C%(2**64)))