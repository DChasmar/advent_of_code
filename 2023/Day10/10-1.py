import sys
sys.setrecursionlimit(100000)

with open('./Day10/input/10.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    array = [str(val).strip() for val in line]
    grid.append(array)

def check_char(x, y, last_x, last_y, num_set):
    # You could add (x, y) to num_set as a tuple;
    # I choose to create a unique number: 1000 * x + y
    number = 1000 * x + y
    if number in num_set: return num_set, True
    else: num_set.add(number)
    if grid[x][y] == 'L':
        if (last_y > y): return check_char(x - 1, y, x, y, num_set)
        elif (last_x < x): return check_char(x, y + 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == 'F':
        if (last_y > y): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x, y + 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == 'J':
        if (last_y < y): return check_char(x - 1, y, x, y, num_set)
        elif (last_x < x): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '7':
        if (last_y < y): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '|':
        if (last_x < x): return check_char(x + 1, y, x, y, num_set)
        elif (last_x > x): return check_char(x - 1, y, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '-':
        if (last_y < y): return check_char(x, y + 1, x, y, num_set)
        elif (last_y > y): return check_char(x, y - 1, x, y, num_set)
        else: return num_set, False
    elif grid[x][y] == '.':
        if len(num_set) > 10: print(x, y)
        return num_set, False

total = 0

S_INDEX = (50, 39)

num_set = set()
num_set.add(S_INDEX[0]*1000 + S_INDEX[1])
num_set, loop = check_char(50, 40, 50, 39, num_set)
if loop: total = max(total, len(num_set))

print(total // 2)