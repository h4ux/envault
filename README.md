```
                                    __   __   
    ____   _______  _______   __ __|  |_/  |_ 
  _/ __ \ /    \  \/ /\__  \ |  |  \  |\   __\
  \  ___/|   |  \   /  / __ \|  |  /  |_|  |  
    \___  >___|  /\_/  (____  /____/|____/__|  
           \/     \/           \/              
 
``` 


![Release](https://github.com/h4ux/envault/actions/workflows/release.yml/badge.svg)
[![Linux](https://svgshare.com/i/Zhy.svg)](https://svgshare.com/i/Zhy.svg)
[![macOS](https://svgshare.com/i/ZjP.svg)](https://svgshare.com/i/ZjP.svg)

![envault](https://user-images.githubusercontent.com/77572830/187724006-91ccf9f7-52b3-4f6d-8a91-015bf7d6e5dd.svg)

envualt is a cli tool for injecting env variables from hashicorp vault kv storage securly

Ex:

```
  -configure
        create configuration for the vault to be used
  -d    envault debug (verbose)
  -list string {json, table, yaml}
        List key value secrets
  -run
        set env variables and run command after double dash Ex. envault run -- npm run dev
  -set string
        Set key value secret Ex: envault -set=KEY=VALUE
  -status
        Get vault status
  -v    envault version
```

## Installation via install.sh

```bash
# binary will be in $(go env GOPATH)/bin/envault
curl -sSfL https://raw.githubusercontent.com/h4ux/envault/main/install.sh | sh -s -- -b $(go env GOPATH)/bin

# defualt installation into ./bin/
curl -sSfL https://raw.githubusercontent.com/h4ux/envault/main/install.sh | sh -s

```

Once you run envault -configure envault will create .env file in ~/.config/envault/.env

### Help

** Currently supports Mac OS and Linux OS (Windows can be added with very little effort)
