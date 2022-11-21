package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/ixtendio/gofre"
	"github.com/ixtendio/gofre/handler"
	"github.com/ixtendio/gofre/response"
	"github.com/ixtendio/gofre/router/path"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type (
	Route struct {
		Method string
		Path   string
	}
)

var (
	staticRoutes = []*Route{
		{"GET", "/"},
		{"GET", "/cmd.html"},
		{"GET", "/code.html"},
		{"GET", "/contrib.html"},
		{"GET", "/contribute.html"},
		{"GET", "/debugging_with_gdb.html"},
		{"GET", "/docs.html"},
		{"GET", "/effective_go.html"},
		{"GET", "/files.log"},
		{"GET", "/gccgo_contribute.html"},
		{"GET", "/gccgo_install.html"},
		{"GET", "/go-logo-black.png"},
		{"GET", "/go-logo-blue.png"},
		{"GET", "/go-logo-white.png"},
		{"GET", "/go1.1.html"},
		{"GET", "/go1.2.html"},
		{"GET", "/go1.html"},
		{"GET", "/go1compat.html"},
		{"GET", "/go_faq.html"},
		{"GET", "/go_mem.html"},
		{"GET", "/go_spec.html"},
		{"GET", "/help.html"},
		{"GET", "/ie.css"},
		{"GET", "/install-source.html"},
		{"GET", "/install.html"},
		{"GET", "/logo-153x55.png"},
		{"GET", "/Makefile"},
		{"GET", "/root.html"},
		{"GET", "/share.png"},
		{"GET", "/sieve.gif"},
		{"GET", "/tos.html"},
		{"GET", "/articles/"},
		{"GET", "/articles/go_command.html"},
		{"GET", "/articles/index.html"},
		{"GET", "/articles/wiki/"},
		{"GET", "/articles/wiki/edit.html"},
		{"GET", "/articles/wiki/final-noclosure.go"},
		{"GET", "/articles/wiki/final-noerror.go"},
		{"GET", "/articles/wiki/final-parsetemplate.go"},
		{"GET", "/articles/wiki/final-template.go"},
		{"GET", "/articles/wiki/final.go"},
		{"GET", "/articles/wiki/get.go"},
		{"GET", "/articles/wiki/http-sample.go"},
		{"GET", "/articles/wiki/index.html"},
		{"GET", "/articles/wiki/Makefile"},
		{"GET", "/articles/wiki/notemplate.go"},
		{"GET", "/articles/wiki/part1-noerror.go"},
		{"GET", "/articles/wiki/part1.go"},
		{"GET", "/articles/wiki/part2.go"},
		{"GET", "/articles/wiki/part3-errorhandling.go"},
		{"GET", "/articles/wiki/part3.go"},
		{"GET", "/articles/wiki/test.bash"},
		{"GET", "/articles/wiki/test_edit.good"},
		{"GET", "/articles/wiki/test_Test.txt.good"},
		{"GET", "/articles/wiki/test_view.good"},
		{"GET", "/articles/wiki/view.html"},
		{"GET", "/codewalk/"},
		{"GET", "/codewalk/codewalk.css"},
		{"GET", "/codewalk/codewalk.js"},
		{"GET", "/codewalk/codewalk.xml"},
		{"GET", "/codewalk/functions.xml"},
		{"GET", "/codewalk/markov.go"},
		{"GET", "/codewalk/markov.xml"},
		{"GET", "/codewalk/pig.go"},
		{"GET", "/codewalk/popout.png"},
		{"GET", "/codewalk/run"},
		{"GET", "/codewalk/sharemem.xml"},
		{"GET", "/codewalk/urlpoll.go"},
		{"GET", "/devel/"},
		{"GET", "/devel/release.html"},
		{"GET", "/devel/weekly.html"},
		{"GET", "/gopher/"},
		{"GET", "/gopher/appenginegopher.jpg"},
		{"GET", "/gopher/appenginegophercolor.jpg"},
		{"GET", "/gopher/appenginelogo.gif"},
		{"GET", "/gopher/bumper.png"},
		{"GET", "/gopher/bumper192x108.png"},
		{"GET", "/gopher/bumper320x180.png"},
		{"GET", "/gopher/bumper480x270.png"},
		{"GET", "/gopher/bumper640x360.png"},
		{"GET", "/gopher/doc.png"},
		{"GET", "/gopher/frontpage.png"},
		{"GET", "/gopher/gopherbw.png"},
		{"GET", "/gopher/gophercolor.png"},
		{"GET", "/gopher/gophercolor16x16.png"},
		{"GET", "/gopher/help.png"},
		{"GET", "/gopher/pkg.png"},
		{"GET", "/gopher/project.png"},
		{"GET", "/gopher/ref.png"},
		{"GET", "/gopher/run.png"},
		{"GET", "/gopher/talks.png"},
		{"GET", "/gopher/pencil/"},
		{"GET", "/gopher/pencil/gopherhat.jpg"},
		{"GET", "/gopher/:pencil/gopherhat.jpg"},
		{"GET", "/gopher/pencil/gopherhelmet.jpg"},
		{"GET", "/gopher/pencil/gophermega.jpg"},
		{"GET", "/gopher/pencil/gopherrunning.jpg"},
		{"GET", "/gopher/pencil/gopherswim.jpg"},
		{"GET", "/gopher/pencil/gopherswrench.jpg"},
		{"GET", "/play/"},
		{"GET", "/play/fib.go"},
		{"GET", "/play/hello.go"},
		{"GET", "/play/life.go"},
		{"GET", "/play/peano.go"},
		{"GET", "/play/pi.go"},
		{"GET", "/play/sieve.go"},
		{"GET", "/play/solitaire.go"},
		{"GET", "/play/tree.go"},
		{"GET", "/progs/"},
		{"GET", "/progs/cgo1.go"},
		{"GET", "/progs/cgo2.go"},
		{"GET", "/progs/cgo3.go"},
		{"GET", "/progs/cgo4.go"},
		{"GET", "/progs/defer.go"},
		{"GET", "/progs/defer.out"},
		{"GET", "/progs/defer2.go"},
		{"GET", "/progs/defer2.out"},
		{"GET", "/progs/eff_bytesize.go"},
		{"GET", "/progs/eff_bytesize.out"},
		{"GET", "/progs/eff_qr.go"},
		{"GET", "/progs/eff_sequence.go"},
		{"GET", "/progs/eff_sequence.out"},
		{"GET", "/progs/eff_unused1.go"},
		{"GET", "/progs/eff_unused2.go"},
		{"GET", "/progs/error.go"},
		{"GET", "/progs/error2.go"},
		{"GET", "/progs/error3.go"},
		{"GET", "/progs/error4.go"},
		{"GET", "/progs/go1.go"},
		{"GET", "/progs/gobs1.go"},
		{"GET", "/progs/gobs2.go"},
		{"GET", "/progs/image_draw.go"},
		{"GET", "/progs/image_package1.go"},
		{"GET", "/progs/image_package1.out"},
		{"GET", "/progs/image_package2.go"},
		{"GET", "/progs/image_package2.out"},
		{"GET", "/progs/image_package3.go"},
		{"GET", "/progs/image_package3.out"},
		{"GET", "/progs/image_package4.go"},
		{"GET", "/progs/image_package4.out"},
		{"GET", "/progs/image_package5.go"},
		{"GET", "/progs/image_package5.out"},
		{"GET", "/progs/image_package6.go"},
		{"GET", "/progs/image_package6.out"},
		{"GET", "/progs/interface.go"},
		{"GET", "/progs/interface2.go"},
		{"GET", "/progs/interface2.out"},
		{"GET", "/progs/json1.go"},
		{"GET", "/progs/json2.go"},
		{"GET", "/progs/json2.out"},
		{"GET", "/progs/json3.go"},
		{"GET", "/progs/json4.go"},
		{"GET", "/progs/json5.go"},
		{"GET", "/progs/run"},
		{"GET", "/progs/slices.go"},
		{"GET", "/progs/timeout1.go"},
		{"GET", "/progs/timeout2.go"},
		{"GET", "/progs/update.bash"},
	}

	varCaptureRoutes = []*Route{
		// OAuth Authorizations
		{"GET", "/authorizations"},
		{"GET", "/authorizations/{id}"},
		{"POST", "/authorizations"},
		//{"PUT", "/authorizations/clients/{client_id}"},
		//{"PATCH", "/authorizations/{id}"},
		{"DELETE", "/authorizations/{id}"},
		{"GET", "/applications/{client_id}/tokens/{access_token}"},
		{"DELETE", "/applications/{client_id}/tokens"},
		{"DELETE", "/applications/{client_id}/tokens/{access_token}"},

		// Activity
		{"GET", "/events"},
		{"GET", "/repos/{owner}/{repo}/events"},
		{"GET", "/networks/{owner}/{repo}/events"},
		{"GET", "/orgs/{org}/events"},
		{"GET", "/users/{user}/received_events"},
		{"GET", "/users/{user}/received_events/public"},
		{"GET", "/users/{user}/events"},
		{"GET", "/users/{user}/events/public"},
		{"GET", "/users/{user}/events/orgs/{org}"},
		{"GET", "/feeds"},
		{"GET", "/notifications"},
		{"GET", "/repos/{owner}/{repo}/notifications"},
		{"PUT", "/notifications"},
		{"PUT", "/repos/{owner}/{repo}/notifications"},
		{"GET", "/notifications/threads/{id}"},
		//{"PATCH", "/notifications/threads/{id}"},
		{"GET", "/notifications/threads/{id}/subscription"},
		{"PUT", "/notifications/threads/{id}/subscription"},
		{"DELETE", "/notifications/threads/{id}/subscription"},
		{"GET", "/repos/{owner}/{repo}/stargazers"},
		{"GET", "/users/{user}/starred"},
		{"GET", "/user/starred"},
		{"GET", "/user/starred/{owner}/{repo}"},
		{"PUT", "/user/starred/{owner}/{repo}"},
		{"DELETE", "/user/starred/{owner}/{repo}"},
		{"GET", "/repos/{owner}/{repo}/subscribers"},
		{"GET", "/users/{user}/subscriptions"},
		{"GET", "/user/subscriptions"},
		{"GET", "/repos/{owner}/{repo}/subscription"},
		{"PUT", "/repos/{owner}/{repo}/subscription"},
		{"DELETE", "/repos/{owner}/{repo}/subscription"},
		{"GET", "/user/subscriptions/{owner}/{repo}"},
		{"PUT", "/user/subscriptions/{owner}/{repo}"},
		{"DELETE", "/user/subscriptions/{owner}/{repo}"},

		// Gists
		{"GET", "/users/{user}/gists"},
		{"GET", "/gists"},
		//{"GET", "/gists/public"},
		//{"GET", "/gists/starred"},
		{"GET", "/gists/{id}"},
		{"POST", "/gists"},
		//{"PATCH", "/gists/{id}"},
		{"PUT", "/gists/{id}/star"},
		{"DELETE", "/gists/{id}/star"},
		{"GET", "/gists/{id}/star"},
		{"POST", "/gists/{id}/forks"},
		{"DELETE", "/gists/{id}"},

		// Git Data
		{"GET", "/repos/{owner}/{repo}/git/blobs/{sha}"},
		{"POST", "/repos/{owner}/{repo}/git/blobs"},
		{"GET", "/repos/{owner}/{repo}/git/commits/{sha}"},
		{"POST", "/repos/{owner}/{repo}/git/commits"},
		//{"GET", "/repos/{owner}/{repo}/git/refs/*ref"},
		{"GET", "/repos/{owner}/{repo}/git/refs"},
		{"POST", "/repos/{owner}/{repo}/git/refs"},
		//{"PATCH", "/repos/{owner}/{repo}/git/refs/*ref"},
		//{"DELETE", "/repos/{owner}/{repo}/git/refs/*ref"},
		{"GET", "/repos/{owner}/{repo}/git/tags/{sha}"},
		{"POST", "/repos/{owner}/{repo}/git/tags"},
		{"GET", "/repos/{owner}/{repo}/git/trees/{sha}"},
		{"POST", "/repos/{owner}/{repo}/git/trees"},

		// Issues
		{"GET", "/issues"},
		{"GET", "/user/issues"},
		{"GET", "/orgs/{org}/issues"},
		{"GET", "/repos/{owner}/{repo}/issues"},
		{"GET", "/repos/{owner}/{repo}/issues/{number}"},
		{"POST", "/repos/{owner}/{repo}/issues"},
		//{"PATCH", "/repos/{owner}/{repo}/issues/{number}"},
		{"GET", "/repos/{owner}/{repo}/assignees"},
		{"GET", "/repos/{owner}/{repo}/assignees/{assignee}"},
		{"GET", "/repos/{owner}/{repo}/issues/{number}/comments"},
		//{"GET", "/repos/{owner}/{repo}/issues/comments"},
		//{"GET", "/repos/{owner}/{repo}/issues/comments/{id}"},
		{"POST", "/repos/{owner}/{repo}/issues/{number}/comments"},
		//{"PATCH", "/repos/{owner}/{repo}/issues/comments/{id}"},
		//{"DELETE", "/repos/{owner}/{repo}/issues/comments/{id}"},
		{"GET", "/repos/{owner}/{repo}/issues/{number}/events"},
		//{"GET", "/repos/{owner}/{repo}/issues/events"},
		//{"GET", "/repos/{owner}/{repo}/issues/events/{id}"},
		{"GET", "/repos/{owner}/{repo}/labels"},
		{"GET", "/repos/{owner}/{repo}/labels/{name}"},
		{"POST", "/repos/{owner}/{repo}/labels"},
		//{"PATCH", "/repos/{owner}/{repo}/labels/{name}"},
		{"DELETE", "/repos/{owner}/{repo}/labels/{name}"},
		{"GET", "/repos/{owner}/{repo}/issues/{number}/labels"},
		{"POST", "/repos/{owner}/{repo}/issues/{number}/labels"},
		{"DELETE", "/repos/{owner}/{repo}/issues/{number}/labels/{name}"},
		{"PUT", "/repos/{owner}/{repo}/issues/{number}/labels"},
		{"DELETE", "/repos/{owner}/{repo}/issues/{number}/labels"},
		{"GET", "/repos/{owner}/{repo}/milestones/{number}/labels"},
		{"GET", "/repos/{owner}/{repo}/milestones"},
		{"GET", "/repos/{owner}/{repo}/milestones/{number}"},
		{"POST", "/repos/{owner}/{repo}/milestones"},
		//{"PATCH", "/repos/{owner}/{repo}/milestones/{number}"},
		{"DELETE", "/repos/{owner}/{repo}/milestones/{number}"},

		// Miscellaneous
		{"GET", "/emojis"},
		{"GET", "/gitignore/templates"},
		{"GET", "/gitignore/templates/{name}"},
		{"POST", "/markdown"},
		{"POST", "/markdown/raw"},
		{"GET", "/meta"},
		{"GET", "/rate_limit"},

		// Organizations
		{"GET", "/users/{user}/orgs"},
		{"GET", "/user/orgs"},
		{"GET", "/orgs/{org}"},
		//{"PATCH", "/orgs/{org}"},
		{"GET", "/orgs/{org}/members"},
		{"GET", "/orgs/{org}/members/{user}"},
		{"DELETE", "/orgs/{org}/members/{user}"},
		{"GET", "/orgs/{org}/public_members"},
		{"GET", "/orgs/{org}/public_members/{user}"},
		{"PUT", "/orgs/{org}/public_members/{user}"},
		{"DELETE", "/orgs/{org}/public_members/{user}"},
		{"GET", "/orgs/{org}/teams"},
		{"GET", "/teams/{id}"},
		{"POST", "/orgs/{org}/teams"},
		//{"PATCH", "/teams/{id}"},
		{"DELETE", "/teams/{id}"},
		{"GET", "/teams/{id}/members"},
		{"GET", "/teams/{id}/members/{user}"},
		{"PUT", "/teams/{id}/members/{user}"},
		{"DELETE", "/teams/{id}/members/{user}"},
		{"GET", "/teams/{id}/repos"},
		{"GET", "/teams/{id}/repos/{owner}/{repo}"},
		{"PUT", "/teams/{id}/repos/{owner}/{repo}"},
		{"DELETE", "/teams/{id}/repos/{owner}/{repo}"},
		{"GET", "/user/teams"},

		// Pull Requests
		{"GET", "/repos/{owner}/{repo}/pulls"},
		{"GET", "/repos/{owner}/{repo}/pulls/{number}"},
		{"POST", "/repos/{owner}/{repo}/pulls"},
		//{"PATCH", "/repos/{owner}/{repo}/pulls/{number}"},
		{"GET", "/repos/{owner}/{repo}/pulls/{number}/commits"},
		{"GET", "/repos/{owner}/{repo}/pulls/{number}/files"},
		{"GET", "/repos/{owner}/{repo}/pulls/{number}/merge"},
		{"PUT", "/repos/{owner}/{repo}/pulls/{number}/merge"},
		{"GET", "/repos/{owner}/{repo}/pulls/{number}/comments"},
		//{"GET", "/repos/{owner}/{repo}/pulls/comments"},
		//{"GET", "/repos/{owner}/{repo}/pulls/comments/{number}"},
		{"PUT", "/repos/{owner}/{repo}/pulls/{number}/comments"},
		//{"PATCH", "/repos/{owner}/{repo}/pulls/comments/{number}"},
		//{"DELETE", "/repos/{owner}/{repo}/pulls/comments/{number}"},

		// Repositories
		{"GET", "/user/repos"},
		{"GET", "/users/{user}/repos"},
		{"GET", "/orgs/{org}/repos"},
		{"GET", "/repositories"},
		{"POST", "/user/repos"},
		{"POST", "/orgs/{org}/repos"},
		{"GET", "/repos/{owner}/{repo}"},
		//{"PATCH", "/repos/{owner}/{repo}"},
		{"GET", "/repos/{owner}/{repo}/contributors"},
		{"GET", "/repos/{owner}/{repo}/languages"},
		{"GET", "/repos/{owner}/{repo}/teams"},
		{"GET", "/repos/{owner}/{repo}/tags"},
		{"GET", "/repos/{owner}/{repo}/branches"},
		{"GET", "/repos/{owner}/{repo}/branches/{branch}"},
		{"DELETE", "/repos/{owner}/{repo}"},
		{"GET", "/repos/{owner}/{repo}/collaborators"},
		{"GET", "/repos/{owner}/{repo}/collaborators/{user}"},
		{"PUT", "/repos/{owner}/{repo}/collaborators/{user}"},
		{"DELETE", "/repos/{owner}/{repo}/collaborators/{user}"},
		{"GET", "/repos/{owner}/{repo}/comments"},
		{"GET", "/repos/{owner}/{repo}/commits/{sha}/comments"},
		{"POST", "/repos/{owner}/{repo}/commits/{sha}/comments"},
		{"GET", "/repos/{owner}/{repo}/comments/{id}"},
		//{"PATCH", "/repos/{owner}/{repo}/comments/{id}"},
		{"DELETE", "/repos/{owner}/{repo}/comments/{id}"},
		{"GET", "/repos/{owner}/{repo}/commits"},
		{"GET", "/repos/{owner}/{repo}/commits/{sha}"},
		{"GET", "/repos/{owner}/{repo}/readme"},
		//{"GET", "/repos/{owner}/{repo}/contents/*path"},
		//{"PUT", "/repos/{owner}/{repo}/contents/*path"},
		//{"DELETE", "/repos/{owner}/{repo}/contents/*path"},
		//{"GET", "/repos/{owner}/{repo}/{archive_format}/{ref}"},
		{"GET", "/repos/{owner}/{repo}/keys"},
		{"GET", "/repos/{owner}/{repo}/keys/{id}"},
		{"POST", "/repos/{owner}/{repo}/keys"},
		//{"PATCH", "/repos/{owner}/{repo}/keys/{id}"},
		{"DELETE", "/repos/{owner}/{repo}/keys/{id}"},
		{"GET", "/repos/{owner}/{repo}/downloads"},
		{"GET", "/repos/{owner}/{repo}/downloads/{id}"},
		{"DELETE", "/repos/{owner}/{repo}/downloads/{id}"},
		{"GET", "/repos/{owner}/{repo}/forks"},
		{"POST", "/repos/{owner}/{repo}/forks"},
		{"GET", "/repos/{owner}/{repo}/hooks"},
		{"GET", "/repos/{owner}/{repo}/hooks/{id}"},
		{"POST", "/repos/{owner}/{repo}/hooks"},
		//{"PATCH", "/repos/{owner}/{repo}/hooks/{id}"},
		{"POST", "/repos/{owner}/{repo}/hooks/{id}/tests"},
		{"DELETE", "/repos/{owner}/{repo}/hooks/{id}"},
		{"POST", "/repos/{owner}/{repo}/merges"},
		{"GET", "/repos/{owner}/{repo}/releases"},
		{"GET", "/repos/{owner}/{repo}/releases/{id}"},
		{"POST", "/repos/{owner}/{repo}/releases"},
		//{"PATCH", "/repos/{owner}/{repo}/releases/{id}"},
		{"DELETE", "/repos/{owner}/{repo}/releases/{id}"},
		{"GET", "/repos/{owner}/{repo}/releases/{id}/assets"},
		{"GET", "/repos/{owner}/{repo}/stats/contributors"},
		{"GET", "/repos/{owner}/{repo}/stats/commit_activity"},
		{"GET", "/repos/{owner}/{repo}/stats/code_frequency"},
		{"GET", "/repos/{owner}/{repo}/stats/participation"},
		{"GET", "/repos/{owner}/{repo}/stats/punch_card"},
		{"GET", "/repos/{owner}/{repo}/statuses/{ref}"},
		{"POST", "/repos/{owner}/{repo}/statuses/{ref}"},

		// Search
		{"GET", "/search/repositories"},
		{"GET", "/search/code"},
		{"GET", "/search/issues"},
		{"GET", "/search/users"},
		{"GET", "/legacy/issues/search/{owner}/{repository}/{state}/{keyword}"},
		{"GET", "/legacy/repos/search/{keyword}"},
		{"GET", "/legacy/user/search/{keyword}"},
		{"GET", "/legacy/user/email/{email}"},

		// Users
		{"GET", "/users/{user}"},
		{"GET", "/user"},
		//{"PATCH", "/user"},
		{"GET", "/users"},
		{"GET", "/user/emails"},
		{"POST", "/user/emails"},
		{"DELETE", "/user/emails"},
		{"GET", "/users/{user}/followers"},
		{"GET", "/user/followers"},
		{"GET", "/users/{user}/following"},
		{"GET", "/user/following"},
		{"GET", "/user/following/{user}"},
		{"GET", "/users/{user}/following/{target_user}"},
		{"PUT", "/user/following/{user}"},
		{"DELETE", "/user/following/{user}"},
		{"GET", "/users/{user}/keys"},
		{"GET", "/user/keys"},
		{"GET", "/user/keys/{id}"},
		{"POST", "/user/keys"},
		//{"PATCH", "/user/keys/{id}"},
		{"DELETE", "/user/keys/{id}"},
	}
)

