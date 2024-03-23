package demp

import "testing"

func TestDemp(t *testing.T) {
	variables := map[string]string{
		"name":         "exampleName",
		"mail":         "exampleMail",
		"home-address": "ExampleHomeAddress",
		"binary":       "exampleBinary",
		"ldflags":      "-X main.VERSION=1.0.0",
		"main-file":    "main.go",
		"pwd":          "/some/dir",
	}

	tt := []struct {
		template string
		result   string
	}{
		{"${name} ${mail}", "exampleName exampleMail"},
		{"$name ${mail}", "exampleName exampleMail"},
		{"${name} $mail", "exampleName exampleMail"},
		{"$name $mail", "exampleName exampleMail"},
		{"$name$mail", "exampleNameexampleMail"},
		{"$namemail", "namemail"},
		{"$$$name", "$exampleName"},
		{"${name}mail", "exampleNamemail"},
		{"${name}$mail", "exampleNameexampleMail"},

		{`go build -o $binary -ldflags "$ldflags" ${main-file}`, `go build -o exampleBinary -ldflags "-X main.VERSION=1.0.0" main.go`},
		{`docker run --rm -v ${pwd}:/work --user 1000:1000 pspiagicw/doc-generator`, `docker run --rm -v /some/dir:/work --user 1000:1000 pspiagicw/doc-generator`},
	}

	for _, template := range tt {
		result := ResolveTemplate(template.template, variables)
		if result != template.result {
			t.Errorf("Expected %s, got '%s'\n", template.result, result)
		}
	}

}
