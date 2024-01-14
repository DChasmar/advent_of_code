# I chose to solve this problem mathematically, and used code sporadically to assist me.
# I will try to take you through my process.

# Fetch data:
with open('./Day24/input/24.txt', 'r') as file:
    lines = file.readlines()

LINES = 300

positions = []
vectors = []

for line in lines:
    parts = line.split('@')
    positions.append([int(num) for num in parts[0].split(',')])
    vectors.append([int(num) for num in parts[1].split(',')])

pass

# Let our projectile equal: (pos_x, pos_y, pos_z) + time * (vec_x, vec_y, vec_z)

# We know the following to be true for all stones:
# position[0] + time * vector[0] = pos_x + time * vec_x
# position[1] + time * vector[1] = pos_y + time * vec_y
# position[2] + time * vector[2] = pos_z + time * vec_z

# Rearrange to isolate pos_x / pos_y / pos_z:
# pos_x = position[0] + time * vector[0] - time * vec_x
# Factor out time:
# pos_x = position[0] + time * (vector[0] - vec_x)
# pos_y = position[1] + time * (vector[1] - vec_y)
# pos_z = position[2] + time * (vector[2] - vec_z)

# (vector[0] - vec_x) will be the same for different stones IF their vector[0] values are the same

# Use this function to find the stones that share the same x, y, or z vector
def find_shared_index(vectors):
    for i in range(len(vectors)):
        for j in range(i + 1, len(vectors)):
            for index in [0, 1, 2]:
                if vectors[i][index] == vectors[j][index]:
                    pass
                    # print(f"Arrays {i} and {j} share the same value at index {index}")

find_shared_index(vectors)
'''
Taking two stones that share the same vector, beacuse...
pos_x = position[0]_1 + time_1 * (vector[0] - vec_x)
pos_x = position[0]_2 + time_2 * (vector[0] - vec_x)
position[0]_1 + time_1 * (vector[0] - vec_x) = position[0]_2 + time_2 * (vector[0] - vec_x)

Let (vector[0] - vec_x) = n
position[0]_1 + time_1 * n = position[0]_2 + time_2 * n
position[0]_1 - position[0]_2 = time_2 * n - time_1 * n
position[0]_1 - position[0]_2 =  n (time_2 - time_1)

We then do the prime factorization of (position[0]_1 - position[0]_2)
See file 24-2-prime-factorization.py
'''

# Some examples of stones that share the same vector.
# These are their positions
same_vx = [
    265773177829978,
    258406044978185,
    260258023390560,
    267037309903546,
    264322616492155,
    269412247273325,
    259423970886769
] # 19

same_vx2 = [
    257434946624214,
    261704797104278,
    268403596036630,
    260297779253622
] # 22

same_vy = [
    335433658881985,
    150430030914047,
    289447484208054
]

same_vz = [
    314538386160773,
    347681905165705,
    323802359291430,
    305831892945703
] # -18

# Set n to be one less than the length of the same_xyz set
n = 3

for i in range(n):
    for j in range(n - i):
        pass
        # print(abs(same_vz[i] - same_vz[i + 1 + j]))

'''
Thanks to the work done in file 24-2-prime-factorization.py,
we now know the vector of our projectile is [6, 326, 101]

Take any two stones to create a system of equations.
The two unknowns in our two equations are t1 and t2; the times
it takes for the projectile to hit the stone.
'''

# 257520024329236, 69140711609471, 263886787577054 @ 21, 351, 72
# 227164924449606, 333280170830371, 330954002548352 @ 70, -28, -35

'''
We only need to use two of the dimensions. I choose x and y.
'''

# 257520024329236 + 15 * t1 = pos_x
# 227164924449606 + 64 * t2 = pos_x

# 69140711609471 + 25 * t1 = pos_y
# 333280170830371 + -354 * t2 = pos_y

# 257520024329236 + 15 * t1 = 227164924449606 + 64 * t2
# 69140711609471 + 25 * t1 = 333280170830371 + -354 * t2

'''
Scale equations to eliminate t1. Subtract.
'''
    
# 2575200243292360 + 150 * t1 = 2271649244496060 + 640 * t2
# 6 * 69140711609471 + 150 * t1 =  6 * 333280170830371 + -354 * 6 * t2

'''
This is the outcome.
'''

t1 = 891348774638
t2 = 683208304675

# The 0 index vector: 257520024329236, 69140711609471, 263886787577054 @ 21, 351, 72
# 15 = vectors[0][0] - vec_x
# 25 = vectors[0][1] - vec_y
# -29 = vectors[0][2] - vec_z
print((positions[0][0] + 15 * t1) + (positions[0][1] + 25 * t1) + (positions[0][2] + -29 * t1))

# answer is 600352360036779