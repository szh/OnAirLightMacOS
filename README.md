szh/OnAirLightMacOs
=======================

This is a fork of [TravisTX/OnAirLightMacOS](https://github.com/TravisTX/OnAirLightMacOS) which uses a Raspberry Pi Zero W and an RGB LED (or two single color LEDs) instead of a Hue light. It causes a light to turn red when any microphone or camera is active on a Mac.

Build
-----

Run `./build.sh` to compile both parts - the client which runs on Mac OS, and the server which runs on a Raspberry Pi Zero W. This should also work with other Raspberry Pi models - just edit the `server/build.sh` script to build for the correct ARM version.

Server
------

To install the server on your Raspberry Pi Zero W, SSH or SCP into your Pi and copy the files in `bin/server/` to `/home/pi/` or some other directory. Then edit your `/etc/rc.local` file to run the server at boot, by adding this line before `exit 0`:

```bash
(cd /home/pi; sudo ./server > server.log &)
```

**Important:** Make sure to edit the `app.env` file to set the correct GPIO pin numbers and disable test mode!

Now restart your Pi with the command `sudo reboot`.

Client
------

To run the client on your Mac, copy the contents of `bin/macos/` to your user directory or anywhere you wish. Edit the `app.env` file to set the correct IP address of the Raspberry Pi Zero W (you can use a hostname instead if you want). Then set the app to run at startup by going to `System Preferences > Users and Groups > Login Items` and adding the `client` executable to the list of applications to launch at login. You probably want to check the box under "Hide" unless you want to see a terminal window with the output whenever you log in (although that's useful for debugging!).

Todo Items
-----------

- Add a timeout after which point the server resets the status if it hasn't received a request.
