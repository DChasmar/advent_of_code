def getData():
    data = []
    with open('./2022/Day10/input/10.txt', 'r') as file:
        for line in file:
            line = line.strip()  # Check if the line is not empty
            if line == 'noop':
                data.append((line, 0))
            else:
                parts = line.split(' ')
                data.append((parts[0], int(parts[1])))
    return data

def solve():
    data = getData()
    x = 1
    count = 0
    summ = 0
    for tuple in data:
        if count % 40 == 19:
            summ += (count + 1) * x
        if tuple[0] == 'noop':
            count += 1
        elif tuple[0] == 'addx':
            for j in range(2):                
                count += 1
                if count % 40 == 19 and j == 0:
                    summ += (count + 1) * x
            x += tuple[1]
        else:
            print("Error")
        if count > 220:
            return summ

result = solve()
print(result)

# 14340 is too high
# 7539 is too low
# 11721 is too low
# 11820 is correct