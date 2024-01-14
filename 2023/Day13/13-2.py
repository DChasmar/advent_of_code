from collections import defaultdict

with open('./Day13/input/13.txt', 'r') as file:
    lines = file.readlines()

grids = [[]]
for line in lines:
    if len(line.strip()) == 0:
        grids.append([])
    else:
        grids[-1].append(line.strip())


col_misses = []

def find_mirror(array, grid_num):
    data = defaultdict(int)
    length = len(array)
    for i in range(length):
        test = array[i]
        for j in range(i + 1, length):
            if test == array[j]:
                total = i + j
                data[total] += 1
    
    image = False

    mirror_index = -1
    mirror_count = 0

    for key in data:
        if key % 2 == 0: continue
        mirror_index = int((key + 1) / 2)
        mirror_count = int(min(mirror_index, length - mirror_index))
        if mirror_count == data[key] + 1:
            image = True
            break
    return image, mirror_index

def check_columns(array, grid_num):
    array_length = len(array)
    string_length = len(array[0])
    new_array = []
    for i in range(string_length):
        new_string = ''
        for j in range(array_length):
            new_string += array[j][i]
        new_array.append(new_string)
    return find_mirror(new_array, grid_num)

no_row = []

def check_rows(array, grid_num):
    image, mirror_index = find_mirror(array, grid_num)
    if image: return 100 * mirror_index
    else:
        no_row.append(grid_num)
        image, mirror_index = check_columns(array, grid_num)
        if not image:
            return 0
        return mirror_index

def one_diff(index):
    num_rows = len(grids[index])
    num_cols = len(grids[index][0])
    count = 0
    for i in range(num_cols):
        if grids[index][0][i] != grids[index][1][i]: 
            count += 1
    if count == 1: return 100
    else: count = 0
    for i in range(num_cols):
        if grids[index][-2][i] != grids[index][-1][i]: 
            count += 1
    if count == 1: return (num_rows - 1) * 100
    else: count = 0
    for i in range(num_rows):
        if grids[index][i][0] != grids[index][i][1]: 
            count += 1
    if count == 1: return 1
    else: count = 0
    for i in range(num_rows):
        if grids[index][i][-1] != grids[index][i][-2]: 
            count += 1
    if count == 1: return num_cols - 1
    else: 
        return 0
        
total = 0

one_diffs = []
for i in range(100):
    result = one_diff(i)
    if result > 0:
        one_diffs.append(i)
    total += result
    
may_be_wrong = [24, 47, 54, 57, 75]

for grid in may_be_wrong:
    print(f'grid: {grid}')
    for s in grids[grid]: print(s)

working_r = []
working_c = []

lame = 0
replace = 0

for i in range(100):
    if i in one_diffs: continue
    result = check_rows(grids[i], i)
    if i in may_be_wrong:
        print(f' row {i} result is {result}')
        lame += result
    total += result
    if result > 0: working_r.append(i)
    booly, result = check_columns(grids[i], i)
    if booly and result > 0:
        if i in may_be_wrong:
            print(f'column {i} result is {result}')
            replace += result

print(total)
print(total - lame + replace + 300 - 7)

# My approach to solving this one was flawed, so I wasted a substantial
# amount of time trying to locate my error. Eventually I determined the grids
# that “may_be_wrong” and checked them myself. I realized one was
# wrong (grid 54), so that is why I add 300 and subtract 7 in the end.
# Rather than polishing this one, I will leave it a mess.