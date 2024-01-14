# This function checks for an adjacent symbol by searching for a character that is
# not a dot nor a digit.
def find_symbol(x1, y1, x2, y2):
    # These indices track all the points on the grid that surround the number, including the number itself
    for i in range(y1, y2 + 1):
        for j in range(x1, x2 + 1):
            if lines[i][j] == '*':
                # Multiply the row (i) by 1000 and add the column (j) merely to create a unique
                # number for the index in the grid
                return int(1000 * i + j)
    return None

# This function finds the whole number based on the index of the leftmost digit in the number.
def get_number(line, index, grid):
    number = int(grid[line][index])
    # also_index tracks the index of the next digit in the number, and ultimately, the index after the number
    also_index = index + 1
    while grid[line][also_index].isdigit():
        # mulitiply the current number by 10 and add the new digit
        number = 10 * number + int(grid[line][also_index])
        also_index += 1
    return number, find_symbol(index - 1, line - 1, also_index, line + 1)

# This number includes the dots added at the beginning and end of each array (140 + 2)
GRID_SIZE = 142

dots_array = ['.' for _ in range(GRID_SIZE)]
lines = []
lines.append(dots_array)
sumVal = 0
my_dict = {}

with open('./Day3/input/3.txt', 'r') as file:
    for line in file:
        for char in line:
            new_line = ['.'] + list(line.strip()) + ['.']
        lines.append(new_line)
    lines.append(dots_array)

    for row, array in enumerate(lines):
        for col, char in enumerate(array):
            if char.isdigit():
                newNumber, key = get_number(row, col, lines)
                if key != None:
                    # We record the occurence of a number adjacent to symbol '*';
                    # This symbol may or may not be adjacent to another number
                    if key in my_dict:
                        # second occurence
                        my_dict[key]['product'] *= newNumber
                        my_dict[key]['touches'] += 1
                    # first occurence
                    else: my_dict[key] = {'product': newNumber, 'touches': 1}
                deleteIndex = col
                # Avoid counting a number multiple times;
                # We want just 325, not 325, 25 and 5
                while lines[row][deleteIndex].isdigit():
                    lines[row][deleteIndex] = '.'
                    deleteIndex += 1
    
    for key in my_dict:
        if my_dict[key]['touches'] == 2:
            sumVal += my_dict[key]['product']

print(sumVal)