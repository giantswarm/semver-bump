# Semantic Versioning Bumper

A simple semantic versioning file bumper to keep your version files in line
built by [Giant Swarm](http://giantswarm.io).

## Installing

You can install semver-bump from source using `go get`:

    go get github.com/giantswarm/semver-bump

Because go expects all of your libraries to be found in either `$GOROOT` or
 `$GOPATH`, it's helpful to symlink the project to one of the following paths:

 * `ln -s /path/to/your/semver-bump $GOPATH/src/github.com/giantswarm/semver-bump`
 * `ln -s /path/to/your/semver-bump $GOROOT/src/pkg/github.com/giantswarm/semver-bump`

## Running

With semver-bump you can bump your projects version into a `VERSION` file location
in your project. It supports bumping of major, minor and patch versions via the
following subcommands:

 * `semver-bump major-release`
 * `semver-bump minor-release`
 * `semver-bump patch-release`
