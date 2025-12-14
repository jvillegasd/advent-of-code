#!/usr/bin/env python3
"""
Advent of Code 2025 - Day 10 (Part 2): Joltage Configuration

This solution uses the z3-solver constraint satisfaction and optimization library
to solve a linear system with integer constraints.

Problem Summary:
---------------
Each machine has:
- A set of buttons, where each button affects specific joltage counters
- Target joltage values that need to be reached for each counter
- Each button press increases all affected counters by 1

Goal: Find the minimum number of total button presses needed to configure all
      machines to their target joltage levels.

Solution Approach:
-----------------
This is a constrained optimization problem that can be formulated as:

Variables:
    x_i = number of times button i is pressed (non-negative integer)

Constraints:
    For each counter j: Σ(x_i where button i affects counter j) = target_j
    All x_i >= 0

Objective:
    Minimize Σ(x_i) (total button presses)
"""

import os

from z3 import *


def parse_input(filename):
    """Parse input file and return list of machines."""
    machines = []

    with open(filename) as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            parts = line.split()

            # Parse buttons (everything except first and last element)
            buttons = []
            for part in parts[1:-1]:
                # Remove parentheses and split by comma
                indices = part.strip("()").split(",")
                buttons.append([int(idx) for idx in indices])

            # Parse target joltage values (last element)
            joltage_str = parts[-1].strip("{}")
            joltage = [int(x) for x in joltage_str.split(",")]

            machines.append({"buttons": buttons, "joltage": joltage})

    return machines


def solve_machine(machine: dict) -> int:
    """
    Solve a single machine using z3.
    Returns minimum button presses needed, or -1 if impossible.
    """
    buttons = machine["buttons"]
    target = machine["joltage"]
    num_buttons = len(buttons)
    num_counters = len(target)

    # Create z3 optimizer
    opt = Optimize()

    # Create variables for button presses (must be non-negative integers)
    button_vars = [Int(f"button_{i}") for i in range(num_buttons)]

    # Constraint: all button presses must be >= 0
    for var in button_vars:
        opt.add(var >= 0)

    # Constraint: for each counter, sum of effects must equal target
    for counter_idx in range(num_counters):
        counter_sum = 0
        for button_idx, button in enumerate(buttons):
            if counter_idx in button:
                counter_sum += button_vars[button_idx]
        opt.add(counter_sum == target[counter_idx])

    # Objective: minimize total button presses
    total_presses = Sum([var for var in button_vars])
    opt.minimize(total_presses)

    # Solve
    if opt.check() == sat:
        model = opt.model()
        result = sum(model.evaluate(var).as_long() for var in button_vars)
        return result
    else:
        return -1


def solve_part2(machines: list[dict]) -> int:
    """Solve part 2 for all machines."""
    total = 0

    for machine in machines:
        result = solve_machine(machine)
        if result >= 0:
            total += result

    return total


def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    filename = os.path.join(script_dir, "input.txt")

    machines = parse_input(filename)

    part2 = solve_part2(machines)
    print(f"Part 2: {part2}")


if __name__ == "__main__":
    main()
