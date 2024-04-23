# freebsd-pkg

A plugin to install packages using FreeBSD's `pkg(8)` written in rust.

## Module structure

```yaml
Name: PkgModule
Type: Pkg
ExtraFlags:
    - "--debug"
Packages:
    - "bash"
    - "fish"
```

`ExtraFlags` can remain empty, the default flags passed to pkg are `install -y`

