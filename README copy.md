```
            _____           _____           _   _   
           / ____|         |_   _|         (_) | |  
          | |  __    ___     | |    _ __    _  | |_ 
          | | |_ |  / _ \    | |   | '_ \  | | | __|
          | |__| | | (_) |  _| |_  | | | | | | | |_ 
           \_____|  \___/  |_____| |_| |_| |_|  \__|
 
``` 


![Release](https://github.com/h4ux/goinit/actions/workflows/release.yml/badge.svg)
[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)
[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)

![image](https://user-images.githubusercontent.com/77572830/154341792-8fab3f7d-bb0b-4f49-b6c5-25dc706fc167.png)

This is a very simple Go Lang project structure creator and github repo.

I suggest the following directory structure for conveniance

Create a path in your working station containg your Github project structure

Ex:

```
~/Dev/github.com/h4ux/{project name / repository name}
```

## Installation via install.sh

```bash
# binary will be in $(go env GOPATH)/bin/goinit
curl -sSfL https://raw.githubusercontent.com/h4ux/goinit/main/install.sh | sh -s -- -b $(go env GOPATH)/bin

# defualt installation into ./bin/
curl -sSfL https://raw.githubusercontent.com/h4ux/goinit/main/install.sh | sh -s

```

Once you install a file called .env will be created in ~/.config/goinit/.env
Please add the relevant data as described below

You can overide the global .env file by creating a local one on the working directory
### Help

![image](https://user-images.githubusercontent.com/77572830/154342133-1cd8cff5-a8b4-421e-a058-c4675ab16591.png)


### .env file format
```
GH_TOKEN={GITHUB Personal token goes here}
GH_ORG={GITHUB USER NAME or ORG}
GO_FOLDERS=bin,cmd,configs,deployments,docs,internal,pkg,tests #folders that you whould like to create separated by comma
GO_PROJECTS_PATH={your full path for your Go projects}
```

** Currently supports Mac OS and Linux OS (Windows can be added with very little effort)
