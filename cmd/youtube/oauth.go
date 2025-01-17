package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func authConstruct(exc *youtube.Exchange) error {
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   fil, err := os.Open(cac + "/mech/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   return json.NewDecoder(fil).Decode(exc)
}

func authExchange() error {
   oau, err := youtube.NewOAuth()
   if err != nil {
      return err
   }
   fmt.Printf(`1. Go to
%v

2. Enter this code
%v

3. Press Enter to continue`, oau.Verification_URL, oau.User_Code)
   fmt.Scanln()
   exc, err := oau.Exchange()
   if err != nil {
      return err
   }
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   cac += "/mech"
   os.Mkdir(cac, os.ModeDir)
   fil, err := os.Create(cac + "/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   enc := json.NewEncoder(fil)
   enc.SetIndent("", " ")
   return enc.Encode(exc)
}

func authRefresh() error {
   var exc youtube.Exchange
   err := authConstruct(&exc)
   if err != nil {
      return err
   }
   if err := exc.Refresh(); err != nil {
      return err
   }
   cac, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   fil, err := os.Create(cac + "/mech/youtube.json")
   if err != nil {
      return err
   }
   defer fil.Close()
   enc := json.NewEncoder(fil)
   enc.SetIndent("", " ")
   return enc.Encode(exc)
}
