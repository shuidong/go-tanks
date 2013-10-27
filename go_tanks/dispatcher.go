package go_tanks

import (
  "log"
  controllers "./controllers"
)

type Dispatcher struct {
  World   *World
  Client  *Client
}

func ( d *Dispatcher ) run () {
  d.dispatch()
}

func ( d * Dispatcher ) dispatch () {
    controller := controllers.AuthController{Client: d.Client}
    err := controller.Authorize()
    if ( err != nil ) { 
      log.Println("ERROR:\t", err)
      d.Client.Close();
      return
    } 
    //controller := GameController{Client: d.Client}
    //controller.run()
}
