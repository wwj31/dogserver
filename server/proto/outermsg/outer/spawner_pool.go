// Code generated by "spawner -pool=true"; DO NOT EDIT.

package outer

import "sync"

var spawnerPools = map[string]*sync.Pool{
	"outer.ChatResp":            {New: func() interface{} { return spawner["outer.ChatResp"]() }},
	"outer.ChatReq":             {New: func() interface{} { return spawner["outer.ChatReq"]() }},
	"outer.ChatNotify":          {New: func() interface{} { return spawner["outer.ChatNotify"]() }},
	"outer.MailInfo":            {New: func() interface{} { return spawner["outer.MailInfo"]() }},
	"outer.Mail":                {New: func() interface{} { return spawner["outer.Mail"]() }},
	"outer.Unknown":             {New: func() interface{} { return spawner["outer.Unknown"]() }},
	"outer.Ok":                  {New: func() interface{} { return spawner["outer.Ok"]() }},
	"outer.Fail":                {New: func() interface{} { return spawner["outer.Fail"]() }},
	"outer.UseItemResp":         {New: func() interface{} { return spawner["outer.UseItemResp"]() }},
	"outer.UseItemReq":          {New: func() interface{} { return spawner["outer.UseItemReq"]() }},
	"outer.ItemChangeNotify":    {New: func() interface{} { return spawner["outer.ItemChangeNotify"]() }},
	"outer.Ping":                {New: func() interface{} { return spawner["outer.Ping"]() }},
	"outer.Pong":                {New: func() interface{} { return spawner["outer.Pong"]() }},
	"outer.EnterGameResp":       {New: func() interface{} { return spawner["outer.EnterGameResp"]() }},
	"outer.RoleInfoPush":        {New: func() interface{} { return spawner["outer.RoleInfoPush"]() }},
	"outer.LoginReq":            {New: func() interface{} { return spawner["outer.LoginReq"]() }},
	"outer.LoginResp":           {New: func() interface{} { return spawner["outer.LoginResp"]() }},
	"outer.EnterGameReq":        {New: func() interface{} { return spawner["outer.EnterGameReq"]() }},
	"outer.ItemInfoPush":        {New: func() interface{} { return spawner["outer.ItemInfoPush"]() }},
	"outer.ReadMailReq":         {New: func() interface{} { return spawner["outer.ReadMailReq"]() }},
	"outer.ReadMailResp":        {New: func() interface{} { return spawner["outer.ReadMailResp"]() }},
	"outer.ReceiveMailItemReq":  {New: func() interface{} { return spawner["outer.ReceiveMailItemReq"]() }},
	"outer.MailListResp":        {New: func() interface{} { return spawner["outer.MailListResp"]() }},
	"outer.ReceiveMailItemResp": {New: func() interface{} { return spawner["outer.ReceiveMailItemResp"]() }},
	"outer.DeleteMailReq":       {New: func() interface{} { return spawner["outer.DeleteMailReq"]() }},
	"outer.AddMailNotify":       {New: func() interface{} { return spawner["outer.AddMailNotify"]() }},
	"outer.MailListReq":         {New: func() interface{} { return spawner["outer.MailListReq"]() }},
}
