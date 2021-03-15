Lightweight go web server + templating stack.

Born out of frustration spending too much time searching for and trying to install complex CMS systems only to find out they didn't do what I wanted.

Features:

- produces single deployable binary.
- no external dependencies required (unless you want https, in which case install a full fledged web server and set up a reverse proxy + letencrypt).
- simple templating system (gives you a standard way to have a look and feel for your web site without cutting and pasting HTML)
- no javascript required. Very little css. 

Layout:

...

/assets/  <-- edit stylesheets and add images here
/templates/ <-- put your content here.  Convention is that if you want to serve up a page called /foo.html, then create files like this:
          foo.html
          foo.title
... 

Requirements:

go-1.16+

Usage:

...

go build
./grumpyweb >> your.log

...

Yes, this is how grumpy old developers roll. At least this one. YMMV

Future work: 

Rotate log files automatically.
