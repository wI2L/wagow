// Copyright (c) 2017, William Poussier.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/wI2L/wagow"
)

func main() {
	// Create echo instance and apply
	// middlewares before and after router.
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Logger.SetLevel(log.INFO)

	e.POST("/", handler)

	// Start...
	go func() {
		a := net.JoinHostPort("0.0.0.0", os.Getenv("PORT"))
		if err := e.Start(a); err != nil {
			e.Logger.Info("shutting down server")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// Wait for interrupt signal to gracefully shutdown
	// the server with a timeout of 10 seconds.
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Request represents a wake-on-wan request.
type Request struct {
	Address  string `json:"address" form:"address"`
	Target   string `json:"target" form:"target"`
	Password string `json:"password" form:"password"`
}

func handler(c echo.Context) error {
	var req Request

	// Bind the parameters to a Request object.
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	// Parse the target MAC address.
	mac, err := net.ParseMAC(req.Target)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid target parameter")
	}
	// Create a UDP client.
	client, err := wagow.NewUDPClient()
	if err != nil {
		return err
	}
	// Send magic packet.
	err = client.Wake(req.Address, mac, req.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "magic packet successfully sent")
}
