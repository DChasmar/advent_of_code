def fourUnique():
    s = ''
    with open('./2022/Day07/input/7.txt', 'r') as file:
        for line in file:
            s = line.strip()
    
    for i in range(len(s) - 13):
        window = s[i:i+14]
        if len(set(window)) == 14:
            return i + 14
    
    return -1

result = fourUnique()
print(result)