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

package http_auth

import (
	"http";
	"encoding/base64";
	"io";
	"os";
	"strings";
	"net";
	"bufio";
	"strconv";
	"fmt";
	"bytes";
)

type readClose struct {
	io.Reader;
	io.Closer;
}

type badStringError struct {
	what	string;
	str	string;
}

func (e *badStringError) String() string	{ return fmt.Sprintf("%s %q", e.what, e.str) }

// Given a string of the form "host", "host:port", or "[ipv6::address]:port",
// return true if the string includes a port.
func hasPort(s string) bool	{ return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func send(req *http.Request) (resp *http.Response, err os.Error) {
	addr := req.URL.Host;
	if !hasPort(addr) {
		addr += ":http"
	}
	conn, err := net.Dial("tcp", "", addr);
	if err != nil {
		return nil, err
	}

	err = req.Write(conn);
	if err != nil {
		conn.Close();
		return nil, err;
	}

	reader := bufio.NewReader(conn);
	resp, err = http.ReadResponse(reader);
	if err != nil {
		conn.Close();
		return nil, err;
	}

	r := io.Reader(reader);
	if v := resp.GetHeader("Content-Length"); v != "" {
		n, err := strconv.Atoi64(v);
		if err != nil {
			return nil, &badStringError{"invalid Content-Length", v}
		}
		r = io.LimitReader(r, n);
	}
	resp.Body = readClose{r, conn};

	return;
}

func encodedUsernameAndPassword(user, pwd string) string {
	bb := &bytes.Buffer{};
	encoder := base64.NewEncoder(base64.StdEncoding, bb);
	encoder.Write(strings.Bytes(user + ":" + pwd));
	encoder.Close();
	return bb.String();
}

func Get(url, user, pwd string) (r *http.Response, err os.Error) {
	var req http.Request;

	req.Header = map[string]string{"Authorization": "Basic " +
		encodedUsernameAndPassword(user, pwd)};
	if req.URL, err = http.ParseURL(url); err != nil {
		return
	}
	if r, err = send(&req); err != nil {
		return
	}
	return;
}


// Post issues a POST to the specified URL.
//
// Caller should close r.Body when done reading it.
func Post(url, user, pwd, bodyType string, body io.Reader) (r *http.Response, err os.Error) {
	var req http.Request;
	req.Method = "POST";
	req.Body = body;
	req.Header = map[string]string{
		"Content-Type": bodyType,
		"Transfer-Encoding": "chunked",
		"X-Twitter-Client": "gotweet",
		"X-Twitter-Version": "0.1",
		"Authorization": "Basic " + encodedUsernameAndPassword(user, pwd),
	};

	req.URL, err = http.ParseURL(url);
	if err != nil {
		return nil, err
	}

	return send(&req);
}
