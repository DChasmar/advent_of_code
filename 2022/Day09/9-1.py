'''
    How many positions does the tail of the rope visit at least once?
'''

# Find the max movements in either direction and choose a grid position (in you want to see the grid as a 2d array)
# But it is better to just start at (0, 0) and model the movement of head and tail
# Keep track of positions of the tail (as a tuple) in a set

def get_directions():
    directions = []
    numbers = []
    with open('./2022/Day09/input/9.txt', 'r') as file:
        for line in file:
            line = line.strip()
            direction, number = line.split(" ")
            directions.append(direction)
            numbers.append(int(number))
    return directions, numbers

def up(head_row, head_col, tail_row, tail_col, moves):
    new_positions = set()
    while moves > 0:
        if head_row < tail_row:
            head_row -= 1
            tail_row -= 1
            tail_col = head_col
        else:
            head_row -= 1
        new_positions.add((tail_row, tail_col))
        moves -= 1
    return head_row, head_col, tail_row, tail_col, new_positions

def down(head_row, head_col, tail_row, tail_col, moves):
    new_positions = set()
    while moves > 0:
        if head_row > tail_row:
            head_row += 1
            tail_row += 1
            tail_col = head_col
        else:
            head_row += 1
        new_positions.add((tail_row, tail_col))
        moves -= 1
    return head_row, head_col, tail_row, tail_col, new_positions

def left(head_row, head_col, tail_row, tail_col, moves):
    new_positions = set()
    while moves > 0:
        if head_col < tail_col:
            head_col -= 1
            tail_col -= 1
            tail_row = head_row
        else:
            head_col -= 1
        new_positions.add((tail_row, tail_col))
        moves -= 1
    return head_row, head_col, tail_row, tail_col, new_positions

def right(head_row, head_col, tail_row, tail_col, moves):
    new_positions = set()
    while moves > 0:
        if head_col > tail_col:
            head_col += 1
            tail_col += 1
            tail_row = head_row
        else:
            head_col += 1
        new_positions.add((tail_row, tail_col))
        moves -= 1
    return head_row, head_col, tail_row, tail_col, new_positions

d, n = get_directions()
MOVES = len(d)
assert MOVES == len(n)
head_row, head_col, tail_row, tail_col = 1000, 1000, 1000, 1000
tail_positions = set()
tail_positions.add((1000, 1000))
for move in range(MOVES):
    if d[move] == 'U':
        head_row, head_col, tail_row, tail_col, new_positions = up(head_row, head_col, tail_row, tail_col, n[move])
        tail_positions.update(new_positions)
    elif d[move] == 'D':
        head_row, head_col, tail_row, tail_col, new_positions = down(head_row, head_col, tail_row, tail_col, n[move])
        tail_positions.update(new_positions)
    elif d[move] == 'L':
        head_row, head_col, tail_row, tail_col, new_positions = left(head_row, head_col, tail_row, tail_col, n[move])
        tail_positions.update(new_positions)
    elif d[move] == 'R':
        head_row, head_col, tail_row, tail_col, new_positions = right(head_row, head_col, tail_row, tail_col, n[move])
        tail_positions.update(new_positions)

print(len(tail_positions))