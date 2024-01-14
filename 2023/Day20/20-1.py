from queue import Queue

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

low, high = 0, 0

def process_flip_flop(item, pulse):
    global low
    global high
    if pulse: return
    elif not pulse:
        if ff_memory[item]:
            ff_memory[item] = False
            pulse_memory[item] = False
            for val in flip_flops[item]:
                q.put([val, False])
        elif not ff_memory[item]:
            ff_memory[item] = True
            pulse_memory[item] = True # This was mistakenly recorded as False for part 1; 
                                      # it had not effect on the outcome of part 1, but it does impact 
                                      # part 2, and so caused great annoyance and confusion
            for val in flip_flops[item]:
                q.put([val, True])

def process_conjunction(item, pulse):
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
        for val in conjunctions[item]: q.put([val, True])

highs, lows = [], []

for i in range(1000):
    low += 1
    for val in broadcaster: q.put([val, False])
    while not q.empty():
        result = q.get()
        next, pulse = result[0], result[1]
        if pulse: high += 1
        elif not pulse: low += 1
        if next in flip_flops: process_flip_flop(next, pulse)
        elif next in conjunctions: process_conjunction(next, pulse)
        elif next == 'rx': pulse_memory[next] = pulse
        else: print('error')
    highs.append(high)
    lows.append(low)

print(highs[-1] * lows[-1])


# last ouput: highs: 41290 * lows: 17806 = 735209740 (too low)
# last ouput: highs: 41290 * lows: 24365 = 1006030850 (too high)
# last ouput: highs: 44089 * lows: 21566 = 950823374 (too high)
# last ouput: highs: 44089 * lows: 22566 = 994912374 (should be too high as well)
# last ouput: highs: 47849 * lows: 18806 = 899848294 (correct! after realizing a misunderstanding of the rules for pulse memory)