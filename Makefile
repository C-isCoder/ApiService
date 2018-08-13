all: gotool
^I@go build -v .
clean:
^Irm -f apiservice
^Ifind . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
^Igofmt -w .
^Igo tool vet . |& grep -v vendor;ture
# ca:
^I# openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"...

help:
^I@echo "make - compile the source code"
^I@echo "make clean - remove binary file and vim swp files"
^I@echo "make gotool - run go tool 'fmt' and 'vet'"
^I# @echo "make ca - generate ca files"

.PHONY: clean gotool ca help
