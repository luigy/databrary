{-# LANGUAGE OverloadedStrings, TemplateHaskell, QuasiQuotes #-}
module Databrary.Model.Token
  ( module Databrary.Model.Token.Types
  , loginTokenId
  , lookupLoginToken
  , createLoginToken
  , removeLoginToken
  , lookupSession
  , createSession
  , removeSession
  , lookupUpload
  , createUpload
  , removeUpload
  , cleanTokens
  ) where

import Control.Monad (when, void, (<=<))
import Control.Monad.IO.Class (MonadIO, liftIO)
import qualified Data.ByteString as BS
import qualified Data.ByteString.Base64.URL as Base64
import Data.Int (Int64)
import Database.PostgreSQL.Typed (pgSQL)
import Database.PostgreSQL.Typed.Query (simpleQueryFlags)

import Databrary.Ops
import Databrary.Has (view, peek, peeks)
import Databrary.Files (removeFile)
import Databrary.Service.Types
import Databrary.Service.Entropy
import Databrary.Service.Crypto
import Databrary.Service.DB
import Databrary.Store.Types
import Databrary.Store.Upload
import Databrary.Model.SQL (selectQuery)
import Databrary.Model.SQL.Select (makeQuery, selectOutput)
import Databrary.Model.Offset
import Databrary.Model.Id.Types
import Databrary.Model.Identity.Types
import Databrary.Model.Volume.Types
import Databrary.Model.Party
import Databrary.Model.Token.Types
import Databrary.Model.Token.SQL

useTPG

loginTokenId :: (MonadHasService c m, MonadIO m) => LoginToken -> m (Id LoginToken)
loginTokenId tok = Id <$> sign (unId (view tok :: Id Token))

lookupLoginToken :: (MonadDB m, MonadHasService c m) => Id LoginToken -> m (Maybe LoginToken)
lookupLoginToken =
  flatMapM (\t -> dbQuery1 $(selectQuery selectLoginToken "$!WHERE login_token.token = ${t} AND expires > CURRENT_TIMESTAMP"))
    <=< unSign . unId

lookupSession :: MonadDB m => BS.ByteString -> m (Maybe Session)
lookupSession tok =
  dbQuery1 $(selectQuery selectSession "$!WHERE session.token = ${tok} AND expires > CURRENT_TIMESTAMP")

lookupUpload :: (MonadDB m, MonadHasIdentity c m) => BS.ByteString -> m (Maybe Upload)
lookupUpload tok = do
  auth <- peek
  dbQuery1 $ fmap ($ auth) $(selectQuery selectUpload "$!WHERE upload.token = ${tok} AND expires > CURRENT_TIMESTAMP AND upload.account = ${view auth :: Id Party}")

createToken :: (MonadHasService c m, MonadDB m) => (Id Token -> DBM a) -> m a
createToken insert = do
  e <- peek
  let loop = do
        tok <- liftIO $ Id . Base64.encode <$> entropyBytes 24 e
        r <- dbQuery1 [pgSQL|SELECT token FROM token WHERE token = ${tok}|]
        case r `asTypeOf` Just tok of
          Nothing -> insert tok
          Just _ -> loop
  dbTransaction $ do
    _ <- dbExecuteSimple "LOCK TABLE token IN SHARE ROW EXCLUSIVE MODE"
    loop

createLoginToken :: (MonadHasService c m, MonadDB m) => SiteAuth -> Bool -> m LoginToken
createLoginToken auth passwd = do
  when passwd $ void $ dbExecute [pgSQL|DELETE FROM login_token WHERE account = ${view auth :: Id Party} AND password|]
  (tok, ex) <- createToken $ \tok ->
    dbQuery1' [pgSQL|INSERT INTO login_token (token, account, password) VALUES (${tok}, ${view auth :: Id Party}, ${passwd}) RETURNING token, expires|]
  return $ LoginToken
    { loginAccountToken = AccountToken
      { accountToken = Token tok ex
      , tokenAccount = auth
      }
    , loginPasswordToken = passwd
    }

sessionDuration :: Bool -> Offset
sessionDuration False = 7*24*60*60
sessionDuration True = 30*60

createSession :: (MonadHasService c m, MonadDB m) => SiteAuth -> Bool -> m Session
createSession auth su = do
  e <- peek
  (tok, ex, verf) <- createToken $ \tok -> do
    verf <- liftIO $ Base64.encode <$> entropyBytes 12 e
    dbQuery1' [pgSQL|INSERT INTO session (token, expires, account, superuser, verf) VALUES (${tok}, CURRENT_TIMESTAMP + ${sessionDuration su}::interval, ${view auth :: Id Party}, ${su}, ${verf}) RETURNING token, expires, verf|]
  return $ Session
    { sessionAccountToken = AccountToken
      { accountToken = Token tok ex
      , tokenAccount = auth
      }
    , sessionSuperuser = su
    , sessionVerf = verf
    }

createUpload :: (MonadHasService c m, MonadDB m, MonadHasIdentity c m) => Volume -> BS.ByteString -> Int64 -> m Upload
createUpload vol name size = do
  auth <- peek
  (tok, ex) <- createToken $ \tok ->
    dbQuery1' [pgSQL|INSERT INTO upload (token, account, volume, filename, size) VALUES (${tok}, ${view auth :: Id Party}, ${volumeId vol}, ${name}, ${size}) RETURNING token, expires|]
  return $ Upload
    { uploadAccountToken = AccountToken
      { accountToken = Token tok ex
      , tokenAccount = auth
      }
    , uploadFilename = name
    , uploadSize = size
    }

removeLoginToken :: MonadDB m => LoginToken -> m Bool
removeLoginToken tok =
  dbExecute1 [pgSQL|DELETE FROM login_token WHERE token = ${view tok :: Id Token}|]

removeSession :: (MonadDB m) => Session -> m Bool
removeSession tok =
  dbExecute1 [pgSQL|DELETE FROM session WHERE token = ${view tok :: Id Token}|]

removeUploadFile :: (MonadStorage c m) => Upload -> m Bool
removeUploadFile tok = liftIO . removeFile =<< peeks (uploadFile tok)

removeUpload :: (MonadDB m, MonadStorage c m) => Upload -> m Bool
removeUpload tok = do
  r <- dbExecute1 [pgSQL|DELETE FROM upload WHERE token = ${view tok :: Id Token}|]
  when r $ void $ removeUploadFile tok
  return r

cleanTokens :: (MonadDB m, MonadStorage c m) => m ()
cleanTokens = do
  toks <- dbQuery $ ($ nobodySiteAuth) <$> $(makeQuery simpleQueryFlags ("DELETE FROM upload WHERE expires < CURRENT_TIMESTAMP RETURNING " ++) (selectOutput selectUpload))
  mapM_ removeUploadFile toks
  dbExecute_ "DELETE FROM token WHERE expires < CURRENT_TIMESTAMP"
