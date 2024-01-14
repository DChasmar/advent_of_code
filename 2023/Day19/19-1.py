workflows = {}

with open('./Day19/inputs/19-workflows.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        p = line.split('{')
        p[1] = p[1].strip('}\n')
        workflow = p[1].split(',')
        workflows[p[0]] = workflow

parts = []

with open('./Day19/inputs/19-parts.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        line = line.replace('{', '').replace('}', '')
        array = [item.strip() for item in line.replace(',', '=').split('=')]
        part = []
        for i in range(1, 8, 2):
            part.append(int(array[i]))
        parts.append(part)

# [x, m, a, s]

def test_condition(part, string, more, index):
    pieces = string.split(':')
    if more:
        if part[index] > int(pieces[0][2:]): return True, pieces[1]
        else: return False, None
    elif not more:
        if part[index] < int(pieces[0][2:]): return True, pieces[1]
        else: return False, None
    else: print(f'error: neither test_condition passed: part: {part}, string: {string}, more: {more}, index: {index}')

def get_index(step):
    if step[0] == 'x': return 0
    elif step[0] == 'm': return 1
    elif step[0] == 'a': return 2
    elif step[0] == 's': return 3

def traverse_workflows(part, key):
    steps = workflows[key]
    for step in steps:
        if len(step) == 1:
            if step == 'A': return sum(part)
            elif step == 'R': return 0
            else: print('error on line 31')
        elif step[1] in ['>', '<']:
            if step[1] == '>': test, result = test_condition(part, step, True, get_index(step))
            elif step[1] == '<': test, result = test_condition(part, step, False, get_index(step))
            if not test: continue
            if result == 'A': return sum(part)
            elif result == 'R': return 0
            else:
                if type(part) is list:
                    stack.append([part, result])
                    return 0
                else: print(f'error: part: {part}')
        elif len(step) > 1 and len(step) < 4:
            stack.append([part, step])
            return 0
        else:
            print(f'error: part: {part}, key: {key}')
            return 0
    print(f'error: all steps done: part: {part}, key: {key}')

total = 0

stack = []

for part in parts:
    stack.append([part, 'in'])

while len(stack) > 0:
    latest = stack.pop()
    total += traverse_workflows(*latest)

print(total)