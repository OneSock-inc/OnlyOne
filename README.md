[![backend](https://github.com/OneSock-inc/OnlyOne/actions/workflows/backend.yml/badge.svg?branch=main)](https://github.com/OneSock-inc/OnlyOne/actions/workflows/backend.yml)
[![frontend](https://github.com/OneSock-inc/OnlyOne/actions/workflows/frontend.yml/badge.svg)](https://github.com/OneSock-inc/OnlyOne/actions/workflows/frontend.yml)


![image of the app running on an iphone13](./screenshots/onlyone.png)

# About OnlyOne   
OnlyOne wishes to solve the problem of the "lonely sock who is alone after having washed your clothes".
The app is a webapp (SPA) with angular for the frontend and golang for the REST API (backend). Therefore once running, it can be used by any browser. We wanted to do a mobile-first app but also support desktop that's the reason we chose [Angular](https://angular.io/).

## How does it work ?
OnlyOne lets you create a profile, adds your socks and matches to win other socks ! It's that simple !
If you accept another sock and the owner of the other sock accepts yours, then you have a match (denoted by the little heart in the UI). Having a match means that you have 50% of chance to get his sock and 50% chance of lossing yours.

If you loose, you'll have instruction on how to send your sock to the winner.

Don't play if you are not willing to loose your sock.

## OnlyOne Documentation
The setup instruction and documentation are split in two : the [frontend instruction](./frontend/README.md) and the [backend instruction](./backend/README.md).

The thorough documentation can be downloaded from the teams channel 

## The production environment
The workflow to deploy is as follows :
- A PR is merged in main
- Github notifies the [webhook server](https://github.com/OneSock-inc/Infra) of the merge
- The webhook downloads this repo and run a docker-compose up command
- The services are updated with the latest version of the code

You can find more about the watcher in the infrastructure repository of our organisation.
