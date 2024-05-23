FROM golang:latest

RUN apt update

RUN apt install grepcidr jq git libpcap-dev moreutils -y

RUN go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
RUN go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest
RUN go install -v github.com/projectdiscovery/katana/cmd/katana@latest
RUN go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
RUN go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest

RUN go install github.com/tomnomnom/gf@latest
WORKDIR "/root"
RUN mkdir .gf
RUN git clone https://github.com/1ndianl33t/Gf-Patterns
RUN mv ~/Gf-Patterns/*.json ~/.gf

RUN mkdir "/config"

WORKDIR "/mount"

ENTRYPOINT ["bash", "-c"]