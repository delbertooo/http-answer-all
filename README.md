# http-answer-all

This is a tiny web-server answering any request with a status `200 OK` and a given response content (empty by default).

The content and content-type can be modified on a global basis.

## Usage

```sh
./http-answer-all
```

... and open your browser on [http://localhost:8090/](http://localhost:8090/). 
Or navigate to

* [http://localhost:8090/foo/bar.json](http://localhost:8090/foo/bar.json) or
* [http://localhost:8090/wp-admin/foo.php?user=bar](http://localhost:8090/wp-admin/foo.php?user=bar) or
* any other URL - it doesn't matter :wink:.


### All available options

You can alter the configuration using environment variables:

```sh
APP_RESPONSE='console.log("hi");' \
  APP_CONTENT_TYPE='text/javascript' \
  APP_LISTEN_ADDR='127.0.0.1:9090' \
  ./http-answer-all
```
