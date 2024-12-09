'''
    The scenic score is determined by the number of trees it can see
    up, down, left, and right, and those values multiplied to one another.
    So long as a tree is less than my tree, I can see beyond it.
    Trees on the edge will have a score of zero, since in one direction,
    the will be able to see zero trees.
'''

def makeGrid():
    new_grid = []
    with open('./2022/Day08/input/8.txt', 'r') as file:
        for line in file:
            array = [int(num) for num in line.strip()]
            new_grid.append(array)
    return new_grid


def checkUp(row, col, grid):
    num = grid[row][col]
    up_limit = row - 1
    while up_limit > 0 and grid[up_limit][col] < num:
        up_limit -= 1
    return abs(row - up_limit)

def checkDown(row, col, grid, length):
    num = grid[row][col]
    down_limit = row + 1
    while down_limit < length - 1 and grid[down_limit][col] < num:
        down_limit += 1
    return abs(row - down_limit)

def checkLeft(row, col, grid):
    num = grid[row][col]
    left_limit = col - 1
    while left_limit > 0 and grid[row][left_limit] < num:
        left_limit -= 1
    return abs(col - left_limit)

def checkRight(row, col, grid, length):
    num = grid[row][col]
    right_limit = col + 1
    while right_limit < length - 1 and grid[row][right_limit] < num:
        right_limit += 1
    return abs(col - right_limit)

grid = makeGrid()
LENGTH = 99
max_score = 0
for row in range(1, LENGTH - 1, 1):
    for col in range(1, LENGTH - 1, 1):
        score = checkUp(row, col, grid) * checkDown(row, col, grid, LENGTH) * checkLeft(row, col, grid) * checkRight(row, col, grid, LENGTH)
        max_score = max(score, max_score)
print(max_score)