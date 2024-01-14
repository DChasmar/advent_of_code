dir = []
dis = []
colors = []

with open('./Day18/input/18.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        line = line.strip()
        parts = line.split(' ')
        dir.append(parts[0])
        dis.append(int(parts[1]))
        colors.append(parts[2].strip('()'))

# directions are [LEFT, DOWN, RIGHT, UP]
        
directions = [0, 0, 0, 0]
max_x = -1000
min_x = 1000
max_y = -1000
min_y = 1000
for i in range(len(dir)):
    if dir[i] == 'L': directions[0] += dis[i]
    elif dir[i] == 'D': directions[1] += dis[i]
    elif dir[i] == 'R': directions[2] += dis[i]
    elif dir[i] == 'U': directions[3] += dis[i]
    max_x = max(max_x, directions[2] - directions[0])
    min_x = min(min_x, directions[2] - directions[0])
    max_y = max(max_y, directions[1] - directions[3])
    min_y = min(min_y, directions[1] - directions[3])

# print(max_x) # Output: 200
# print(min_x) # Output: -130
# print(max_y) # Output: 161
# print(min_y) # Output: -184
# print(directions) # Output: [854, 841, 854, 841]

# directions reveals we end up back where we start
# the size grid is max_x - min_x and max_y - min_y

NUM_ROWS, NUM_COLS = 346, 331

grid = [[2 for _ in range(NUM_COLS)] for _ in range(NUM_ROWS)]

row_index, col_index = 184, 130

grid[row_index][col_index] = 0

for i in range(len(dir)):
    for j in range(dis[i]):
        if dir[i] == 'L': col_index -= 1
        elif dir[i] == 'D': row_index += 1
        elif dir[i] == 'R': col_index += 1
        elif dir[i] == 'U': row_index -= 1
        grid[row_index][col_index] = 0

stack = []

for h in range(0, NUM_ROWS, NUM_ROWS - 1):
    outside = False
    for i in range(NUM_COLS):
        if not outside and grid[h][i] == 2:
            stack.append([h,i])
            outside = True
        elif outside and grid[h][i] == 0: outside = False

for i in range(0, NUM_COLS, NUM_COLS - 1):
    outside = False
    for j in range(NUM_ROWS):
        if not outside and grid[j][i] == 2:
            stack.append([j,i])
            outside = True
        elif outside and grid[j][i] == 0: outside = False

while len(stack) > 0:
    latest = stack.pop()
    row = latest[0]
    col = latest[1]
    grid[row][col] = 1
    if row < NUM_ROWS - 1 and grid[row + 1][col] == 2: stack.append([row + 1, col])
    if row > 0 and grid[row - 1][col] == 2: stack.append([row - 1, col])
    if col < NUM_COLS - 1 and grid[row][col + 1] == 2: stack.append([row, col + 1])
    if col > 0 and grid[row][col - 1] == 2: stack.append([row, col - 1])

count = 0
for row in grid:
    for cell in row:
        if cell == 0 or cell == 2:
            count += 1

print(count)