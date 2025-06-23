#!/usr/bin/env bash

go build -o write_battery_level ./main.go

mkdir -p "$HOME/Dev/arch_virt_machine/jni_ext"

cp write_battery_level "$HOME/Dev/arch_virt_machine/jni_ext"

cp com.jni.updatebatterylevel.plist "$HOME/Library/LaunchAgents"
launchctl load "$HOME/Library/LaunchAgents/com.jni.updatebatterylevel.plist"

echo "Battery level agent installed"

# Check with launchctl list | grep jni
