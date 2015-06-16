{-# LANGUAGE TemplateHaskell, OverloadedStrings, DeriveDataTypeable #-}
{-# OPTIONS_GHC -fno-warn-orphans #-}
module Databrary.Model.Release.Types 
  ( Release(..)
  ) where

import Data.Maybe (fromMaybe)
import Language.Haskell.TH.Lift (deriveLift)

import Databrary.Has (Has(..))
import Databrary.Service.DB
import Databrary.Model.Enum

useTPG

makeDBEnum "release" "Release"

instance Has Release (Maybe Release) where
  view = fromMaybe ReleasePRIVATE

deriveLift ''Release