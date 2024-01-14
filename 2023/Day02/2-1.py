sum = 0

with open('./Day02/input/2.txt', 'r') as file:
    for line in file:
        # Remove ':', ';', and ',' from the line
        cleaned_line = line.replace(':', '').replace(';', '').replace(',', '')

        # Split the cleaned line by space ' '
        elements = cleaned_line.split()

        # Convert string numbers to integers
        elements = [int(elem) if elem.isdigit() else elem for elem in elements]

        terminate_loop = False

        for i, val in enumerate(elements):
            if (i != 1) and isinstance(val, int):
                if val >= 15:
                    terminate_loop = True
                    break
                elif val >= 14 and (elements[i + 1] == 'red' or elements[i + 1] == 'green'):
                    terminate_loop = True
                    break
                elif val >= 13 and (elements[i + 1] == 'red'):
                    terminate_loop = True
                    break
        
        if not terminate_loop:
            sum += elements[1]

print(sum)