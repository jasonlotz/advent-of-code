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


# Given a hand counter, return the hand value
def calc_hand_counter_value(hand_counter):
    hand_value = HAND_TYPES[tuple(sorted(hand_counter.values(), reverse=True))]
    return hand_value


def calc_high_card_score(hand, jokers=False):
    card_values = CARD_VALUES.copy()
    if jokers:
        card_values["J"] = 1

    high_card_score = 0
    multiplier = 10000000000
    for card in hand:
        high_card_score += card_values[card] * multiplier
        multiplier /= 100
    return high_card_score


def hand_value(hand: str) -> int:
    hand_counter = Counter(hand)
    hand_value = calc_hand_counter_value(hand_counter)

    high_card_score = calc_high_card_score(hand)

    return hand_value, high_card_score


def hand_value_with_jokers(hand: str) -> int:
    hand_counter = Counter(hand)
    hand_value = calc_hand_counter_value(hand_counter)
    high_card_score = calc_high_card_score(hand, True)
    new_hand = hand

    possible_hands = []
    if "J" in hand_counter.keys():
        print(f"Joker found in hand: {hand}")

        # Loop over all possible hands for the joker and add to possible hands
        for card in CARD_VALUES.keys():
            possible_hand = hand.replace("J", card)
            possible_hand_counter = Counter(possible_hand)
            possible_hand_value = calc_hand_counter_value(
                possible_hand_counter)

            possible_hands.append((possible_hand, possible_hand_value))

        # Sort the possible hands by hand value
        possible_hands.sort(key=lambda x: (x[1]), reverse=True)
        new_hand = possible_hands[0][0]
        hand_value = possible_hands[0][1]

    return new_hand, hand_value, high_card_score


def main():
    inputs = open(INPUT_PATH).read().splitlines()

    hands = []
    hands_with_jokers = []
    total_winnings = 0
    total_winnings_with_jokers = 0

    for inputs in inputs:
        hand, bid = inputs.split()

        value, high_card_score = hand_value(hand)
        hands.append((hand, value, high_card_score, int(bid)))
        hands.sort(key=lambda x: (x[1], x[2]))

        new_hand, value_with_jokers, high_card_score_with_jokers = hand_value_with_jokers(
            hand)
        hands_with_jokers.append(
            (hand, new_hand, value_with_jokers, high_card_score_with_jokers, int(bid)))
        hands_with_jokers.sort(key=lambda x: (x[2], x[3]))

    print("Scored Hand (No Jokers)")
    for i, hand in enumerate(hands, 1):
        # print(i, hand, hand[3] * i)
        total_winnings += hand[3] * i

    print("Scored Hand (Jokers)")
    for i, hand in enumerate(hands_with_jokers, 1):
        print(i, hand, hand[4] * i)
        total_winnings_with_jokers += hand[4] * i

    print(f"Total winnings: {total_winnings}")
    print(f"Total winnings with jokers: {total_winnings_with_jokers}")


if __name__ == "__main__":
    main()
