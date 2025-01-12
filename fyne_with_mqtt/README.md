
# Experimantal Fyne App for IOS
* Note that this does not yet work! *


## Creating the iOS app
On the iPhone enable Developer Mode in Settings --> Pricacy & Security
This is quite at the bottom of the list and will require a restart
and furhter confirmation after the restart.

After installing xcode and correcting the path to the command line tools
Open Xcode and install the iOS development tools.
Restart Xcode.

In Xcode login with your Apple ID as the developer account (which needs to be
payed for).

Create a certificate (this will be specific for your local dev device).

In the apple developer portl create a "profile" for the app.
https://developer.apple.com/account/resources/profiles/list

You will need an app ID (a created org.mopore.test) before.
In the next step you shulld be able to attach the certificate generated before.

For the next steup you will need your device ID.
In Xcode go to "Window" --> "Devices and Simulators"
It will then take some time to have the iPhone fully connected since 
data seems to be transferred via the USB cable.

After everything is done download your certificat, import it in your keychain 
and set it to "always trust".

Install the fyne command line tool by running:
```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
```

Inside the project folder create the iOS application bundle by running:
```shell
fyne package -os ios -appID org.mopore.test -icon ./icon.png
```

Should you get an error that "xcode is required" reset the path to the
command line tools by:
```shell
sudo xcode-select --switch /Applications/Xcode.app/Contents/Developer
```

Finding a good app icon which gets accepted seems to be difficult.


## Installing the app on the iPhone
Via xcode find your iPhone in the "Devices and Simulators" window.
Then click on the iPhone and select "Install App" and select the 
earlier generated app file.