func benchmarkRoutes(b *testing.B, router http.Handler, useStaticRoutes bool) {
	var r *http.Request
	if useStaticRoutes {
		r = httptest.NewRequest("GET", "https://www.domain.com/gopher/pencil/gopherhelmet.jpg", nil)
	} else {
		r = httptest.NewRequest("GET", "https://www.domain.com/repos/owner/repo/commits/sha", nil)
	}
	w := httptest.NewRecorder()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		//we don't reset the headers because the header writing produces heap allocations and, it is out of the benchmark scope
		//w.Header().Del("Content-Type")
		//w.Header().Del("X-Content-Type-Options")
		w.Body.Reset()

		router.ServeHTTP(w, r)

		if w.Result().StatusCode != 200 {
			b.Fatalf("got %d for %s", w.Result().StatusCode, r.URL.Path)
		}
	}
}

func benchmarkRoutesConcurrent(b *testing.B, router http.Handler, useStaticRoutes bool) {
	var routes []*Route
	if useStaticRoutes {
		routes = staticRoutes
	} else {
		routes = varCaptureRoutes
	}

	routesLen := len(routes)
	requests := make([]*http.Request, routesLen)
	for i := 0; i < routesLen; i++ {
		requests[i] = httptest.NewRequest("GET", "https://www.domain.com/legacy/issues/search/owner/repository/state/keyword", nil)
	}

	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	b.ReportAllocs()
	b.SetParallelism(routesLen)
	b.RunParallel(func(pb *testing.PB) {
		w := httptest.NewRecorder()
		for pb.Next() {
			//we don't reset the headers because the header writing produces heap allocations and, it is out of the benchmark scope
			//w.Header().Del("Content-Type")
			//w.Header().Del("X-Content-Type-Options")
			w.Body.Reset()

			req := requests[rand.Intn(routesLen)]
			router.ServeHTTP(w, req)
			if w.Result().StatusCode != 200 {
				b.Fatalf("got %d for %s", w.Result().StatusCode, req.URL.Path)
			}
		}
	})
}

