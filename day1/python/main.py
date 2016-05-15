#!/usr/bin/env python3

def parse_input(inputString):
    floor = 0
    position = 1
    firstBasementFloor = None

    for c in inputString:
        if c == ')':
            floor -= 1
            if firstBasementFloor is None and floor == -1:
                firstBasementFloor = position
        if c == '(':
            floor += 1
        position += 1

    return floor, firstBasementFloor

txt = open("input.txt")

inputString = txt.read()

outputFloor, outputFirstBasementFloor = parse_input(inputString)

print("The floor is {} and first basement position is {}".format(outputFloor, outputFirstBasementFloor))


