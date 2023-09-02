```
 __  __  ___ _____ _____   ____                        _      
|  \/  |/ _ \_   _|_   _| / ___|  __ _ _ __ ___  _ __ | | ___ 
| |\/| | | | || |   | |   \___ \ / _` | '_ ` _ \| '_ \| |/ _ \
| |  | | |_| || |   | |    ___) | (_| | | | | | | |_) | |  __/
|_|  |_|\__\_\|_|   |_|   |____/ \__,_|_| |_| |_| .__/|_|\___|
                                                |_|           
```

How was it created?
===================

```bash
mkdir mqtt_example mqttSimpleSample
cd mqttSimpleSample
go mod init mqttSimpleSample
go get -u github.com/eclipse/paho.mqtt.golang
vim main.go  # Add your code
cd main
go run .
```

To make a build from main directory:
```bash
go build -o mqtt_executable main/main.go
```

Note that we run `go run .` to include all files in the current directory.
