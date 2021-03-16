package main

import (
	"github.com/a-ok123/go-psl/internal/common"
	"github.com/a-ok123/go-psl/internal/pastelclient"
	"github.com/labstack/echo/v4"
)

type TicketProc struct {
	PslNode pastelclient.PslNode
	config  common.Config
	logger  common.Logger
}

func (p *TicketProc) Init(app *common.Application) {
	p.config = app.Cfg
	p.logger = app.Log
}

func (p *TicketProc) RegisterArtTicket(c echo.Context) error {
	/*conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}

	go func() {
		defer conn.Close()

		wg.Add(2)
		eg, egCtx := errgroup.WithContext(context.Background())

		results := make(chan int)
		eg.Go( func() error {
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					// handle error
				}
				select {
				case results <- current:
					return nil
				// Close out if another error occurs.
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
		go func() {
			g.Wait()
			close(results)
		}()

		for result := range results {
			fmt.Println("processed", result)
		}

		// Wait for all goroutines to complete.
		eg.Wait()
	}()*/
	return nil
}