package main

import "fmt"
import "github.com/pspiagicw/demp"

func main() {

	variables := map[string]string{
		"name":         "exampleName",
		"mail":         "exampleMail",
		"home-address": "ExampleHomeAddress",
	}

	tt := []struct {
		template string
		result   string
	}{
		{"${name} ${mail}", "exampleName exampleMail"},
		{"$name ${mail}", "exampleName exampleMail"},
		{"$name ${mail}", "exampleName exampleMail"},
		{"$name $mail", "exampleName exampleMail"},
		{"$name$mail", "exampleNameexampleMail"},
		{"$namemail", "$namemail"},
		{"${name}mail", "exampleNamemail"},
		{"${name}$mail", "exampleNameexampleMail"},
	}

	for _, t := range tt {
		result := demp.ResolveTemplate(t.template, variables)
		if result != t.result {
			fmt.Printf("Expected %s, got %s\n", t.result, result)
		}
	}
}
