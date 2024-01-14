# This function checks for an adjacent symbol by searching for a character that is
# not a dot nor a digit.
def find_symbol(x1, y1, x2, y2):
    # These indices track all the points on the grid that surround the number, including the number itself
    for i in range(y1, y2 + 1):
        for j in range(x1, x2 + 1):
            # ignore digits and dots
            if not lines[i][j].isdigit() and not lines[i][j] == '.':
                return True
    return False

# This function finds the whole number based on the index of the leftmost digit in the number.
def get_number(row, index, grid):
    number = int(grid[row][index])
    # also_index tracks the index of the next digit in the number, and ultimately, the index after the number
    also_index = index + 1
    while grid[row][also_index].isdigit():
        # mulitiply the current number by 10 and add the new digit
        number = 10 * number + int(grid[row][also_index]) 
        also_index += 1
    is_adjacent = find_symbol(index - 1, row - 1, also_index, row + 1)
    if is_adjacent: return number
    else: return 0

# This number includes the dots added at the beginning and end of each array (140 + 2)
GRID_SIZE = 142

dots_array = ['.' for _ in range(GRID_SIZE)]
lines = []
# add an array of dots at the BEGINNING to avoid out-of-bounds error
lines.append(dots_array)
sumVal = 0

with open('./Day3/input/3.txt', 'r') as file:
    for line in file:
        for char in line:
            # add dots at the BEGINNING and END of the array to avoid out-of-bounds error
            new_line = ['.'] + list(line.strip()) + ['.'] 
        lines.append(new_line)
    # add an array of dots at the END to avoid out-of-bounds error
    lines.append(dots_array)

    for row, array in enumerate(lines):
        for col, char in enumerate(array):
            if char.isdigit():
                newSum = sumVal + get_number(row, col, lines)
                if newSum > sumVal:
                    deleteIndex = col
                    # Avoid counting a number multiple times;
                    # We want just 325, not 325, 25 and 5
                    while lines[row][deleteIndex].isdigit():
                        lines[row][deleteIndex] = '.'
                        deleteIndex += 1
                    sumVal = newSum

print(sumVal)