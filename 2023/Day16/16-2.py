import sys
sys.setrecursionlimit(1000000)

with open('./Day16/input/16.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    line = line.strip()
    grid.append(list(line))

def move(c, r, prev_c, prev_r, from_left, from_right, from_up, from_down):
    value = 1000 * r + c
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
            if c > prev_c and c < 109: move(c + 1, r, c, r, from_left, from_right, from_up, from_down)
            elif c < prev_c and c > 0: move(c - 1, r, c, r, from_left, from_right, from_up, from_down)
        elif c == prev_c:
            if r > prev_r and r < 109: move(c, r + 1, c, r, from_left, from_right, from_up, from_down)
            elif r < prev_r and r > 0: move(c, r - 1, c, r, from_left, from_right, from_up, from_down)
    elif grid[r][c] == '/':
        if r == prev_r:
            if c > prev_c and r > 0: move(c, r - 1, c, r, from_left, from_right, from_up, from_down)
            elif c < prev_c and r < 109: move(c, r + 1, c, r, from_left, from_right, from_up, from_down)
        elif c == prev_c:
            if r > prev_r and c > 0: move(c - 1, r, c, r, from_left, from_right, from_up, from_down)
            elif r < prev_r and c < 109: move(c + 1, r, c, r, from_left, from_right, from_up, from_down)
    elif grid[r][c] == "\\":
        if r == prev_r:
            if c > prev_c and r < 109: move(c, r + 1, c, r, from_left, from_right, from_up, from_down)
            elif c < prev_c and r > 0: move(c, r - 1, c, r, from_left, from_right, from_up, from_down)
        elif c == prev_c:
            if r > prev_r and c < 109: move(c + 1, r, c, r, from_left, from_right, from_up, from_down)
            elif r < prev_r and c > 0: move(c - 1, r, c, r, from_left, from_right, from_up, from_down)
    elif grid[r][c] == '|':
        if r == prev_r:
            if r > 0: move(c, r - 1, c, r, from_left, from_right, from_up, from_down)
            if r < 109: move(c, r + 1, c, r, from_left, from_right, from_up, from_down)
        elif c == prev_c:
            if r > prev_r and r < 109: move(c, r + 1, c, r, from_left, from_right, from_up, from_down)
            elif r < prev_r and r > 0: move(c, r - 1, c, r, from_left, from_right, from_up, from_down)
    elif grid[r][c] == '-':
        if r == prev_r:
            if c > prev_c and c < 109: move(c + 1, r, c, r, from_left, from_right, from_up, from_down)
            elif c < prev_c and c > 0: move(c - 1, r, c, r, from_left, from_right, from_up, from_down)
        elif c == prev_c:
            if c > 0: move(c - 1, r, c, r, from_left, from_right, from_up, from_down)
            if c < 109: move(c + 1, r, c, r, from_left, from_right, from_up, from_down)

best = 0

for i in range (110):
    from_left = set()
    from_left.add(1000 * i)
    from_right = set()
    from_up = set()
    from_down = set()
    move(0, i, -1, i, from_left, from_right, from_up, from_down)
    result = len(from_left | from_right | from_up | from_down)
    best = max(best, result)

for i in range (110):
    from_left = set()
    from_right = set()
    from_right.add(1000 * i + 109)
    from_up = set()
    from_down = set()
    move(109, i, 110, i, from_left, from_right, from_up, from_down)
    result = len(from_left | from_right | from_up | from_down)
    best = max(best, result)

for i in range (110):
    from_left = set()
    from_right = set()
    from_up = set()
    from_up.add(i)
    from_down = set()
    move(i, 0, i, -1, from_left, from_right, from_up, from_down)
    result = len(from_left | from_right | from_up | from_down)
    best = max(best, result)

for i in range (110):
    from_left = set()
    from_right = set()
    from_up = set()
    from_down = set()
    from_down.add(109000 + i)
    move(i, 109, i, 110, from_left, from_right, from_up, from_down)
    result = len(from_left | from_right | from_up | from_down)
    best = max(best, result)

print(best)