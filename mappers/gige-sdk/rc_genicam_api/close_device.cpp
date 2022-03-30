#include "close_device.h"

extern "C" void close_device(MyDevice myDevice) {
	myDevice.dev->close();
}