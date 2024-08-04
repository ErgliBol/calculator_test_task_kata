The calculator should be able to perform the following operations with two numbers: addition, subtraction, multiplication, and division, i.e., a + b, a - b, a * b, a / b. The data should be provided in a single line (see example below). Solutions where each number and the arithmetic operation are provided on a new line are considered incorrect.

The calculator should work with both Arabic (1, 2, 3, 4, 5...) and Roman numerals (I, II, III, IV, V...).

The calculator should accept input numbers ranging from 1 to 10 inclusive. The output numbers are not limited in magnitude and can be any value.

The calculator can only work with integers.

The calculator can only work with either Arabic or Roman numerals at a time. If the user inputs a mix of numerals, such as "3 + II", the calculator should trigger a panic and terminate.

When Roman numerals are input, the result should be output in Roman numerals, and similarly, when Arabic numerals are input, the result should be expected in Arabic numerals.

If the user inputs invalid numbers, the application should trigger a panic and terminate.

If the user inputs a string that does not match one of the aforementioned arithmetic operations, the application should trigger a panic and terminate.

The result of the division operation should be an integer, with any remainder discarded.

The result of the calculator's operations with Arabic numbers can be negative or zero. However, the result of the calculator's operations with Roman numerals can only be positive numbers. If the result is less than one, the program should trigger a panic.
