package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	stdLog "log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ranco-dev/gbms/api/modules/account"
	"github.com/Ranco-dev/gbms/api/modules/group"
	"github.com/Ranco-dev/gbms/api/modules/permission"
	"github.com/Ranco-dev/gbms/pkg/config"
	"github.com/Ranco-dev/gbms/pkg/db"
	"github.com/Ranco-dev/gbms/pkg/log"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var router = testSetupRouter()

func testSetupRouter() (router *gin.Engine) {
	gin.SetMode(gin.TestMode)
	router = gin.New()

	routeURL(router)

	return
}

func testSetupDB() {
	conf := config.GetConfig()
	dbName := "gbms_test"
	dbHost := conf.GetString("db.host")
	dbPort := conf.GetString("db.port")
	dbUser := conf.GetString("db.user")
	dbPassword := conf.GetString("db.pass")

	if err := db.PgDBConnection(dbHost, dbPort, dbUser, dbPassword, dbName); err != nil {
		stdLog.Panic(err.Error())
	}

	err := db.Conn.Ping(context.Background())
	if err != nil {
		stdLog.Panic(err.Error())
	}
}

func init() {
	log.InitLogger("./logs/gbms.log", "debug", 1, 5, 30)
	defer log.Sugar.Sync()
	testSetupDB()
}

