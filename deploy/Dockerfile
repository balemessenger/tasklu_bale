FROM golang:1.12

ARG docker_version=0

COPY ./build/release/taskulu_$docker_version /bin/taskulu

CMD ["/bin/taskulu"]