# acari

# ><(((º>


## A bruteforce of directories with a low level of requests and that passes the endpoints in a list of URLs, one endpoint for all URLs and then passes the second endpoint and so on until the list of endpoints runs out

### Install:
```bash
go install  -v github.com/eikehacker1/acari@latest 
```

```bash
sudo cp ~/go/bin/acari /usr/bin
```
acari -h

-s string

     URL list file
   -w string
   
     Endpoint list file

##### acari -s listurls.txt -w your-wordlist

# Yes it only takes 200 but use httpx to make sure
```bash
cari -s listurls.txt -w your-wordlist | httpx -status-code -mc 20
```

