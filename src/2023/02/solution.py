input_path = "src/2023/02/input.txt"


class Game:
    def __init__(self):
        self.id = None
        self.rounds = []

    def parse(self, line):
        splitLine = line.strip().split(":")  # 'Game 90', ' 4 green,...

        self.id = int(splitLine[0].split(" ")[1])  # 90

        rounds = splitLine[1].strip().split(";")  # [' 4 green,...', ...

        self.rounds = list(map(lambda round: round.strip().split(","), rounds))

    def validate(self, max_red, max_green, max_blue):
        for round in self.rounds:
            for color_count in round:
                count, color = color_count.strip().split(" ")
                count = int(count)

                if ((color == "red" and count > max_red) or
                        (color == "green" and count > max_green) or
                        (color == "blue" and count > max_blue)):
                    return False
        return True

    def minimum_cubes_required(self):
        result = {"red": 0, "green": 0, "blue": 0}

        for round in self.rounds:
            for color_count in round:
                count, color = color_count.strip().split(" ")
                count = int(count)

                result[color] = max(result[color], count)

        return result

    def __repr__(self):
        return f"Game(num={self.id}, rounds={self.rounds})"


games: list[Game] = []
total_valid = 0
total_invalid = 0
sum_valid_ids = 0
total_cubes = 0


with open(input_path, "r") as input_file:
    for line in input_file:
        game = Game()
        game.parse(line)
        if game.validate(12, 13, 14):
            total_valid += 1
            sum_valid_ids += game.id
        else:
            total_invalid += 1

        # Part 2
        min_cubes_required = game.minimum_cubes_required()
        total_cubes += min_cubes_required["red"] * \
            min_cubes_required["green"] * min_cubes_required["blue"]


print(f"Total valid: {total_valid}")
print(f"Total invalid: {total_invalid}")
print(f"Sum of valid ids: {sum_valid_ids}")
print(f"Total cubes: {total_cubes}")
