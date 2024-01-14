sum = 0

forward_dict = ['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']

def find_first(word):
    for i, char in enumerate(word):
        # Check for digit
        if char.isdigit():
            return int(char)

        # Check for specific characters
        elif char in ['o', 't', 'f', 's', 'e', 'n']:
            substrings = [word[i:i + length] for length in [3, 4, 5] if i + length <= len(word)]

            for substring in substrings:
                if substring in forward_dict:
                    return forward_dict.index(substring) + 1

def find_last(word):
    for i in range(len(word) - 1, -1, -1):
        # Check for digit
        if word[i].isdigit():
            return int(word[i])

        # Check for specific characters
        elif word[i] in ['o', 'e', 'r', 'x', 'n', 't']:
            substrings = [word[i - length + 1 :i + 1] for length in [3, 4, 5] if i - length >= -1]

            for substring in substrings:
                if substring in forward_dict:
                    return forward_dict.index(substring) + 1

# Read the .txt file
with open('./Day1/input/1.txt', 'r') as words_file:
    words = words_file.read().split()

for word in words:
    total = find_first(word) * 10 + find_last(word)
    sum += total

print(sum)