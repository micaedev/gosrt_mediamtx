Implementation of the SRT protocol in pure go

- [SRT reference implementation](https://github.com/Haivision/srt)
- [SRT RFC](https://haivision.github.io/srt-rfc/draft-sharabayko-mops-srt.txt)
- [SRT Technical Overview](https://github.com/Haivision/srt/files/2489142/SRT_Protocol_TechnicalOverview_DRAFT_2018-10-17.pdf)

# TODO

Everything

- nicer API <- done
- less CPU <- the ticker per connection uses up a lot of CPU because it is in a tight loop

# Try it out

Do this only on your own risk. You should have:

## Setup

- macOS (something recent)
- Homebrew
- go

Install ffmpeg and libsrt:

```
brew install ffmpeg
brew install srt
```

Build the binary of the server:

```
cd server
go build
```

You should end up with the `server` binary.

Build the client:

```
cd contrib
clang -Wall -O2 -o client client.c -lsrt -L/usr/local/Cellar/openssl@1.1/1.1.1h/lib -lssl -lcrypto
```

You may need to adjust the path to the OpenSSL libs as the version might change meanwhile.

Alternatively you can use the pure golang client:

```
cd client
go build
```

## Connect the pieces

Start ffmpeg pumping a mpegts stream to an UDP address:

```
ffmpeg -f lavfi -re -i testsrc2=rate=25:size=640x360 -codec:v libx264 -b:v 1024k -maxrate:v 1024k -bufsize:v 1024k -preset ultrafast -r 25 -g 50 -pix_fmt yuv420p -vsync 1 -flags2 local_header -f mpegts "udp://127.0.0.1:6000?pkt_size=1316"
```

Then start the SRT server:

```
./server -addr :6001
```

The server will listen on udp://127.0.0.1:6001

Now send the video data to the server:

```
srt-live-transmit udp://:6000 'srt://127.0.0.1:6001?streamid=publish:/live/stream' -v
```

You should see some messages on the screen like

```
Media path: 'udp://:6000' --> 'srt://127.0.0.1:6001?streamid=publish:/live/stream'
SRT parameters specified:

    streamid = 'publish:/live/stream'
Opening SRT target caller on 127.0.0.1:6001
Connecting to 127.0.0.1:6001
SRT target connected 
```

The console where the SRT server is running should also show something about handshake and so on.

Alternatively your can use the golang client:

```
cd client
./client -from udp://:6000 -to "srt://127.0.0.1:6001/?streamid=publish:/live/stream"
```

Now start the client (C, based on libsrt) to get the stream from the server and pipe it into ffplay:

```
cd contrib
./client 127.0.0.1 6001 /live/stream | ffplay -f mpegts -i -
```

or with the golang client:

```
cd client
./client -from "srt://127.0.0.1:6001/?streamid=/live/stream" -to - | ffplay -f mpegts -i -
```

You will first see some error messages from ffplay because the stream will most likely not start at a key frame. But then the window
with the video stream should pop up.

## Encryption

The stream can be encrypted with a passphrase. First start the server with a passphrase (the passphrase has to be at least 10 characters long
otherwise `srt-live-transmit` will not accept it):

```
./server -addr :6001 -passphrase foobarfoobar
```

Send an encrpyted stream to the server:

```
srt-live-transmit udp://:6000 'srt://127.0.0.1:6001?streamid=publish:/live/stream&passphrase=foobarfoobar' -v
```

Receive an encrypted stream from the server:

```
cd client
./client -from "srt://127.0.0.1:6001/?streamid=/live/stream&passphrase=foobarfoobar" -to - | ffplay -f mpegts -i -
``` 
