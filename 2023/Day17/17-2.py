from queue import Queue

with open('./Day17/input/17.txt', 'r') as file:
    lines = file.readlines()

SIZE = 141
MIN = 4
MAX = 10

grid = [[int(digit) for digit in line.strip()] for line in lines]

# Up is 0, Down is 1, Left is 2, Right is 3
direction_matrices = {0: [[929 for _ in range(SIZE)] for _ in range(SIZE)],
                      1: [[929 for _ in range(SIZE)] for _ in range(SIZE)],
                      2: [[929 for _ in range(SIZE)] for _ in range(SIZE)],
                      3: [[929 for _ in range(SIZE)] for _ in range(SIZE)]}

def check_min(row, col, total, moves, prev):
    if moves > 400: return False

    matrix = direction_matrices[prev]

    if total < matrix[row][col]:
        matrix[row][col] = total
        return True
    else:
        return False

min_total = 929

def move_123(row, col, total, moves, prev):
    global min_total
    if total >= min_total: return
    if row == SIZE - 1 and col == SIZE - 1:
        min_total = min(min_total, total)
        return
    if prev in [0, 1]:
        for offset in range(MIN, MAX + 1):
            if col > offset - 1:
                left_total = sum(grid[row][col - i] for i in range(1, offset + 1))
                q.put([row, col - offset, total + left_total, moves + offset, 2])
        
        for offset in range(MIN, MAX + 1):
            if col < SIZE - offset:
                right_total = sum(grid[row][col + i] for i in range(1, offset + 1))
                q.put([row, col + offset, total + right_total, moves + offset, 3])

    elif prev in [2, 3]:
        for offset in range(MIN, MAX + 1):
            if row > offset - 1:
                up_total = sum(grid[row - i][col] for i in range(1, offset + 1))
                q.put([row - offset, col, total + up_total, moves + offset, 0])

        for offset in range(MIN, MAX + 1):
            if row < SIZE - offset:
                down_total = sum(grid[row + i][col] for i in range(1, offset + 1))
                q.put([row + offset, col, total + down_total, moves + offset, 1])

q = Queue()
q.put([0, 0, 0, 0, 0])
q.put([0, 0, 0, 0, 2])

while not q.empty():
    next = q.get()
    result = check_min(*next)
    if result:
        print(*next)
        move_123(*next)

print(min_total)