#!/bin/sh
#
meson setup --wipe build
cd build
meson compile
cd -
