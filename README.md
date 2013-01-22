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
