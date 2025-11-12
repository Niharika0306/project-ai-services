# Tutorial

// TO-DO - make the tutorial use case specific

This tutorial guides you through listing available templates, creating an application from a template, monitoring its status, and cleaning up when you're done. It assumes your environment is already configured and validated.

Tip: If you have not yet set up your environment, complete the steps in `Installation` first.

### List all available templates

```bash
ai-services application templates
```

### Create an application

Choose a template and a unique name for your application.

```bash
ai-services application create <app_name> --template-name=<template_name>
```

Deployment can take a few minutes. You can relax while the resources are provisioned.

```bash
the successful deployment should have some sort of message displayed
```

Once deployment completes, you can monitor pod status and readiness with the `ai-services application ps` command.

### Cleanup

When you're finished, delete the application to free resources.

```bash
ai-services application stop <app_name>
ai-services application delete <app_name>  
```
