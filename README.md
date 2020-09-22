
# Go No IP
A lookup service for Hostname and dynamic IP addresses

This program runs on a well-known static IP.

Site registration
- a service sets its hostname and ip via a POST to https://.../api/address/a.b.com&ip=a.b.c.d

- a client requests the service IP via a GET to https://.../api/address/a.b.com


## Add dist path
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

## build
go install

## run
go-noip