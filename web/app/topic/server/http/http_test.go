package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Seven4X/link/web/library/api"
	setup "github.com/Seven4X/link/web/library/echo"
	mydb "github.com/Seven4X/link/web/library/store/db"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	topicJson    = `{"name":"t1","refTopicId":12,"position":4}`
	topicJsonErr = `{"name":"t1","refTopicId":12,"position":5}`
)

//使用go-sqlmock仅测试web层
//fixme https://github.com/DATA-DOG/go-sqlmock/issues/239 sqlmock应该解析SQL而不是根据是否通过insert或者query调用，因为query也能做插入操作
func TestCreateTopicSucc(t *testing.T) {
	// 1.Setup
	e := setup.NewEcho()

	req := httptest.NewRequest(http.MethodPost, "/topic", strings.NewReader(topicJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//2. mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.MatchExpectationsInOrder(false)

	mock.ExpectExec(`INSERT INTO "topic"`).
		WithArgs("t1", "", "", 0, 0, 0).
		WillReturnResult(sqlmock.NewResult(15, 1))
	//3.设置mock数据库
	mydb.SetMockDb(db)
	//4.Assertions
	if assert.NoError(t, newTopic(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		suc := api.Result{}
		_ = json.Unmarshal(rec.Body.Bytes(), &suc)
		println(rec.Body.String())
		assert.Equal(t, 0, suc.Ok)
		assert.Equal(t, "15", suc.Data)
	}
}

func TestCreateTopicFailed(t *testing.T) {
	// Setup
	e := setup.NewEcho()

	req := httptest.NewRequest(http.MethodPost, "/topic", strings.NewReader(topicJsonErr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, newTopic(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		suc := api.Result{}
		json.Unmarshal(rec.Body.Bytes(), &suc)
		println(rec.Body.String())
		assert.Equal(t, 1, suc.Ok)
	}
}

//参考：https://learnku.com/go/t/26711 如果需要可以和配合testfixtures进行环境构造
func TestCreateTopicSuccUseTxdb(t *testing.T) {
	// 1.Setup
	e := setup.NewEcho()

	req := httptest.NewRequest(http.MethodPost, "/topic", strings.NewReader(topicJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//2.设置mock数据库
	mydb.RegisterMockDriver()
	//3.Assertions
	if assert.NoError(t, newTopic(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		suc := api.Result{}
		_ = json.Unmarshal(rec.Body.Bytes(), &suc)
		println(rec.Body.String())
		assert.Equal(t, 0, suc.Ok)
		assert.Equal(t, "1", suc.Data)
	}
}
