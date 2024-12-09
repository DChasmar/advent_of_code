def checkAllCalorieLoads():
    all_loads = []
    current_load = 0
    with open('./2022/Day01/input/1.txt', 'r') as file:
        for line in file:
            if line.strip():  # Check if the line is not empty
                current_load += int(line)
            else:
                all_loads.append(current_load)
                current_load = 0
    all_loads.sort(reverse=True)  # Sort the list in descending order
    return sum(all_loads[:3])

result = checkAllCalorieLoads()
print(result)