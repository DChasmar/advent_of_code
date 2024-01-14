def solve(line):
    key, values = line.strip().split(" ")
    key_five = [key] * 5
    key = '?'.join(key_five)
    key = '.' + key + '.'
    print(key)
    nums = list(map(int, values.split(',')))
    nums *= 5
    print(nums)

    memo = {}

    # key_index is index of the char in the string we are now processing
    # next_index is the index of the next number in nums that we are processing
    # remaining is what remains of the number in nums currently being processed
    def process(index, next_index, remaining):
        if (index, next_index, remaining) in memo:
            return memo[(index, next_index, remaining)]
        if index == len(key):
            if next_index == len(nums) and remaining == 0: return 1
            return 0
        if next_index > len(nums): return 0

        total = 0
        if key[index] == '.':
            if remaining == 0:
                if next_index < len(nums):
                    total += process(index + 1, next_index + 1, nums[next_index])
                total += process(index + 1, next_index, 0)
        elif key[index] == '#':
            if remaining > 0:
                total += process(index + 1, next_index, remaining - 1)
        else:
            # this is a ?
            # choose '.' if the streak is over
            if remaining == 0:
                if next_index < len(nums):
                    total += process(index + 1, next_index + 1, nums[next_index])
                total += process(index + 1, next_index, 0)
            # choose '#' if there is still a streak left
            if remaining > 0:
                total += process(index + 1, next_index, remaining - 1)
        
        memo[(index, next_index, remaining)] = total
        return total
    
    return process(0, 0, 0)

final_total = 0
count = 0

with open('./Day12/input/12.txt', 'r') as file:
    lines = file.readlines()

for line in lines:
    count += 1
    print(count)
    result = solve(line)
    print(result)
    final_total += result

print(final_total)

# 1: 23570904
# 2: 2444990