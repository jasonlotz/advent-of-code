INPUT_PATH = "input-files/2023/04/input.txt"


def find_matching_numbers(line: str) -> set:
    _, numbers = line.split(":")
    winning_numbers, my_numbers = numbers.split("|")

    winning_numbers = winning_numbers.strip().split()
    my_numbers = my_numbers.strip().split()

    matching_numbers = set(winning_numbers) & set(my_numbers)

    return matching_numbers


def score_matching_numbers(matching_numbers: set) -> int:
    return 2 ** (len(matching_numbers) - 1)


def win_cards(all_card_matches: list) -> list:
    won_cards = [1] * len(all_card_matches)

    for i, card_matches in enumerate(all_card_matches):
        if card_matches:
            cards_to_win = len(card_matches)
            cards_per_win = won_cards[i]

            for j in range(i + 1, i + cards_to_win + 1):
                won_cards[j] += cards_per_win

    return won_cards


def main():
    total = 0
    all_card_matches = []

    with open(INPUT_PATH, "r") as input_file:
        for line in input_file:
            matching_numbers = find_matching_numbers(line)
            if len(matching_numbers) > 0:
                line_score = score_matching_numbers(matching_numbers)
                total += line_score
            all_card_matches.append(matching_numbers)

        won_cards = win_cards(all_card_matches)

    print(f"Total score: {total}")

    print(f"Won cards:{sum(won_cards)}")


if __name__ == "__main__":
    main()
