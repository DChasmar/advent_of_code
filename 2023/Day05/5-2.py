with open('./Day05/inputs/5-seeds.txt', 'r') as file:
    for line in file:
        seeds = [int(value) for value in line.split() if value.isdigit()]

step_arrays = [[] for _ in range(7)]
location_low = float('inf')

for i in range(7):
    with open(f'./Day05/inputs/5-s{i + 1}.txt', 'r') as file:
        for line in file:
            new_array = [int(value) for value in line.split() if value.isdigit()]
            step_arrays[i].append(new_array)

def search_steps(value, step):
    if step > 6: return value
    for array in step_arrays[step]:
        if value >= array[1] and value < array[1] + array[2]:
            difference = array[0] - array[1]
            return search_steps(value + difference, step + 1)
    return search_steps(value, step + 1)

ranges = []
for i in range(0, len(seeds), 2):
    ranges.append((seeds[i], seeds[i] + seeds[i + 1]))

def find_min(range, step):
    closest = float('inf')
    for array in step_arrays[step]:
        if range[0] <= array[1] and range[1] > array[1]:
            closest = min(closest, array[1])
    if closest == float('inf'):
        return [range, step + 1]
    else:
        queue.append([(closest, range[1]), step])
        return [(range[0], closest - 1), step + 1]

def form_ranges(range, step):
    if step > 6: return range[0], range[1]
    for array in step_arrays[step]:
        if range[0] >= array[1] and range[0] < array[1] + array[2]:
            difference = array[0] - array[1]
            if range[1] < array[1] + array[2]:
                return form_ranges((range[0] + difference, range[1] + difference), step + 1)
            else:
                queue.append([(array[1] + array[2], range[1]), step])
                return form_ranges((range[0] + difference, array[1] + array[2] - 1 + difference), step + 1)
    array = find_min(range, step)
    return form_ranges(array[0], array[1])

queue = []

for range in ranges:
    queue.append([range, 0])

while len(queue) > 0:
    array = queue.pop()
    low, high = form_ranges(array[0], array[1])
    location_low = min(location_low, low, high)

print(location_low)

# With this one, one’s initial thought might be that the low in part 2
# must be less than the low in part 1, since there are so many more
# inputs in part 2. However, that is not necessarily the case. It could
# be that the low in part 1 was derived from a seed that we do not
# have in part 2, since that seed value now accounts for the range
# from the prior seed.

# We solve this problem by accounting for all the ranges of values (initially 10)
# that undergo the same mapping from one stage to the next. Once we have all
# the ranges that undergo the same mapping, we test the low numbers of each
# range, and take the lowest location that results from starting at those lows.