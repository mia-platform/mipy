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
            "azureProject": "string"
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
- --forward-env (-f): Forward environment to pipeline as ENVIRONMENT_TO_DEPLOY variable
- --cr-list
- --parallel
- --error-code
- --debug
- --dry-run

#### Forward Environment Variable

When using the `--forward-env` (or `-f`) flag, the environment value specified with `--environment` will be passed as an environment variable named `ENVIRONMENT_TO_DEPLOY` to the pipeline execution. This is useful when your pipeline needs to know which environment it's deploying to.

Example:
```bash
mipy launch -e production -f -u myuser -p mypassword
```

This will set `ENVIRONMENT_TO_DEPLOY=production` as a variable in the pipeline run.
