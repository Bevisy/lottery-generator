FROM alpine:3.11
ADD lottery-generator /lottery-generator
ENTRYPOINT ["/lottery-generator"]
