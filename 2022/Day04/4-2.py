import re

def findOverlap():
    total = 0
    with open('./2022/Day04/input/4.txt', 'r') as file:
        for line in file:
            line = line.strip()
            parts = [int(x) for x in re.split(',|-', line)]
            print(parts)
            if parts[1] < parts[2] or parts[0] > parts[3]:
                continue
            total += 1  
    return total

result = findOverlap()
print(result)