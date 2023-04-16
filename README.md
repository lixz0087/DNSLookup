# DNS Lookup

This app allow users to perform DNS lookup via API as well as store DNS queries history. It is a containerized app.

## Endpoints:

```POST localhost:8080/api/v1/dig```
Post with a json body containing: domain or IP to lookup, DNS type (e.g. A, MX, PTR), DNS server for query (e.g. 8.8.8.8, or tesla.com). Returns DNS records from DNS server.

Example input:
```
{
	"domain": "www.google.com",
	"dns_type": "A",
	"dns_server": "8.8.8.8"
}
```

```GET localhost:8080/api/v1/dig/history``` Returns history of last 30 DNS queries in json format with timestamp

## Supported DNS Types:
```CNAME, NSEC, NSEC3, ATMA, CERT, MR, RP, RT, SOA, TXT, DHCID, DNAME, NSAPPTR, MX, LOC, DS, AFSDB, A, DNSKEY, MINFO, AAAA, NS, ISDN, TLSA, MB, MG, PTR, SRV, RRSIG, CAA, HINFO, NAPTR, X25```

## Steps to Run The APP
1. Make sure docker is installed https://docs.docker.com/desktop/install/mac-install/
2. Build the docker image with command ` docker build -t dnslookup .`
3. Run the docker image on port `8080` with command `docker run -p 8080:8081 -it dnslookup`
4. Now you can send HTTP request to different endpoints of the app.

## Unit Tests
To run all unit tests for the project, run this command in the root directory:
```go test -v ./...```