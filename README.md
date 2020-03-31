# gh-action-with-emoji

Some actions when commit message's prefx is emoji

# How to use  
  
### First  
Add `mizukisonoko/gh-action-memo-with-emoji@master` to your workflows yaml  

```yml
name: my actions in own repository
on:
  push:
    branches:
      - master
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2


    - uses: 'mizukisonoko/gh-action-memo-with-emoji@master'
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

```

### Second  
  
commit & push  

```shell
$ git add code.go
$ git commit -m":imp: add temporary document"
$ git push origin master
```

# Current supported emoji
- :imp:

# Will supported emoji
- :tada:
- :fire:
- :books:
- :bug: