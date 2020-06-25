# kickstart-generator
Kickstart File Generator in Go

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

```
curl 'localhost:8080/ks_generate/?os=centos7&version=7.3.1611&fqdn=myhost.example.com'
```
or just go to the browser and type the same url

### Add your own parameters

Modify the ks.tmpl file by replacing the value you want to parameterized with {{.your_variable}} and then just use it in the url 
```
curl 'localhost:8080/ks_generate/?os=centos7&version=7.3.1611&fqdn=myhost.example.com&your_variable=foo'
```
