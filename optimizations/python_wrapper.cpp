// python_wrapper.cpp
#include <Python.h>
#include "physics_optimization.h"

// Wrapper function for Python
static PyObject* py_calculate_position(PyObject* self, PyObject* args) {
    double velocity, acceleration, time;

    // Parse the Python arguments
    if (!PyArg_ParseTuple(args, "ddd", &velocity, &acceleration, &time)) {
        return NULL;
    }

    // Call the C++ function
    double result = calculate_position(velocity, acceleration, time);
    
    // Return the result as a Python float
    return Py_BuildValue("d", result);
}

// Method definition for the module
static PyMethodDef OptimizationMethods[] = {
    {"calculate_position", py_calculate_position, METH_VARARGS, "Calculate position."},
    {NULL, NULL, 0, NULL} // Sentinel
};

// Module definition
static struct PyModuleDef optimizationmodule = {
    PyModuleDef_HEAD_INIT,
    "optimizations", // Module name
    NULL,           // Module documentation
    -1,             // Size of per-interpreter state of the module
    OptimizationMethods
};

// Module initialization function
PyMODINIT_FUNC PyInit_optimizations(void) {
    return PyModule_Create(&optimizationmodule);
}
