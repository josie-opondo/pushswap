# Push-Swap Project

## Overview

The **Push-Swap** project is a challenge to implement a non-comparative sorting algorithm using two stacks and a set of specific instructions. The goal is to sort a given stack (`a`) in ascending order using the fewest possible operations. The project consists of two programs:

1. **push-swap**: Calculates and displays the smallest sequence of instructions to sort the stack.
2. **checker**: Verifies if a given sequence of instructions correctly sorts the stack.

---

## Instructions

### Installation

1. Clone the repository:
    ```bash
    git clone https://learn.zone01kisumu.ke/git/nymaina/push-swap.git
    ```

2. Navigate to the project directory:
    ```bash
    cd push-swap
    ```

3. Building the Binaries:
    ```bash
    make build
    ```

4. To remove the generated binaries, use:
    ```bash
    make clean
    ```

### Available Operations
The following operations can be performed on stacks `a` and `b`:

- **`pa`**: Push the top element of stack `b` to stack `a`.
- **`pb`**: Push the top element of stack `a` to stack `b`.
- **`sa`**: Swap the first two elements of stack `a`.
- **`sb`**: Swap the first two elements of stack `b`.
- **`ss`**: Perform `sa` and `sb` simultaneously.
- **`ra`**: Rotate stack `a` (shift all elements up; the first becomes the last).
- **`rb`**: Rotate stack `b`.
- **`rr`**: Perform `ra` and `rb` simultaneously.
- **`rra`**: Reverse rotate stack `a` (shift all elements down; the last becomes the first).
- **`rrb`**: Reverse rotate stack `b`.
- **`rrr`**: Perform `rra` and `rrb` simultaneously.

---

## Programs

### push-swap

#### Description
The **push-swap** program receives a list of integers (stack `a`) as input and outputs the smallest list of instructions to sort the stack. The output is optimized to use the fewest possible operations.

#### Usage
```bash
$ ./push-swap "<stack_elements>"
```

#### Example
```bash
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
```

#### Error Handling
- Non-integer inputs or duplicate values result in:
  ```
  Error
  ```
- No arguments result in no output.

---

### checker

#### Description
The **checker** program takes a list of integers (stack `a`) and a sequence of instructions as input. It simulates the execution of the instructions and verifies whether stack `a` is sorted and stack `b` is empty.

#### Usage
```bash
$ ./checker "<stack_elements>"
<instructions>
```

#### Example
```bash
$ ./checker "3 2 1 0"
sa
rra
pb
KO

$ echo -e "rra\npb\nsa\n" | ./checker "3 2 1 0"
Error

$ echo -e "rra\npb\nsa\npa" | ./checker "3 2 1 0"
OK
```

#### Error Handling
- Non-integer inputs, duplicate values, or invalid instructions result in:
  ```
  Error
  ```
- No arguments result in no output.

---

## Implementation

### Requirements
- **Language**: Go
- **Packages**: Only standard Go packages are allowed.
- **Executables**:
  - `push-swap`
  - `checker`

### Goals
- Sort stack `a` using the minimum number of operations.
- Verify the correctness of sorting instructions.

---

## Testing

### Examples
```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l
6

$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

### Error Cases
- Non-integer input:
  ```bash
  $ ./checker "3 2 one 0"
  Error
  ```
- Duplicate values:
  ```bash
  $ ./push-swap "3 3 2"
  Error
  ```
- Invalid instructions:
  ```bash
  $ echo "invalid_instruction" | ./checker "3 2 1"
  Error
  ```

---

## Learning Objectives
- Understand and implement basic algorithms.
- Apply sorting algorithms.
- Work with stacks and stack operations.
- Enhance problem-solving skills in Go.
- Write clean, modular, and testable code.

---

### Tests
Unit tests are recommended to verify:
- Stack operations.
- Instruction parsing and execution.
- Sorting algorithm correctness.

---

