from queue import Queue
import math

flip_flops = {}
conjunctions = {}
broadcaster = ['gz', 'xg', 'cd', 'sg']

with open('./Day20/input/20.txt', 'r') as file:
    lines = file.readlines()
    for line in lines:
        line = line.strip('\n')
        line = line.replace(' ', '')
        parts = line.split('->')
        destinations = parts[1].split(',')
        if parts[0][0] == '&': conjunctions[parts[0][1:]] = destinations
        elif parts[0][0] == '%': flip_flops[parts[0][1:]] = destinations
        else: continue

ff_memory = {key: False for key in flip_flops} # False means off; True means on
pulse_memory = {key: False for key in (flip_flops.keys() | conjunctions.keys())} # False means low; True means high (most recent pulse received)
pulse_memory['rx'] = False
c_inputs = {conjunction: [] for conjunction in conjunctions} # inputs for each conjunction

for conjunction in conjunctions: c_inputs[conjunction] = []
    
for key in flip_flops:
    for value in flip_flops[key]:
        if value in conjunctions:
            c_inputs[value].append(key)

for key in conjunctions:
    for value in conjunctions[key]:
        if value in conjunctions:
            c_inputs[value].append(key)

q = Queue()

press_counts = []

def process_flip_flop(item, pulse):
    if pulse: return
    elif not pulse:
        if ff_memory[item]:
            ff_memory[item] = False
            pulse_memory[item] = False
            for val in flip_flops[item]:
                q.put([val, False])
        elif not ff_memory[item]:
            ff_memory[item] = True
            pulse_memory[item] = True
            for val in flip_flops[item]:
                q.put([val, True])

def process_conjunction(item, pulse, i):
    all = True
    for input in c_inputs[item]:
        if not pulse_memory[input]:
            all = False
            break
    if all:
        pulse_memory[item] = False
        for val in conjunctions[item]: q.put([val, False])
    elif not all:
        pulse_memory[item] = True
        if item in ['dd', 'fh', 'xp', 'fc']:
            # print(f'{item} sends high in this interval: {i}')
            press_counts.append(i + 1)
        for val in conjunctions[item]: q.put([val, True])

# Each conjunction that is an input for &dn, which is the only input 
# for rx sends one high pulse between 0 and 5000 button presses:
# fc sends high in this interval: 3917
# xp sends high in this interval: 3919
# dd sends high in this interval: 4003
# fh sends high in this interval: 4027
# the final solution is the lowest common multiple of these numbers
# since each number is prime, the solution is also the product of the four numbers

for i in range(5000):
    for val in broadcaster: q.put([val, False])
    while not q.empty():
        result = q.get()
        next, pulse = result[0], result[1]
        if next in flip_flops: process_flip_flop(next, pulse)
        elif next in conjunctions: process_conjunction(next, pulse, i)
        elif next == 'rx':
            if not pulse: print(f'{next} this many times: {i}')
        else: print('error')
    
print(math.lcm(*press_counts))