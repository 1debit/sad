# Search And Do Something

Use one of the pre-compiled binaries in the [releases](https://github.com/yammine/sad/releases)
or build it yourself:

```bash
go get github.com/yammine/sad
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
