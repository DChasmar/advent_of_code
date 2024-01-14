array = []
with open('./Day6/input/6.txt', 'r') as file:
    for line in file:
        # Split the modified line by ':' to extract numerical values
        parts = line.split(':')

        # Iterate through parts to find values after ':'
        for part in parts[1:]:
            # Extract numerical values and add to the array
            numbers = [int(value) for value in part.split() if value.isdigit()]
            array.append(numbers)

product = 1

for i, val in enumerate(array[0]):
    count = 0
    for j in range(val):
        if j * (val - j) > array[1][i]:
            count += 1
    product *= count


print(product)