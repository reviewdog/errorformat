"""Small test script taken from https://wiki.python.org/moin/SimplePrograms"""

import random
import sys # F401 'os' imported but unused
import os # F401 'os' imported but unused

### E265 block comment should start with '# '
print("Hello from reviewdog!")
print("Let's play a small number guessing game to test the flake8 github action.")
print("This game is taken from https://wiki.python.org/moin/SimplePrograms.")

guesses_made = 0

name = input("Hello! What is your name?\n")

number = random.randint(1, 20)
print("Well, {0}, I am thinking of a number between 1 and 20.".format(name)) # E501 line too long (80 > 79 characters)

while guesses_made < 6:

    guess = int(input("Take a guess: "))

    guesses_made += 1

    if guess < number:
        print("Your guess is too low.")

    if guess > number:
        print("Your guess is too high.")

    if guess == number:
        break

if guess == number:
    print(
        "Good job, {0}! You guessed my number in {1} guesses!".format(
            name, guesses_made
        )
    )
else:
    print("Nope. The number I was thinking of was {0}".format(number))

import itertools # E402 module level import not at top of file