def search_steps(value, step):
    if step == 7: return value
    for array in step_arrays[step]:
        if value >= array[1] and value < array[1] + array[2]:
            difference = array[0] - array[1]
            return search_steps(value + difference, step + 1)
    return search_steps(value, step + 1)

with open('./Day5/inputs/5-seeds.txt', 'r') as file:
    for line in file:
        seeds = [int(value) for value in line.split() if value.isdigit()]

step_arrays = [[] for _ in range(7)]

for i in range(7):
    with open(f'./Day5/inputs/5-s{i + 1}.txt', 'r') as file:
        for line in file:
            new_array = [int(value) for value in line.split() if value.isdigit()]
            step_arrays[i].append(new_array)

location_low = float('inf')

for seed in seeds:
    result = search_steps(seed, 0)
    if result is not None:
        location_low = min(location_low, result)

print(location_low)