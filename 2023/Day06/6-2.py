array = []
with open('./Day06/input/6.txt', 'r') as file:
    for line in file:
        # Split the modified line by ':' to extract numerical values
        parts = line.split(':')

        # Iterate through parts to find values after ':'
        for part in parts[1:]:
            # Extract numerical values and add to the array
            number = part.replace(" ", "")
            array.append(int(number))

count = 0

val = array[0] // 2

while val * (array[0] - val) > array[1]:
    count += 1
    val -= 1

val = array[0] // 2 + 1

while val * (array[0] - val) > array[1]:
    count += 1
    val += 1

print(count)

# There is no doubt a more economical method of solving than just incrementing
# by one. A binary search approach would be better. This code takes a few 
# seconds to provide the solution.