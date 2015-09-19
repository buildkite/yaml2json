# yaml2json

![image](http://www.weivo.com/meat_grinder.files/meat_grinder.jpg)

A command line tool to easily convert YAML to JSON

### Usage

Pass it a YAML file and it'll output JSON

```
$ yaml2json examples/basic.yml
{"age":30,"name":"John"}
```

It will return an exit code if it fails to succesfully parse the YAML.

### Building

Running the build script will create 2 files in the `dist` folder. One for OSX and one for Linux.

```
$ ./scripts/build
Building yaml2json 💨

Compiling for OSX
👍  dist/yaml2json-darwin-amd64

Compiling for Linux
👍  dist/yaml2json-linux-amd64

All done! ✅
```

### Development

```
git clone git@github.com:buildkite/yaml2json.git
cd yaml2json
direnv allow
go run main.go examples/basic.yml
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Copyright

Copyright (c) 2015 Keith Pitt, Buildkite Pty Ltd. See LICENSE for details.
