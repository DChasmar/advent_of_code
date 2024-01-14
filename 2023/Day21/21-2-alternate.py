from queue import Queue

# Calculate based on ACCESSIBLE plots

with open('./Day21/input/21.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    line = line.strip()
    grid.append(list(line))

S_INDEX = 65
SIZE = 131
STEPS = 26501365
REMAINDER = STEPS % SIZE # this is 65
# An important detail here (that was also the case in part 1) is that the plots on the perimeter
# do not have rocks close to them. If they did, it might be possible that a plot was not a rock
# nor surrounded, yet it still could not be reached because there are rocks just before it when
# we are about to run out of steps to take.
N = STEPS // SIZE # this is 202300

odd_reachable = set()
even_reachable = set()
odd_corners = set()
even_corners = set()

def reachable(row, col):
    return not grid[row][col] == '#' and not all(grid[x][y] == '#' for x, y in [(row + 1, col), (row - 1, col), (row, col + 1), (row, col - 1)])

def corner(row, col):
    return (abs(S_INDEX - row) + abs(S_INDEX - col)) > REMAINDER

def perimeter(row, col):
    return row in [0, 130] or col in [0, 130]

for row in range(SIZE):
    for col in range(SIZE):
        tup = (row, col)
        if (row + col) % 2 == 0:
            if perimeter(row, col):
                even_reachable.add(tup)
                if corner(row, col): even_corners.add(tup)
            elif reachable(row, col):
                even_reachable.add(tup)
                if corner(row, col): even_corners.add(tup)
        elif (row + col) % 2 == 1:
            if perimeter(row, col):
                odd_reachable.add(tup)
                if corner(row, col):odd_corners.add(tup)
            elif reachable(row, col):
                odd_reachable.add(tup)
                if corner(row, col): odd_corners.add(tup)

ODDS = len(odd_reachable)
ODD_C = len(odd_corners)
EVENS = len(even_reachable)
EVEN_C = len(even_corners)

print(ODDS)
print(ODD_C)
print(EVENS)
print(EVEN_C)

even_corners_test = set()
odd_corners_test = set()

test_grid  = [[float('inf') for _ in range(131)] for _ in range(131)]

def bfs(row, col, steps):
    if steps >= test_grid[row][col] or steps > 64: return
    test_grid[row][col] = steps
    if (row + col) % 2 == 0: even_corners_test.add((row, col))
    elif (row + col) % 2 == 1: odd_corners_test.add((row, col))
    if row > 0 and grid[row - 1][col] != '#': q.put([row - 1, col, steps + 1])
    if col > 0 and grid[row][col - 1] != '#': q.put([row, col - 1, steps + 1])
    if row < 130 and grid[row + 1][col] != '#': q.put([row + 1, col, steps + 1])
    if col < 130 and grid[row][col + 1] != '#': q.put([row, col + 1, steps + 1])

q = Queue()

q.put([0, 0, 0])
q.put([130, 0, 0])
q.put([0, 130, 0])
q.put([130, 130, 0])

while not q.empty():
    next = q.get()
    # print(*next)
    bfs(*next)

# print(sorted(even_corners_test))
EVEN_C_TEST = len(even_corners_test)
ODD_C_TEST = len(odd_corners_test)

print(ODD_C_TEST)
print(EVEN_C_TEST)

for val in even_corners:
    if val not in even_corners_test: print(val)

for val in odd_corners:
    if val not in odd_corners_test: print(val)

result = (N + 1) * (N + 1) * ODDS + N * N * EVENS - (N + 1) * ODD_C + N * EVEN_C_TEST

print(result)

# The important detail regarding the plots on the very edge of the of our area of search
# is that we should consider which odd plots (row + col is odd) we can reach from the CENTER
# with up to 65 steps (this includes all the plots that pass the 'reachable' test).
# Conversely, we should consider all the even plots (row + col is even) we can reach from
# the CORNERS with up to 64 steps.
# This is why the result calculation subtracts ODD_C and adds EVEN_C_TEST (the latter of
# which tests which plots we can reach from the corner)