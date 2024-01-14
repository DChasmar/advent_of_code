hands = []
bets = []
cards = ['J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A']
# Change J to having the lowest value
cards_object = {'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}
with open('./Day7/input/7.txt', 'r') as file:
    for line in file:
        # Split the modified line by ':' to extract numerical values
        parts = line.split(' ')

        # Iterate through parts to find values after ':'
        for i, part in enumerate(parts):
            if i == 0: hands.append(part)
            elif i == 1: bets.append(int(part))

# Value sections
high_card = []
pair = []
two_pair = []
triple = []
full_house = []
quads = []
fives = []

# add jack_count to count the number of Js in each hand
def jack_count(hand):
    count = 0
    for char in hand:
        if char == 'J': count += 1
    return count

# Check the occurrences of a card value in the hand
def check_char_counts(hand, number):
    char_counts = {}
    for char in hand:
        if char in char_counts:
            char_counts[char] += 1
        else:
            char_counts[char] = 1
    for char, count in char_counts.items():
        if count == number: return True
    return False

# Determine the value section to which each hand of cards belongs
# with the help of function check_char_counts and jack_count
for hand in hands:
    unique_cards = len(set(hand))
    if unique_cards == 5:
        if jack_count(hand) > 0: pair.append(hand)
        else: high_card.append(hand)
    elif unique_cards == 4:
        if jack_count(hand) > 0: triple.append(hand)
        else: pair.append(hand)
    elif unique_cards == 3:
        if check_char_counts(hand, 3):
            if jack_count(hand) > 0: quads.append(hand)
            else: triple.append(hand)
        else:
            if jack_count(hand) == 2: quads.append(hand)
            elif jack_count(hand) == 1: full_house.append(hand)
            else: two_pair.append(hand)        
    elif unique_cards == 2:
        if check_char_counts(hand, 4):
            if jack_count(hand) > 0: fives.append(hand)
            else: quads.append(hand)
        else:
            if jack_count(hand) > 0: fives.append(hand)
            else: full_house.append(hand) 
    elif unique_cards == 1:
        fives.append(hand)

def convert_sort_return_hand(card_array):
    # We convert the card values to letters, so we can sort them by alphabetical order
    # in order to determine their final order.
    converted_hands = [[chr(ord('a') + cards_object[char]) for char in hand] for hand in card_array]
    joined_strings = [''.join(inner_array) for inner_array in converted_hands]
    joined_strings.sort()
    separated_strings = [list(s) for s in joined_strings]
    for hand in separated_strings:
        for i, char in enumerate(hand):
            temp = cards[ord(char) - ord('a')]
            hand[i] = temp
    rejoined_strings = [''.join(inner_array) for inner_array in separated_strings]
    return rejoined_strings

final_order = []

# Create the final order with the help of function convert_sort_return_hand
for hand in convert_sort_return_hand(high_card):
    final_order.append(hand)
for hand in convert_sort_return_hand(pair):
    final_order.append(hand)
for hand in convert_sort_return_hand(two_pair):
    final_order.append(hand)
for hand in convert_sort_return_hand(triple):
    final_order.append(hand)
for hand in convert_sort_return_hand(full_house):
    final_order.append(hand)
for hand in convert_sort_return_hand(quads):
    final_order.append(hand)
for hand in convert_sort_return_hand(fives):
    final_order.append(hand)

sum_val = 0

for i, hand in enumerate(final_order):
    hand_index = hands.index(hand)
    sum_val += (i + 1) * bets[hand_index]

print(sum_val)