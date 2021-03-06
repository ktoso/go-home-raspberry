go home raspberry
=================
Is a simple app to control home devices by turning them on/off, using a simple HTTP call.
The app should be running on a raspberry pi (least energy consuption) and be equiped with an telldus tellstick.

what you'll need
----------------
You'll need:

* raspberry pi, 
* telldus tellstick,
* sd card,
* radio controlled power outlets,
* ethernet / usb micro cables

<img src="https://raw.github.com/ktoso/go-home-raspberry/master/device.png" alt="the device"/>

demo
----
Here's a demo of the app in action: http://www.youtube.com/watch?v=HdCBzMMkDDY&list=UUYSzMKdiOqRHcGiWcORlKeQ&index=1

PS: Yeap, it's not restful (YET).   

to install go
-------------

    hg clone -u default https://code.google.com/p/go /opt/go
    cd /opt
    chown -R pi.pi go
    su - pi 
    cd /opt/go/src
    ./make.bash
    # may take 40 minutes

http://dave.cheney.net/tag/go-golang-raspberrypi

to install telldus tools
------------------------

http://elinux.org/R-Pi_Tellstick_core

    wget http://download.telldus.se/TellStick/Software/telldus-core/telldus-core-2.1.1.tar.gz
    sudo apt-get install libftdi1 libftdi-dev libconfuse0 libconfuse-dev cmake
    cd /usr/src
    sudo tar xzf ~/download/telldus-core-2.1.1.tar.gz
    cd telldus-core-2.1.1
    sudo cmake .
    sudo make
    sudo make install

then configure `/etc/tellstick.conf`

    user = "nobody"
    group = "plugdev"
    deviceNode = "/dev/tellstick"
    ignoreControllerConfirmation = "false"
    
    device {
      id = 1
      name = "C"
      controller = 0
      protocol = "sartano"
      model = "codeswitch:elro"
      parameters {
        # devices = ""
        house = "A"
        unit = "1"
        code = "1111100100"
        system = "1"
        # units = ""
        fade = "false"
      }
    }
    
    device {
      id = 2
      name = "B"
      controller = 0
      protocol = "sartano"
      model = "codeswitch:elro"
      parameters {
        # devices = ""
        house = "A"
        unit = "1"
        code = "1111101000"
        system = "1"
        # units = ""
        fade = "false"
      }
    }

then configure the daemon...

to install the go library
-------------------------

    go get code.google.com/p/gorest
    
    go run main.go

http://code.google.com/p/gorest/wiki/GettingStarted?tm=6


license
=======

I hereby release this software on the terms of the **beerware v2 open source license**.
To make it short, use it however you want, but if we meet and you like dthe software - buy me a beer (or slice of pizza).
