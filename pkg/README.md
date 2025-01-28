# pkg Directory

The `pkg` directory is used to store compiled package objects. This directory serves as an intermediate compilation cache and has the following purposes:

- Stores compiled package files (`.a` files)
- Caches compiled packages to speed up subsequent builds
- Contains architecture-specific compiled packages (e.g., linux_amd64, windows_amd64)

Note: Like the `bin` directory, this directory is managed automatically by Go's build system. The contents are used to speed up compilation by avoiding recompilation of unchanged packages. 