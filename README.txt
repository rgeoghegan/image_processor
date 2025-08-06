Image Processor
===============

Welcome to the image processor project!

How to build
------------

The easiest way to build is:

1. Make sure you have the [vips library installed](https://www.libvips.org/install.html) (for example, do `brew install vips` on mac os x if using homebrew)
1. Run the following command in the project to build `bin/image_processor`:

        make build

1. The run the binary to have it running on `localhost:8000`:

        ./bin/image_processor


Tests
-----

To run the tests, simply do:

    make test

Using the processor
-------------------

The image_processor comes with three endpoints for your convenience:

* `/convert`: convert a png to a jpeg

    Simply post the content of the png to the endpoint.

* `/resize?width=600&height=480`: Resize an image to the given size

    Post the content of the image to the endpoint, and use the url params to set the desired size.

* `/compress?level=10`: Compress an image to the given level

    Post the content of the image to the endpoint, and use the url params to set the desired compression level.
