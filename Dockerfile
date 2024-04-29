FROM golang:latest

RUN apt update

RUN apt install jq -y
RUN apt install grepcidr -y

RUN go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
RUN go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest
RUN go install github.com/projectdiscovery/katana/cmd/katana@latest
RUN go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

WORKDIR "/mount"

ENTRYPOINT ["bash", "-c"]