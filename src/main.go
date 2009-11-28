// Copyright 2009 Dmitry Chestnykh <dmitry@codingrobots.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os";
	"flag";
	"fmt";
	"./twitter";
)

var username = flag.String("u", "", "username (Twitter login)")
var password = flag.String("p", "", "password")

func requireLogin() {
	if *username == "" || *password == "" {
		flag.Usage();
		fmt.Fprintf(os.Stderr, "Username and password required for this function!\n");
		os.Exit(1);
	}
}

func checkForError(s string, err os.Error) string {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err);
		os.Exit(1);
	}
	return s;
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options...] <action> ...\n"
						   "Options:\n", os.Args[0]);
	flag.PrintDefaults();
	fmt.Fprintf(os.Stderr, "Actions:\n"
		"  post		Post status update (followed by status). Alias: p\n"
		"  user		Show user timeline. Alias: u\n"
		"  friends	Show friends timeline. Alias: (nothing)\n"
		"  mentions	Show mentions. Alias: @\n"
		"  public	Show public timeline\n");
}

func main() {
	flag.Usage = Usage;
	flag.Parse();

	tw := twitter.NewTwitter(*username, *password);

	switch flag.Arg(0) {
	case "@", "mentions":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.Mentions()));
	case "", "friends":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.FriendsTimeline()));
	case "u", "user":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.UserTimeline()));
	case "public":
		os.Stdout.WriteString(checkForError(tw.PublicTimeline()))
	case "p", "post":
		requireLogin();
		s := "";
		for i := 1; i < flag.NArg(); i++ {
			if i > 1 {
				s += " "
			}
			s += flag.Arg(i);
		}
		err := tw.UpdateStatus(s);
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err);
			os.Exit(1);			
		}
	}
	os.Exit(0);
}
