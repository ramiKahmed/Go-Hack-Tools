# Go-Hack-Tools
lightweight tools  written in Golang to aid in the pentesting &amp; red teaming phases .


## certNamesLookup
- Retrieves hostnames that are related to the main company ( the domain provided ) included in the SSL certificates. ( Assets Discovery )
```

Usage: go run certNamesLookup.go -domain google.com 
```


## SSLScraper
- Fast concurrent SSL Scraper that Extract Common names from SSL certificates by providing a file containing ip:port format.
```
Usage: go run SSLScraper.go -file <filename> 
```

