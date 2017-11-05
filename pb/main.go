package main

import (
	pb "./pb3"
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func promptForAddress(r io.Reader) (*pb.Person, error) {
	p := &pb.Person{}
	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID number: ")
	// An int32 field in the .proto file is represented as an int32 field
	// in the generated Go struct.
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	// A string field in the .proto file results in a string field in Go.
	// We trim the whitespace because rd.ReadString includes the trailing
	// newline character in its output.
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}
		// The
		// PhoneNumber
		// message
		// type
		// is
		// nested
		// within
		// the
		// Person
		// message
		// in
		// the
		// .proto
		// file.
		// This
		// results
		// in
		// a
		// Go
		// struct
		// named
		// using
		// the
		// name
		// of
		// the
		// parent
		// prefixed
		// to
		// the
		// name
		// of
		// the
		// nested
		// message.
		// Just
		// as
		// with
		// pb.Person,
		// it
		// can
		// be
		// created
		// like
		// any
		// other
		// struct.
		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		ptype = strings.TrimSpace(ptype)

		// A
		// proto
		// enum
		// results
		// in
		// a
		// Go
		// constant
		// for
		// each
		// enum
		// value.
		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q.  Using default.\n", ptype)
		}

		// A
		// repeated
		// proto
		// field
		// maps
		// to
		// a
		// slice
		// field
		// in
		// Go.
		// We
		// can
		// append
		// to
		// it
		// like
		// any
		// other
		// slice.
		p.Phones = append(p.Phones, pn)
	}

	return p, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <addressboo_file>\n", os.Args[0])
	}
	fname := os.Args[1]

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found. Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}
	book := &pb.AddressBook{}
	if err = proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error with address:", err)
	}
	book.People = append(book.People, addr)
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
