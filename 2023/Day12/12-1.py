def solve(line):
    key, values = line.strip().split(" ")
    # Add a period at the beginning and end of the key string
    key = '.' + key + '.'
    nums = list(map(int, values.split(',')))

    # key_index is index of the char in the string we are now processing
    # next_index is the index of the next number in nums (relative to the one we are processing)
    # remaining is what remains of the number in nums currently being processed (we decrement this value)
    
    # process is a dynamic programming function that goes through each character in the string 'key'
    def process(index, next_index, remaining):
        # Base Case 1: We are at the end of the word
        if index == len(key):
            # Have we used up all the numbers in nums?
            if next_index == len(nums) and remaining == 0: return 1
            # If not, this arrangement is not possible
            return 0
        # Base Case 2: We have exhausted the numbers in nums; this arrangement is not possible
        # This line is actually unnecessary given conditions in code below
        if next_index > len(nums): return 0

        total = 0
        if key[index] == '.':
            if remaining == 0:
                # Is there another num left in nums?
                if next_index < len(nums):
                    total += process(index + 1, next_index + 1, nums[next_index])
                total += process(index + 1, next_index, 0)
        elif key[index] == '#':
            if remaining > 0:
                total += process(index + 1, next_index, remaining - 1)
        else:
            # this is a '?'
            # choose '.' if the streak is over
            if remaining == 0:
                if next_index < len(nums):
                    total += process(index + 1, next_index + 1, nums[next_index])
                total += process(index + 1, next_index, 0)
            # choose '#' if there is still a streak (remaining is greater than 0)
            if remaining > 0:
                total += process(index + 1, next_index, remaining - 1)
        
        return total
    
    # Starts the processing of the characters in the string
    return process(0, 0, 0)

final_total = 0

with open('./Day12/input/12.txt', 'r') as file: lines = file.readlines()
for line in lines: final_total += solve(line)

print(final_total)