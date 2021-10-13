package main

import (
	"database/sql"
	"errors"
	"html/template"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Regular expression to check if email is of valid format
var emailRegex = regexp.MustCompile(
	"^[\\w!#$%&'*+/=?`{|}~^-]+(?:\\.[\\w!#$%&'*+/=?`{|}~^-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,6}$") // regular expression

// Checks if email is of valid format, return true if valid else false
func isEmailFormatValid(e string) bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

/*
// Tim to Jason: Let you decide whether you want to implement yours or mine.
// validateUserName validates any input named "username"
func isValidUserName(u string) (string, error) {
	// escape strings first
	u = template.HTMLEscapeString(u) // done at server side

	// trim trailing spaces left & right of username
	u = strings.TrimSpace(u)
	// make username lowercase
	u = strings.ToLower(u)

	// if field is empty
	if len(u) == 0 {
		return "", errors.New("empty username")
	}
	// check if field has English chars, dash, or numbers, and has to be 3-32 chars long
	if m, _ := regexp.MatchString("^[a-z0-9-]{3,32}$", u); !m {
		return "", errors.New("invalid username")
	}
	return u, nil
}
*/

// validatePassword validates any input named "password"
// used for adminRESTAPI.go (in conj with promotionaladminconsole)
func validatePassword(pw string) (string, error) {
	// escape strings first
	pw = template.HTMLEscapeString(pw) // done at server side

	// if field is empty
	if len(pw) == 0 {
		return "", errors.New("empty password")
	}
	return pw, nil
}

// NB: For signups, all passwords, after validation here, go through the following fn below!
// Password validates plain password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.

// Not used for daily logins. Only for signups.
func isValidPassword(pw string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pw {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

// Working
// This function checks whether username is available in buyer table. Returns true if available, else return false
func isBuyerUserNameExists(db *sql.DB, userName string) bool {
	sqlStmt := `SELECT username FROM crazyformasks_db.BuyerTable where Username=?`
	err := db.QueryRow(sqlStmt, userName).Scan(&userName)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}
	return true
}

// Working
// This function checks whether email is available in buyer table. Returns true if available, else return false
func isBuyerEmailExists(db *sql.DB, email string) bool {
	sqlStmt := `SELECT email FROM crazyformasks_db.BuyerTable where Email=?`
	err := db.QueryRow(sqlStmt, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}
	return true
}

// got problem
// This function checks whether email is available in seller table. Returns true if available, else return false
func isSellerEmailExists(db *sql.DB, email string) bool {
	sqlStmt := `SELECT email FROM crazyformasks_db.SellerTable where email = ?`
	err := db.QueryRow(sqlStmt, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}
	return true
}

// Working
// This function checks whether username is available in seller table. Returns true if available, else return false
func isSellerUserNameExists(db *sql.DB, userName string) bool {
	sqlStmt := `SELECT username FROM crazyformasks_db.SellerTable where Username=?`
	err := db.QueryRow(sqlStmt, userName).Scan(&userName)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}
	return true
}

// check if user inputs valid string onto form fields
// filter out strange characters (e.g. <script>)
// if cannotBeANumber, returns false for "20340" or "2340", but true for "ABC123"
func isValidString(e string, minLength, maxLength int, cannotBeANumber bool) bool {

	// trim trailing spaces left & right of username
	e = strings.TrimSpace(e)

	e = template.HTMLEscapeString(e)

	inputStringRegex := regexp.MustCompile(
		"^[a-zA-Z0-9-.,!?() ]{" +
			strconv.Itoa(minLength) + "," +
			strconv.Itoa(maxLength) + "}$")

	if !inputStringRegex.MatchString(e) {
		return false
	}

	if cannotBeANumber {
		// string should not be a number
		_, err := strconv.Atoi(e)
		return err != nil
	}

	return true
}

// check if user inputs valid floats onto form fields such as price
func isValidFloat(e string) bool {
	e = template.HTMLEscapeString(e)
	_, err := strconv.ParseFloat(e, 64)
	return err == nil
}

// check if user inputs valid int onto form fields such as quantity
func isValidInt(e string) bool {
	e = template.HTMLEscapeString(e)
	_, err := strconv.Atoi(e)
	return err == nil
}

// check if user inputs valid product filters
// differs from isValidString in that it checks also for commas
func isValidStringOfFilters(e string, minLength, maxLength int) bool {

	// trim trailing spaces left & right of username
	e = strings.TrimSpace(e)

	e = template.HTMLEscapeString(e)

	inputStringRegex := regexp.MustCompile(
		"^[a-zA-Z0-9-, ]{" +
			strconv.Itoa(minLength) + "," +
			strconv.Itoa(maxLength) + "}$")

	return inputStringRegex.MatchString(e)
}

// Regex validation for credit card
func isValidCreditCard(cardNum string) bool {
	amexCard := regexp.MustCompile("^3[47][0-9]{13}$")
	//masterCard := regexp.MustCompile("^5[1-5][0-9]{14}$")
	masterCard := regexp.MustCompile("^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$")
	visaCard := regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$")
	visaMasterCard := regexp.MustCompile("^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$")

	if amexCard.MatchString(cardNum) || masterCard.MatchString(cardNum) || visaCard.MatchString(cardNum) || visaMasterCard.MatchString(cardNum) {
		return true
	}
	return false
}

/*
Amex Card: ^3[47][0-9]{13}$
BCGlobal: ^(6541|6556)[0-9]{12}$
Carte Blanche Card: ^389[0-9]{11}$
Diners Club Card: ^3(?:0[0-5]|[68][0-9])[0-9]{11}$
Discover Card: ^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$
Insta Payment Card: ^63[7-9][0-9]{13}$
JCB Card: ^(?:2131|1800|35\d{3})\d{11}$
KoreanLocalCard: ^9[0-9]{15}$
Laser Card: ^(6304|6706|6709|6771)[0-9]{12,15}$
Maestro Card: ^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$
Mastercard: ^5[1-5][0-9]{14}$
Solo Card: ^(6334|6767)[0-9]{12}|(6334|6767)[0-9]{14}|(6334|6767)[0-9]{15}$
Switch Card: ^(4903|4905|4911|4936|6333|6759)[0-9]{12}|(4903|4905|4911|4936|6333|6759)[0-9]{14}|(4903|4905|4911|4936|6333|6759)[0-9]{15}|564182[0-9]{10}|564182[0-9]{12}|564182[0-9]{13}|633110[0-9]{10}|633110[0-9]{12}|633110[0-9]{13}$
Union Pay Card: ^(62[0-9]{14,17})$
Visa Card: ^4[0-9]{12}(?:[0-9]{3})?$
Visa Master Card: ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$
*/
