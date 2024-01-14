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

print(empty_columns)
print(empty_rows)

for i, val in enumerate(empty_columns): empty_columns[i] = val + 1
for i, val in enumerate(empty_rows): empty_rows[i] = val + 1

print(empty_columns)
print(empty_rows)

row_length = len(lines) + len(empty_columns)

new_grid = []
for i, array in enumerate(grid):
    if i in empty_rows: new_grid.append(['.' for _ in range(row_length)])
    new_array = []
    for j, item in enumerate(array):
        if j in empty_columns: new_array.append('.')
        if item != '': new_array.append(item)
    new_grid.append(new_array)

rows = []
columns = []

for i, array in enumerate(new_grid):
    for j, item in enumerate(array):
        if item == '#':
            rows.append(j)
            columns.append(i)

total = 0

for i in range(len(rows)):
    for j in range(i + 1, len(rows)):
        total += abs(rows[i] - rows[j])

for i in range(len(columns)):
    for j in range(i + 1, len(columns)):
        total += abs(columns[i] - columns[j])

print(total)

new_empty_rows = []
new_empty_columns = [num for num in range(151)]
for i, array in enumerate(new_grid):
    count = 0
    for j, item in enumerate(array):
        if item == '#':
            count += 1
            if j in new_empty_columns: new_empty_columns.remove(j)
    if count == 0: new_empty_rows.append(i)

print(new_empty_columns)
print(new_empty_rows)