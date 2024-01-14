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

location = 'AAA'

count = 0
length = len(LR)

while location != 'ZZZ':
    index = 0 if LR[count % length] == 'L' else 1
    location = data[location][index]
    count += 1

print(count)