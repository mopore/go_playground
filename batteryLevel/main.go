package main

/*
#cgo LDFLAGS: -framework IOKit -framework CoreFoundation
#include <IOKit/ps/IOPowerSources.h>
#include <IOKit/ps/IOPSKeys.h>
#include <CoreFoundation/CoreFoundation.h>

int getBatteryLevel() {
    CFTypeRef powerSourcesInfo = IOPSCopyPowerSourcesInfo();
    if (powerSourcesInfo == NULL) {
        return -1;
    }

    CFArrayRef powerSourcesList = IOPSCopyPowerSourcesList(powerSourcesInfo);
    if (powerSourcesList == NULL || CFArrayGetCount(powerSourcesList) == 0) {
        CFRelease(powerSourcesInfo);
        return -1;
    }

    CFDictionaryRef powerSource = IOPSGetPowerSourceDescription(powerSourcesInfo, CFArrayGetValueAtIndex(powerSourcesList, 0));
    if (powerSource == NULL) {
        CFRelease(powerSourcesInfo);
        CFRelease(powerSourcesList);
        return -1;
    }

    CFNumberRef capacity = CFDictionaryGetValue(powerSource, CFSTR(kIOPSCurrentCapacityKey));
    int level = 0;
    CFNumberGetValue(capacity, kCFNumberIntType, &level);

    CFRelease(powerSourcesInfo);
    CFRelease(powerSourcesList);

    return level;
}
*/
import "C"
import (
	"fmt"
	"log"
	"os"
)

const (
	batFile = "/Users/jni/arch_share/virt_machine/jni_ext/battery_level_percent.txt"
)


func main() {
	batteryLevel := C.getBatteryLevel()
	text := fmt.Sprintf("%d", batteryLevel)

	file, err := os.OpenFile(batFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
		panic(err)
	}
	defer file.Close()

	file.WriteString(text)
}
