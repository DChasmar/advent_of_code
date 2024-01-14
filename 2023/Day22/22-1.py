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

grid = [[0 for _ in range(10)] for _ in range(10)]
supporting_index = [[-1 for _ in range(10)] for _ in range(10)]

necessary = set()
unnecessary = set()

fallen = []

for i in range(1231):
    if bricks[i][2] != bricks[i][5]:
        x = bricks[i][0]
        y = bricks[i][1]
        current_z = grid[x][y]
        z_difference = bricks[i][5] - bricks[i][2]
        new_fallen = [x, y, current_z + 1, x, y, current_z + 1 + z_difference]
        fallen.append(new_fallen)
        grid[x][y] += z_difference + 1
        if supporting_index[x][y] != -1: necessary.add(supporting_index[x][y])
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
        if len(supporting_indices) > 1: unnecessary.update(supporting_indices)
        elif len(supporting_indices) == 1: necessary.update(supporting_indices)
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
        if len(supporting_indices) > 1: unnecessary.update(supporting_indices)
        elif len(supporting_indices) == 1: necessary.update(supporting_indices)
    elif bricks[i][0] == bricks[i][3] and bricks[i][1] == bricks[i][4] and bricks[i][2] == bricks[i][5]:
        x = bricks[i][0]
        y = bricks[i][1]
        z = bricks[i][2]
        current_z = grid[x][y]
        new_fallen = [x, y, current_z + 1, x, y, current_z + 1]
        fallen.append(new_fallen)
        supporting_indices = set()
        if supporting_index[x][y] != -1: necessary.add(supporting_index[x][y])
        supporting_index[x][y] = i
        grid[x][y] = current_z + 1

print(1231 - len(necessary))

both  = set()
neither = set()
for i in range(1231):
    if i in necessary and i in unnecessary: both.add(i)
    elif i not in necessary and i not in unnecessary: neither.add(i)

# The array fallen and sets unnecessary, both and neither were not essential.
# I included them in case I wanted to use them for part 2 or try a different
# approach for part 1.