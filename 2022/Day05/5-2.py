import re

def moveBoxes():
    boxes = [
    ['G', 'T', 'R', 'W'], 
    ['G', 'C', 'H', 'P', 'M', 'S', 'V', 'W'], 
    ['C', 'L', 'T', 'S', 'G', 'M'], 
    ['J', 'H', 'D', 'M', 'W', 'R', 'F'], 
    ['P', 'Q', 'L', 'H', 'S', 'W', 'F', 'J'], 
    ['P', 'J', 'D', 'N', 'F', 'M', 'S'],
    ['Z', 'B', 'D', 'F', 'G', 'C', 'S', 'J'],
    ['R', 'T', 'B'],
    ['H', 'N', 'W', 'L', 'C']
    ]
    with open('./2022/Day05/input/5.txt', 'r') as file:
        for line in file:
            line = line.strip()
            numbers = [int(x) for x in re.findall(r'\b\d+\b', line)]
            new_box = boxes[numbers[1] - 1][-numbers[0]:].copy()
            boxes[numbers[1] - 1] = boxes[numbers[1] - 1][:-numbers[0]]
            boxes[numbers[2] - 1] = boxes[numbers[2] - 1] + new_box
    final = []
    for array in boxes:
        final.append(array[-1])
    return ''.join(final)

result = moveBoxes()
print(result)