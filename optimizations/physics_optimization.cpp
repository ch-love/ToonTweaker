// physics_optimization.cpp
#include <cmath>

extern "C" {
    // Function to calculate position using velocity and acceleration
    double calculate_position(double velocity, double acceleration, double time) {
        return velocity * time + 0.5 * acceleration * pow(time, 2);
    }
}
