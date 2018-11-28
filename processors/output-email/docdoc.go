// Code generated by "bitfanDoc "; DO NOT EDIT
package email

import "bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
		Name:       "email",
		ImportPath: "bitfan/processors/output-email",
		Doc:        "Send email when an output is received. Alternatively, you may include or exclude the email output execution using conditionals.",
		DocShort:   "Sends email to a specified address when output is received",
		Options: &doc.ProcessorOptions{
			Doc: "",
			Options: []*doc.ProcessorOption{
				&doc.ProcessorOption{
					Name:           "Host",
					Alias:          "address",
					Doc:            "The address used to connect to the mail server",
					Required:       false,
					Type:           "string",
					DefaultValue:   "\"localhost\"",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Port",
					Alias:          "port",
					Doc:            "Port used to communicate with the mail server",
					Required:       false,
					Type:           "int",
					DefaultValue:   "25",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Username",
					Alias:          "username",
					Doc:            "Username to authenticate with the server",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Password",
					Alias:          "password",
					Doc:            "Password to authenticate with the server",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "From",
					Alias:          "from",
					Doc:            "The fully-qualified email address for the From: field in the email",
					Required:       false,
					Type:           "string",
					DefaultValue:   "\"bitfan@nowhere.com\"",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Replyto",
					Alias:          "replyto",
					Doc:            "The fully qualified email address for the Reply-To: field",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "replyto => \"test@nowhere.com\"",
				},
				&doc.ProcessorOption{
					Name:           "To",
					Alias:          "to",
					Doc:            "The fully-qualified email address to send the email to.\n\nThis field also accepts a comma-separated string of addresses, for example: `\"me@host.com, you@host.com\"`\n\nYou can also use dynamic fields from the event with the %{fieldname} syntax",
					Required:       true,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "to => \"me@host.com, you@host.com\"",
				},
				&doc.ProcessorOption{
					Name:           "Cc",
					Alias:          "cc",
					Doc:            "The fully-qualified email address(es) to include as cc: address(es).\n\nThis field also accepts a comma-separated string of addresses, for example: `\"me@host.com, you@host.com\"`",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "cc => \"me@host.com, you@host.com\"",
				},
				&doc.ProcessorOption{
					Name:           "Bcc",
					Alias:          "bcc",
					Doc:            "The fully-qualified email address(es) to include as bcc: address(es).\n\nThis field also accepts a comma-separated string of addresses, for example: `\"me@host.com, you@host.com\"`",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "bcc => \"me@host.com, you@host.com\"",
				},
				&doc.ProcessorOption{
					Name:           "Subject",
					Alias:          "subject",
					Doc:            "Subject: for the email\n\nYou can use template",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "subject => \"message from {{.host}}\"",
				},
				&doc.ProcessorOption{
					Name:           "Subjectfile",
					Alias:          "subjectfile",
					Doc:            "Path to Subject template file for the email",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "HTMLBody",
					Alias:          "htmlbody",
					Doc:            "HTML Body for the email, which may contain HTML markup",
					Required:       false,
					Type:           "location",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "htmlBody => \"<h1>Hello</h1> message received : {{.message}}\"",
				},
				&doc.ProcessorOption{
					Name:           "Body",
					Alias:          "body",
					Doc:            "Body for the email - plain text only.",
					Required:       false,
					Type:           "location",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "body => \"message : {{.message}}. from {{.host}}.\"",
				},
				&doc.ProcessorOption{
					Name:           "Attachments",
					Alias:          "attachments",
					Doc:            "Attachments - specify the name(s) and location(s) of the files",
					Required:       false,
					Type:           "array",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "AttachEventData",
					Alias:          "attachments_with_event",
					Doc:            "Use event field's values as attachment content\neach pair is  : event field's path => attachment's name",
					Required:       false,
					Type:           "hash",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      " attachments_with_event=>{\"mydata\"=>\"myimage.jpg\"}",
				},
				&doc.ProcessorOption{
					Name:           "Images",
					Alias:          "images",
					Doc:            "Images - specify the name(s) and location(s) of the images",
					Required:       false,
					Type:           "array",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "EmbedB64Images",
					Alias:          "embed_b64_images",
					Doc:            "Search for img:data in HTML body, and replace them to a reference to inline attachment",
					Required:       false,
					Type:           "bool",
					DefaultValue:   "false",
					PossibleValues: []string{},
					ExampleLS:      "",
				},
			},
		},
		Ports: []*doc.ProcessorPort{},
	}
}
