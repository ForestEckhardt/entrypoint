# Entrypoint Cloud Native Buildpack

This buildpack is meant to be used at the end of the buildpack order definition and will write a start command that is generated from the contents of an `entrypoint.toml` in the base of the application directory. The `entrypoint.toml` is meant to reflect the contents of [`launch.toml`](https://github.com/buildpacks/spec/blob/master/buildpack.md#launchtoml-toml) that is currently supported by `packit`.

## Usage

To package this buildpack for consumption:

```
$ ./scripts/package.sh
```

This builds the buildpack's Go source using `GOOS=linux` by default. You can
supply another value as the first argument to `package.sh`.

## `entrypoint.toml`

```toml
[[processes]]
type = "<process type>"
command = "<command>"
args = ["<arguments>"]
direct = false
```

If you are looking for concrete definitions on what these fields do inside of `packit` you can check the documentation [here](https://godoc.org/github.com/paketo-buildpacks/packit#Process). For the definition from the Cloud Native Buildpack specification itself you can check out the documentation [here](https://github.com/buildpacks/spec/blob/master/buildpack.md#launchtoml-toml).
