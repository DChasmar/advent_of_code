# Read the .txt file

with open('./Day1/input/1.txt', 'r') as words_file:
    words = words_file.read().split()

sum = 0

def find_first(word):
    for i, char in enumerate(word):
        # Check for digit
        if char.isdigit():
            return int(char)

def find_last(word):
    for i in range(len(word) - 1, -1, -1):
        # Check for digit
        if word[i].isdigit():
            return int(word[i])

for word in words:
    total = find_first(word) * 10 + find_last(word)
    sum += total

print(sum)