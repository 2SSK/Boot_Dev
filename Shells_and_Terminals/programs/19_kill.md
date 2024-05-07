# KILL

Sometimes a porgram is in such a bad state (or is so malicious) that it doesn't respond to **SIGNIT**, in which case the best option is to use another shell session (new terminal window) to manually <span style="color:violet">**kill**</span> the program.

## SYNTAX

    kill PID

**PID** stads for "process ID". Every process that's running on your machine has a unique ID.<br>
The <span style="color:violet">**ps**</span>, "process status" command can be used to list the processes running on your machine, and their IDs:

    ps aux

The "aux" options just mean "show all processes, inculding those owned by other users, and show extra information about each process".
