# Brute Force Password Calculator

## Description

This is a simple brute force password calculator. It will calculate the time it would take to brute force a password of a given length, using a given character set, and the hashing speed of the computer it is running on.

## Usage

```bash
git clone git@github.com:huzaifamk/Go-Brute-Force-Calc.git
cd Go-Brute-Force-Calc
go mod tidy
go run main.go
```

## Example

```bash
$ go run main.go
Brute Force Calculator #Hashcat
Enter the speed in hashes per second:
- 45500
Enter the length of the password:
- 9
Choose from the following options:
1. Numbers only
2. Letters only(lower)
3. Letters only(upper)
4. Letters only(both upper and lower)
5. Letters and numbers(lower)
6. Letters and numbers(upper)
7. Letters and numbers(both upper and lower)
- 4
You chose letters only(both upper and lower)
There are 2779905883635712 possible combinations
It will take 1018280543 minutes to brute force this password
It will take 16971342 hours to brute force this password
That is 707139 days
```

## License

BSD 2-Clause License
Copyright (c) 2023, Huzaifa M.
