# Evernote-ENEX-to-Diigo-Bookmarks
A Go programming language project for adding title, tags, and source url of notes in an Evernote exported ENEX file as Diigo bookmarks via the Diigo API.

# How to Use
1. Get your Diigo API key from https://www.diigo.com/api_keys/new/
2. In the <B>post</B> function, replace <B><...></B> witih your Diigo API key, username, and password, correspondingly.

    	var key string = "<your Diigo API key>"
    	var username string = "<your Diigo username>"
    	var passwd string = "<your Diigo password>"
3. Save and Compile the go file to an executable, e.g. e2d.exe
4. Run the app in command-line interface, for example, > e2d.exe Evernote.enex

# Exception
Some characters in the ENEX file may cause an error in parsing the xml structure, and make it fail to run.
Try to eliminate imcompatiable notes during export your ENEX file from Evernote.
