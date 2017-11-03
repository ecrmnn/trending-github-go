# trending-github-go
> Simple API for getting trending repositories on GitHub

[![Travis](https://img.shields.io/travis/ecrmnn/trending-github-go/master.svg?style=flat-square)](https://travis-ci.org/ecrmnn/trending-github-go/builds)

### Installation
```bash
go get github.com/ecrmnn/trending-github-go
```

### Usage
```
import trending "github.com/ecrmnn/trending-github-go"
```

```go
trending.All("daily")
// Returns []Repository for all languages today

trending.Language("Go", "weekly")
// Returns []Repository for a specific language this week
```

```go
type Repository struct {
	author      string
	name        string
	href        string
	description string
	language    string
	stars       int
	forks       int
}
```

### License
MIT © [Daniel Eckermann](http://danieleckermann.com)
