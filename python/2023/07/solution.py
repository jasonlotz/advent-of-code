from collections import Counter
from pprint import pprint

INPUT_PATH = "input-files/2023/07/input.txt"

HAND_TYPES = {
    (5,): 7,  # Five of a kind
    (4, 1): 6,  # Four of a kind
    (3, 2): 5,  # Full house
    (3, 1, 1): 4,  # Three of a kind
    (2, 2, 1): 3,  # Two pairs
    (2, 1, 1, 1): 2,  # One pair
    (1, 1, 1, 1, 1): 1  # High card
}

CARD_VALUES = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2
}


def hand_value(hand: str) -> int:
    hand_counter = Counter(hand)
    hand_tuple = tuple(sorted(hand_counter.values(), reverse=True))

    # Calculate hand value
    hand_value = HAND_TYPES[hand_tuple]

    # Calculate high card score
    high_card_score = 0
    multiplier = 10000000000
    for card in hand:
        high_card_score += CARD_VALUES[card] * multiplier
        multiplier /= 100

    return hand_value, high_card_score


def main():
    inputs = open(INPUT_PATH).read().splitlines()

    hands = []
    total_winnings = 0

    for inputs in inputs:
        hand, bid = inputs.split()

        value, high_card_score = hand_value(hand)

        hands.append((hand, value, high_card_score, int(bid)))
        hands.sort(key=lambda x: (x[1], x[2]))

    for i, hand in enumerate(hands, 1):
        print(i, hand, hand[3] * i)
        total_winnings += hand[3] * i

    print(total_winnings)


if __name__ == "__main__":
    main()
