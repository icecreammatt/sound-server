# sound-server

This is a web server that uses `mpg123` to play sound files located in the `/root` directory.
The usage for this is to play sound effects when a door in my house is opened. Such as the Seinfield theme song.
I have a separate device that takes care of monitoring when door open and close events happen. See this script [here](https://github.com/icecreammatt/gopherwink/blob/master/scripts/monitor.sh#L21)

`curl <host-ip>:8080/sound-file.mp3` where `sound-file.mp3` is located in the `/root` directory. To play a sound file on the remote device.

## Building

`build-and-deploy.sh` can be used to build and deploy to a target device running ARM linux such as a Raspberry Pi or C.H.I.P.

## Remote device setup

`mpg123` is used to play the sound files but this can easily be swapped out with another CLI audio player

* Linux `apt-get install mpg123`
* Mac `brew install mpg123`

Update the the `build-and-deploy.sh` script so it contains the IP of the device desired to deploy to.
Deploy the binary to the device and then set it up so it starts automatically on device startup.
Ensure that there are audio files in `/root`.

## Testing

Drop some sample sound files in the `/root` directory of the remote device. Then hit the current URL in the browser or command line
`curl <remote-ip>:8080/sound-file.mp3`