//----------------------------------------- GOFRE --------------------------------------

func gofreHandler() handler.Handler {
	return func(ctx context.Context, mc path.MatchingContext) (response.HttpResponse, error) {
		return response.PlainTextHttpResponseOK("ok"), nil
	}
}

func loadGofreRoutes(g *gofre.MuxHandler, useStaticRoutes bool) {
	var routes []*Route
	if useStaticRoutes {
		routes = staticRoutes
	} else {
		routes = varCaptureRoutes
	}

	for _, r := range routes {
		path := r.Path
		switch r.Method {
		case "GET":
			g.HandleGet(path, gofreHandler())
		case "POST":
			g.HandlePost(path, gofreHandler())
		case "PATCH":
			g.HandlePatch(path, gofreHandler())
		case "PUT":
			g.HandlePut(path, gofreHandler())
		case "DELETE":
			g.HandleDelete(path, gofreHandler())
		}
	}
}

func Benchmark_GofreStatic(b *testing.B) {
	gm, _ := gofre.NewMuxHandlerWithDefaultConfig()
	loadGofreRoutes(gm, true)
	benchmarkRoutes(b, gm, true)
}

func Benchmark_GofreVarCapture(b *testing.B) {
	gm, _ := gofre.NewMuxHandlerWithDefaultConfig()
	loadGofreRoutes(gm, false)
	benchmarkRoutes(b, gm, false)
}

