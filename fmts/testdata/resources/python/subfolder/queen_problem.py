"""Small test script taken from https://wiki.python.org/moin/SimplePrograms"""

import pwd # F401 'os' imported but unused
import grp # F401 'os' imported but unused

BOARD_SIZE = 8

### E265 block comment should start with '# '
print("Hello from reviewdog!")
print("Let's play a small queen problem game to test the flake8 github action.")
print("This game is taken from https://wiki.python.org/moin/SimplePrograms.")

class BailOut(Exception):
    pass

def validate(queens):
    left = right = col = queens[-1] # E501 line too long (80 > 79 characters). Long description text
    for r in reversed(queens[:-1]):
        left, right = left-1, right+1
        if r in (left, col, right):
            raise BailOut

def add_queen(queens):
    for i in range(BOARD_SIZE):
        test_queens = queens + [i]
        try:
            validate(test_queens)
            if len(test_queens) == BOARD_SIZE:
                return test_queens
            else:
                return add_queen(test_queens)
        except BailOut:
            pass
    raise BailOut

queens = add_queen([])
print (queens)
print ("\n".join(". "*q + "Q " + ". "*(BOARD_SIZE-q-1) for q in queens))

import dis # E402 module level import not at top of file