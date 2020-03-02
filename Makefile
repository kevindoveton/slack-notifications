bundle:
	make _bundle APP=snotification

_bundle:
	mkdir -p ${APP}.app/Contents/MacOS && \
		mkdir -p ${APP}.app/Contents/Resources && \
		make icons && \
		mv icon.icns ${APP}.app/Contents/Resources && \
		cp assets/Info.plist ${APP}.app/Contents && \
		cp assets/icon.png ${APP}.app/Contents/Resources && \
		make build && \
		mv snotification ${APP}.app/Contents/MacOS

run:
	go run cmd/snotification/main.go

build:
	go build cmd/snotification/main.go && \
		mv main snotification

icons:
	make _icons ICON_PATH=assets/icon.png ICONSET_PATH=/tmp/Icon.iconset

_icons:
	mkdir ${ICONSET_PATH} && \
		sips -z 16 16     ${ICON_PATH} --out ${ICONSET_PATH}/icon_16x16.png && \
		sips -z 32 32     ${ICON_PATH} --out ${ICONSET_PATH}/icon_16x16@2x.png && \
		sips -z 32 32     ${ICON_PATH} --out ${ICONSET_PATH}/icon_32x32.png && \
		sips -z 64 64     ${ICON_PATH} --out ${ICONSET_PATH}/icon_32x32@2x.png && \
		sips -z 128 128   ${ICON_PATH} --out ${ICONSET_PATH}/icon_128x128.png && \
		sips -z 256 256   ${ICON_PATH} --out ${ICONSET_PATH}/icon_128x128@2x.png && \
		sips -z 256 256   ${ICON_PATH} --out ${ICONSET_PATH}/icon_256x256.png && \
		sips -z 512 512   ${ICON_PATH} --out ${ICONSET_PATH}/icon_256x256@2x.png && \
		sips -z 512 512   ${ICON_PATH} --out ${ICONSET_PATH}/icon_512x512.png && \
		cp ${ICON_PATH} ${ICONSET_PATH}/icon_512x512@2x.png && \
		iconutil --convert icns ${ICONSET_PATH} --output icon.icns && \
		rm -R ${ICONSET_PATH}
