'''
    How many positions does the tail of the rope visit at least once?
'''

# Find the max movements in either direction and choose a grid position (in you want to see the grid as a 2d array)
# But it is better to just start at (0, 0) and model the movement of head and tail
# Keep track of positions of the tail (as a tuple) in a set

def get_directions():
    directions = []
    numbers = []
    with open('./2022/Day09/input/9.txt', 'r') as file:
        for line in file:
            line = line.strip()
            direction, number = line.split(" ")
            directions.append(direction)
            numbers.append(int(number))
    return directions, numbers

def up(bits):
    pass
def down(bits):
    pass

def left(bits):
    pass

def right(bits):
    pass

d, n = get_directions()
MOVES = len(d)
assert MOVES == len(n)
bits = {
    0: [1000, 1000],
    1: [1000, 1000],
    2: [1000, 1000],
    3: [1000, 1000],
    4: [1000, 1000],
    5: [1000, 1000],
    6: [1000, 1000],
    7: [1000, 1000],
    8: [1000, 1000],
    9: [1000, 1000]
}
tail_positions = set()
for move in range(MOVES):
    for shift in range(n[move]):
        if d[move] == 'U':
            up(bits.copy())
        elif d[move] == 'D':
            bits[0][0] += 1
        elif d[move] == 'L':
            bits[0][1] -= 1
        elif d[move] == 'R':
            bits[0][1] += 1
        tuple9 = (bits[9][0], bits[9][1])
        print(tuple9)
        tail_positions.add(tuple9)

print(tail_positions)
print(len(tail_positions))