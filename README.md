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

http://developer.telldus.com/wiki/TellStickInstallationUbuntu

    sudo echo 'deb http://download.telldus.com/debian/ stable main' >> /etc/apt/sources.list
    wget -q http://download.telldus.se/debian/telldus-public.key -O- | sudo apt-key add -
    sudo apt-get update
    sudo apt-get install telldus-core    

to install the go library
-------------------------

    go get code.google.com/p/gorest
    
    go run main.go

http://code.google.com/p/gorest/wiki/GettingStarted?tm=6
