## Advent Of Code 2023

> This year I thought I'd give it a ***go*** 

I started learning go earlier this year and really enjoyed it (I'm still working on this cli tool [unq](https://github.com/justjcurtis/unq)) and thought I'd complete this years [AOC](https://adventofcode.com/) with the goal of writing the fastest possible solutions to each problem using `go` to help me get more familiar the language. Anyway here's how it went.

### TOC

* [Installation](#installation)
* [Usage](#usage)
    * [Puzzle Input](#puzzle-input)
    * [Flags](#flags)
* [Testing](#testing)
* [Building](#building)
* [Results](#results)

### Installation
- Install [`go`](https://go.dev/doc/install)

- Clone the repo && cd into the directory

     ```
     git clone https://github.com/justjcurtis/AdventOfCode2023
     cd AdventOfCode2023
     ```

- Run the project

    ```
    go run .
    ```

### Usage

#### Puzzle Input
The files in `./puzzleInput/` are read in and passed to each solution via `./main.go`. If you want to replace those files with your own input data to ensure correct solutions / compare runtimes just replace the file for the corresponding day in the `./puzzleInput/` dir & follow the naming convention in there (`day_{n}.txt`).

#### Flags

| Flag | Description | Example |
| ---- | ----------- | ------- |
| `-n` | How many times to run each solution. | `go run . -n 1000` Run each solution 1000 times and report the average runtime for each solution + the total average runtime. |
| `-min` | Report the minimum time instead of the average. | `go run . -min` Set `-n` to 5000 by default & report the minimum time for each day and the total minimum runtime. |
| `-d` | Only run a single day | `go run . -d 8` Only run day 8. This will only run the solution once unless other flags are set. |

### Testing
Each day is unit tested using the example input from the puzzle fpr that day on adventofcode.com

- Run the unit tests with go:

    ```
    go test ./... -v
    ```

### Building

To build locally:
- Follow the [installation instructions](#installation)
- Then run

    ```
    go build .
    ```
- A new file will be create in the root dir (`AdventOfCode2023`)
- Run the build with

    ```
    ./AdventOfCode2023 
    ```

### Results
Results show are the min runtime for each soltuion taken over 10,000 runs as reported on my desktop (Ryzen 5 7600X). The fastest solution is shown in bold for showing off purposes. Reading the input data from disk is not included as part of the solution here so the runtime you see is the runtime of any parsing & logic requried to solve the puzzle.
>*Your results may vary*

| # | Runtime (both parts) |
| - | -------------------- |
| Day 1 | 46.55µs |
| Day 2 | 17.18µs |
| Day 3 | 97.15µs |
| Day 4 | 119.80µs |
| Day 5 | 27.61µs |
| ***Day 6*** | ***539ns*** |
| Day 7 | 236.16µs |
| Day 8 | 153.82µs |
| ------- | ----------------------------- |
| **Total** | **698.80µs** |


##### [Take Me To The TOP!](#top)
