# yb-http

yb-http is a lightweight and easy use webserver. It is very fast, 
because it is based on fiber. 
It is not made for everyone. It needs an index.html file in every file to run without errors.
It is more like a fun project than a real webserver for production use.


## Usage
yb-http is made for the use behind a reverse proxy. Because yb-http does not include
encryption via SSL or TLS it is insecure to use it without an proxy with encryption.
