# Supported Tags:

This are the current tags supported for Qurl.

For each tag, there is also multiple image versions, following the [releases](https://github.com/repejota/qurl/releases) available at Github.

# latest

The *latest* tag is the most updated docker image for the project.
This image is built from the official 
[docker golang](https://hub.docker.com/_/golang/) image, and it is pushed on
each new release of *Qurl*.

Check the [Dockerfile](https://github.com/repejota/qurl/blob/master/Dockerfile).

# scratch

The *scratch* ttag is the small size image for the project and it is suited to
be used in production. This image is only about 3Mb of size.
This image is built from the official 
[docker scratch](https://hub.docker.com/_/scratch/) image, and it is also
pushed on each new release of *Qurl*.

Check the [Dockerfile](https://github.com/repejota/qurl/blob/master/Dockerfile.scratch  ).