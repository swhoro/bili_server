package router

import (
	"math/rand"
	"strings"
	"time"

	"github.com/goccy/go-json"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	g "b.carriage.fun/server/global"
	response "b.carriage.fun/server/response/error"
	v1logic "b.carriage.fun/server/router/api/v1"
)

func RegisterRouter() {
	// 生成secret
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sb := strings.Builder{}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	sb.Grow(n)
	l := len(charset)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(l)])
	}
	g.Secret = sb.String()

	jwtMiddle := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(g.Secret)},
		ContextKey:  "jwt_token",
		TokenLookup: "cookie:session",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.ReturnWithNotLogin(c)
		},
	})

	g.App = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := g.App.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	bangumi := v1.Group("/bangumi")

	user.Post("/add", v1logic.AddUser)
	user.Post("/login", v1logic.Login)
	user.Get("/restricted", jwtMiddle, v1logic.GetRestricted)
	user.Put("/modify", jwtMiddle, v1logic.ModifyUser)

	bangumi.Post("/add", jwtMiddle, v1logic.AddBangumi)
	bangumi.Get("/all", v1logic.GetAllBangumi)
}
