```
  ____      _                 _____         _   
 / ___|___ | |__  _ __ __ _  |_   _|__  ___| |_ 
| |   / _ \| '_ \| '__/ _` |   | |/ _ \/ __| __|
| |__| (_) | |_) | | | (_| |   | |  __/\__ \ |_ 
 \____\___/|_.__/|_|  \__,_|   |_|\___||___/\__|

```

# Install Cobra
Install Cobra (globally) `go install github.com/spf13/cobra-cli@latest`.
To install cobra you will need to have your $GOPATH/bin in your $PATH.
I added the following lines add the end of my `~/.zshrc` file:
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

# Create a Cobra CLI tool
Create your project directory `mkdir cobraCliTool`.
Initialize your project `go mot init cobra-test`.
Get Cobra for your project `go get -u github.com/spf13/cobra@latest`.
Create your Cobra CLI tool `cobra-cli init`.

# Run and install
Run it with `go run .` or `go run main.go`.
Install with `go install cobra-test`.
After installing you can run it with `cobra-test`.

