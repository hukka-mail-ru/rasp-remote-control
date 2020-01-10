# Raspberry Remote Control Project

This instuction was helpful: https://github.com/ole-vi/bluetooth-server


# Bluetooth preparation. 
## Setup the SPP (serial port profile)

```
apt-get install bluez-tools
```

Open terminal, edit this file

```
sudo nano /etc/systemd/system/dbus-org.bluez.service
```

Add -C at the end of the ExecStart= line, to start the bluetooth daemon in 'compatibility' mode. Add ExecStartPost=/usr/bin/sdptool add SP immediately after that line, to add the SP Profile. The two lines should look like this:

```
ExecStart=/usr/lib/bluetooth/bluetoothd -C

ExecStartPost=/usr/bin/sdptool add SP
```

Reboot RPi


## Step 2 - Start the Buetooth/RFCOMM server on RPi

Set RPi visible:

```
 bt-adapter --set Discoverable on
 bt-adapter --set  Pairable on  
```

Run the RFCOMM server on RPi (and set a big timeout to make the connection alive for a while):

```
 sudo rfcomm watch hci0 -L 86400
```

NOTE:  Observed problem: can't connect to the server the second time without restarting it.  The workaround: the arguments go a bit wrong way in the command above (-L 86400 must go first). It produces a small error inside of rfcomm, but hocus-pocus!  Now rfcomm behaves correctly if I close the connection on the Android side. That's a miracle! See  https://superuser.com/questions/1462985/bluetooth-bluez-rfcomm-listen-command-does-not-terminate/1512828#1512828


## Step 3 - Start the RPi pin-control Application

In another terminal listen the RFCOMM device for incoming data:
```
 cat /dev/rfcomm0
```

Download the code from here: https://github.com/hukka-mail-ru/rasp-remote-control

Compile and run the code:

```
go get
go run rasp-remote-control
```

## Step 4 - Send data
- Connect to the RFCOMM server on RPi via Android app (Serial Bluetooth Terminal)
- Send some data,
- Disconnect

# InfraRed supprot preparation

## Install and configure LIRC 
(Check this: https://github.com/mtraver/rpi-ir-remote)


Install LIRC:
```
sudo aptitude install lirc
```
Add to /boot/config.txt:
```
dtoverlay=gpio-ir,gpio_pin=22
dtoverlay=gpio-ir-tx,gpio_pin=23
```

Edit /etc/lirc/lirc_options.conf :  
```
[lircd]
nodaemon        = False
driver          = default
device          = /dev/lirc0
output          = /var/run/lirc/lircd
pidfile         = /var/run/lirc/lircd.pid
plugindir       = /usr/lib/arm-linux-gnueabihf/lirc/plugins
permission      = 666
allow-simulate  = No
repeat-max      = 600

[lircmd]
uinput          = False
nodaemon        = False
```
Reboot:
```
sudo reboot
```
Check the lirc daemon is running:
```
> ps aux | grep lirc                
root       507  0.0  0.3   7152  2968 ?        Ss   13:39   0:00 /usr/sbin/lircd --nodaemon
pi         949  0.0  0.0   4784   572 pts/0    S+   13:59   0:00 grep lirc
```

## Connect the IR receiver to RPi

(see https://www.instructables.com/id/How-To-Useemulate-remotes-with-Arduino-and-Raspber/)

The IR receiver has got 3 pins. Connect them to RPi pins, beginning with the left:
- 1st (OUT) -> to GPIO 22
- 2nd (GND) -> to GND pin
- 3rd (VCC) -> to 5V pin

## Test the IR receiver
```
cat /dev/lirc1
```
...and press any key on a IR remote control, pointed to the IR receiver.






