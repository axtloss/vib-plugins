# Shim
A simple plugin that executes another program.

`Shim` takes the type of the module and executes the apropriate file in the plugins directory.

Any output over stdout will be returned to vib as the commands to include in the Containerfile.


# Building
This plugin requires [yyjson](https://github.com/ibireme/yyjson).
```
gcc plugin.c -lyyjson -fPIC -shared -o shim.so
```

# Usage
Due to the way vib loads plugins, the .so file has to be the same name as the program shim has to load.
In addition, the program also needs to be in the recipe plugin directory.

An example setup:
```
.
|- recipe.yml
\- plugins
    |- myPlugin.so   <-- This is the shim
    \- myPlugin      <-- This is the actual program
```
