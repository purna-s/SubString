package SubString

import (

	"fmt"
	"strings"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-SubString")

// MyActivity is a stub for your Activity implementation
type SubString struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SubString{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *SubString) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SubString) Eval(ctx activity.Context) (done bool, err error) {
	

	str := ctx.GetInput("InputString").(string)
	sep := ctx.GetInput("Separator").(string)

    i := strings.Index(str, sep)
    //fmt.Println("Index: ", i)
    if i > -1 {
        a := str[:i]
        b := str[i+1:]
        ctx.SetOutput("SubStringBefore", a)
        ctx.SetOutput("SubStringAfter", b)
    } else {
        fmt.Println("Index not found")
        fmt.Println(str)
    }

}
