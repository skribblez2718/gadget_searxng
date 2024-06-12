# Gadget SearXNG
gadget_searxng provides an automated way to get URLs from SearXNG search results. These URLs can then be used to enumerate subdomains, endpoints, etc. and be fed into other tools for additional scope discovery during the information gathering phase of a penetration test, red team engagement or bug bounty. 

##  Install
```sh
git clone https://github.com/skribblez2718/gadget_searxng.git
cd gadget_searxng
go build
```

## Usage
```sh
gadget_searxng  --host "<searxng_host>" --searchTerm "<search_term>" --output "<output_file>" [--maxPages <int>]
```

## References
[SearXNG Docs](https://docs.searxng.org/)
