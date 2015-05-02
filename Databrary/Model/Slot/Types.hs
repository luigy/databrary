{-# LANGUAGE OverloadedStrings, TemplateHaskell, TypeFamilies #-}
module Databrary.Model.Slot.Types
  ( SlotId(..)
  , Slot(..)
  , slotId
  , containerSlotId
  , containerSlot
  , MonadHasSlotId
  , MonadHasSlot
  ) where

import Databrary.Has (makeHasRec)
import Databrary.Model.Id
import Databrary.Model.Kind
import Databrary.Model.Segment
import Databrary.Model.Container.Types

data SlotId = SlotId
  { slotContainerId :: !(Id Container)
  , slotSegmentId :: !Segment
  } deriving (Show)

type instance IdType Slot = SlotId

containerSlotId :: Id Container -> Id Slot
containerSlotId c = Id $ SlotId c fullSegment

data Slot = Slot
  { slotContainer :: !Container
  , slotSegment :: !Segment
  }

slotId :: Slot -> Id Slot
slotId (Slot c s) = Id $ SlotId (containerId c) s

containerSlot :: Container -> Slot
containerSlot c = Slot c fullSegment

instance Kinded Slot where
  kindOf _ = "slot"

makeHasRec ''SlotId ['slotContainerId, 'slotSegmentId]
makeHasRec ''Slot ['slotContainer, 'slotSegment]