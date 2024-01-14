from queue import Queue

with open('./Day17/input/17.txt', 'r') as file:
    lines = file.readlines()

SIZE = 141

grid = []
for line in lines:
    line = line.strip()
    grid.append([])
    for digit in line:
        grid[-1].append(int(digit))

# Up is 0, Down is 1, Left is 2, Right is 3
        
# less than 1338
        
up = [[1338 for _ in range(141)] for _ in range(141)]
down = [[1338 for _ in range(141)] for _ in range(141)]
left = [[1338 for _ in range(141)] for _ in range(141)]
right = [[1338 for _ in range(141)] for _ in range(141)]

def check_min(row, col, total, prev):
    if prev == 0:
        if total < up[row][col]:
            up[row][col] = total
            return True
        else: return False
    elif prev == 1:
        if total < down[row][col]:
            down[row][col] = total
            return True
        else: return False
    elif prev == 2:
        if total < left[row][col]:
            left[row][col] = total
            return True
        else: return False
    elif prev == 3:
        if total < right[row][col]:
            right[row][col] = total
            return True
        else: return False
    else:
        print('error')
        return False

min_total = 1338

def move_123(row, col, total, prev):
    if row == 140 and col == 140:
        global min_total
        min_total = min(min_total, total)
        return
    if prev in [0, 1]:
        if col > 0: q.put([row, col - 1, total + grid[row][col - 1], 2])
        if col > 1: q.put([row, col - 2, total + grid[row][col - 1] + grid[row][col - 2], 2])
        if col > 2: q.put([row, col - 3, total + grid[row][col - 1] + grid[row][col - 2] + grid[row][col - 3], 2])
        if col < 140: q.put([row, col + 1, total + grid[row][col + 1], 3])
        if col < 139: q.put([row, col + 2, total + grid[row][col + 1] + grid[row][col + 2], 3])
        if col < 138: q.put([row, col + 3, total + grid[row][col + 1] + grid[row][col + 2] + grid[row][col + 3], 3])
    elif prev in [2, 3]:
        if row > 0: q.put([row - 1, col, total + grid[row - 1][col], 0])
        if row > 1: q.put([row - 2, col, total + grid[row - 1][col] + grid[row  - 2][col - 2], 0])
        if row > 2: q.put([row - 3, col, total + grid[row - 1][col] + grid[row - 2][col] + grid[row - 3][col], 0])
        if row < 140: q.put([row + 1, col, total + grid[row + 1][col], 1])
        if row < 139: q.put([row + 2, col, total + grid[row + 1][col] + grid[row + 2][col], 1])
        if row < 138: q.put([row + 3, col, total + grid[row + 1][col] + grid[row + 2][col] + grid[row + 3][col], 1])

q = Queue()
q.put([0, 0, 0, 0])
q.put([0, 0, 0, 2])

while not q.empty():
    next = q.get()
    result = check_min(*next)
    if result:
        print(*next)
        move_123(*next)

print(min_total)

# This code takes a long time to run (8 to 10 minutes)