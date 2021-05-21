# Tools

`make install-tools`

Tools repository based on this blog post:

https://marcofranssen.nl/manage-go-tools-via-go-modules/

## Rename binary
Setting a comment in the import line with the format `#name:<NEW_NAME>#` will
rename the binary file to `<NEW_NAME>` when running `make install-tools`
