with open('./Day21/input/21.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    line = line.strip()
    grid.append(list(line))

S_INDEX = 65

dest = set()

new_grid = grid

x = 129
y = 65

for i in range(65):
    new_x = x - i
    new_y1 = y + i
    new_y2 = y - i
    while new_y1 >= y:
        if grid[new_x][new_y1] == '.': dest.add(1000 * new_x + new_y1)
        new_y1 -= 2
    while new_y2 <= y:
        if grid[new_x][new_y2] == '.': dest.add(1000 * new_x + new_y2)
        new_y2 += 2
    
x = 1
y = 65

for i in range(64):
    new_x = x + i
    new_y1 = y + i
    new_y2 = y - i
    while new_y1 >= y:
        if grid[new_x][new_y1] == '.': dest.add(1000 * new_x + new_y1)
        new_y1 -= 2
    while new_y2 <= y:
        if grid[new_x][new_y2] == '.': dest.add(1000 * new_x + new_y2)
        new_y2 += 2

for val in dest:
    x = int(val / 1000)
    y = val % 1000
    new_grid[x][y] = '?'

# new_grid[65][65] = 'S'

print(len(dest))

# I got the correct answer, but missed the starting point 'S' and counted grid[80][70] which could not be accessed
# as it is surrounded on four sides by rocks '#'

# Here is code to find the plots surrounded by rocks, which are not rocks themselves

# for i in range(1, 130):
#     for j in range(1, 130):
#         if grid[i + 1][j] == '#' and grid[i - 1][j] == '#' and grid[i][j + 1] == '#' and grid[i][j - 1] == '#' and grid[i][j] != '#':
#             if (i + j) % 2 == 0 and abs(65 - i) + abs(65 - j) < 65: print(i, j)