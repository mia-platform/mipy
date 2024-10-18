# mipy 🐘
Cli helper to help in IaC provisioning through the Console pipelines

## Setup

The mipy cli makes use of a configmap file as this one:
```json
{
    "basePath": "string",
    "templates": [
        {
            "type": "enum", 
            "id": "string",
            "cicdProvider": "string", // for now only "azure" is supported
            "cicdProviderBaseUrl": "string",
            "azureOrganization": "string",
            "azureProject": "string",
            "terraformPipelineId": "string"
        }
    ],
    "logLevel": "string"
}
```

## Commands

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

### launch

flags are:
- --environment (-e): required
- --cr-list
- --parallel
- --error-code
- --debug
- --dry-run
