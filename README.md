# fnew
[![Build Status](https://travis-ci.com/ncipollo/fnew.svg?branch=master)](https://travis-ci.com/ncipollo/fnew)

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
* Projects define in your local configuration manifest (these will always take precedence over other projects).
* Projects found in the manifest repository defined in your configuration file.
* Projects found in the default manifest.
