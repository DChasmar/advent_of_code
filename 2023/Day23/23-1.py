colors = []
dir = []
dis = []

with open('./Day23/input/23.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    array = [str(val).strip() for val in line]
    grid.append(array)

    
intersections = set()
tallies = []
    
def check_adjacents(row, col, prev_row, prev_col, tally, seen):
    number = 1000 * row + col
    if number in seen: return
    else: seen.add(number)
    if row == 0 and col == 1:
        tallies.append(tally)
        return
    count = 0
    if grid[row + 1][col] in ['.', '^', 'v', '<', '>']: count += 1
    if grid[row - 1][col] in ['.', '^', 'v', '<', '>']: count += 1
    if grid[row][col + 1] in ['.', '^', 'v', '<', '>']: count += 1
    if grid[row][col - 1] in ['.', '^', 'v', '<', '>']: count += 1
    if grid[row + 1][col] in ['.', '^'] and row + 1 != prev_row: stack.append([row + 1, col, row, col, tally + 1, seen.copy()])
    if grid[row - 1][col] in ['.', 'v'] and row - 1 != prev_row: stack.append([row - 1, col, row, col, tally + 1, seen.copy()])
    if grid[row][col + 1] in ['.', '<'] and col + 1 != prev_col: stack.append([row, col + 1, row, col, tally + 1, seen.copy()])
    if grid[row][col - 1] in ['.', '>'] and col - 1 != prev_col: stack.append([row, col - 1, row, col, tally + 1, seen.copy()])
    if count > 2: intersections.add(1000 * row + col)

# Rather than start at the beginning, start at the end.
stack = [[139, 139, 140, 139, 1, {140139}]]

while len(stack) > 0:
    latest = stack.pop()
    check_adjacents(latest[0], latest[1], latest[2], latest[3], latest[4], latest[5])

tallies.sort()

# Code takes five seconds to execute
print(tallies[-1])