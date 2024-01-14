with open('./Day09/input/9.txt', 'r') as file:
    lines = file.readlines()

total = 0
for line in lines:
    numbers =[[int(num) for num in line.split()]]
    while len(set(numbers[-1])) != 1:
        numbers.append([])
        for i in range(len(numbers[-2]) - 1):
            numbers[-1].append(numbers[-2][i + 1] - numbers[-2][i])
    for i in range(len(numbers) - 1):
        numbers[-2 - i].append(numbers[-1 - i][-1] + numbers[-2 - i][-1])
    total += numbers[0][-1]

print(total)