func Benchmark_GofreVarCapture_Concurrent(b *testing.B) {
	gm, _ := gofre.NewMuxHandlerWithDefaultConfig()
	loadGofreRoutes(gm, false)
	benchmarkRoutesConcurrent(b, gm, false)
}

//----------------------------------------- ECHO --------------------------------------

func loadEchoRoutes(e *echo.Echo, useStaticRoutes bool) {
	var routes []*Route
	if useStaticRoutes {
		routes = staticRoutes
	} else {
		routes = varCaptureRoutes
	}

	for _, r := range routes {
		path := strings.ReplaceAll(strings.ReplaceAll(r.Path, "/{", "/:"), "}", "")
		switch r.Method {
		case "GET":
			e.GET(path, echoHandler(path))
		case "POST":
			e.POST(path, echoHandler(path))
		case "PATCH":
			e.PATCH(path, echoHandler(path))
		case "PUT":
			e.PUT(path, echoHandler(path))
		case "DELETE":
			e.DELETE(path, echoHandler(path))
		}
	}
}

func echoHandler(path string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}
}

func Benchmark_EchoStatic(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, true)
	benchmarkRoutes(b, e, true)
}

func Benchmark_EchoVarCapture(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, false)
	benchmarkRoutes(b, e, false)
}

func Benchmark_EchoVarCapture_Concurrent(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, false)
	benchmarkRoutesConcurrent(b, e, false)
}

