FROM golang:1.12.4

LABEL AUTHOR="Suhail Abood" \
  DESCRIPTION="Go builder for datastore service" \
  USAGE="docker build -t suhdev/nutrient-datastore ."

COPY ./nutrient-datastore nutrient-datastore
ENTRYPOINT ["./nutrient-datastore", "--server_address=0.0.0.0:8002", "--broker_address=0.0.0.0:10002"]