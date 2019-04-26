FROM golang:1.12.4

LABEL AUTHOR="Suhail Abood" \
  DESCRIPTION="Go builder for business service" \
  USAGE="docker build -t suhdev/nutrient-business ."

COPY nutrient-business /nutrient-business

CMD ["/nutrient-business", "--server_address=0.0.0.0:8000", "--broker_address=0.0.0.0:10000"]