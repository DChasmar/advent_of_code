import re

def parse_monkey_data(file_path):
    with open(file_path, 'r') as file:
        lines = file.readlines()

    monkey_data = {}
    current_monkey = None

    for line in lines:
        line = line.strip()
        if line.startswith('Monkey'):
            current_monkey = int(line.split(' ')[1][:-1])
            monkey_data[current_monkey] = {}
        elif line.startswith('Starting items:'):
            items = list(map(int, line.split(':')[1].split(',')))
            monkey_data[current_monkey]['starting_items'] = items
        elif line.startswith('Operation:'):
            operation = line.split(':')[1].strip()
            monkey_data[current_monkey]['operation'] = operation
        elif line.startswith('Test:'):
            test = line.split(':')[1].strip()
            monkey_data[current_monkey]['test'] = test
        elif line.startswith('If true:'):
            if_true = int(re.search(r'\d+', line).group())
            monkey_data[current_monkey]['if_true'] = if_true
        elif line.startswith('If false:'):
            if_false = int(re.search(r'\d+', line).group())
            monkey_data[current_monkey]['if_false'] = if_false

    return monkey_data

file_path = './2022/Day11/input/11.txt'
monkey_data = parse_monkey_data(file_path)
print(monkey_data)