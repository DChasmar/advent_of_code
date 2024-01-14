from math import lcm

with open('./Day8/inputs/8.txt', 'r') as file:
    lines = file.readlines()

# Create a dictionary from the lines
data = {}
for line in lines:
    key, values = line.strip().split(" = ")
    values_tuple = tuple(values[1:-1].split(", "))
    data[key] = values_tuple

with open('./Day8/inputs/8-LR.txt', 'r') as file:
    for line in file:
        LR = str(line)

# Add all nodes ending with A to starts
starts = []
for key in data:
    if key.endswith('A'): starts.append(key)

count = 0
length = len(LR)
z_counts = []

# There are 6 nodes that end in A and 6 nodes that end in Z
# Before 23000 cycles, each node beginning with A maps to a node ending
# with Z once (after x maps) and will return to that node every x times
while count < 23000:
    for i, location in enumerate(starts):
        index = 0 if LR[count % length] == 'L' else 1
        newLocation = data[location][index]
        if newLocation[2] =='Z': z_counts.append(count + 1)
        starts[i] = newLocation
    count += 1

# Take the lowest common multiple of the 6 z-counts to get the count when each will
# map to a node ending with Z
result = lcm(*z_counts)
print(result)