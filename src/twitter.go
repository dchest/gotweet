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

package twitter

import (
	"./http_auth";
	"io";
	"fmt";
	"json";
	"os";
	"regexp";
	"http";
	"bytes";
	"strings";
)

type Twitter struct {
	Account Account;
}

type Account struct {
	Username	string;
	Password	string;
}

type User struct {
	Screen_name string;
}

type Tweet struct {
	Text	string;
	Source	string;
	User	User;
}

const (
	mentionsURL		= "http://twitter.com/statuses/mentions.json";
	friendsTimelineURL	= "http://twitter.com/statuses/friends_timeline.json";
	userTimelineURL		= "http://twitter.com/statuses/user_timeline";	// no .json!
	publicTimelineURL	= "http://twitter.com/statuses/public_timeline.json";
	updateURL		= "http://twitter.com/statuses/update.json";
)

func NewTwitter(user, pwd string) *Twitter {
	var acc Account;
	acc.Username = user;
	acc.Password = pwd;
	return &Twitter{acc};
}

func (t *Twitter) getTimeline(url string) (out string, err os.Error) {
	var r *http.Response;	
	r, err = http_auth.Get(url, t.Account.Username, t.Account.Password);

	if err != nil {
		return
	}

	if r.StatusCode != 200 {
		err = os.ErrorString("Twitter returned: " + r.Status);
		return
	}

	b, _ := io.ReadAll(r.Body);
	r.Body.Close();

	tweets := make([]Tweet, 20);
	json.Unmarshal(string(b), tweets);

	re, _ := regexp.Compile("<a[^>]*>(.*)</a>");
	for _, t := range tweets {
		// Source may be a link: <a href="...">source</a>
		// Extract text of the link with regexp.
		matches := re.MatchStrings(t.Source);
		if matches != nil && len(matches) > 0 {
			t.Source = matches[1]
		}

		if t.Text != "" {
			out = out + fmt.Sprintf("%v: %v (%v)\n\n", t.User.Screen_name,
				t.Text, t.Source)
		}
	}
	return;
}

func (t *Twitter) post(url, s string) (our string, err os.Error) {
	bb := &bytes.Buffer{};
	bb.Write(strings.Bytes(s));

	var r *http.Response;	
	r, err = http_auth.Post(url, t.Account.Username, t.Account.Password,
		"application/x-www-form-urlencoded", bb);

	if err != nil {
		return
	}

	b, _ := io.ReadAll(r.Body);
	r.Body.Close();

	if r.StatusCode != 200 {
		err = os.ErrorString("Twitter returned: " + r.Status);
		return
	}

	return string(b), nil;
}


func (t *Twitter) Mentions() (string, os.Error) { 
	return t.getTimeline(mentionsURL)
}

func (t *Twitter) FriendsTimeline() (string, os.Error) {
	return t.getTimeline(friendsTimelineURL)
}

func (t *Twitter) UserTimeline() (string, os.Error) {
	return t.getTimeline(userTimelineURL + "/" + t.Account.Username + ".json")
}

func (t *Twitter) PublicTimeline() (string, os.Error) { 
	return t.getTimeline(publicTimelineURL) 
}

func (t *Twitter) UpdateStatus(s string) os.Error {
	s = "status=" + http.URLEscape(s);
	_, err := t.post(updateURL, s);
	return err;
}
