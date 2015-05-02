{-# LANGUAGE TypeOperators, QuasiQuotes #-}
module Databrary.Model.Id 
  ( module Databrary.Model.Id.Types
  , idIso
  , pathId
  ) where

import qualified Databrary.Iso as I
import Databrary.Iso.TH
import Databrary.HTTP.Route.Path
import Databrary.HTTP.Route.PathParser
import Databrary.Model.Kind
import Databrary.Model.Id.Types

idIso :: IdType a I.<-> Id a
idIso = [iso|a <-> Id a|]

pathId :: forall a . (PathDynamic (IdType a), Kinded a) => PathParser (Id a)
pathId = PathFixed (kindOf (undefined :: a)) >/> idIso I.<$> PathDynamic