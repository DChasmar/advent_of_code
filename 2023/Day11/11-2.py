import bisect

with open('./Day11/input/11.txt', 'r') as file:
    lines = file.readlines()

grid = []
for line in lines:
    array = [str(val).strip() for val in line]
    grid.append(array)

empty_rows = []
empty_columns = [num for num in range(140)]
for i, array in enumerate(grid):
    count = 0
    for j, item in enumerate(array):
        if item == '#':
            count += 1
            if j in empty_columns: empty_columns.remove(j)
    if count == 0: empty_rows.append(i)

# print(empty_columns)
# print(empty_rows)

rows = []
columns = []

for i, array in enumerate(grid):
    for j, item in enumerate(array):
        if item == '#':
            rows.append(i)
            columns.append(j)

total = 0
millions = 0

row_gaps = {}

for i in range(len(rows)):
    for j in range(i + 1, len(rows)):
        total += abs(rows[i] - rows[j])
        pos_i = bisect.bisect_left(empty_rows, rows[i])
        pos_j = bisect.bisect_left(empty_rows, rows[j])
        gap = abs(pos_i - pos_j)
        row_gaps[gap] = row_gaps.get(gap, 0) + 1
        millions += gap

col_gaps = {}

for i in range(len(columns)):
    for j in range(i + 1, len(columns)):
        total += abs(columns[i] - columns[j])
        pos_i = bisect.bisect_left(empty_columns, columns[i])
        pos_j = bisect.bisect_left(empty_columns, columns[j])
        gap = abs(pos_i - pos_j)
        col_gaps[gap] = col_gaps.get(gap, 0) + 1
        millions += gap

# print("Total:", total)
# print("Total:", total + millions)

total += 999999 * millions

print("Total:", total)
# print("Millions:", millions)
# print("Row Gaps:", row_gaps)
# print("Column Gaps:", col_gaps)