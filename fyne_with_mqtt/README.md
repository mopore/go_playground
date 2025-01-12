Install the fyne command line tool by running:
```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
```

Ensure to to login with your Apple ID in Xcode (Xcode --> Settings)
Also create a new certificate under "manage certificates".


Create the iOS application bundle by running:
```shell
fyne package -os ios -appID org.mopore.fyne_mqtt_test -icon ./icon.png
```

Should you get an error that "xcode is required" reset the path to the
command line tools by:
```shell
sudo xcode-select --switch /Applications/Xcode.app/Contents/Developer
```

Finding a good app icon which gets accepted seems to be difficult.

## Install app on iPHone
On the iPhone enable Developer Mode in Settings --> Pricacy & Security
This is quite at the bottom of the list and will require a restart
and furhter confirmation after the restart.


In Xcode go to "Window" --> "Devices and Simulators"
It will then take some time to have the iPhone fully connected since 
data seems to be transferred via the USB cable.

Then click on the iPhone and select "Install App" and select the 
earlier generated app file.

Export the certificate from Xcode and import it into Apple Keychain.
Open Keychain Access and double click on the certificate and set 
the certificate to "Always Trust".







