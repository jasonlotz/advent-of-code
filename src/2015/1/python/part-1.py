# Run: python src/2015/1/python/part-1.py

import re

input_string = 'src/2015/1/input.txt'

with open(input_string, 'r') as input_file:
    for line in input_file:
        left = len(line.split('('))
        right = len(line.split(')'))

        print("Floor #: " + str(left - right))
