[![Monocle Service Score Badge](https://monocle.prod.chmfin.com/badge/sad/service?s=JDJhJDEyJFpyZ3U4bGJtQVRGUXlhQmRWUmNFWWVWbF)](https://monocle.prod.chmfin.com/1debit/sad)

# Search And Do Something

Use one of the pre-compiled binaries in the [releases](https://github.com/1debit/sad/releases)
or build it yourself:

```bash
go get github.com/1debit/sad
```

### Examples
```
sad --token=<your personal gh access token> --org=github --language | jq
```

Would look like:
```json
{
  "language": null,
  "url": "git@github.com:github/media.git"
}
{
  "language": "Ruby",
  "url": "git@github.com:github/albino.git"
}
{
  "language": "Ruby",
  "url": "git@github.com:github/hubahuba.git"
}
{
  "language": "JavaScript",
  "url": "git@github.com:github/jquery-hotkeys.git"
}
{
  "language": "JavaScript",
  "url": "git@github.com:github/jquery-relatize_date.git"
}
...
```
### Macos
Install:
```
brew install go
git clone git@github.com:yammine/sad.git
cd sad
go run main.go --token=your_github_personal_access_token(not_oauth) --org=1debit --language --forks
```

If you want to use the script outside the cloned folder, make a build and use the binary instead of "go run main.go":
```
try:
go build
./sad or /script_location/sad/sad
```

Compare the number of repos run:
```
go run main.go --token=your_github_personal_access_token(not_oauth) --org=1debit --language --forks | nl
```
Clone all repos:
```
for url in $(go run main.go --token=your_github_personal_access_token(not_oauth) --org=1debit --language --forks | jq '.url' | tr -d '"'); do git clone $url; done
```
If that does not work, try this:
```
go run main.go --token=your_github_personal_access_token(not_oauth) --org=1debit --language --forks > sad.out
cd ...
cat .../sad.out | jq '.url' | tr -d '"' | (while read url; do git clone $url; done)
```
Get the language stack (list of unique languages):
```
go run main.go --token=your_github_personal_access_token(not_oauth) --org=1debit --language --forks | jq '.language' | tr -d '"' | sort | uniq
```
