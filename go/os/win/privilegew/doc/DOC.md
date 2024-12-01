# How to Obtain Administrator Privileges for Windows Programs

## Creating a Manifest File

Create a file named manifest.xml with the following content:


```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
  <trustInfo xmlns="urn:schemas-microsoft-com:asm.v2">
    <security>
      <requestedPrivileges>
        <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
      </requestedPrivileges>
    </security>
  </trustInfo>
</assembly>

```

## Embedding the Manifest File into a Go Program

You can use the rsrc tool to embed the manifest file into your Go program.

Install the rsrc tool:

```shell
go install github.com/akavel/rsrc@latest
```

Generate a .syso file using the rsrc tool in the main package directory:

```shell
rsrc -manifest manifest.xml
```

Compile your Go program in the main package directory:

```shell
go build -o your_program.exe
```
