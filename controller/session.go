package controller

import (
	// "archive/zip"
	// "encoding/json"
	// "fmt"
	// "github.com/eaciit/colony-core/v0"
	"github.com/eaciit/acl"
	"github.com/eaciit/colony-manager/helper"
	"github.com/eaciit/dbox"
	// _ "github.com/eaciit/dbox/dbc/jsons"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	// "io"
	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "strings"
	// "time"
	// "reflect"
)

type SessionController struct {
	App
}

func CreateSessionController(s *knot.Server) *SessionController {
	var controller = new(SessionController)
	controller.Server = s
	return controller
}

func (a *SessionController) GetSession(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	a.InitialSetDatabase()
	tSession := new(acl.Session)

	arrm := make([]toolkit.M, 0, 0)
	c, err := acl.Find(tSession, nil, toolkit.M{}.Set("take", 0))
	if err == nil {
		err = c.Fetch(&arrm, 0, false)
	}

	if err != nil {
		return helper.CreateResult(true, nil, err.Error())
	} else {
		return helper.CreateResult(true, arrm, "")
	}

}
func (a *SessionController) FindSession(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	a.InitialSetDatabase()
	payload := map[string]interface{}{}
	err := r.GetPayload(&payload)
	tSession := new(acl.Session)
	err = acl.FindByID(tSession, payload["_id"].(string))
	if err != nil {
		return helper.CreateResult(true, nil, err.Error())
	} else {
		return helper.CreateResult(true, tSession, "")
	}

}
func (a *SessionController) DeleteSession(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	a.InitialSetDatabase()
	tSession := new(acl.Session)
	payload := map[string]interface{}{}
	err := r.GetPayload(&payload)
	err = acl.FindByID(tSession, payload["_id"].(string))

	if err != nil {
		return helper.CreateResult(true, nil, err.Error())
	}

	err = acl.Delete(tSession)
	if err != nil {
		return helper.CreateResult(true, nil, err.Error())
	}

	return helper.CreateResult(true, nil, "success")
}

func (a *SessionController) prepareconnection() (conn dbox.IConnection, err error) {
	conn, err = dbox.NewConnection("mongo",
		&dbox.ConnectionInfo{"localhost:27017", "valegrab", "", "", toolkit.M{}.Set("timeout", 3)})
	if err != nil {
		return
	}
	err = conn.Connect()
	return
}

func (a *SessionController) InitialSetDatabase() error {
	conn, err := a.prepareconnection()

	if err != nil {
		return err
	}

	err = acl.SetDb(conn)
	if err != nil {
		return err
	}
	return nil
}