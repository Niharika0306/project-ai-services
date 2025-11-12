# Commands

This section lists all available commands and subcommands supported by AI-services

### Command: `ai-services version`

**Description:**  
Displays detailed version information 

**Usage**

```bash
ai-services version
ai-services [command] --version
ai-services [command] -v
```

```bash
% ai-services version
Version: 732d48a
GitCommit: 732d48a
BuildDate: 2025-11-04T07:33:42Z
% ai-services --version
ai-services version 732d48a
% ai-services -v
ai-services version 732d48a
```

### Command: `ai-services application`

**Description:**  
Manage application configuration.

**Subcommands:**

| Subcommand | Description |
|-------------|--------------|
| `ai-services application create` | Create application |
| `ai-services application templates` | List available application templates |
| `ai-services application model list --template <template_name>` | List models used by templates|
| `ai-services application model download --template <template_name>` | Download models used by templates|
| `ai-services application image list --template <template_name>` | Lists container images used by the template |
| `ai-services application image pull --template <template_name>` | Pulls container images used by the template |
| `ai-services application ps` | Similar to `podman ps` - lists containers deployed by AI-services |
| `aiservices application stop <application_name>` | Stops all pods in the application |
| `aiservices application stop <application_name> --pod-name <pod_name>` | Stops specific pod in the application |
| `aiservices application start <application_name>` | Starts all pods in the application |
| `aiservices application start <application_name> --pod-name <pod_name>` | Starts specific pod in the application |
| `ai-services application delete <application_name>` | Delete application |

### Command: `ai-services completion`

**Description:**  
Generates the autocompletion script for specified shell

**Subcommands:**

| Subcommand | Description |
|-------------|--------------|
| `ai-services completion bash` |  Generate the autocompletion script for bash |
| `ai-services completion fish` |  Generate the autocompletion script for fish |
| `ai-services completion powershell` |  Generate the autocompletion script for powershell |
| `ai-services completion zsh` |  Generate the autocompletion script for zsh |

### Command: `ai-services help`

**Description:**  
Displays a list of all available commands and their brief descriptions.
The help command can also be used to view detailed usage information for a specific command or subcommand.

**Usage**: 

```bash
ai-services [command] help
ai-services [command] --help
ai-services help
```
