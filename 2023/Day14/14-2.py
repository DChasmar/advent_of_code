with open('./Day14/input/14.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    line = line.strip()
    grid.append(list(line))

GRID_SIZE = 100

def north():
    for i in range(GRID_SIZE):
        most = 0
        for j in range(GRID_SIZE):
            if grid[j][i] == 'O':
                grid[j][i], grid[most][i] = grid[most][i], grid[j][i]
                most += 1
            elif grid[j][i] == '#':
                most = j + 1

def west():
    for i in range(GRID_SIZE):
        most = 0
        for j in range(GRID_SIZE):
            if grid[i][j] == 'O':
                grid[i][j], grid[i][most] = grid[i][most], grid[i][j]
                most += 1
            elif grid[i][j] == '#':
                most = j + 1

def south():
    for i in range(GRID_SIZE - 1, -1, -1):
        most = GRID_SIZE - 1
        for j in range(GRID_SIZE - 1, -1, -1):
            if grid[j][i] == 'O':
                grid[j][i], grid[most][i] = grid[most][i], grid[j][i]
                most -= 1
            elif grid[j][i] == '#':
                most = j - 1

def east():
    for i in range(GRID_SIZE - 1, -1, -1):
        most = GRID_SIZE - 1
        for j in range(GRID_SIZE - 1, -1, -1):
            if grid[i][j] == 'O':
                grid[i][j], grid[i][most] = grid[i][most], grid[i][j]
                most -= 1
            elif grid[i][j] == '#':
                most = j - 1


def weight():
    total = 0
    for i in range(GRID_SIZE):
        for j in range(GRID_SIZE):
            if grid[i][j] == 'O':
                total += 100 - i
    return total

weights = {}

# 1000 is an arbitrary nuber of cycles chosen to then determine when we reach a repeated cycle
# 200 would be sufficient
for i in range (1000):
    north()
    west()
    south()
    east()
    result = weight()
    if result in weights: weights[result].append(i)
    else: weights[result] = [i]

print(weights)

# After printing the weights after certain numbers of cycles, I realized a 94876 is the weight
# after 108 cycles, and again every 27 cycles after that. So, after 999 999 cycles (I tracked
# the indices), that weight will recur. 