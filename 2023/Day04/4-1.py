numbers_array = []
sumVal = 0
with open('./Day04/input/4.txt', 'r') as file:
    for line in file:

        # Split the modified line by ':' to extract numerical values
        parts = line.split(':')

        # Iterate through parts to find values after ':'
        for part in parts[1:]:
            # Extract numerical values and add to the array
            numbers = [int(value) for value in part.split() if value.isdigit()]
            numbers_array.append(numbers)

for array in numbers_array:
    array_length = len(array)
    set_length = len(set(array))
    difference = array_length - set_length
    if difference > 0:
        sumVal += 2 ** (difference - 1)

print(sumVal)