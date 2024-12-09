def rockPaperScissors():
    rps1 = {"AX": 4, "AY": 8, "AZ": 3, "BX": 1, "BY": 5, "BZ": 9, "CX": 7, "CY": 2, "CZ": 6}
    rps2 = {"AX": 3, "AY": 4, "AZ": 8, "BX": 1, "BY": 5, "BZ": 9, "CX": 2, "CY": 6, "CZ": 7}
    total1 = 0
    total2 = 0
    with open('./2022/Day02/input/2.txt', 'r') as file:
        for line in file:
            line = line.strip()
            line = line.replace(" ", "")
            total1 += rps1[line]
            total2 += rps2[line]

    return total1, total2

part1, part2 = rockPaperScissors()
print(f"Part 1: {part1}")
print(f"Part 2: {part2}")