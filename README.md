# Hashmod

This program will take a hostname, IP address, or other unique host
identifier, to distribute an enabled / disabled flag across a cluster
without coordination, using consistent hashing.

## Use Case

The motiviation behind this project was to provide a stateless way to
enable New Relic on a portion of the nodes in a cluster of web servers
in order to reduce costs, but still have metrics on some of the
servers. It was also an excuse to study a topic of interest.

By installing this binary in a Docker container and setting an
environment variable based on the output in a container startup
script, I was able to disable the New Relic agent on a portion of the
servers, without requiring a centralized service.

## Examples

```
# Use hostname (default) (use hostname, divide cluster into 2 buckets)
$ ./hashmod
false

# Use hostname, with options (use hostname, divide cluster into 4 buckets)
$ ./hashmod -input hostname -buckets 4
false

# Use input from stdin (test behavior)
echo -e "abc.example.com\ndef.example.com" | ./hashmod_export -input stdin -buckets 2 -hashalg xxhash64
2018/12/09 23:45:47 index: 0
2018/12/09 23:45:47 line: abc.example.com
2018/12/09 23:45:47 enabled: false
2018/12/09 23:45:47 ----------
2018/12/09 23:45:47 index: 1
2018/12/09 23:45:47 line: def.example.com
2018/12/09 23:45:47 enabled: true
2018/12/09 23:45:47 ----------
2018/12/09 23:45:47 buckets: 2
2018/12/09 23:45:47 enabledCount: 1
2018/12/09 23:45:47 totalCount: 2
2018/12/09 23:45:47 percent: 50.00%
```
