# Docker good practices
Examples of docker good practices with examples starting from a simple 
container build to a production ready build.

Links
* https://docs.openshift.com/container-platform/4.9/openshift_images/create-images.html
* https://sysdig.com/blog/dockerfile-best-practices/
* https://snyk.io/blog/10-docker-image-security-best-practices/

There are examples for
* golang
* nodejs
* python

We demonstrate different good practices for building containers by 
building a container for a simple demo application. To begin with 
the build is as simple as it can be and we add good practices one 
by one. The goal is to end up with a production ready container.

Obligatory statement that good practices and security is not a destination,
it is a continual process. Whenever we discover or learn of a new good 
practice we should incorporate it.

Good practices that will be implemented.
* Do not run as root
* Use a minimal init system such as dumb-init
* Remove temoprary files to avoid wasting space
* Use multi-stage builds to separate build and runtime environments
* Make code and executables owned by root

Good practices that are only used when necessary
* Mark used port by using EXPOSE
* Use a base image from a trusted source
* Run several related commands in the same RUN
* Use exec in entrypoint script (if one is used)
* Set defaults for enironment variables used by the container
* Write temporary data to /tmp
* Do not use ADD keyword unless absolutely necessary
* Use .dockerignore file to avoid copying unnecessary files

