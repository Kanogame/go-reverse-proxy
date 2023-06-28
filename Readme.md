# Reverse proxy, Load balancer, Http server or just another NGINX clone..

## Installation
 you can start using go-reverse-proxy by just downloading latest release and changing some stuff in config.txt

## Config
config starts from http, everything must be inside them, otherwise there will be a lot of errors)
```
http {
    ....
}
```

### WARNING: config supports only inline {}
this means you **cant** use something like
```
http 
{
    ....
}
```

### Firstly you need to set up basic:
* port
* log folder _(optional)_
* custom_404 _(optional)_
```
http 
{
    port: "8080"; 
	log: "./latest.log"; 
	custom_404: "./404/404.html";
}
```
_(semicolons are optional too)_

### Locations
there are three types of locations:
* static
* proxy
* proxy_load

```
location(/) { 
    type: "static"; 
    path: "./static"; 
} 

location(/proxy/) { 
    type: "proxy";
    path: "http://127.0.0.1:12312/";
} 

location(/app/) { 
type: "proxy_load"; 
path: ["http://127.0.0.1:12312/", "http://127.0.0.1:12322/"];
} 
```
## Thats all, I hope you like my project and that it was helpful)