#!/usr/bin/env python
import ctypes

lib = ctypes.CDLL('./sum.so')
print(lib.Sum(7, 11))
