with open('./Day15/input/15.txt', 'r') as file:
    for line in file:
        parts = line.split(',')

ss = []
for part in parts:
    ss.append(part)

def ops(total, val):
    total += ord(val)
    total *= 17
    total %= 256
    return total

total = 0

for s in ss:
    calc = 0
    for char in s:
        calc = ops(calc, char)
    total += calc

print(total)