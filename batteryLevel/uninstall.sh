#!/usr/bin/env bash

launchctl unload "$HOME/Library/LaunchAgents/com.jni.updatebatterylevel.plist"
rm "$HOME"/arch_share/virt_machine/jni_ext/write_battery_level*
rm "$HOME"/arch_share/virt_machine/jni_ext/battery_level_percent.txt

echo "Battery Level agent uninstalled."

# Check with launchctl list | grep jni
