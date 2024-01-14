import sys
sys.setrecursionlimit(1000000)

with open('./Day16/input/16.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    line = line.strip()
    grid.append(list(line))
    
from_left = set()
from_left.add(0)
from_right = set()
from_up = set()
from_down = set()
order = []

def move(c, r, prev_c, prev_r):
    value = 1000 * r + c
    order.append(value)
    if c > prev_c:
        if value in from_down: return
        else: from_down.add(value)
    elif c < prev_c:
        if value in from_up: return
        else: from_up.add(value)
    elif r > prev_r:
        if value in from_left: return
        else: from_left.add(value)
    elif r < prev_r:
        if value in from_right: return
        else: from_right.add(value)
    if grid[r][c] == '.':
        if r == prev_r:
            if c > prev_c and c < 109: move(c + 1, r, c, r)
            elif c < prev_c and c > 0: move(c - 1, r, c, r)
        elif c == prev_c:
            if r > prev_r and r < 109: move(c, r + 1, c, r)
            elif r < prev_r and r > 0: move(c, r - 1, c, r)
    elif grid[r][c] == '/':
        if r == prev_r:
            if c > prev_c and r > 0: move(c, r - 1, c, r)
            elif c < prev_c and r < 109: move(c, r + 1, c, r)
        elif c == prev_c:
            if r > prev_r and c > 0: move(c - 1, r, c, r)
            elif r < prev_r and c < 109: move(c + 1, r, c, r)
    elif grid[r][c] == "\\":
        if r == prev_r:
            if c > prev_c and r < 109: move(c, r + 1, c, r)
            elif c < prev_c and r > 0: move(c, r - 1, c, r)
        elif c == prev_c:
            if r > prev_r and c < 109: move(c + 1, r, c, r)
            elif r < prev_r and c > 0: move(c - 1, r, c, r)
    elif grid[r][c] == '|':
        if r == prev_r:
            if r > 0: move(c, r - 1, c, r)
            if r < 109: move(c, r + 1, c, r)
        elif c == prev_c:
            if r > prev_r and r < 109: move(c, r + 1, c, r)
            elif r < prev_r and r > 0: move(c, r - 1, c, r)
    elif grid[r][c] == '-':
        if r == prev_r:
            if c > prev_c and c < 109: move(c + 1, r, c, r)
            elif c < prev_c and c > 0: move(c - 1, r, c, r)
        elif c == prev_c:
            if c > 0: move(c - 1, r, c, r)
            if c < 109: move(c + 1, r, c, r)


move(0, 1, 0, 0)

num_set = from_left.union(from_right, from_up, from_down)
print(len(num_set))