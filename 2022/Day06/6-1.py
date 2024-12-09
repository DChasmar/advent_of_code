def fourUnique():
    s = ''
    with open('./2022/Day06/input/6.txt', 'r') as file:
        for line in file:
            s = line.strip()
    
    for i in range(len(s) - 3):
        window = s[i:i+4]
        if len(set(window)) == 4:
            return i + 4
    
    return -1

result = fourUnique()
print(result)