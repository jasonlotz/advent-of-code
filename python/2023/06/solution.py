INPUT_PATH = "input-files/2023/06/input.txt"

races = []


def calc_winning_hold_times(time, record_distance) -> int:
    winning_button_hold_times = 0
    won_last_trial = False

    # print(f"Calculating winning button hold times for time={
    #      time}, distance={record_distance}")
    for hold_time in range(1, time):

        trial_distance = hold_time * (time - hold_time)
        # print(f"Hold time: {
        #      hold_time} = Trial distance: {trial_distance}")

        if trial_distance > record_distance:
            winning_button_hold_times += 1
        else:
            if won_last_trial:
                break

    return winning_button_hold_times


def part_1():
    total_margin_of_error = 1
    times, distances = open(INPUT_PATH).read().splitlines()

    times = list(map(int, times.split()[1:]))
    distances = list(map(int, distances.split()[1:]))

    races = zip(times, distances)

    for race in races:
        time, distance = race
        # print(f"Processing race with time={time}, distance={distance}")

        winning_hold_times = calc_winning_hold_times(
            time, distance)

        total_margin_of_error *= winning_hold_times

    print(f"Total margin of error (part 1): {total_margin_of_error}")


def part_2():
    time, distance = open(INPUT_PATH).read().splitlines()

    time = int("".join(time.split()[1:]))
    distance = int("".join(distance.split()[1:]))

    winning_hold_times = calc_winning_hold_times(
        time, distance)

    print(f"Winning hold times (part 2): {winning_hold_times}")


def main():
    part_1()
    part_2()


if __name__ == "__main__":
    main()
