import sys
sys.setrecursionlimit(100000)

with open('./Day10/input/10.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    array = [str(val).strip() for val in line]
    grid.append(array)

def check_char(x, y, last_x, last_y, num_set):
    number = 1000*x + y
    if number in num_set: return num_set, True
    else: num_set.add(number)
    if grid[x][y] == 'L':
        if (last_y > y): return check_char(x - 1, y, x, y, num_set)
        elif (last_x < x): return check_char(x, y + 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == 'F':
        if (last_y > y): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x, y + 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == 'J':
        if (last_y < y): return check_char(x - 1, y, x, y, num_set)
        elif (last_x < x): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '7':
        if (last_y < y): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '|':
        if (last_x < x): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x - 1, y, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '-':
        if (last_y < y): return check_char(x, y + 1, x, y, num_set)
        elif (last_y > y): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '.':
        if len(num_set) > 10: print(x, y)
        return num_set, False

total = 0

S_INDEX = (50, 39)
num_set = set()
num_set.add(S_INDEX[0]*1000 + S_INDEX[1])
num_set, loop = check_char(50, 40, 50, 39, num_set)
# The code above is all from part 1
# Its value is finding all the indices of the pieces in the S loop (nums_set)

# Create a 140 x 140 grid of only dots
new_grid = [['.' for _ in range(140)] for _ in range(140)]
# Add back the symbols of the original grid if the indices are in num_set
for num in num_set:
    x = int(num / 1000)
    y = num % 1000
    new_grid[x][y] = grid[x][y]

# Create a big grid that adds a spaces between each symbol of the original grid
# The spaces are represented by commas ','
big_grid = []

for array in new_grid:
    # an extra row at the top to avoid out-of bounds error
    big_grid.append([',' for _ in range(281)])
    big_grid.append([])
    for char in array:
        big_grid[-1].append(',')
        if char != 'S': big_grid[-1].append(char)
        # In order to complete the grid properly, the 'S' should be a '-'
        # I could have made this edit above when creating the new_grid to avoid
        # checking for 'S' 140 x 140 times
        else:big_grid[-1].append('-')
    big_grid[-1].append(',')

# an extra row at the bottom to avoid out-of bounds error
big_grid.append([',' for _ in range(281)])

# This code is necessary to close up the loop
for i, array in enumerate(big_grid):
    if i == 0 or i == 280: continue
    for j, char in enumerate(array):
        if j == 0 or j == 280: continue
        if char == ',':
            # close the loop horizontally
            if big_grid[i][j - 1] in ['-', 'F', 'L'] and big_grid[i][j + 1] in ['-', 'J', '7']:
                big_grid[i][j] = '-'
            # close the loop vertically
            if big_grid[i - 1][j] in ['|', 'F', '7'] and big_grid[i + 1][j] in ['|', 'J', 'L']:
                big_grid[i][j] = '|'

# Recursive function to find all the indices inside the loop
def find_inside(x, y, prev_x, prev_y):
    count = 0
    number = 1000 * x + y
    if number in found: return count
    else: found.add(number)
    if big_grid[x][y] == '.': count += 1
    if big_grid[x + 1][y] in [',', '.']:
        if prev_x != x + 1: stack.append([x + 1, y, x, y])
    if big_grid[x - 1][y] in [',', '.']:
        if prev_x != x - 1: stack.append([x - 1, y, x, y])
    if big_grid[x][y + 1] in [',', '.']:
        if prev_y != y + 1: stack.append([x, y + 1, x, y])
    if big_grid[x][y - 1] in [',', '.']:
        if prev_y != y - 1: stack.append([x, y - 1, x, y])
    return count

stack = []
found = set()
tally = 0

# starting_index is 16, 66;
# I identified one index that could work by looking at big_grid exported to a .txt file
# See 10-test2 in outputs
stack.append([16, 66, 15, 66])

# Use a stack with the recursive function to avoid a maximum recursion depth error
while len(stack) > 0:
    latest = stack.pop()
    tally += find_inside(latest[0], latest[1], latest[2], latest[3])

print(tally)