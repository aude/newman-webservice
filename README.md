newman-webservice
=================

Simply a tool to use [Newman](https://github.com/postmanlabs/newman) as a web service.

Examples
--------

### Server

	$ docker run --rm -p 8080:8080 aude/newman-webservice

### Client

	$ curl -v "http://localhost:8080/collections/b93e12d-2bde1187-2651-4f81-94dd-2f55e324be3d?environment=b93e12d-d640c4f2-4a6b-46cd-8a41-b3ab4c852d50&apikey=$POSTMAN_API_KEY"

Usage
-----

	usage: newman-webservice

Install
-------

	$ docker pull aude/newman-webservice

Uninstall
---------

	$ docker rmi aude/newman-webservice

Dev
---

See `Makefile`.

Configure
---------

This web service is configured using environment variables.

See [Usage](#usage).
