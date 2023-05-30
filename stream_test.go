package p4

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStream_Streams(t *testing.T) {
	var (
		si      *StreamInfo
		streams []*StreamInfo
		stream  = "//DM99.ZGame.Project/Development/ZGame_ArtDev"
	)
	conn, err := setup(t)
	Convey("test Stream functions", t, func() {
		So(err, ShouldBeNil)

		Convey("List streams", func() {
			streams, err = conn.Streams()
			So(err, ShouldBeNil)
			So(len(streams), ShouldBeGreaterThanOrEqualTo, 0)
		})

		Convey("Get stream info", func() {
			si, err = conn.Stream(stream)
			So(err, ShouldBeNil)
			So(si.Stream, ShouldEqual, stream)
		})

		Convey("Delete stream", func() {
			message, err := conn.DeleteStream(stream, true)
			So(err, ShouldBeNil)
			So(message, ShouldEqual, fmt.Sprintf("Stream %s deleted.", stream))
		})

		Convey("Create stream", func() {
			var (
				name       = "ZGame_ArtDev"
				parent     = "//DM99.ZGame.Project/Main/ZGame_Mainline"
				streamType = "development"
			)
			message, err := conn.CreateStream(name, streamType, parent, stream, true)
			So(err, ShouldBeNil)
			So(message, ShouldEqual, fmt.Sprintf("Stream %s saved.", stream))
		})

		Convey("Create stream with invalid stream type", func() {
			var (
				name       = "ZGame_ArtDev"
				parent     = "//DM99.ZGame.Project/Main/ZGame_Mainline"
				streamType = "abc"
			)
			_, err := conn.CreateStream(name, streamType, parent, stream, true)
			So(err, ShouldNotBeNil)
		})
	})
}
