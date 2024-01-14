colors = []
dir = []
dis = []

with open('./Day18/input/18.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        line = line.strip()
        parts = line.split(' ')
        colors.append(parts[2].strip('(#)'))

hex_by_index = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f']

def numify(hex):
    total = 0
    for i in range(-1, -6, -1): total += 16 ** (abs(i) - 1) * hex_by_index.index(hex[i])
    return total

for i in range(len(colors)):
    dir.append(int(colors[i][-1]))
    result = numify(colors[i][:-1])
    dis.append(result)

# 0 is R, 1 is D, 2 is L, 3 is U
    
# There are 660 distance moves

directions = [0, 0, 0, 0]
max_x = float('-inf')
min_x = float('inf')
max_y = float('-inf')
min_y = float('inf')
x = 0
y = 0
positions = []
for i in range(len(dir)):
    if dir[i] == 0:
        directions[0] += dis[i]
        x += dis[i]
    elif dir[i] == 1:
        directions[1] += dis[i]
        y += dis[i]
    elif dir[i] == 2:
        directions[2] += dis[i]
        x -= dis[i]
    elif dir[i] == 3:
        directions[3] += dis[i]
        y -= dis[i]
    max_x = max(max_x, directions[0] - directions[2])
    min_x = min(min_x, directions[0] - directions[2])
    max_y = max(max_y, directions[1] - directions[3])
    min_y = min(min_y, directions[1] - directions[3])
    positions.append((x, y))

# The code below applies the shoelace formula for solving the area of a polygon:

area = 0
for i in range(0, 660, 1):
    area += positions[(i + 1) % 660][1] * positions[i][0]
    area -= positions[(i + 1) % 660][0] * positions[i][1]

# sum(dis) // 2 accounts for the other half of the perimeter
# + 1 accounts for the other corner

print(area // 2 + (sum(dis) // 2) + 1)

# This is an alternative form fo the code above.

# area = 0
# for i in range(0, 660, 2):
#     area += positions[i + 1][1] * positions[i][0]
#     area -= positions[i + 1][0] * positions[i][1]

# print(area + (sum(dis) // 2) + 1)

# print(f'X: {max_x}')
# print(f'X: {min_x}')
# print(f'Y: {max_y}')
# print(f'Y: {min_y}')
# print(directions)

# print(f'rows: {num_rows}')
# print(f'cols: {num_cols}')
# print(positions)

# latest guess: 1995383379590 and 1995383380249 are too low
# correct: 104454050898331