# `init` - create a new project with templates

- A cli tool to create a new project with templates from
  [csd1100/templates](https://github.com/csd1100/templates).

## Capabilities

- Create a new project from available templates
- Initialize git repository for project<sup>\*</sup>
- Automatically install dependencies from templates using dependency manager
  like pnpm / cargo<sup>\*</sup>

\* - Can be disabled via flag.

## Requirements

- `git` installed.
- The programming language installed depending on project template used.
- Package / Dependency managers like pnpm installed depending on project template
  used.

## Usage

`init -t <template_name> -n <project_name> [OPTIONS]`

- **Arguments:**
    - `-t|--template <value>` - Templates from repo. Required. Currently available:
      `go`, `js`, `rust`, `ts`, `electron-react-ts`.
    - `-c|--current` - create project in current directory. Existing files won't
      be replaced and if git repo is already initialized won't be done again.
      project name will be name of current directory if not specified using `-n`
    - `-n|--name <value>` - Name of the project. Required if `-c` is not used.
    - `-p|--path <value>` - Path to directory where project will be created.
    - `-o|--options <values>` - Special options for that template.
      Values are comma-separated `key1=value1,key2=value2`
    - `-h|--help` - Display help.
    - `-G|--no-git` - Do not initialize git repository
    - `-S|--no-sync` - Do not run commands like `pnpm install`.
    - `-v|--verbosity <0-5>` - Change the level of logging to STDOUT.  
      Where 0 is PANIC level and 5 is TRACE level.

## Template Specific Info

### go

- For go template if `packageName` option is passed using `-o packageName=github.com/csd1100`
  then this will create go module with name `packageName\project_name`.  
  i.e. if you use options `-t go -n my-app -o packageName=github.com/csd1100`
  the project module name will be `github.com/csd1100/my-app`

## For developers

- If Following environment variables are set local templates can be used:

```sh
export DEV=true
export INIT_DEV_REPO_PATH=`<local path to the templates repo>`
export INIT_DEV_BRANCH_NAME=`<name of the branch to pull>`
```

- Using only `INIT_DEV_BRANCH_NAME` will use branch from
  https://github.com/csd1100/templates repository.
