# `init` - create a new project with templates

- A cli tool to create a new project with templates from
  [csd1100/templates](https://github.com/csd1100/templates).

## Capabilities

- Create a new project with name from available templates
- Initialize git repository for project<sup>\*</sup>
- Automatically install dependencies from templates using dependency manager
  like npm / cargo<sup>\*</sup>

\* - Can be disabled via flag.

## Requirements

- `git` installed.
- The programming language installed depending on project template used.
- Package / Dependency managers like npm installed depending on project template
  used.

## Usage

`init -t <template_name> -n <project_name> [OPTIONS]`

- **Arguments:**
  - `-t|--template <value>` - Templates from repo. Required. Currently available:
    go, js.
  - `-n|--name <value>` - Name of the project. Required.
  - `-p|--path <value>` - Path to directory where project will be created.
  - `-o|--options <values>` - Special options for that template.
    Values are comma-separated `key1=value1,key2=value2`
  - `-h|--help` - Display help.
  - `-G|--no-git` - Do not initialize git repository
  - `-S|--no-sync` - Do not run commands like `npm install`.
  - `-v|--verbosity <0-5>` - Change the level of logging to STDOUT.  
    Where 0 is PANIC level and 5 is TRACE level.

## Template Specific Info

### go

- For go template if `packageName` option is passed using `-o packageName=github.com/csd1100`
  then this will create go module with name `packageName\project_name`.  
   i.e. if you use options `-t go -n my-app -o packageName=github.com/csd1100`
  the project module name will be `github.com/csd1100/my-app`
