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

def changeImage(image, x, count):
    row = count // 40
    col = count % 40
    if abs(col - x) <= 1:
        image[row][col] = '#'
    return image


def solve():
    data = getData()
    image = [40 * ['.'] for _ in range(6)]
    x = 1
    count = 0
    for tuple in data:
        image = changeImage(image, x, count)
        if tuple[0] == 'noop':
            count += 1
        elif tuple[0] == 'addx':
            for j in range(2):                
                count += 1
                if j == 0:
                    image = changeImage(image, x, count)
            x += tuple[1]
        else:
            print("Error")
    return image

result = solve()
file_path = './2022/Day10/output/image.txt'
with open(file_path, 'w') as file:
    for row in result:
        row_str = ''.join(map(str, row))
        file.write(row_str + '\n')