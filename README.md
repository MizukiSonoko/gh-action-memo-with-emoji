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


    - uses: 'mizukisonoko/gh-action-with-emoji@master'
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

### Third When emoji is IMP
  
Created an issue   
![issue](https://raw.githubusercontent.com/MizukiSonoko/gh-action-with-emoji/master/images/img.png)

# Current supported emoji
- :imp: `:imp:`  
Make an issue about this commit  
  
- :no_entry: `:no_entry:`  
Cancel job  
  
# Will supported emoji
- :tada: `:tada:`
- :fire: `:fire:`
- :books: `:books:`
- :bug: `:bug:`
- :mega: `:mega:`

# License

MIT