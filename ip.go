package gokit

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetIp(ctx *fiber.Ctx) string {
	ips := ctx.IPs()
	if len(ips) == 0 {
		return ""
	}
	return strings.Join(ips, ",")
}
