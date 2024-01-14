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

boxes = [[[],[]] for _ in range(256)]

def calc_me(chars):
    calc = 0
    for char in chars:
        calc = ops(calc, char)
    return calc

def add_lens(chars, box, lens):
    if chars in boxes[box][0]:
        index = boxes[box][0].index(chars)
        boxes[box][1][index] = lens
    else:
        boxes[box][0].append(chars)
        boxes[box][1].append(lens)

def remove_lens(chars, box):
    if chars in boxes[box][0]:
        index = boxes[box][0].index(chars)
        del boxes[box][0][index]
        del boxes[box][1][index]

for s in ss:
    if s[-1] == '-':
        chars = s[:-1]
        box = calc_me(chars)
        remove_lens(chars, box)
    elif s[-2] == '=':
        pieces = s.split('=')
        chars = pieces[0]
        box = calc_me(chars)
        lens = int(pieces[1])
        add_lens(chars, box, lens)

total = 0

for i, box in enumerate(boxes):
    for j, lens in enumerate(box[1]):
        total += (i + 1) * (j + 1) * lens

print (total)