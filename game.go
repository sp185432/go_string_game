package main

import (
	"fmt"
	
	"sort"
	"strconv"
	"strings"
	"time"
	"github.com/inancgumus/screen"
)

func main() {

	//declaration of all the variiables that are used
	var choice bool
	var name1, name2, Userstring, guessword, activeperson, inactiveperson string
	var active, inactive, wins1, wins2, j int
	var comparestring, noofrounds, count1, flag int
	var choiceStart, output string
	var temp1, temp2 int64

	//printing the welcome statement and starting statements
	fmt.Println("\n\t\t\t----------WELCOME TO STRING GAME----------")
	fmt.Println("\n\t\t\tNote: This is a 2 player game, only 2 can play!")
	fmt.Println("\n\t\t\tDo you want to start the game? ")
	fmt.Println("\n\t\t\tEnter 1 for YES and 0 for NO       ")
	fmt.Scan(&choice)

	if choice == true {

		//details of the players
		fmt.Println("\nEnter the name of player 1: ")

		fmt.Scan(&name1)

		fmt.Println("\nEnter the name of player 2: ")
		fmt.Scan(&name2)

		fmt.Println("\nHow many rounds do each one of you want to play?")
		fmt.Scan(&noofrounds)

		//giving each player certain points before itself based on noofrounds they want to play
		points1 := 100 * noofrounds
		points2 := 100 * noofrounds

		//printing few rules and also the points credited to players
		fmt.Println("\n\t\t NOTE: Please enter meaningful words as input strings")
		fmt.Println("\n\t\t NOTE: TIME IS ALSO RECORDED. GUESS THE WORD SOON!")
		fmt.Println("\n\t\tInitially points that both the players have are!!")
		fmt.Println("\n\t\t Player 1\t\t\t\t\t Player 2")
		fmt.Println("\t\t", strings.ToUpper(name1), "\t\t\t\t\t", strings.ToUpper(name2))
		fmt.Println("\t\t ------\t\t\t\t\t\t ------")
		fmt.Println("\t\t", points1, "\t\t\t\t\t\t", points2)

	Second:
		for j = 1; j <= noofrounds; j++ {

			flag = 0
			count1 = 0
			fmt.Println("\n\t\t\tWho wants to start first?")
			fmt.Println("\n\t\tchoose 1 for", strings.ToUpper(name1), "and choose 2 for", strings.ToUpper(name2))

			fmt.Scan(&choiceStart)

			if choiceStart != "1" && choiceStart != "2" {
				fmt.Println("\t\tEnter a valid choice!!")
				flag = 1
				continue

			}

		First:
			if choiceStart == "1" {
				count1++
				//storing a copy of entered names in respective required variables according to the condition
				active = 1
				activeperson = name1
				inactiveperson = name2

				fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), "has started the game")

			} else {
				count1++
				active = 2
				activeperson = name2
				inactiveperson = name1
				fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), "has started the game")

			}

			fmt.Println("\nEnter the input string: ")

			fmt.Scan(&Userstring)

			//clearingScreen()

			screen.Clear()

			screen.Clear()

			screen.Clear()

			//calling the function rearrage which will rearrange the string entered by one player
			output = ReArrange(Userstring)

			fmt.Println("\n\t\tThe string entered by player", active, " is rearranged as: ")

			fmt.Println("\n\t\t\t\t\t", output)

			fmt.Println("\n\t\t=========================================================")

			if active == 1 {
				inactive = 2
				fmt.Println("\n\t\tNow, It's player 2 turn to guess the correct word")
				fmt.Println("\n\t\tPlayer 2 has 3 chances to guess the correct word")
			} else {
				inactive = 1
				fmt.Println("\n\t\tNow, It's player 1 turn to guess the correct word")
				fmt.Println("\n\t\tPlayer 1 has 3 chances to guess the correct word")
			}

			i := 1
			for {
				if i >= 4 {
					break
				}
				fmt.Println("\t\t=========================================================")
				fmt.Println("\n\t\tPlayer ", inactive, ": Attempt---", i)

				//recording the time taken by user to enter a string(starting point)
				t0 := time.Now()

				fmt.Scan(&guessword)

				//records the end point
				t1 := time.Now()

				//difference of two times which gives the time taken by user to enter a input
				t2 := t1.Sub(t0)

				//calling the function comparestring
				comparestring = CompareString(guessword, Userstring)

				if comparestring == 0 {
					i = i + 10

					fmt.Println("\n\t\t", strings.ToUpper(inactiveperson), " has guessed the word entered by ", strings.ToUpper(activeperson), "corrrectly")

					fmt.Println("\n\t\tThe time taken by", strings.ToUpper(inactiveperson), "to guess the word is ", t2)

					if inactiveperson == name1 {

						temp1 = int64(t2)
						points2 = points2 - 100

					} else {

						temp2 = int64(t2)
						points1 = points1 - 100

					}

				} else {
					i++
					fmt.Println("\n\t\tThe word entered doesn't match")
					fmt.Println("\n\t\tPlayer ", inactive, "has only", (4 - i), "attempts left")
				}

			}

			if (4 - i) == 0 {

				fmt.Println("\n\t\t", strings.ToUpper(inactiveperson), " didn't guess the word entered by ", strings.ToUpper(activeperson), "corrrectly")

				if activeperson == name1 {
					//temp1 = 0
					temp2 = 0
					//incorrect2++
					points2 = points2 - 100

				} else {
					temp1 = 0
					//temp2 = 0
					//incorrect1++
					points1 = points1 - 100

				}
			}

			//giving the second player chnace to play in that round
			choiceStart = strconv.Itoa(inactive)

			count1++

			if count1 > 2 {

				fmt.Println("\n\t\t=========================================================")
				fmt.Println("\n\tBy the end of round ", j, "The points of the players are: ")
				fmt.Println("\t\t\tPlayer 1: ", points1)
				fmt.Println("\t\t\tPlayer 2: ", points2)
				fmt.Println("\n\t\t=========================================================")

			} else {

				goto First

			}

			if temp1 != 0 {
				if temp2 == 0 {

					//fmt.Println("Inside first loop")
					//fmt.Println("number of incorrect1: ", incorrect1)
					//fmt.Println("number of incorrect2: ", incorrect2)
					wins1++
					fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), "has won the round ", j)
					goto Result

				}

			} else if temp2 != 0 {
				if temp1 == 0 {

					//fmt.Println("Inside second loop")
					//fmt.Println("number of incorrect1: ", incorrect1)
					//fmt.Println("number of incorrect2: ", incorrect2)

					wins2++
					fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), "has won the round ", j)
					goto Result

				}

			} else if temp1 == 0 {
				if temp2 == 0 {
					fmt.Println("\n\t\tBoth of them didn't guess the word correctly")
				}
			} else {
				fmt.Println("\n\t\tBoth of them didn't guess the word correctly")
			}

			if temp1 < temp2 {
				wins1++
				fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), "has won the round ", j)

			} else if temp1 > temp2 {
				wins2++
				//fmt.Println(temp1)
				//fmt.Println(temp2)
				fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), "has won the round ", j)

			} else {
				fmt.Println("\n\t\t None of them has won this round")
			}
		Result:
			fmt.Println("\n\t\t=========================================================")
			fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), " No. of WINS: ", wins1)
			fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), " No. of WINS: ", wins2)

		}

		if flag == 0 {

			if points1 > points2 {

				fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), "has won the game")

			} else if points2 > points1 {

				fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), "has won the game")

			} else if wins1 > wins2 {

				fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), "has won the game")

			} else if wins1 < wins2 {

				fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), "has won the game")

			} else {

				fmt.Println("\n\t\tIt's a tie!!! Both the players have equal scores")
				fmt.Println("\n\t\tNow both the players will get another turn each")

				j = (noofrounds - (noofrounds - 1))
				noofrounds = (noofrounds - (noofrounds - 1))
				goto Second

			}
		} else {

			fmt.Println("\n\tYou have exhausted your chances to select who wants to play first. You can restart the game when you are ready")
			fmt.Println("\n\t\t=========================================================")
			fmt.Println("\n\t\t FINAL RESULT:")
			fmt.Println("\n\t\tPlayer 1", strings.ToUpper(name1), " No. of WINS: ", wins1)
			fmt.Println("\n\t\tPlayer 2", strings.ToUpper(name2), " No. of WINS: ", wins2)
		}

	} else {

		fmt.Println("\n\t\tYou can start the game when you are ready to play")

	}

}

//CompareString is a function used for the comparison of two strings
func CompareString(input1 string, input2 string) int {

	value := strings.Compare(strings.ToUpper(input1), strings.ToUpper(input2))
	return value
}

//ReArrange is a function to rearrange the input string
func ReArrange(Userstring string) string {

	//the entered input string will be split into each character and will be stored as an array in newArray
	newArray := strings.Split(Userstring, "")

	//fmt.Println(newArray)

	//fmt.Println(len(newArray))

	//to sort the elements in an array in ascending order
	//this is displayed as the rearranged string
	//opponent should find out the actual string from this rearranged string

	sort.Slice(newArray, func(i, j int) bool {

		return newArray[i] < newArray[j]

	})

	/*printing the rearranged string
	for _, value := range newArray {

		fmt.Printf(value)
	}*/

	justString := strings.Join(newArray, "")

	//fmt.Println(justString)

	return justString
}

/*func clearingScreen() {

	screen.Clear()
	fmt.Println("Hello World!")

}*/
