package main

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	c, err := client.DialTLS("imap.host.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	if err = c.Login("USER", "PASS"); err != nil {
		log.Fatal(err)
	}

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println(m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("messages:")
	for msg := range messages {
		log.Println(msg.SeqNum, msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
