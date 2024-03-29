package main

import (
	"fmt"
	"strings"
)



const conferenceTickets int = 50
var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

	
func main() {
	
	greetUsers()

	for{
		
	  firstName, lastName, email, userTickets := getUserInput()
	
      isValidName, isValidEmail, isValidTicketNumber :=  validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			

		//calling first names
	    firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
         
		
		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
		} else {
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
			
		}
		 

		

		
	
		
		
		

	}

	

	

}

func greetUsers(){
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available.\n" , conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

	
		return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for name 
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
		
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string ){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets . you will recieve a confirmation email at %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}