# DerpBearAPIGo

playing with golang and making a service to add some functionality to derp-bear

### Build Status

[ ![Codeship Status for orieken/derp_bear_api_go](https://codeship.com/projects/7f537d00-cd10-0132-f567-767651b3a193/status?branch=master)](https://codeship.com/projects/76218)


# Deps

A Yeoman generator for Go: [Generator-go](https://www.npmjs.com/package/generator-go)
BDD-style Golang testing framework: [Ginkgo](http://onsi.github.io/ginkgo/)

## useful posts

* [Deploying a Go app to Heroku](https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html)
* [Go by Example](https://gobyexample.com/)
* [Go Koans](https://github.com/cdarwin/go-koans)


#Notes for running on Codeship


####notes for me
Set the root to my brew go dir

```
export GOROOT=/usr/local/Cellar/go/1.4.2/libexec
```

Set the path to my local bin

```
export GOPATH=$HOME/Projects/Learning/Golang/derp_bear_api_go/
export PATH=$PATH:$GOPATH/bin
```