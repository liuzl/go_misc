package main

import (
	"fmt"

	"github.com/nyaruka/ezconf"
)

// Config is our apps configuration struct, we can use the `help` tag to add usage information
type Config struct {
	AWSRegion     string `help:"the aws region that S3 buckets are in"`
	DB            string `help:"the url describing how to connect to the database"`
	EC2InstanceID string `help:"the id of the ec2 instance we are running on"`
	NumWorkers    int    `help:"the number of workers to start"`
}

func main() {
	// instantiate our config with our defaults
	config := &Config{
		AWSRegion:     "us-west-2",
		DB:            "postgres://user@secret:rds-internal.foo.bar/courier",
		EC2InstanceID: "i-12355111134a",
		NumWorkers:    32,
	}

	// create our loader object, configured with configuration struct (must be a pointer), our name
	// and description, as well as any files we want to search for
	loader := ezconf.NewLoader(
		config,
		"courier", "Courier - a fast message broker for IP and SMS messages",
		[]string{"courier.toml"},
	)

	// load our configuration, exiting if we encounter any errors
	loader.MustLoad()

	// our settings have now been loaded into our config struct
	fmt.Printf("Final Settings:\n%+v\n", *config)

	// if we wish we can also validate our config using our favorite validation library
}
