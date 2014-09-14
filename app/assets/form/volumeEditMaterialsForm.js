'use strict';

module.directive('volumeEditMaterialsForm', [
  'pageService', function (page) {
    var link = function ($scope) {
      var volume = $scope.volume;
      var slot = volume.top;
      var form = $scope.volumeEditMaterialsForm;

      function materialData(asset) {
	return {
	  name: asset.asset.name,
	  classification: asset.asset.classification+'',
	  excerptOn: asset.excerpt !== undefined,
	  excerpt: (asset.excerpt || 0)+''
	};
      }

      form.materials = slot.assets.map(function (asset) {
	asset.asset.get(['creation']);
	return {
	  asset: asset,
	  data: materialData(asset)
	};
      });

      function materialName(material) {
	return material.file && material.file.file.name || material.asset && material.asset.name || page.constants.message('file');
      }

      form.fileAdded = function (file, event) {
	var material = angular.element(event.srcElement).scope().material;

	if (material)
	  delete material.replace;
	else {
	  material = {
	    data: {
	      classification: page.classification.RESTRICTED+'',
	    },
	  };
	  form.materials.push(material);
	  /* need a better scroll target: */
	  page.display.scrollTo(page.$d.height());
	}
	material.file = file;
	material.progress = 0;
	file.material = material;

	page.assets.assetStart(file).then(function () {
	  file.resume();
	}, function (res) {
	  form.messages.addError({
	    type: 'red',
	    body: page.constants('asset.upload.rejected', materialName(material)), 
	    report: res,
	  });
	  delete material.file;
	});
      };

      form.fileSuccess = function (file) {
	form.save(file.material);
	form.totalProgress();
      };

      form.fileProgress = function (file) {
        file.material.progress = file.progress();
      };

      form.totalProgress = function () {
        form.progress = $scope.$flow.progress();
      };

      form.excerptOptions = function (material) {
	var l = {};
	for (var i = page.constants.classification.length-1; i > material.data.classification; i --)
	  l[i] = page.constants.classification[i];
	l[0] = page.constants.classification[0];
        return l;
      };

      form.save = function (material) {
	if (!material.data.excerptOn)
	  material.data.excerpt = '';

	var act;
	if (material.file) {
	  material.data.upload = material.file.uniqueIdentifier;
	  act = material.asset ? material.asset.replace(material.data) : slot.createAsset(material.data);
	} else
	  act = material.asset.asset.save(material.data);

	act.then(function (asset) {
	  if (asset instanceof page.models.SlotAsset)
	    material.asset = asset;
	  else
	    material.asset.asset = asset;
	  material.data = materialData(material.asset);
	  if (material.file) {
	    if (!('creation' in material.asset.asset))
	      material.asset.asset.creation = {date: Date.now(), name: material.file.file.name}; 
	    material.file.cancel();
	    delete material.file;
	  }

	  form.messages.add({
	    type: 'green',
	    countdown: 3000,
	    body: page.constants.message('volume.edit.materials.update.success', materialName(material)),
	  });

	  material.form.$setPristine();
	}, function (res) {
	  form.messages.addError({
	    type: 'red',
	    body: page.constants.message('volume.edit.materials.update.error', materialName(material)),
	    report: res,
	  });
	});
      };

      form.saveAll = function () {
	form.materials.forEach(function (material) {
	  if (material.form.$dirty)
	    form.save(material);
        });
      };

      form.replace = function (material) {
	material.replace = true;
      };

      form.remove = function (material) {
	if (material.replace) {
	  delete material.replace;
	  return;
	}
	if (material.file) {
	  material.file.cancel();
	  delete material.file;
	  if (!material.asset)
	    form.materials.splice(form.materials.indexOf(material), 1);
	  return;
	}

	material.asset.remove().then(function () {
	  form.messages.add({
	    type: 'green',
	    countdown: 3000,
	    body: page.constants.message('volume.edit.materials.remove.success', material.asset.name || page.constants.message('file')),
	  });

	  form.materials.splice(form.materials.indexOf(material), 1);
	}, function (res) {
	  form.messages.addError({
	    type: 'red',
	    body: page.constants.message('volume.edit.materials.remove.error', material.asset.name || page.constants.message('file')),
	    report: res,
	  });
	});
      };

      var $float = $('.vem-float');
      var $floater = $('.vem-float-floater');
      form.scrollFn = page.display.makeFloatScrollFn($float, $floater, 24*1.5);
      page.$w.scroll(form.scrollFn);
    };

    //

    return {
      restrict: 'E',
      templateUrl: 'volumeEditMaterialsForm.html',
      replace: true,
      link: link
    };
  }
]);
