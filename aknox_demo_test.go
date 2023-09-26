package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ondbyte/aknox_demo/auth"
	"github.com/ondbyte/aknox_demo/notes"
	"github.com/ondbyte/aknox_demo/response"
	"github.com/ondbyte/aknox_demo/router"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"AKNOX DEMO",
	func() {

		firstNote := "first note"
		router := router.NewSessionsRouter()
		auth.InitRoutes(router, auth.NewService(auth.NewRepo()))
		notes.InitRoutes(router, notes.NewService(notes.NewRepo()))
		var res *httptest.ResponseRecorder
		var sid string
		var cookies = make([]*http.Cookie, 0)

		BeforeEach(func() {
			res = httptest.NewRecorder()
		})

		AfterEach(func() {
			newCookies := res.Result().Cookies()
			for _, c := range newCookies {
				cookies = append(cookies, c)
			}
		})

		It(
			"should fail to login as service doesnt have a signed up user",
			func() {
				body := `
					{
						"email":"yadu@email.com",
						"password":"yadu@123"
					}
				`
				req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusTeapot))
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).NotTo(BeEmpty())
			},
		)

		It(
			"should signup",
			func() {
				body := `
				{
					"name":"yadu",
					"email":"yadu@email.com",
					"password":"yadu@123"
				}
			`
				req := httptest.NewRequest("POST", "/signup", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
			},
		)

		It(
			"should login",
			func() {
				body := `
				{
					"email":"yadu@email.com",
					"password":"yadu@123"
				}
			`
				req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf(map[string]interface{}{}))
				Expect(resBody.Data.(map[string]interface{})).To(HaveKey("sid"))
				Expect(resBody.Data.(map[string]interface{})["sid"]).To(BeAssignableToTypeOf(""))
				sid = resBody.Data.(map[string]interface{})["sid"].(string)
				Expect(sid).NotTo(Equal(""), "expected sid to be non empty but it is empty")
				cookies = res.Result().Cookies()
			},
		)

		It(
			"should create a note",
			func() {
				body := fmt.Sprintf(`
				{
					"sid":"%v",
					"note":"%v"
				}
			`, sid, firstNote,
				)
				req := httptest.NewRequest("POST", "/notes", strings.NewReader(body))
				for _, c := range cookies {
					req.AddCookie(c)
				}
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf(map[string]interface{}{}))
				data := resBody.Data.(map[string]interface{})
				Expect(data).To(HaveLen(2))
				Expect(data).To(HaveKey("id"))
				Expect(data).To(HaveKeyWithValue("note", firstNote))

			},
		)

		It(
			"should create a note",
			func() {
				body := fmt.Sprintf(`
				{
					"sid":"%v",
					"note":"%v"
				}
			`, sid, firstNote,
				)
				req := httptest.NewRequest("POST", "/notes", strings.NewReader(body))
				for _, c := range cookies {
					req.AddCookie(c)
				}
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf(map[string]interface{}{}))
				data := resBody.Data.(map[string]interface{})
				Expect(data).To(HaveLen(2))
				Expect(data).To(HaveKey("id"))
				Expect(data).To(HaveKeyWithValue("note", firstNote))

			},
		)

		It(
			"should list 2 notes",
			func() {
				body := fmt.Sprintf(`
			{
				"sid":"%v"
			}
		`, sid,
				)
				req := httptest.NewRequest("GET", "/notes", strings.NewReader(body))
				for _, c := range cookies {
					req.AddCookie(c)
				}
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf([]interface{}{}))
				data := resBody.Data.([]interface{})
				Expect(data).To(HaveLen(2))
				Expect(data[0]).To(BeAssignableToTypeOf(map[string]interface{}{}))
				note1 := data[0].(map[string]interface{})
				Expect(note1).To(HaveKey("id"))
				Expect(note1).To(HaveKeyWithValue("note", firstNote))
			},
		)

		It(
			"should delete a note",
			func() {
				body := fmt.Sprintf(`
				{
					"sid":"%v",
					"id":"0"
				}
			`, sid,
				)
				req := httptest.NewRequest("DELETE", "/notes", strings.NewReader(body))
				for _, c := range cookies {
					req.AddCookie(c)
				}
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf(map[string]interface{}{}))
				data := resBody.Data.(map[string]interface{})
				Expect(data).To(HaveLen(2))
				Expect(data).To(HaveKey("id"))
				Expect(data).To(HaveKeyWithValue("note", firstNote))

			},
		)

		It(
			"should list 1 notes",
			func() {
				body := fmt.Sprintf(`
			{
				"sid":"%v"
			}
		`, sid,
				)
				req := httptest.NewRequest("GET", "/notes", strings.NewReader(body))
				for _, c := range cookies {
					req.AddCookie(c)
				}
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK), "body=", res.Body.String())
				resBody, err := response.ReadResponse(res)
				Expect(err).To(BeNil())
				Expect(resBody.Error).To(BeEmpty())
				Expect(resBody.Data).To(BeAssignableToTypeOf([]interface{}{}))
				data := resBody.Data.([]interface{})
				Expect(data).To(HaveLen(1))
				Expect(data[0]).To(BeAssignableToTypeOf(map[string]interface{}{}))
				note1 := data[0].(map[string]interface{})
				Expect(note1).To(HaveKey("id"))
				Expect(note1).To(HaveKeyWithValue("note", firstNote))
			},
		)
	},
)
