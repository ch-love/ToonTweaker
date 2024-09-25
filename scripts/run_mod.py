# run_mod.py
import ctypes
import os

# Load the shared library
lib = ctypes.CDLL(os.path.join(os.path.dirname(__file__), '../optimizations/optimizations.so'))

# Define the argument and return types for the C++ function
lib.calculate_position.argtypes = (ctypes.c_double, ctypes.c_double, ctypes.c_double)
lib.calculate_position.restype = ctypes.c_double

def main():
    # Example usage of the C++ function
    velocity = 10.0
    acceleration = 2.0
    time = 5.0
    position = lib.calculate_position(velocity, acceleration, time)
    print(f"Calculated Position: {position}")

if __name__ == "__main__":
    main()
