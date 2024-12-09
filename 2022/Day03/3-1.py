def commonCharacter():
    total = 0
    with open('./2022/Day03/input/3.txt', 'r') as file:
        for line in file:
            s = line.strip()
            length = len(s)
            midpoint = length // 2

            part1 = s[:midpoint]
            part2 = s[midpoint:]

            common_values = set(part1) & set(part2)
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