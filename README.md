# Student Management System (Go)

A command-line Student Management System written in Go. This is my first proper Go mini-project, built to practise the language fundamentals: structs, slices, maps, functions, multiple return values and error handling.

## What it does

- Stores students (ID, name, age, address, subject scores)
- Lists all registered students
- Finds a student by ID (returns an error if not found)
- Calculates a student's average score
- Decides whether a student has passed (passing average: 50)
- Displays only the students who passed
- Adds a new student after validating the data

## Project structure

```
student-management/
├── go.mod
└── main.go
```

Everything lives in one `main.go` for now. I'll split it into packages once I've learned them properly.

## How to run

Requires Go installed (https://go.dev/dl/).

```bash
git clone <your-repo-url>
cd student-management
go run main.go
```

## Example output

```
<paste the real output of your program here>
```

## Go concepts practised

| Concept | Where it's used |
|---|---|
| Constants | `passingMark`, the passing average rule |
| Structs and nested structs | `Student` contains an `Address` |
| Slices | The collection of students, which grows with `append` |
| Maps | Subject scores (`map[string]float64`, subject to score) |
| Functions and multiple return values | `findStudent`, `calculateAverage`, `addStudent` |
| `error` and `nil` | Not-found, empty-scores and validation failures |
| `for` and `range` | Listing students, summing scores, searching |
| `if` / `continue` | Skipping failed or invalid students when displaying passes |
| Comparison and logical operators | Validation and the pass check |

## Design decisions

**Scores are a map, not a slice.** A score belongs to a subject, so `map[string]float64` models the data more honestly than a plain list of numbers. Map iteration order is random in Go, so subjects may print in a different order each run.

**`calculateAverage` returns `(float64, error)`.** An empty set of scores is different from scoring zero, and dividing by zero is meaningless. Returning an error lets the caller tell the two cases apart. It also works on one student's scores, so it can be reused anywhere.

**`hasPassed` returns a `bool`.** A yes/no question gets a yes/no answer. The pass rule lives in one place (the `passingMark` constant), so changing it means editing one line.

**`findStudent` returns `(Student, error)`.** Go has no exceptions, so failure is reported through a returned `error`. When nothing matches it returns an empty `Student{}` with a non-nil error, and the caller checks `err` first.

**`addStudent` returns the updated slice.** `append` may allocate a new underlying array, so the new slice must be returned and assigned back by the caller. If validation fails, the original slice is returned unchanged, so a bad record never alters the list.

**Validation is its own function.** `validateStudent` checks the name, the age (must not be negative) and the ID (not empty and not already used). `addStudent` calls it and only appends if it returns `nil`.

## What I learned

- Go's `if err != nil` pattern, and why errors should never be silently ignored
- The difference between `return`, `break` and `continue` inside loops
- Why a function that only prints should not return a value
- Integer division truncates in Go, so values must be converted to `float64` before dividing
- Unused variables and imports are compile errors, not warnings
- `append` is a built-in function, not a method, and its result must be used

## Possible next steps

- Add an interactive menu
- Split the code into packages
- Add unit tests for the average, validation and search functions
- Persist students to a file or database

## Author

Losika Nicholas Losio