//----------------------------------------- GIN --------------------------------------

func loadGinRoutes(g *gin.Engine, useStaticRoutes bool) {
	var routes []*Route
	if useStaticRoutes {
		routes = staticRoutes
	} else {
		routes = varCaptureRoutes
	}

	for _, r := range routes {
		path := strings.ReplaceAll(strings.ReplaceAll(r.Path, "/{", "/:"), "}", "")
		switch r.Method {
		case "GET":
			g.GET(path, ginHandler(path))
		case "POST":
			g.POST(path, ginHandler(path))
		case "PATCH":
			g.PATCH(path, ginHandler(path))
		case "PUT":
			g.PUT(path, ginHandler(path))
		case "DELETE":
			g.DELETE(path, ginHandler(path))
		}
	}
}

func ginHandler(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}
}

func Benchmark_GinStatic(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, true)
	benchmarkRoutes(b, g, true)
}

func Benchmark_GinVarCapture(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, false)
	benchmarkRoutes(b, g, false)
}

func Benchmark_GinVarCapture_Concurrent(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, false)
	benchmarkRoutesConcurrent(b, g, false)
}

//----------------------------------------- GORILLA --------------------------------------

func loadGorillaRoutes(g *mux.Router, useStaticRoutes bool) {
	var routes []*Route
	if useStaticRoutes {
		routes = staticRoutes
	} else {
		routes = varCaptureRoutes
	}

	for _, r := range routes {
		path := r.Path
		switch r.Method {
		case "GET":
			g.HandleFunc(path, gorillaHandler()).Methods("GET")
		case "POST":
			g.HandleFunc(path, gorillaHandler()).Methods("POST")
		case "PATCH":
			g.HandleFunc(path, gorillaHandler()).Methods("PATCH")
		case "PUT":
			g.HandleFunc(path, gorillaHandler()).Methods("PUT")
		case "DELETE":
			g.HandleFunc(path, gorillaHandler()).Methods("DELETE")
		}
	}
}

func gorillaHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

func Benchmark_GorillaStatic(b *testing.B) {
	g := mux.NewRouter()
	loadGorillaRoutes(g, true)
	benchmarkRoutes(b, g, true)
}

func Benchmark_GorillaVarCapture(b *testing.B) {
	g := mux.NewRouter()
	loadGorillaRoutes(g, false)
	benchmarkRoutes(b, g, false)
}

func Benchmark_GorillaVarCapture_Concurrent(b *testing.B) {
	g := mux.NewRouter()
	loadGorillaRoutes(g, false)
	benchmarkRoutesConcurrent(b, g, false)
}
