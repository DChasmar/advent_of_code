workflows = {}

with open('./Day19/inputs/19-workflows.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        p = line.split('{')
        p[1] = p[1].strip('}\n')
        workflow = p[1].split(',')
        workflows[p[0]] = workflow

# [x, m, a, s]

def test_condition(part, string, more, index, current_workflow, step):
    pieces = string.split(':')
    test_number = int(pieces[0][2:])
    low = part[index][0]
    high = part[index][1]
    next_workflow = pieces[1]
    if more:
        if low > test_number:
            stack.append([part, next_workflow, 0])
            return
        elif high <= test_number:
            stack.append([part, current_workflow, step + 1])
            return
        elif high > test_number and low <= test_number:
            new_part1 = part.copy()
            new_part1[index] = (low, test_number)
            new_part2 = part.copy()
            new_part2[index] = (test_number + 1, high)
            stack.append([new_part1, current_workflow, step + 1])
            stack.append([new_part2, next_workflow, 0])
            return
        else: print(f'error: ???')
    elif not more:
        if high < test_number:
            stack.append([part, next_workflow, 0])
            return
        elif low >= test_number:
            stack.append([part, current_workflow, step + 1])
            return
        elif high >= test_number and low < test_number:
            new_part1 = part.copy()
            new_part1[index] = (low, test_number - 1)
            new_part2 = part.copy()
            new_part2[index] = (test_number, high)
            stack.append([new_part1, next_workflow, 0])
            stack.append([new_part2, current_workflow, step + 1])
            return
        else: print(f'error: ?????')
    else: print(f'error: neither test_condition passed: part: {part}, string: {string}, more: {more}, index: {index}')

def get_index(step):
    if step[0] == 'x': return 0
    elif step[0] == 'm': return 1
    elif step[0] == 'a': return 2
    elif step[0] == 's': return 3

def traverse_workflows(part, key, step):
    if key == 'A':
        good.append(part)
        return
    elif key == 'R':
        bad.append(part)
        return
    else:
        steps = workflows[key]
        s = steps[step]
        if s == 'A':
            good.append(part)
            return
        elif s == 'R':
            bad.append(part)
            return
        if s[1] in ['>', '<']:
            if s[1] == '>': test_condition(part, s, True, get_index(s), key, step)
            elif s[1] == '<': test_condition(part, s, False, get_index(s), key, step)
        elif len(s) > 1 and len(s) < 4:
            stack.append([part, s, 0])
            return
        else:
            print(f'error: part: {part}, key: {key}')
            return

total = 0

stack = [[[(1, 4000), (1, 4000), (1, 4000), (1, 4000)], 'in', 0]]

good = []
bad = []

while len(stack) > 0:
    latest = stack.pop()
    traverse_workflows(latest[0], latest[1], latest[2])

for array in good:
    total += (array[0][1] - array[0][0] + 1) * (array[1][1] - array[1][0] + 1) * (array[2][1] - array[2][0] + 1) * (array[3][1] - array[3][0] + 1)

print(total)