func Test_CreatePermission(t *testing.T) {
	Convey("Test Create Permission1", t, func() {
		w := httptest.NewRecorder()
		var active int = 1
		createPermission := permission.Permission{
			Title:       "單元測試" + randStringRunes(6),
			Slug:        "/go/unit/" + randStringRunes(6),
			Description: "go unit test",
			Active:      &active,
		}
		jsonValue, _ := json.Marshal(createPermission)
		req, _ := http.NewRequest("POST", "/api/v1/permission", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})
}

func Test_GetPermissionList(t *testing.T) {
	type response struct {
		Permissions []permission.Permission
	}
	var permissions response

	Convey("Test Get Permission List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/permissions", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &permissions)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Permission", func() {
			So(permissions, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_GetPermission(t *testing.T) {
	Convey("Test Get Permission", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/permission/2", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Permission permission.Permission
		}
		var permission response
		err := json.Unmarshal(w.Body.Bytes(), &permission)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as Permission", func() {
			So(permission, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_UpdatePermission(t *testing.T) {
	Convey("Test Update Permission", t, func() {
		w := httptest.NewRecorder()
		var active int = 0
		updatePermission := permission.UpdatePermission{
			ID:          2,
			Title:       "單元測試" + randStringRunes(6),
			Slug:        "/go/unit/" + randStringRunes(6),
			Description: "go unit test2",
			Active:      &active,
		}
		jsonValue, _ := json.Marshal(updatePermission)
		req, _ := http.NewRequest("PATCH", "/api/v1/permission", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_CreateGroup(t *testing.T) {
	Convey("Test Create Group", t, func() {
		w := httptest.NewRecorder()

		createGroup := group.Group{
			Name:        "gotestGroup_" + randStringRunes(6),
			Permissions: []int{2},
			SubGroups:   []int{2},
			Description: "gotestGroup",
		}
		jsonValue, _ := json.Marshal(createGroup)
		req, _ := http.NewRequest("POST", "/api/v1/group", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})
}

func Test_GetGroupList(t *testing.T) {
	type response struct {
		Groups []group.Group
	}
	var groups response

	Convey("Test Get Group List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/groups", nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &groups)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Group", func() {
			So(groups, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_GetGroup(t *testing.T) {
	type response struct {
		Group group.Group
	}
	var group response

	Convey("Test Get Group", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/group/2", nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &group)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as Group", func() {
			So(group, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_UpdateGroup(t *testing.T) {
	Convey("Test Update Group", t, func() {
		w := httptest.NewRecorder()

		updateGroup := group.UpdateGroup{
			ID:          2,
			Name:        "gotestGroup2_" + randStringRunes(6),
			Permissions: []int{2},
			Description: "gotestGroup2",
		}
		jsonValue, _ := json.Marshal(updateGroup)
		req, _ := http.NewRequest("PATCH", "/api/v1/group", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_CreateAccount(t *testing.T) {
	Convey("Test Create Account", t, func() {
		w := httptest.NewRecorder()

		var groups = []int{2}
		var active = 1
		createAccount := account.CreateInfo{
			Login:    "gotest_" + randStringRunes(6),
			Username: "gotest",
			Email:    "gotest_" + randStringRunes(6) + "@gotest.com",
			Password: "gotest",
			Remark:   "gotest",
			Active:   &active,
			Group:    &groups,
		}
		jsonValue, _ := json.Marshal(createAccount)
		req, _ := http.NewRequest("POST", "/api/v1/account", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})
}

func Test_GetAccountList(t *testing.T) {
	type response struct {
		Accounts []account.User
	}
	var users response

	Convey("Test Get Account List", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/accounts", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &users)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []User", func() {
			So(users, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_GetAccount(t *testing.T) {
	type response struct {
		Account account.User
	}
	var user response

	Convey("Test Get Account", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/account/2", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &user)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as User", func() {
			So(user, ShouldHaveSameTypeAs, response{})
		})
	})
}

func Test_UpdateAccount(t *testing.T) {
	Convey("Test Update Account", t, func() {
		w := httptest.NewRecorder()

		var groups = []int{2}
		var active = 0
		var remark = "gotest2"
		updateAccount := account.UpdateUser{
			ID:       "2",
			Username: "gotest2_" + randStringRunes(6),
			Email:    "gotest2_" + randStringRunes(6) + "@gotest.com",
			Password: "gotest2",
			Remark:   &remark,
			Active:   &active,
			Group:    &groups,
		}
		jsonValue, _ := json.Marshal(updateAccount)
		req, _ := http.NewRequest("PATCH", "/api/v1/account", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_GrantPermission(t *testing.T) {
	Convey("Test Grant Permission", t, func() {
		w := httptest.NewRecorder()
		grantPermission := permission.GrantPermission{
			GroupID:       2,
			PremissionsID: []int{2},
		}
		jsonValue, _ := json.Marshal(grantPermission)
		req, _ := http.NewRequest("PUT", "/api/v1/permission/grant", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_RemovePermission(t *testing.T) {
	Convey("Test Remove Permission", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", "/api/v1/permission/3", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusNoContent {
			stdLog.Println(w.Body)
		}
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_RemoveGroup(t *testing.T) {
	Convey("Test Remove Group", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", "/api/v1/group/remove_group/3", nil)
		router.ServeHTTP(w, req)
		if w.Code != http.StatusNoContent {
			stdLog.Println(w.Body)
		}
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func Test_RemoveAccount(t *testing.T) {
	Convey("Test Remove Account", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", "/api/v1/account/3", nil)

		router.ServeHTTP(w, req)
		if w.Code != http.StatusNoContent {
			stdLog.Println(w.Body)
		}
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

// Test_CRUD_PGA is testing permission, group and account CRUD
func Test_CRUD_PGA(t *testing.T) {
	router := testSetupRouter()
	testSetupDB()

	//
	// Permission
	//
	type permissionResponse struct {
		Permissions []permission.Permission
	}
	var permissions permissionResponse

	Convey("Test Create Permission1", t, func() {
		w := httptest.NewRecorder()
		var active int = 1
		createPermission := permission.Permission{
			Title:       "單元測試",
			Slug:        "/go/unit/test",
			Description: "go unit test",
			Active:      &active,
		}
		jsonValue, _ := json.Marshal(createPermission)
		req, _ := http.NewRequest("POST", "/api/v1/permission", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})

	Convey("Test Get Permission List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/permissions", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &permissions)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Permission", func() {
			So(permissions, ShouldHaveSameTypeAs, permissionResponse{})
		})
	})

	pid1 := permissions.Permissions[len(permissions.Permissions)-1].ID

	Convey("Test Get Permission", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/permission/%d", pid1), nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Permission permission.Permission
		}
		var permission response
		err := json.Unmarshal(w.Body.Bytes(), &permission)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as Permission", func() {
			So(permission, ShouldHaveSameTypeAs, response{})
		})
	})

	Convey("Test Update Permission", t, func() {
		w := httptest.NewRecorder()
		var active int = 0
		updatePermission := permission.UpdatePermission{
			ID:          pid1,
			Title:       "單元測試2",
			Slug:        "/go/unit/test2",
			Description: "go unit test2",
			Active:      &active,
		}
		jsonValue, _ := json.Marshal(updatePermission)
		req, _ := http.NewRequest("PATCH", "/api/v1/permission", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Create Permission2", t, func() {
		w := httptest.NewRecorder()
		var active int = 1
		createPermission := permission.Permission{
			Title:       "單元測試3",
			Slug:        "/go/unit/test3",
			Description: "go unit test3",
			Active:      &active,
		}
		jsonValue, _ := json.Marshal(createPermission)
		req, _ := http.NewRequest("POST", "/api/v1/permission", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})

	permissions.Permissions = nil
	Convey("Test Get Permission List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/permissions", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &permissions)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Permission", func() {
			So(permissions, ShouldHaveSameTypeAs, permissionResponse{})
		})
	})

	pid2 := permissions.Permissions[len(permissions.Permissions)-1].ID

	//
	// Group
	//
	Convey("Test Create Group", t, func() {
		w := httptest.NewRecorder()

		createGroup := group.Group{
			Name:        "gotestGroup",
			Permissions: []int{pid1, pid2},
			Description: "gotestGroup",
		}
		jsonValue, _ := json.Marshal(createGroup)
		req, _ := http.NewRequest("POST", "/api/v1/group", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})

	type response struct {
		Groups []group.Group
	}
	var groups response

	Convey("Test Get Group List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/groups", nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &groups)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Group", func() {
			So(groups, ShouldHaveSameTypeAs, response{})
		})
	})

	gid1 := groups.Groups[len(groups.Groups)-1].ID
	Convey("Test Get Group", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/group/%d", gid1), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Group group.Group
		}
		var group response
		err := json.Unmarshal(w.Body.Bytes(), &group)
		So(err, ShouldBeNil)

		if len(group.Group.Permissions) != 2 {
			So(err, ShouldBeIn)
		}

		Convey("The response should have the same type as Group", func() {
			So(group, ShouldHaveSameTypeAs, response{})
		})
	})

	Convey("Test Update Group", t, func() {
		w := httptest.NewRecorder()

		updateGroup := group.UpdateGroup{
			ID:          gid1,
			Name:        "gotestGroup2",
			Permissions: []int{pid1},
			Description: "gotestGroup2",
		}
		jsonValue, _ := json.Marshal(updateGroup)
		req, _ := http.NewRequest("PATCH", "/api/v1/group", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Get Group", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/group/%d", gid1), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Group group.Group
		}
		var group response
		err := json.Unmarshal(w.Body.Bytes(), &group)
		So(err, ShouldBeNil)

		if len(group.Group.Permissions) != 1 {
			So(err, ShouldBeIn)
		}
	})

	Convey("Test Create Group2", t, func() {
		w := httptest.NewRecorder()

		createGroup := group.Group{
			Name:        "gotestGroup3",
			Permissions: []int{pid1, pid2},
			Description: "gotestGroup3",
		}
		jsonValue, _ := json.Marshal(createGroup)
		req, _ := http.NewRequest("POST", "/api/v1/group", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})

	groups.Groups = nil
	Convey("Test Get Group List", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/api/v1/groups", nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &groups)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []Group", func() {
			So(groups, ShouldHaveSameTypeAs, response{})
		})
	})

	gid2 := groups.Groups[len(groups.Groups)-1].ID

	//
	// Account
	//

	type accountResponse struct {
		Accounts []account.User
	}
	var users accountResponse

	Convey("Test Create Account", t, func() {
		w := httptest.NewRecorder()

		var groups = []int{gid1, gid2}
		var active = 1
		createAccount := account.CreateInfo{
			Login:    "gotest",
			Username: "gotest",
			Email:    "gotest@gotest.com",
			Password: "gotest",
			Remark:   "gotest",
			Active:   &active,
			Group:    &groups,
		}
		jsonValue, _ := json.Marshal(createAccount)
		req, _ := http.NewRequest("POST", "/api/v1/account", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusCreated)
	})

	Convey("Test Get Account List", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/accounts", nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		err := json.Unmarshal(w.Body.Bytes(), &users)
		So(err, ShouldBeNil)

		Convey("The response should have the same type as []User", func() {
			So(users, ShouldHaveSameTypeAs, accountResponse{})
		})
	})

	var uid = users.Accounts[len(users.Accounts)-1].ID

	Convey("Test Get Account", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/account/"+uid, nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Account account.User
		}
		var user response
		err := json.Unmarshal(w.Body.Bytes(), &user)
		So(err, ShouldBeNil)

		if len(user.Account.Group) != 2 {
			So(err, ShouldBeIn)
		}

		Convey("The response should have the same type as User", func() {
			So(user, ShouldHaveSameTypeAs, response{})
		})
	})

	Convey("Test Update Account", t, func() {
		w := httptest.NewRecorder()

		var groups = []int{gid1}
		var active = 0
		var remark = "gotest2"
		updateAccount := account.UpdateUser{
			ID:       uid,
			Username: "gotest2",
			Email:    "gotest2@gotest.com",
			Password: "gotest2",
			Remark:   &remark,
			Active:   &active,
			Group:    &groups,
		}
		jsonValue, _ := json.Marshal(updateAccount)
		req, _ := http.NewRequest("PATCH", "/api/v1/account", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Get Account", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/account/"+uid, nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)

		type response struct {
			Account account.User
		}
		var user response
		err := json.Unmarshal(w.Body.Bytes(), &user)
		So(err, ShouldBeNil)

		if len(user.Account.Group) != 1 {
			So(err, ShouldBeIn)
		}
	})

	Convey("Test Grant Permission", t, func() {
		w := httptest.NewRecorder()
		grantPermission := permission.GrantPermission{
			GroupID:       gid1,
			PremissionsID: []int{pid1},
		}
		jsonValue, _ := json.Marshal(grantPermission)
		req, _ := http.NewRequest("PUT", "/api/v1/permission/grant", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Remove Permission", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/permission/%d", pid1), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)

		req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/v1/permission/%d", pid2), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Remove Group", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/group/remove_group/%d", gid1), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)

		req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/v1/group/remove_group/%d", gid2), nil)
		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})

	Convey("Test Remove Account", t, func() {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("DELETE", "/api/v1/account/"+uid, nil)

		router.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusNoContent)
	})
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
