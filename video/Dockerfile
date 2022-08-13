FROM golang:1.18.5-bullseye AS builder

WORKDIR /app

COPY ./video/ .

RUN go build -o video .

FROM debian:bullseye

WORKDIR /setup

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl tar gzip bzip2

RUN apt-get install -y autoconf automake build-essential gcc g++ cmake git-core \
    libass-dev libfreetype6-dev libgnutls28-dev libmp3lame-dev libsdl2-dev libtool \
    libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev \
    pkg-config texinfo wget yasm zlib1g-dev

# Installing meson & ninja build tools
RUN apt-get install -y python3 python3-pip python3-setuptools python3-wheel

WORKDIR /setup/ninja

RUN curl -LO https://github.com/ninja-build/ninja/archive/refs/tags/v1.11.0.tar.gz && \
    tar -zxf v1.11.0.tar.gz && \
    cd ninja-1.11.0 && \
    python3 configure.py --bootstrap && \
    cp -v ninja /usr/bin/ninja

WORKDIR /setup/meson

RUN curl -LO https://github.com/mesonbuild/meson/releases/download/0.63.0/meson-0.63.0.tar.gz && \
    tar -zxf meson-0.63.0.tar.gz && \
    cd meson-0.63.0 && \
    ./packaging/create_zipapp.py --outfile meson --interpreter '/usr/bin/env python3' . && \
    cp -v meson /usr/bin/meson && \
    mkdir -vp ~/.local/bin && \
    cp -v meson ~/.local/bin/meson

WORKDIR /setup/nasm

RUN wget https://www.nasm.us/pub/nasm/releasebuilds/2.15.05/nasm-2.15.05.tar.bz2 && \
    tar xjvf nasm-2.15.05.tar.bz2 && \
    cd nasm-2.15.05 && \
    ./autogen.sh && \
    ./configure --prefix="/usr" --bindir="/usr/bin" && \
    make && \
    make install

WORKDIR /setup/x264

RUN git -C x264 pull 2> /dev/null || git clone --depth 1 https://code.videolan.org/videolan/x264.git && \
    cd x264 && \
    PKG_CONFIG_PATH="/usr/lib/pkgconfig" ./configure --prefix="/usr" --bindir="/usr/bin" --enable-static --enable-pic && \
    make && \
    make install

WORKDIR /setup/x265

RUN apt-get install -y libnuma-dev && \
    wget -O x265.tar.bz2 https://bitbucket.org/multicoreware/x265_git/get/master.tar.bz2 && \
    tar xjvf x265.tar.bz2 && \
    cd multicoreware*/build/linux && \
    cmake -G "Unix Makefiles" -DCMAKE_INSTALL_PREFIX="/usr" -DENABLE_SHARED=off ../../source && \
    make && \
    make install

WORKDIR /setup/libvpx

RUN git -C libvpx pull 2> /dev/null || git clone --depth 1 https://chromium.googlesource.com/webm/libvpx.git && \
    cd libvpx && \
    ./configure --prefix="/usr" --disable-examples --disable-unit-tests --enable-vp9-highbitdepth --as=yasm && \
    make && \
    make install

WORKDIR /setup/opus

RUN git -C opus pull 2> /dev/null || git clone --depth 1 https://github.com/xiph/opus.git && \
    cd opus && \
    ./autogen.sh && \
    ./configure --prefix="/usr" --disable-shared && \
    make && \
    make install

WORKDIR /setup/vmaf

RUN wget https://github.com/Netflix/vmaf/archive/v2.1.1.tar.gz && \
    tar xvf v2.1.1.tar.gz && \
    mkdir -p vmaf-2.1.1/libvmaf/build &&\
    cd vmaf-2.1.1/libvmaf/build && \
    meson setup -Denable_tests=false -Denable_docs=false --buildtype=release --default-library=static .. --prefix "/usr" --bindir="/usr/bin" --libdir="/usr/lib" && \
    ninja && \
    ninja install

WORKDIR /setup/dav1d

RUN git -C dav1d pull 2> /dev/null || git clone --depth 1 https://code.videolan.org/videolan/dav1d.git && \
    mkdir -p dav1d/build && \
    cd dav1d/build && \
    meson setup -Denable_tools=false -Denable_tests=false --default-library=static .. --prefix "/usr" --libdir="/usr/lib" && \
    ninja && \
    ninja install

WORKDIR /setup/ffmpeg

RUN apt-get install -y gnutls-bin libgnutls30 libunistring-dev && \
    wget -O ffmpeg-5.1.tar.bz2 https://ffmpeg.org/releases/ffmpeg-5.1.tar.bz2 && \
    tar xjvf ffmpeg-5.1.tar.bz2 && \
    cd ffmpeg-5.1 && \
    PKG_CONFIG_PATH="/usr/lib/pkgconfig" LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH ./configure \
    --prefix="/usr" \
    --pkg-config-flags="--static" \
    --extra-cflags="-I/usr/bin/include" \
    --extra-ldflags="-L/usr/bin/lib" \
    --extra-libs="-lpthread -lm" \
    --ld="g++" \
    --bindir="/usr/bin" \
    --enable-gpl \
    --enable-gnutls \
    --enable-libass \
    --enable-libfreetype \
    --enable-libmp3lame \
    --enable-libopus \
    --enable-libdav1d \
    --enable-libvorbis \
    --enable-libvpx \
    --enable-libx264 \
    --enable-libx265 && \
    make && \
    make install && \
    hash -r

WORKDIR /setup/grpcurl

RUN curl -LO https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && \
    tar -zxf grpcurl_1.8.6_linux_x86_64.tar.gz && \
    mv grpcurl /usr/bin/grpcurl && \
    chmod +x /usr/bin/grpcurl && \
    rm grpcurl_1.8.6_linux_x86_64.tar.gz LICENSE

WORKDIR /app

RUN rm -rf /setup

COPY --from=builder /app/video /app/video

COPY ./proto/video.proto .

CMD [ "./video" ]
