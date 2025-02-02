package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unicode/utf8"
)

const DefaultCost = 12

var (
	re      = regexp.MustCompile(`^[a-zA-Z0-9$]+$`)
	version string
)

func main() {
	p := flag.String("p", "", "Raw password to be hashed with bcrypt (insecure)")
	c := flag.Int("c", 0, fmt.Sprintf("The bcrypt cost factor, where a higher value means more computation and greater security (min: %d max: %d default: %d)", bcrypt.MinCost, bcrypt.MaxCost, DefaultCost))
	i := flag.Bool("i", false, "Prompt the user to enter the password securely from the console (recommended)")
	n := flag.Bool("pretty", false, "Ensure the resulting bcrypt hash does not contain special characters like dot or slash (recommended)")
	t := flag.Bool("trim", false, "Do not add a newline at the end of the hash (useful for redirecting output)")
	v := flag.String("verify", "", "The input hash value for verifying against the plaintext password")

	flag.Parse()

	if flag.NFlag() == 0 {
		printUsageAndExit()
	}

	if *c == 0 {
		*c = DefaultCost
	}

	var pass string

	if *i {
		pass = readPassword()
	} else {
		pass = *p
	}

	if pass == "" {
		printUsageAndExit()
	}

	if *v != "" {
		err := bcrypt.CompareHashAndPassword([]byte(*v), []byte(pass))
		if err == nil {
			fmt.Println("Password verification succeeded")
			os.Exit(0)
		} else {
			fmt.Println("Password verification failed")
			os.Exit(1)
		}
	}

	if *c > bcrypt.MaxCost || *c < bcrypt.MinCost {
		printUsageAndExit()
	}

	var bytes []byte
	bytesPass := []byte(pass)

	for k := 0; k < 25; k++ {
		bytes, _ = bcrypt.GenerateFromPassword(bytesPass, *c)
		if !*n || re.MatchString(string(bytes)) {
			break
		}
	}

	if !*t {
		bytes = append(bytes, []byte("\n")...)
	}

	os.Stdout.Write(bytes)
	os.Exit(0)
}

func readPassword() string {
	t := "Enter password: "
	fmt.Print(t)
	fd := int(syscall.Stdin)
	b, _ := term.ReadPassword(fd)
	fmt.Print("\r" + strings.Repeat(" ", utf8.RuneCountInString(t)) + "\r")
	return string(b)
}

func printUsageAndExit() {
	fmt.Println("bpasswd v" + version)
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	os.Exit(1)
}
