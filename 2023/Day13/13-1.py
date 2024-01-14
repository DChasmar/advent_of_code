from collections import defaultdict

with open('./Day13/input/13.txt', 'r') as file:
    lines = file.readlines()

grids = [[]]
for line in lines:
    if len(line.strip()) == 0:
        grids.append([])
    else:
        grids[-1].append(line.strip())

def find_mirror(array):
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
        mirror_index = int((key + 1) / 2)
        mirror_count = int(min(mirror_index, length - mirror_index))
        if mirror_count == data[key]:
            image = True
            break
    return image, mirror_index


def check_columns(array):
    array_length = len(array)
    string_length = len(array[0])
    new_array = []
    for i in range(string_length):
        new_string = ''
        for j in range(array_length):
            new_string += array[j][i]
        new_array.append(new_string)
    return find_mirror(new_array)


def check_rows(array):
    image, mirror_index = find_mirror(array)
    if image: return 100 * mirror_index
    else:
        image, mirror_index = check_columns(array)
        return mirror_index

total = 0
for array in grids:
    total += check_rows(array)

print(total)