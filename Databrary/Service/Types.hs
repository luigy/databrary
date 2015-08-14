{-# LANGUAGE OverloadedStrings, TemplateHaskell, RecordWildCards #-}
module Databrary.Service.Types
  ( Secret(..)
  , Service(..)
  , MonadHasService
  ) where

import qualified Data.ByteString as BS

import Databrary.Has (makeHasRec)
import Databrary.Service.DB (DBConn)
import Databrary.Service.Entropy (Entropy)
import Databrary.HTTP.Client (HTTPClient)
import Databrary.Store.Types (Storage)
import Databrary.Store.AV (AV)
import Databrary.Service.Passwd (Passwd)
import Databrary.Service.Log (Logs)
import Databrary.Service.Messages (Messages)
import Databrary.Web.Types (Web)
import Databrary.Static.Service (Static)
import Databrary.Search.Service (Solr)
import Databrary.Ingest.Service (Ingest)
import Databrary.Model.Time

newtype Secret = Secret BS.ByteString

data Service = Service
  { serviceStartTime :: !Timestamp
  , serviceSecret :: !Secret
  , serviceEntropy :: !Entropy
  , servicePasswd :: !Passwd
  , serviceLogs :: !Logs
  , serviceMessages :: !Messages
  , serviceDB :: !DBConn
  , serviceStorage :: !Storage
  , serviceAV :: !AV
  , serviceWeb :: !Web
  , serviceHTTPClient :: !HTTPClient
  , serviceStatic :: !Static
  , serviceSolr :: !Solr
  , serviceIngest :: !Ingest
  }

makeHasRec ''Service ['serviceSecret, 'serviceEntropy, 'servicePasswd, 'serviceLogs, 'serviceMessages, 'serviceDB, 'serviceStorage, 'serviceAV, 'serviceWeb, 'serviceHTTPClient, 'serviceStatic, 'serviceSolr, 'serviceIngest]
