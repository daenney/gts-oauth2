package generates_test

import (
	"context"
	"testing"
	"time"

	"github.com/gotosocial/oauth2/v4"
	"github.com/gotosocial/oauth2/v4/generates"
	"github.com/gotosocial/oauth2/v4/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthorize(t *testing.T) {
	Convey("Test Authorize Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Client: models.New("123456", "123456", "", ""),
			UserID:   "000000",
			CreateAt: time.Now(),
		}
		gen := generates.NewAuthorizeGenerate()
		code, err := gen.Token(context.Background(), data)
		So(err, ShouldBeNil)
		So(code, ShouldNotBeEmpty)
		Println("\nAuthorize Code:" + code)
	})
}
