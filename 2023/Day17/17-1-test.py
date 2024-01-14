from queue import Queue

with open('./Day17/input/17.txt', 'r') as file:
    lines = file.readlines()

SIZE = 141

grid = [[int(digit) for digit in line.strip()] for line in lines]

# Up is 0, Down is 1, Left is 2, Right is 3
directions = [(0, -1, 2), (0, 1, 3), (-1, 0, 0), (1, 0, 1)]

up = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
down = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
left = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
right = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]

def check_min(row, col, total, prev, dp):
    if total < dp[row][col]:
        dp[row][col] = total
        return True
    return False

def move(row, col, total, prev, q, dp):
    if row == SIZE - 1 and col == SIZE - 1:
        global min_total
        min_total = min(min_total, total)
        return
    for dr, dc, new_prev in directions:
        new_row, new_col = row + dr, col + dc
        if 0 <= new_row < SIZE and 0 <= new_col < SIZE:
            new_total = total + grid[new_row][new_col]
            if check_min(new_row, new_col, new_total, new_prev, dp):
                q.put([new_row, new_col, new_total, new_prev])

q = Queue()
min_total = 1338

dp_up = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
dp_down = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
dp_left = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]
dp_right = [[1338 for _ in range(SIZE)] for _ in range(SIZE)]

q.put([0, 0, 0, 0])
q.put([0, 0, 0, 2])

while not q.empty():
    next_state = q.get()
    move(*next_state, q, dp_up)
    move(*next_state, q, dp_down)
    move(*next_state, q, dp_left)
    move(*next_state, q, dp_right)

print(min_total)

# This code takes a long time to run (8 to 10 minutes)