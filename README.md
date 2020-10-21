# Medium RSS API
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-3-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

A REST API wrapper for [Medium RSS Feed](https://help.medium.com/hc/en-us/articles/214874118-RSS-feeds) with built in cache mechanism and [HTML Tokenizer](https://godoc.org/golang.org/x/net/html) that parses Medium's plain HTML string into DOM objects. Just set your medium's user profile name or publication and you're good to go!  

Built and tested using Golang v1.14.4.

## Configuration

There are some configurations which takes place in environment variable that needed to be setup first before you can start using the API which are: 
| Name | Mandatory | Default | Description |
| --- | --- | --- | --- |
| MEDIUM_PROFILE | Y | - | User or publication name |
| MEDIUM_RSSFEED_URL | N | `https://medium.com/feed` | Medium's feed base url |
| BASE_URL | N | localhost | Host of the REST API server | 
| PORT | N | 8080 | Port of the REST API server |
| GO_ENV | N | development | Environment stage |

## Installation & Usage
To build, simply clone this repository and run 
```
make build
```
You can run `watcher.sh` to run the server and re-build automatically everytime the go file changes.

## Installation 
To install simply pull the docker image.

## Usage

Currently, available endpoints are :  

`/version` - to check the application version  
`/health` - to check the application health  
`/medium` - to get the medium's RSS feed  

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/wahyudibo"><img src="https://avatars2.githubusercontent.com/u/4588408?v=4" width="100px;" alt=""/><br /><sub><b>Wahyudi Wibowo</b></sub></a><br /><a href="https://github.com/ByteSchneiderei/medium-rss-api/commits?author=wahyudibo" title="Code">💻</a></td>
    <td align="center"><a href="https://twitter.com/t_strohmeier"><img src="https://avatars3.githubusercontent.com/u/13830953?v=4" width="100px;" alt=""/><br /><sub><b>Thomas Strohmeier</b></sub></a><br /><a href="#platform-tstrohmeier" title="Packaging/porting to new platform">📦</a></td>
    <td align="center"><a href="https://github.com/SebiSpace"><img src="https://avatars2.githubusercontent.com/u/10374656?v=4" width="100px;" alt=""/><br /><sub><b>Sebastian Schaffer</b></sub></a><br /><a href="#platform-SebiSpace" title="Packaging/porting to new platform">📦</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

## License
[MIT License](https://github.com/ByteSchneiderei/medium-rss-api/blob/master/LICENSE). Copyright (c) 2020 ByteSchneiderei GmbH
