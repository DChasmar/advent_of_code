def commonCharacter():
    total = 0
    lines = []
    with open('./2022/Day03/input/3.txt', 'r') as file:
        for line in file:
            lines.append(line.strip())
    print(lines)
    print(len(lines))
    for i in range(0, 300, 3):
        common_values = set(lines[i]).intersection(lines[i + 1], lines[i + 2])
        if len(common_values) == 1:
            common_value = common_values.pop()
            if common_value.isupper():
                total += ord(common_value) - 38
            else:
                total += ord(common_value) - 96
        else: print("Error")

    return total

result = commonCharacter()
print(result)