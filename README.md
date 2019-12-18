# Raspberry Remote Control Project

This instuction was helpful: https://github.com/ole-vi/bluetooth-server


## Bluetooth preparation. Setup the SPP (serial port profile)

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


