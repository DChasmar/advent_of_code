def makeGrid():
    new_grid = []
    with open('./2022/Day08/input/8.txt', 'r') as file:
        for line in file:
            array = [int(num) for num in line.strip()]
            new_grid.append(array)
    return new_grid


def checkUp(grid, length):
    up_visible = set()
    for col in range(length):
        tallest = -1
        for row in range(length - 1, -1, -1):
            num = grid[row][col]
            if num > tallest:
                up_visible.add((row, col))
                tallest = num
    return up_visible

def checkDown(grid, length):
    down_visible = set()
    for col in range(length):
        tallest = -1
        for row in range(length):
            num = grid[row][col]
            if num > tallest:
                down_visible.add((row, col))
                tallest = num
    return down_visible

def checkLeft(grid, length):
    left_visible = set()
    for row in range(length):
        tallest = -1
        for col in range(length):
            num = grid[row][col]
            if num > tallest:
                left_visible.add((row, col))
                tallest = num
    return left_visible

def checkRight(grid, length):
    right_visible = set()
    for row in range(length):
        tallest = -1
        for col in range(length - 1, -1, -1):
            num = grid[row][col]
            if num > tallest:
                right_visible.add((row, col))
                tallest = num
    return right_visible

grid = makeGrid()
LENGTH = 99
visible = set()
visible.update(checkUp(grid, LENGTH))
visible.update(checkDown(grid, LENGTH))
visible.update(checkLeft(grid, LENGTH))
visible.update(checkRight(grid, LENGTH))
print(len(visible))