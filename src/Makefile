include $(GOROOT)/src/Make.$(GOARCH)

all: gotwitter

gotwitter: http_auth.$(O) twitter.$(O) main.$(O)
					 $(LD) -o gotweet main.$(O)

main.$(O): main.go
	$(GC) main.go

http_auth.$(O): http_auth.go
	$(GC) http_auth.go

twitter.$(O): twitter.go
	$(GC) twitter.go

clean:
	rm *.$(O) gotweet
