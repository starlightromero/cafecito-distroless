# Distroless Golang IP Geolocator

This application serves to showcase small Docker images via a simple API that has one route. A GET request to the / route will make an API call to in order to get the user's IP address, ISP, and location data, including the user's latitude and longitude to the millionth decimal place. All this data will be returned.

If you want to download the code and try it out yourself, you will need to sign up for a free API key from:

```
https://geo.ipify.org/
```

## Needed Software

- docker
- docker-compose
- make

## How to Run

Clone the repo
```zsh
git clone git@github.com:starlightromero/distroless-go-ip.git
```

cd into the directory
```zsh
cd distroless-go-ip
```

Rename `.env.sample` to `.env`
```zsh
mv .env.sample .env
```

Edit the `.env` file to contain your API Key
```zsh
vim .env
```

Run the application!


## Makefile Commands

`build-alpine-single`: Build Geolocator app with an Alpine base image using a single stage build

`start-alpine-single`: Start Geolocator app with an Alpine base image using a single stage build

`stop-alpine-single`: Stop Geolocator app with an Alpine base image using a single stage build

`build-alpine`: Build Geolocator app with an Alpine base image using a multi-stage build

`start-alpine`: Start Geolocator app with an Alpine base image using a multi-stage build

`stop-alpine`: Stop Geolocator app with an Alpine base image using a multi-stage build

`build-buster-single`: Build Geolocator app with an Buster base image using a single stage build

`start-buster-single`: Start Geolocator app with an Buster base image using a single stage build

`stop-buster-single`: Stop Geolocator app with an Buster base image using a single stage build

`build-buster`: Build Geolocator app with an Buster base image using a multi-stage build

`start-buster`: Start Geolocator app with an Buster base image using a multi-stage build

`stop-buster`: Stop Geolocator app with an Buster base image using a multi-stage build

`build-distroless`: Build Geolocator app with an Distroless base image using a multi-stage build

`start-distroless`: Start Geolocator app with an Distroless base image using a multi-stage build

`stop-distroless`: Stop Geolocator app with an Distroless base image using a multi-stage build

`build-scratch`: Build Geolocator app with an Scratch base image using a multi-stage build

`start-scratch`: Start Geolocator app with an Scratch base image using a multi-stage build

`stop-scratch`: Stop Geolocator app with an Scratch base image using a multi-stage build

`build-all`: Build all images