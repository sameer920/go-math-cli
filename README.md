# Go-Math CLI

Go-Math CLI is a command-line tool designed to perform basic mathematical operations such as addition, subtraction, multiplication, division, square, square root, and percentage calculations. This project was created for educational purposes to explore Go programming, CLI development, and modular code organization. However, if you find it useful or want to extend its functionality, feel free to do so!

---

## Supported Commands:

- **Addition**: Add two numbers.
- **Subtraction**: Subtract one number from another.
- **Multiplication**: Multiply two numbers with optional rounding.
- **Division**: Divide one number by another with optional rounding.
- **Square**: Calculate the square of a number.
- **Square Root**: Calculate the square root of a number.
- **Percentage**: Calculate a percentage of a given number.

---

## Installation

Before installing the code make sure you have go installed on your system.

1. Clone the repository:
   ```bash
   git clone https://github.com/sameer920/go-math.git
   cd go-math
   ```
2. Build the project:
```go build```
3. Install the binary:
```go install```

---

## Usage

Run the `go-math` command followed by the desired operation and its arguments. Below are some examples:

### Addition
```bash
go-math add 5 10
# Output: Addition of 5 and 10 = 15
```

### Subtraction
```bash
go-math subtract 20 5
# Output: Subtracting 5 from 20 = 15
```

### Multiplication
```bash
go-math multiply 3 4
# Output: Multiplication of 3 and 4 = 12
```

### Division
```bash
go-math divide 10 2
# Output: Division of 10 by 2 = 5
```

### Square
```bash
go-math square 5
# Output: Square of 5 = 25
```

### Square Root
```bash
go-math sqrt 16
# Output: Square root of 16 = 4
```

### Percentage
```bash
go-math percent 200 10
# Output: 10% of 200 = 20
```

---

## Project Structure

- **`cmd/`**: Contains all the CLI commands (e.g., `add`, `subtract`, `multiply`, etc.).
- **`go-math.go`**: The entry point of the application.
- **`math.go`**: Core mathematical functions used by the CLI commands.

---

## Contributing

Since the project was created for educational value, I would probably not be maintaining it. If you want to use it in anyway, you can fork or clone the repo and use it however you want to.

---

## License

This project is released into the public domain under the [Unlicense](LICENSE). You are free to use, modify, publish, and distribute this software for any purpose, without any restrictions.

For more details, see the [LICENSE](LICENSE) file.

---

## Disclaimer

This project was created for **educational purposes** and may not be production-ready. Use it at your own discretion. While I have taken care to ensure the code is functional, there may still be edge cases or bugs. If you encounter any issues, feel free to report them or submit a fix.

---

## Acknowledgments

Special thanks to the developers of the [Cobra CLI library](https://github.com/spf13/cobra) for making CLI development in Go so seamless.

I would also like to acknowledge this [amazing article from JetBrains](https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/conclusion/) that I followed initially to make this project.
---

Happy coding! 😊