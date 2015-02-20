{-# LANGUAGE OverloadedStrings #-}
module Databrary.Controller.Login
  ( postLogin
  , viewLogin
  ) where

import qualified Blaze.ByteString.Builder as Blaze
import Control.Monad (when, unless)
import Control.Monad.Trans.Class (lift)
import qualified Crypto.BCrypt as BCrypt
import qualified Data.Foldable as Fold
import Data.Maybe (fromMaybe)
import Data.Monoid (mempty)
import qualified Data.Text.Encoding as TE
import qualified Text.Blaze.Html5 as Html

import Control.Applicative.Ops
import Control.Has (view)
import Databrary.Action
import Databrary.Web.Form
import Databrary.Web.Cookie
import Databrary.Model.Id.Types
import Databrary.Model.Party
import Databrary.Model.Permission
import Databrary.Model.Token
import Databrary.Web.Form.Deform
import Databrary.Controller.Form

loginAccount :: SiteAuth -> Bool -> AppAction
loginAccount auth su = do
  sess <- createSession auth su
  let Token (Id tok) ex = view sess
  cook <- setSignedCookie "session" tok ex
  okResponse [cook] (mempty :: Blaze.Builder)

displayLogin :: FormData -> FormErrors -> Html.Html
displayLogin fd fe = undefined

viewLogin :: AppRAction
viewLogin = action GET ["login"] $
  okResponse [] $ displayLogin mempty mempty

postLogin :: Bool -> AppRAction
postLogin api = action POST (apiRoute api ["login"]) $ do
  (Just auth, su) <- runForm (api ?!> displayLogin) $ do
    email <- "email" .:> emailTextForm
    password <- "password" .:> deform
    superuser <- "superuser" .:> deform
    auth <- lift $ lookupSiteAuthByEmail email
    let p = view <$> auth
        a = partyAccount =<< p
        su = superuser && Fold.any ((PermissionADMIN ==) . view) auth
    attempts <- lift $ maybe (return 0) recentAccountLogins p
    let pass = maybe False (flip BCrypt.validatePassword (TE.encodeUtf8 password)) (accountPasswd =<< a)
        block = attempts > 4
    lift $ auditAccountLogin pass (fromMaybe nobodyParty p) email
    when block $ "email" .:> deformError "Too many login attempts. Try again later."
    unless pass $ "password" .:> deformError "Incorrect login."
    return (auth, su)
  loginAccount auth su