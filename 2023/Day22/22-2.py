import copy
from queue import Queue

with open('./Day22/input/22.txt', 'r') as file:
    lines = file.readlines()

bricks = []
for line in lines:
    line = line.strip()
    parts = line.split('~')
    array = [int(num) for part in parts for num in part.split(',')]
    bricks.append(array)

bricks = sorted(bricks, key=lambda x: x[2])

X_MAX = 10
Y_MAX = 10
Z_MAX = 256
KEYS = 1231

grid = [[0 for _ in range(10)] for _ in range(10)]
supporting_index = [[-1 for _ in range(10)] for _ in range(10)]

# These dictionaries were added for part 2
supporting = {} # when a block is the only block supporting another
co_supporting = {} # when the block is NOT the only block supporting another
supported_by = {} # when a block is supported by multiple blocks

fallen = []

for i in range(KEYS):
    if bricks[i][2] != bricks[i][5]:
        x = bricks[i][0]
        y = bricks[i][1]
        current_z = grid[x][y]
        z_difference = bricks[i][5] - bricks[i][2]
        new_fallen = [x, y, current_z + 1, x, y, current_z + 1 + z_difference]
        fallen.append(new_fallen)
        grid[x][y] += z_difference + 1
        if supporting_index[x][y] != -1:
            if supporting_index[x][y] in supporting: supporting[supporting_index[x][y]].add(i)
            else: supporting[supporting_index[x][y]] = {i}
        supporting_index[x][y] = i
    elif bricks[i][1] != bricks[i][4]:
        x = bricks[i][0]
        z = bricks[i][2]
        y1 = bricks[i][1]
        y2 = bricks[i][4]
        z_max = 0
        for j in range(y1, y2 + 1): z_max = max(z_max, grid[x][j])
        new_fallen = [x, y1, z_max + 1, x, y2, z_max + 1]
        fallen.append(new_fallen)
        supporting_indices = set()
        for j in range(y1, y2 + 1):
            if supporting_index[x][j] != -1 and grid[x][j] == z_max : supporting_indices.add(supporting_index[x][j])
            supporting_index[x][j] = i
            grid[x][j] = z_max + 1
        if len(supporting_indices) > 1:
            supported_by[i] = supporting_indices.copy()
            for val in supporting_indices:
                if val in co_supporting: co_supporting[val].add(i)
                else: co_supporting[val] = {i}
        elif len(supporting_indices) == 1: 
            lone_index = list(supporting_indices)[0]
            if lone_index in supporting: supporting[lone_index].add(i)
            else: supporting[lone_index] = {i}
    elif bricks[i][0] != bricks[i][3]:
        x1 = bricks[i][0]
        x2 = bricks[i][3]
        y = bricks[i][1]
        z = bricks[i][2]
        z_max = 0
        for j in range(x1, x2 + 1): z_max = max(z_max, grid[j][y])
        new_fallen = [x1, y, z_max + 1, x2, y, z_max + 1]
        fallen.append(new_fallen)
        supporting_indices = set()
        for j in range(x1, x2 + 1):
            if supporting_index[j][y] != -1 and grid[j][y] == z_max: supporting_indices.add(supporting_index[j][y])
            supporting_index[j][y] = i
            grid[j][y] = z_max + 1
        if len(supporting_indices) > 1:
            supported_by[i] = supporting_indices.copy()
            for val in supporting_indices:
                if val in co_supporting: co_supporting[val].add(i)
                else: co_supporting[val] = {i}
        elif len(supporting_indices) == 1:
            lone_index = list(supporting_indices)[0]
            if lone_index in supporting: supporting[lone_index].add(i)
            else: supporting[lone_index] = {i}
    elif bricks[i][0] == bricks[i][3] and bricks[i][1] == bricks[i][4] and bricks[i][2] == bricks[i][5]:
        x = bricks[i][0]
        y = bricks[i][1]
        z = bricks[i][2]
        current_z = grid[x][y]
        new_fallen = [x, y, current_z + 1, x, y, current_z + 1]
        fallen.append(new_fallen)
        supporting_indices = set()
        if supporting_index[x][y] != -1:
            if supporting_index[x][y] in supporting: supporting[supporting_index[x][y]].add(i)
            else: supporting[supporting_index[x][y]] = {i}
        supporting_index[x][y] = i
        grid[x][y] = current_z + 1

supporting_copy = copy.deepcopy(supporting)

q = Queue()

def check_co_supporting(key, index):
    if key not in co_supporting: return
    values = co_supporting[key]
    for value in values:
        all_supports = supported_by[value]
        if all(support in supporting[index] for support in all_supports):
            supporting[index].add(value)
            q.put([value, index])

def check_supporting(key, index):
    if key not in supporting_copy: return
    values = supporting_copy[key]
    for value in values:
        supporting[index].add(value)
        q.put([value, index])
    
def traverse_data(index):
    if index not in supporting_copy: return 0
    values = supporting_copy[index]
    for value in values:
        q.put([value, index])
    while not q.empty():
        key, index = q.get()
        check_supporting(key, index)
        check_co_supporting(key, index)
    return len(supporting[index])

count = 0

for i in range(KEYS):
    count += traverse_data(i)

# Code takes five seconds to execute
print(count)

# I found a way to traverse the data of what supported what both on their own and with
# other blocks in order to get the set of blocks that would fall for each block removed.
# I just took the length of the set and incremented a counter