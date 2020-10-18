# Golang for H264 decode to JPG

Prepare .mp4 or .h264 file

```
cd example
```

## mp4 to .h264

```
ffmpeg -i test.mp4 -an -vcodec libx264 -crf 23 test.h264
```

## play

Transfer .h264 to *.jpg

```
cd example
go build
./example
```


