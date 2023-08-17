package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Brute Force Calculator #Hashcat")

	fmt.Println("Enter the speed in hashes per second:")
	var speed int
	fmt.Scanln(&speed)

	fmt.Println("Enter the length of the password:")
	var length int
	fmt.Scanln(&length)

	fmt.Println("Choose from the following options: \n1. Numbers only \n2. Letters only(lower) \n3. Letters only(upper) \n4. Letters only(both upper and lower) \n5. Letters and numbers(lower) \n6. Letters and numbers(upper) \n7. Letters and numbers(both upper and lower)")
	var choice int
	fmt.Scanln(&choice)

	var combinations int

	if choice == 1 {

		fmt.Println("You chose numbers only")
		combinations = 10
		for i := 1; i < length; i++ {
			combinations = combinations * 10
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 2 {

		fmt.Println("You chose letters only(lower)")
		combinations = 26
		for i := 1; i < length; i++ {
			combinations = combinations * 26
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 3 {

		fmt.Println("You chose letters only(upper)")
		combinations = 26
		for i := 1; i < length; i++ {
			combinations = combinations * 26
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 4 {

		fmt.Println("You chose letters only(both upper and lower)")
		combinations = 52
		for i := 1; i < length; i++ {
			combinations = combinations * 52
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 5 {

		fmt.Println("You chose letters and numbers(lower)")
		combinations = 36
		for i := 1; i < length; i++ {
			combinations = combinations * 36
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 6 {

		fmt.Println("You chose letters and numbers(upper)")
		combinations = 36
		for i := 1; i < length; i++ {
			combinations = combinations * 36
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}

	if choice == 7 {

		fmt.Println("You chose letters and numbers(both upper and lower)")
		combinations = 62
		for i := 1; i < length; i++ {
			combinations = combinations * 62
		}

		fmt.Println("There are", combinations, "possible combinations")

		time := combinations / speed
		minutes := time / 60
		if minutes != 0 {
			fmt.Println("It will take", minutes, "minutes to brute force this password")
		}

		hours := time / 3600
		if hours != 0 {
			fmt.Println("It will take", hours, "hours to brute force this password")
		}

		days := hours / 24
		if days != 0 {
			fmt.Println("That is", days, "days")
		}
	}
	for {
		fmt.Println("Do you want to try again? (y/n) OR Do you want to change the hashing speed? (s)")
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			main()
		}

		if answer == "n" {
			fmt.Println("Thank you for using this program")
			os.Exit(0)
		}

		if answer == "s" {
			fmt.Println("Enter the speed in hashes per second:")
			fmt.Scanln(&speed)
			time := combinations / speed
			minutes := time / 60
			if minutes != 0 {
				fmt.Println("It will take", minutes, "minutes to brute force this password")
			}

			hours := time / 3600
			if hours != 0 {
				fmt.Println("It will take", hours, "hours to brute force this password")
			}

			days := hours / 24
			if days != 0 {
				fmt.Println("That is", days, "days")
			}
		}
	}
}
