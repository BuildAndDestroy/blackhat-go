# Keylogger example

Chapter 4 goes over an example of phishing web credentials. In this case, Roundcube.

## Prerequisites

* Roundcube docker image
```
docker pull robbertkl/roundcube
```

* Source code
```
docker run --rm -it -p 80:80 robbertkl/roundcube
```
Visit http://127.0.0.1 using Chrome browser. Cntrl + s to save the website and all of it's js files.
Name the file "index.html" and "Webpage,Complete. Your tree structure will look like below
but make sure you put them in the public directory you will need to make:
```
$ tree 
.
└── public
   ├── index_files
   │   ├── app.js
   │   ├── common.js
   │   ├── jquery.min.js
   │   ├── jquery-ui.css
   │   ├── jquery-ui.min.js
   │   ├── jstz.min.js
   │   ├── roundcube_logo.png
   │   ├── styles.css
   │   └── ui.js
   └── index.html

```

* Modify the form parameters to match your local server, not the 127 address, like so:
```
FROM   <form name="form" method="post" action="http://127.0.0.1/?_task=login">
TO     <form name="form" method="post" action="/login">
```
I have already done this for you, just keep this in mind when you create/mirror your own landing pages.

## Send it

```
go run main.go

OR

go build -o keyloggertest && ./keyloggertest
```

A "credentials.txt" file will be created and store our creds and client information.

## Note

This isn't exactly a keylogger, seems to be more of a phishing page. Still, the concepts
will help us get what we need.

