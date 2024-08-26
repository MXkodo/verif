package pkg

import "fmt"

func SendEmailWarning(userID string) {
	fmt.Printf("Sending email warning to user %s\n", userID)
}
