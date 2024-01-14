import random

def gcd(a, b):
    while b:
        a, b = b, a % b
    return a

def pollards_rho(n):
    if n % 2 == 0:
        return 2

    x = random.randint(1, n-1)
    y = x
    c = random.randint(1, n-1)
    d = 1

    f = lambda x: (x**2 + c) % n

    while d == 1:
        x = f(x)
        y = f(f(y))
        d = gcd(abs(x - y), n)

    return d

def prime_factorization_advanced(n):
    factors = []
    
    while n > 1:
        factor = pollards_rho(n)
        factors.append(factor)
        n //= factor
    
    return factors

'''
The values in num represent position[0, 1, or 2]_i - position[0, 1, or 2]_j
for all values that share the same vector[0, 1, or 2]
'''
for num in [
    33143519004932,
    9263973130657,
    8706493215070,
    23879545874275,
    41850012220002,
    17970466345727
]:
    result = prime_factorization_advanced(num)
    print(f"Prime factorization of {num}: {sorted(result)}")

'''
Here are the results; the GCF of position[0]_1 - position[0]_2
The GCF will be n (the closing distance between the stone and our projectile per second)

vector[0]: 19 => GCF: 13
vector[0]: 21 => GCF: 15
vector[0]: 22 => GCF: 32 (or 16 ?)
therefore vec_x = 6

vector[1]: -72 => GCF: 199 (-398 ?)
vector[1]: 15 => GCF: 311
therefore vec_y = GCF: 326

vector[2]: -38 => GCF: 139 (-139 ?)
vector[2]: -18 => 17 x 7 = 119 (-119 ?)
therefore vec_z = 101

Vec = [6, 326, 101]
'''