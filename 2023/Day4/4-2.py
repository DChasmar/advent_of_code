numbers_array = []
cards_array = [1] * 213
with open('./Day4/input/4.txt', 'r') as file:
    for line in file:

        # Split the modified line by ':' to extract numerical values
        parts = line.split(':')

        # Iterate through parts to find values after ':'
        for part in parts[1:]:
            # Extract numerical values and add to the array
            numbers = [int(value) for value in part.split() if value.isdigit()]
            numbers_array.append(numbers)

for i, array in enumerate(numbers_array):
    array_length = len(array)
    set_length = len(set(array))
    difference = array_length - set_length
    
    while difference > 0:
        cards_array[i + difference] += cards_array[i]
        difference -= 1

print(sum(cards_array))