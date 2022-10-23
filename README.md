
# Gopher Bot

This is a Discord bot that share some cool gopher images. It read massages from Discord channel and share images on commant "!gopher" and "!random", and shares all available gophers on "!gophers" comment


## Demo

![Screenshot from 2022-10-23 19-45-46](https://user-images.githubusercontent.com/101032025/197399116-b62d3ba6-4811-46e8-ad45-5701690c42a6.png)

![Screenshot from 2022-10-23 19-45-57](https://user-images.githubusercontent.com/101032025/197399158-feeb2e92-bf97-4107-bef5-f0a7b88bcdb2.png)

![Screenshot from 2022-10-23 19-39-02](https://user-images.githubusercontent.com/101032025/197399175-c04e672a-d73e-4285-bd82-349c1b64355f.png)


## Getting started

In this project we need to call KuteGoAPI(https://github.com/gaelleacas/kutego-api) to get those images. So first we need to run the KuteGoAPI on our local machine, to run the api please go through the KuteGoAPI project. Then create a new application on your Discord channel and add a bot, copy the bot token and paste that into the config.json file in the token section.  

### using docker - 

To run the bot on docker you need change the Kute_go_APIURL from "http://localhost:8080" to Kute_go_APIURL = "http://host.docker.internal:8080" (in bot/bot.go file) and then build the docker image and run it.

```bash
  docker build -t gopher-discord:1.0 . 
```

```bash
  docker run --rm gopher-discord:1.0  
```
### run locally - 

Enter the gopher-discord directory and run -  

```bash
  go build 
```
```bash
  go run main.go
```
