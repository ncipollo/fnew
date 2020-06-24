# fnew
fnew allows you to quickly fetch a project, then apply a series of transforms on it. After the transforms complete the project should be setup and ready to be imported into your favorite IDE!

# Usage
```
Usage: fnew [OPTIONS] <source project> <project location>
  -h  prints fnew usage (shorthand)
  -help  prints fnew usage
  -l	lists the available fnew projects (shorthand)
  -list
    	lists the available fnew projects
  -u	updates the fnew project list (shorthand)
  -update
    	updates the fnew project list
  -v	enables verbose output (shorthand)
  -verbose
    	enables verbose output
  -version
    	prints fnew version
```

# Project Manifest
fnew keeps track of projects to fetch in manifest files (or repositories). The defauly manifest repository may be found here: 
[fnew-manifest](https://github.com/file-new/fnew-manifest). Any project found in this repository's manifest may be used without any additional configuration of fnew.

You may define your own projects for fnew to fetch & setup by adding your own manifest. To add your own manifest open or create the fnew configuration file (this will be created automatically if you run `fnew --list`). The configuration file is located at `~/.fnew/config.json`.

The format of the configuration file is as follows:
```json
{
    "repo": "https://url.to.manifest.repo.org",
    "manifest": {
        "projectName1": "https://url.to.project.org",
        "projectName2": "https://url.to.project.org"
    }
}
```
Here's what each key represents:
* repo: (Optional) git url to a manifest repository. The format of the manifest file is the same as define above in the configuration under the "manifest" key. Check out [fnew-manifest](https://github.com/file-new/fnew-manifest) for an example.
* manifest: (Optional) local manifest. The left hand key represents the project name (what you will pass to fnew) and the right hand value points to the git url of the project.

If a project is defined in multiple places the following order of precedence takes place:
* Projects found in your local configuration manifest (these will always take precedence over other projects).
* Projects found in the manifest repository defined in your configuration file.
* Projects found in the default manifest.

# Project Transforms
A project may be configured to apply a series of transforms after it is fetched. The transforms are define in a file named `.fnew ` and is located in the root of the project repository. And example may be found here: [.fnew file](https://github.com/file-new/fnew-test-project/blob/main/.fnew)

The project configuration file has the following format:
```json
{
    "transforms": [
        {
            "transform_argument": "value",
            "type": "type_of_transform"
        },
        {
           "transform_argument": "value",
            "type": "type_of_transform"
        }
    ]
}
```

Transform properties can vary between different transform types. You may use variables for many property values by prefixing the property name with `$` (ex- `$FOO`). Properties will be discussed in detail below.

The next sections will cover the types of transforms supported by fnew.

## File Move Transform
**Description**: This transform will move the file or folder found at the specified input path to the specified output path.

```json
{
    "input_path": "path/to/input",
    "output_path": "path/to/output",
    "type": "file_move"
}
```
Keys:
* **input_path**: The path of the file or folder to move. This path is relative to the project folder. This property supports variables (ex: `$package_name`)
* **output_path**: The destination of the move. This path is relative to the project folder. This property supports variables (ex: $`PACKAGE_PATH`)

## File String Replace Transform
**Description**: This transform will apply a string find and replace to the file(s) specified at the input path. 

```json
{
    "input_path": "path/to/files/*.java",
    "string_replace": {
        "old": "fnew.fnew",
        "new": "$package_name"
    },
    "type": "file_string_replace"
}
```
Keys:
* **input_path**: The path to the file(s) to apply the string replace on. This path is relative to the project and supports wildcards. This variable does *not* currently support variables.
* **string_replace**: Defines the string replace operation. The `old` key defines the string to find and the `new` key defines the string to replace it with. Both `new` and `old` may be set to variables (ex: `package_name`).

## Input Transform
**Description**: This transform allows the user to enter a value for a variable. It will prompt the user to input this during the transformation phase of project setup. The user will be accessible in other transforms via $<variable_name>

```json
{
    "output_variable": "variable_name",
    "skip_if_variable_exists": true,
    "type": "input"
}
```
Keys:
* **output_variable**: The variable for the user to set.
* **skip_if_variable_exists**: When true this transform will be skipped if the variable already exists.

## Run Script Transform
**Description**: This transform will run the shell script found at the specified path. **Note**: This transform currently only supports unix (linux & macOS) scripts. It will not work with Windows scripts.

Note: All variables will be passed into the script via environment variables.

```json
{
    "arguments" : ["foo","bar"],
    "input_path": "path/to/script.sh",
    "type": "run_script"
}
```
Keys:
* **arguments**: Optional arguments to pass into the script.
* **input_path**: The path to the script to run. This path is relative to the project and supports wildcards. This variable does supports variables (ex- `$script_path`).

## Variable Replace Transform
**Description**: This transform will apply a string find and replace to the file(s) specified at the input path. 

```json
{
    "input_variable": "input_variable",
    "output_variable": "output_variable",
    "skip_if_variable_exists": true,
    "string_prefix": "optional.",
    "string_replace": {
        "old": ".",
        "new": "/"
    },
    "string_suffix": ".optional.txt",
    "type": "variable_string_replace"
}
```
Keys:
* **input_variable**: The variable to transform.
* **output_variable**: The variable name where the transform will place the value after replacing the string(s).
* **skip_if_variable_exists**: When true this transform will be skipped if the variable already exists.
* **string_prefix**: An optional prefix to prepend to the variable.
* **string_replace**: Defines the string replace operation. The `old` key defines the string to find and the `new` key defines the string to replace it with. Both `new` and `old` may be set to variables (ex: `package_name`).
* **string_suffix**: An optional suffix to append to the variable.

# Transform Variables
As indicated in the above sections, fnew maintains a list of variables for use within transforms. There are few ways that variables are set in fnew:
* Provided by the user via the input transform (described in the transform section above)
* `$PROJECT_NAME` is set to the project name specified by the user when invoking fnew.
* Environment variables are copied into the internal variables when fnew is run.
