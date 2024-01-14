with open('./Day14/input/14.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    grid.append(line.strip())

GRID_SIZE = 100
total = 0

for i in range(GRID_SIZE):
    most = 0
    for j in range(GRID_SIZE):
        if grid[j][i] == 'O':
            total += GRID_SIZE - most
            most += 1
        elif grid[j][i] == '#':
            most = j + 1

print(total)