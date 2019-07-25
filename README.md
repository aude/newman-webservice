newman-webservice
=================

Simply a tool to use [Newman](https://github.com/postmanlabs/newman) as a web service.

Examples
--------

### Server

	$ docker run --rm -p 8080:8080 aude/newman-webservice

### Client

	$ curl -v "http://localhost:8080/collections/$POSTMAN_COLLECTION_UID?environment=$POSTMAN_ENVIRONMENT_UID&apikey=$POSTMAN_API_KEY"

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
