{-# LANGUAGE OverloadedStrings, TemplateHaskell, QuasiQuotes, RecordWildCards #-}
module Databrary.Model.AssetSlot
  ( module Databrary.Model.AssetSlot.Types
  , lookupAssetSlot
  , lookupAssetAssetSlot
  , lookupSlotAssets
  , lookupContainerAssets
  , lookupVolumeAssetSlots
  , changeAssetSlot
  , changeAssetSlotDuration
  , fixAssetSlotDuration
  , findAssetContainerEnd
  , assetSlotJSON
  ) where

import Control.Applicative ((<*>))
import Control.Monad (when)
import Data.Maybe (fromMaybe, isNothing, catMaybes)
import Database.PostgreSQL.Typed (pgSQL)

import Databrary.Ops
import Databrary.Has (peek, view)
import qualified Databrary.JSON as JSON
import Databrary.Service.DB
import Databrary.Model.Offset
import Databrary.Model.Permission
import Databrary.Model.Segment
import Databrary.Model.Id
import Databrary.Model.Party.Types
import Databrary.Model.Identity.Types
import Databrary.Model.Volume.Types
import Databrary.Model.Container.Types
import Databrary.Model.Slot.Types
import Databrary.Model.Asset
import Databrary.Model.Audit
import Databrary.Model.SQL
import Databrary.Model.AssetSlot.Types
import Databrary.Model.AssetSlot.SQL

lookupAssetSlot :: (MonadHasIdentity c m, MonadDB m) => Id Asset -> m (Maybe AssetSlot)
lookupAssetSlot ai = do
  ident <- peek
  dbQuery1 $(selectQuery (selectAssetSlot 'ident) "$WHERE asset.id = ${ai}")

lookupAssetAssetSlot :: (MonadDB m) => Asset -> m AssetSlot
lookupAssetAssetSlot a = fromMaybe assetNoSlot
  <$> dbQuery1 $(selectQuery selectAssetSlotAsset "$WHERE slot_asset.asset = ${assetId a} AND container.volume = ${volumeId $ assetVolume a}")
  <*> return a

lookupSlotAssets :: (MonadDB m) => Slot -> m [AssetSlot]
lookupSlotAssets (Slot c s) =
  dbQuery $ ($ c) <$> $(selectQuery selectContainerSlotAsset "$WHERE slot_asset.container = ${containerId c} AND slot_asset.segment && ${s} AND asset.volume = ${volumeId $ containerVolume c}")

lookupContainerAssets :: (MonadDB m) => Container -> m [AssetSlot]
lookupContainerAssets = lookupSlotAssets . containerSlot

lookupVolumeAssetSlots :: (MonadDB m) => Volume -> Bool -> m [AssetSlot]
lookupVolumeAssetSlots v top =
  dbQuery $ ($ v) <$> $(selectQuery selectVolumeSlotAsset "$WHERE asset.volume = ${volumeId v} AND (container.top OR ${not top}) ORDER BY container.id")

changeAssetSlot :: (MonadAudit c m) => AssetSlot -> m Bool
changeAssetSlot as = do
  ident <- getAuditIdentity
  if isNothing (assetSlot as)
    then dbExecute1 $(deleteSlotAsset 'ident 'as)
    else do
      (r, _) <- updateOrInsert
        $(updateSlotAsset 'ident 'as)
        $(insertSlotAsset 'ident 'as)
      when (r /= 1) $ fail $ "changeAssetSlot: " ++ show r ++ " rows"
      return True

changeAssetSlotDuration :: MonadDB m => Asset -> m Bool
changeAssetSlotDuration a
  | Just dur <- assetDuration a =
    dbExecute1 [pgSQL|UPDATE slot_asset SET segment = segment(lower(segment), lower(segment) + ${dur}) WHERE asset = ${assetId a}|]
  | otherwise = return False

fixAssetSlotDuration :: AssetSlot -> AssetSlot
fixAssetSlotDuration as 
  | Just dur <- assetDuration (slotAsset as) = as{ assetSlot = (\s -> s{ slotSegment = segmentSetDuration dur (slotSegment s) }) <$> assetSlot as }
  | otherwise = as

findAssetContainerEnd :: MonadDB m => Container -> m (Maybe Offset)
findAssetContainerEnd c = 
  dbQuery1' [pgSQL|SELECT max(upper(segment)) FROM slot_asset WHERE container = ${containerId c}|]

assetSlotJSON :: AssetSlot -> JSON.Object
assetSlotJSON as@AssetSlot{..} = assetJSON slotAsset JSON..++ catMaybes
  [ segmentJSON . slotSegment =<< assetSlot
  , Just $ "permission" JSON..= p
  , p > PermissionNONE ?> "size" JSON..= assetSize slotAsset
  ] where p = dataPermission as
