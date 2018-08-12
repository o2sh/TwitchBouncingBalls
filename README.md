# TwitchBouncingBalls

Project made during the Hackalong Hackathon. It uses the Twitch APIs to display the top 20 live streams for a specific category as bouncing balls 🎈🎈


<p align="center">
<img src="https://github.com/02sh/TwitchBouncingBalls/blob/master/pub/img/demo.gif">
</p>

# Try it

In main.go, make sure to feed your own Client-id. (check this [link](https://dev.twitch.tv/docs/v5/) for more info):

```shell
os.Setenv("GO_TWITCH_CLIENTID", "5hk0e3wz856a198lypq5bai57nf13u")
```

Then open bash:

```shell
go get "github.com/02sh/TwitchBouncingBalls/"
cd $GOPATH/src/github.com/02sh/TwitchBouncingBalls
go run *go
```
