all: gotwitter

gotwitter: http_auth.6 twitter.6 main.6
					 6l -o gotweet http_auth.6 twitter.6 main.6

main.6: main.go
	6g main.go

http_auth.6: http_auth.go
	6g http_auth.go

twitter.6: twitter.go
	6g twitter.go

clean:
	rm *.6 gotweet