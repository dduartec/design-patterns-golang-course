package main

import (
	"fmt"
	"strings"
)

// hide domain object
type email struct {
	from, to, subject, body string
}

// exportable builder struct
type EmailBuilder struct {
	email email
}

// Build methods
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("no @")
	}

	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("no @")
	}

	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {

	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {

	b.email.body = body
	return b
}

// hid real send implementation
func sendEmailImp(email *email) {
	fmt.Println(email)
}

type build func(*EmailBuilder)

// SendEmail using the builder instead of the domain object
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailImp(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar").
			To("bar@foo").
			Subject("Meeting").
			Body("Lets meet")
	})
}
