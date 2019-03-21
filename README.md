![](./animations/banner.png)

# Description

# SAM Architecture
The SAM reservation system consists of 4 processes:
1. User interaction with the website (Reservation Request),
1. Handling the request, persisting it into a database,
1. Notify other services that a new reservation has entered the system,
1. Listen for the event and act as required.

The GIF bellow illustrates the journey that a single reservation request makes. 

![](animations/stopsito-sam-animation.gif)

# How to use
To run the whole infrastructure: 
```
$ ./deploy.sh {aws_profile}
```

Obviously, some names and the profile used in cloudformation will need to change in order to accommodate your application. 

# License 
No license. You can re-use or change this as much as you like! 