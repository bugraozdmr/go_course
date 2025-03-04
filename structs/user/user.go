package user

import (
	"errors"
	"fmt"
	"time"
	"regexp"
)

// eğer dışardan erişim istemediğin metodlar varsa küçük harf kalabilirler


// mesela direkt user'a dışardan erişmek istiyorsun o zaman User olcak
// ayrıca içerdeki değerler içinde büyük harf olmalı yoksa erişmez yazmaz vs.
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// eğer (u user) bu bir kopyasını alır onun yerine pointer mantığını kullan
func (u *user) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

// embedded struct
type Admin struct {
	Email string
	Password string
	User user
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email: email,
		Password: password,
		User: user{
			firstName: "Grant",
			lastName: "Wick",
			birthdate: "?",
			createdAt: time.Now(),
		},
	}
}

// construct -- copy olusturmıcak artık
func NewUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}

	if !IsValidDateFormat(birthdate) {
		return nil, errors.New("invalid birthdate format, expected MM/DD/YYYY")
	}

	if _, err := time.Parse("01/02/2006", birthdate); err != nil {
		return nil, errors.New("invalid birthdate, please enter a real date")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func IsValidDateFormat(date string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])/\d{4}$`)
	return re.MatchString(date)
}

func OutputUserDetails(u *user) {
	fmt.Println("\nUser Details:")
	// normalde böyle değil (*u).firstName şeklinde erişilirdi struct olmasaydı
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Account Created At:", u.createdAt.Format("2006-01-02 15:04:05"))
}

func (u *user) OutputUserDetailsAttached() {
	fmt.Println("\nUser Details:")
	// normalde böyle değil (*u).firstName şeklinde erişilirdi struct olmasaydı
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
	fmt.Println("Account Created At:", u.createdAt.Format("2006-01-02 15:04:05"))
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
