# init
Create a new project with templates - a cli tool to create a new project with templates from [csd1100/templates](https://github.com/csd1100/templates).

## Capabilities
- Create a new project with name from available templates
- Initialize git repository for project<sup>*
- Automatically install dependencies from templates using dependency manager like npm / cargo<sup>*

`*` - Can be disabled via flag.
## Requirements
- `git` installed.
- The programming language installed depending on project template used.
- Package / Dependency managers like npm installed depending on project template used.

## Usage
`init -t <template_name> -n <project_name> [OPTIONS]`

- **Options**
    - `-t|--template`: templates from repo. Currently available: go, js
    - `-n|--name`: name of the project
    - `-G|--no-git`: Do not initialize git repository
    - `-S|--no-sync`: Do not run commands like `npm install`.
