colors = []
dir = []
dis = []

with open('./Day23/input/23.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    array = [str(val).strip() if val not in ['^', 'v', '<', '>'] else '.' for val in line]
    grid.append(array)
    
intersections = set()
tallies = []

for row in range(1, 140):
    for col in range(1, 140):
        if grid[row][col] == '.':
            count = 0
            if grid[row + 1][col] == '.': count += 1
            if grid[row - 1][col] == '.': count += 1
            if grid[row][col + 1] == '.': count += 1
            if grid[row][col - 1] == '.': count += 1
            if count > 2:
                number = 1000 * row + col
                intersections.add(number)
                grid[row][col] = 'X'

data = {}
    
def check_adjacents(row, col, prev_row, prev_col, tally, origin_number):
    if grid[row][col] == 'X':
        number = 1000 * row + col
        if origin_number == 140139: data[number] = [(origin_number, tally + 1)]
        elif origin_number in data: data[origin_number].append((number, tally + 1))
        else: data[origin_number] = [(number, tally + 1)]
        return
    if grid[row + 1][col] in ['.', 'X'] and row + 1 != prev_row: stack.append([row + 1, col, row, col, tally + 1, origin_number])
    if grid[row - 1][col] in ['.', 'X'] and row - 1 != prev_row: stack.append([row - 1, col, row, col, tally + 1, origin_number])
    if grid[row][col + 1] in ['.', 'X'] and col + 1 != prev_col: stack.append([row, col + 1, row, col, tally + 1, origin_number])
    if grid[row][col - 1] in ['.', 'X'] and col - 1 != prev_col: stack.append([row, col - 1, row, col, tally + 1, origin_number])

stack = [[139, 139, 140, 139, 0, 140139]] # end to final intersection

while len(stack) > 0:
    latest = stack.pop()
    check_adjacents(*latest)

stack = [[1, 1, 0, 1, 0, 1]] # start to first intersection

while len(stack) > 0:
    latest = stack.pop()
    check_adjacents(*latest)

grid[1][0] = 'X'
grid[140][139] = 'X'

for number in intersections:
    col = number % 1000
    row = int(number / 1000)
    if grid[row + 1][col] in ['.', 'X']: stack.append([row + 1, col, row, col, 0, number])
    if grid[row - 1][col] in ['.', 'X']: stack.append([row - 1, col, row, col, 0, number])
    if grid[row][col + 1] in ['.', 'X']: stack.append([row, col + 1, row, col, 0, number])
    if grid[row][col - 1] in ['.', 'X']: stack.append([row, col - 1, row, col, 0, number])

while len(stack) > 0:
    latest = stack.pop()
    check_adjacents(*latest)

max_length = 0  # Global variable to track the maximum length

def track_max(new_length):
    global max_length
    if new_length > max_length:
        # Print length if it is greater than max_length (the final (max) value printed will be the answer)
        print(new_length)
        max_length = new_length

def dfs(current_key, current_length, visited):
    visited.add(current_key)
    for neighbor_key, edge_length in data.get(current_key, []):
        if neighbor_key not in visited:
            if neighbor_key == 140139:
                new_length = current_length + edge_length
                track_max(new_length)
            else: stack.append([neighbor_key, current_length + edge_length, visited.copy()])

max_length = 0
stack = [[1, 0, set()]]

# Code takes 10 to 15 seconds to execute
# Instead of traversing the grid one plot at a time, we move from one intersection
# to the next, making sure we have not visited that intersection once before.
# We cannot revisit a path on a square without first visiting an intersection twice.
while len(stack) > 0:
    latest = stack.pop()
    dfs(*latest)