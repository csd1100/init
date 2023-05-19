# `init` - create a new project with templates
- A cli tool to create a new project with templates from [csd1100/templates](https://github.com/csd1100/templates).

## Capabilities
- Create a new project with name from available templates
- Initialize git repository for project<sup>*</sup>
- Automatically install dependencies from templates using dependency manager like npm / cargo<sup>*</sup>

\* - Can be disabled via flag.
## Requirements
- `git` installed.
- The programming language installed depending on project template used.
- Package / Dependency managers like npm installed depending on project template used.

## Usage
`init -t <template_name> -n <project_name> [OPTIONS]`

- **Arguments:**
    - `-t|--template <value>` -  Templates from repo. Required. Currently available: go, js.
    - `-n|--name <value>` -      Name of the project. Required.
    - `-p|--path <value>` -      Path to directory where project will be created.
    - `-h|--help` -              Display help.
    - `-G|--no-git` -            Do not initialize git repository
    - `-S|--no-sync` -           Do not run commands like `npm install`.
