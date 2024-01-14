# Calculate based on INACCESSIBLE plots

# This problem has not yet been solved. See 21-2-alternate for actual solve

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

# These are inaccessible plots
odd_rock_tuples = set()
odd_surrounded_tuples = set()
even_rock_tuples = set()
even_surrounded_tuples = set()

for i in range(1, SIZE - 1):
    for j in range(1, SIZE - 1):
        if (i + j) % 2 == 1:
            if grid[i][j] == '#': odd_rock_tuples.add((i, j)) # plot is a rock
            # it is important that this next line is elif, not if, because we do not want to count rocks
            # surrounded by rocks twice.
            elif all(grid[x][y] == '#' for x, y in [(i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1)]):
                odd_surrounded_tuples.add((i, j)) # plot is surrounded by rocks, and cannot be accessed
        elif (i + j) % 2 == 0:
            if grid[i][j] == '#': even_rock_tuples.add((i, j)) # plot is a rock
            elif all(grid[x][y] == '#' for x, y in [(i + 1, j), (i - 1, j), (i, j + 1), (i, j - 1)]):
                even_surrounded_tuples.add((i, j)) # plot is surrounded by rocks, and cannot be accessed

# These are inaccessible and in the corners of the grid
odd_corners = set()
even_corners = set()

for val in odd_rock_tuples or odd_surrounded_tuples:
    if (val[0] + val[1]) % 2 == 1 and abs(65 - val[0]) + abs(65 - val[1]) > 65: odd_corners.add(val)
for val in even_rock_tuples or even_surrounded_tuples:
    if (val[0] + val[1]) % 2 == 0 and abs(65 - val[0]) + abs(65 - val[1]) > 65: even_corners.add(val)

# def check_corners_and_add(value, corners_set):
#     if value[0] + value[1] < 65 or value[0] + value[1] > 195 or abs(value[0] - value[1]) > 65:
#         corners_set.add(value)

# for val in odd_rock_tuples | odd_surrounded_tuples:
#     check_corners_and_add(val, odd_corners)

# for val in even_rock_tuples | even_surrounded_tuples:
#     check_corners_and_add(val, even_corners)

ODD_INACCESSIBLE = len(odd_rock_tuples) + len(odd_surrounded_tuples)
EVEN_INACCESSIBLE = len(even_rock_tuples) + len(even_surrounded_tuples)
ODD_CORNER_INACCESSIBLE = len(odd_corners)
EVEN_CORNER_INACCESSIBLE = len(even_corners)

print(ODD_INACCESSIBLE)
print(EVEN_INACCESSIBLE)
print(ODD_CORNER_INACCESSIBLE)
print(EVEN_CORNER_INACCESSIBLE)

INACCESSIBLE = (N + 1)**2 * ODD_INACCESSIBLE + N**2 * EVEN_INACCESSIBLE - (N + 1) * ODD_CORNER_INACCESSIBLE + N * EVEN_CORNER_INACCESSIBLE

result = (STEPS) ** 2 - INACCESSIBLE

# The value of the remainder is ideal to solve this problem, as it becomes a math calculation.
# Exactly...
                
print(result)
            
# 652516024613173 is too high
# 301354851181561 is too low
# 604510412173596 is too low
# 604510465580930 is incorrect
# 604510465176327 is incorrect
# 604592315756327 is incorrect
# 604592315755155 is incorrect
# 604592262753596 is incorrect