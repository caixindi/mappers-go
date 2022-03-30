#pragma once
#include <rc_genicam_api/device.h>
#include "const.h"

/**
  Turn off a Genicam device.

  @param mydevice  The MyDevice struct encapsulates a Genicam device
*/
extern "C" void close_device(MyDevice myDevice);
