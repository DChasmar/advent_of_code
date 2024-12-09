def checkAllCalorieLoads():
    max_load = 0
    current_load = 0
    with open('./2022/Day01/input/1.txt', 'r') as file:
        for line in file:
            if line.strip():  # Check if the line is not empty
                current_load += int(line)
            else:
                max_load = max(max_load, current_load)
                current_load = 0
    return max_load

result = checkAllCalorieLoads()
print(result)