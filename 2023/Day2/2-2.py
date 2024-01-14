sum = 0

with open('./Day2/input/2.txt', 'r') as file:
    for line in file:
        # Remove ':', ';', and ',' from the line
        cleaned_line = line.replace(':', '').replace(';', '').replace(',', '')

        # Split the cleaned line by space ' '
        elements = cleaned_line.split()

        # Convert string numbers to integers
        elements = [int(elem) if elem.isdigit() else elem for elem in elements]
        
        maxR = 0
        maxG = 0
        maxB = 0
        
        for i, val in enumerate(elements):
            
            if (i != 1) and isinstance(val, int):
                if elements[i + 1] == 'red':
                    maxR = max(maxR, val)
                elif elements[i + 1] == 'green':
                    maxG = max(maxG, val)
                elif elements[i + 1] == 'blue':
                    maxB = max(maxB, val)
        
        power = maxR * maxG * maxB
        sum += power

print(sum)