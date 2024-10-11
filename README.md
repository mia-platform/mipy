# mipy
Cli helper to help in IaC provisioning through the Console pipelines

## Setup

The mipy cli makes use of a configmap file as this one:
```json
{
    "basePath": "string",
    "templates": [
        {
            "type": "enum", 
            "path": "string"
        }
    ],
    "logLevel": "string"
}
```

## Commands

### version

```
mipy version
```

### help

```
mipy help
```

### config

flags are:
- get
- set [PATH]
- --help

```
mipy config set mipy.json
```

### init

```
mipy init
```

`mipy init` lets you preview the CRs that mipy plans to execute.

### launch

flags are:
- --cr-list
- --parallel
- --error-code
- debug
- dry-run