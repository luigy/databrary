'use strict';

app.controller('party/profile', [
  '$scope', 'displayService', 'party', 'pageService','constantService',
  function ($scope, display, party, page, constants) {
    var getUsers = function(volumes){

      var users = {};
      
      users.sponsors = [];
      users.labGroupMembers = [];
      users.nonGroupAffiliates = [];
      users.otherCollaborators = [];

      _(volumes).pluck('access').flatten().uniq(function(i){
        return i.party.id; 
      }).each(function(v){
        console.log("V: ", v); 
        if(v.children > 0 ){
          users.sponsors.push(v);
        } else if(v.parents && v.parents.length) {
          users.labGroupMembers.push(v);
        } else if(v.parents && v.parents.length) {
          users.nonGroupAffiliates.push(v);
        } else {
          users.otherCollaborators.push(v);
        }

      }).value();
      // The "value()" call is to actually force the chain to work.

      
      console.log("Users: ", users);
      return users;
    };

    // This is a quick helper function to make sure that the isSelected
    // class is set to empty and to avoid repeating code. 
    var unSetSelected = function(v){
      v.isSelected = '';
      if(v.v !== undefined){
        v.v = _.map(v.v, function(a){
          a.isSelected = '';
          return a; 
        });
      }         

      return v; 
    };
      
    $scope.unselectAll = function(){

      $scope.volumes.individual = _.map($scope.volumes.individual, unSetSelected);

      $scope.volumes.collaborator = _.map($scope.volumes.collaborator, unSetSelected);

      $scope.volumes.inherited = _.map($scope.volumes.inherited, unSetSelected);

      $scope.users.sponsors = _.map($scope.users.sponsors, unSetSelected);

      $scope.users.nonGroupAffiliates = _.map($scope.users.nonGroupAffiliates, unSetSelected);

      $scope.users.labGroupMembers = _.map($scope.users.labGroupMembers, unSetSelected);

      $scope.users.otherCollaborators = _.map($scope.users.otherCollaborators, unSetSelected);

    }; 
    
    $scope.clickVolume = function(volume) {

      $scope.unselectAll();
      volume.isSelected = 'volumeSelected';
      volume = _.flatten([volume.v || volume]);
      for(var k = 0; k < volume.length; k++){
        for(var i = 0; i < volume[k].access.length; i++){

          var j;

          for (j = 0; j < $scope.users.sponsors.length; j++){
            if($scope.users.sponsors[j].party.id === volume[k].access[i].party.id){
              $scope.users.sponsors[j].isSelected = 'userSelected';
            }
          }

          for (j = 0; j < $scope.users.labGroupMembers.length; j++) {
            if($scope.users.labGroupMembers[j].party.id === volume[k].access[i].party.id){
              $scope.users.labGroupMembers[j].isSelected = 'userSelected';
            }
          }

          for (j = 0; j < $scope.users.nonGroupAffiliates.length; j++) {
            if($scope.users.nonGroupAffiliates[j].party.id === volume[k].access[i].party.id){
              $scope.users.nonGroupAffiliates[j].isSelected = 'userSelected';
            }
          }

          for (j = 0; j < $scope.users.otherCollaborators.length; j++) {
            if($scope.users.otherCollaborators[j].party.id === volume[k].access[i].party.id){
              $scope.users.otherCollaborators[j].isSelected = 'userSelected';
            }
          }
        }
      }
    };

    // This should take in a user, then select volumes on each thing. 
    $scope.clickUser = function(user){
      $scope.unselectAll();
      user.isSelected = 'userSelected';

      var i, j, k; 

      for(i = 0; i < $scope.volumes.individual.length; i++){
        for(j = 0; j < $scope.volumes.individual[i].v.length; j++){
          for(k = 0; k < $scope.volumes.individual[i].v[j].access.length; k++){
            if($scope.volumes.individual[i].v[j].access[k].party.id == user.party.id){
              $scope.volumes.individual[i].v[j].isSelected = 'volumeSelected';
              $scope.volumes.individual[i].isSelected = 'volumeSelected';              
            }
          }
        }
      }

      for(i = 0; i < $scope.volumes.collaborator.length; i++){
        for(j = 0; j < $scope.volumes.collaborator[i].v.length; j++){
          for(k = 0; k < $scope.volumes.collaborator[i].v[j].access.length; k++){
            if($scope.volumes.collaborator[i].v[j].access[k].party.id == user.party.id) {
              $scope.volumes.collaborator[i].v[j].isSelected = 'volumeSelected';
              $scope.volumes.collaborator[i].isSelected = 'volumeSelected';              
            }
          }
        }
      }

      for(i = 0; i < $scope.volumes.inherited.length; i++){
        for(j = 0; j < $scope.volumes.inherited[i].v.length; j++){
          for(k = 0; k < $scope.volumes.inherited[i].v[j].access.length; k++){
            if($scope.volumes.inherited[i].v[j].access[k].party.id == user.party.id) {
              $scope.volumes.inherited[i].v[j].isSelected = 'volumeSelected';
              $scope.volumes.inherited[i].isSelected = 'volumeSelected';
            }
          }
        }
      }
    };

    var getParents = function(parents) {
      var tempParents = [];
      _.each(parents, function(p){
        if(p.member){
          var v = [];
          var tempThing = {
            p: p,
            v: v
          };
          tempParents.push(tempThing);
        }
      });
      return tempParents;
    };

    var getVolumes = function(volumes) {
      var tempVolumes = {};
      tempVolumes.individual = [];
      tempVolumes.collaborator = [];

      tempVolumes.inherited = getParents(party.parents);

      _.each(volumes, function(v){
        if(!v.party){
          tempVolumes.individual.push({v: [v]});
        } else if(tempVolumes.isCollaborator){
          tempVolumes.collaborator.push({v: [v]});
        } else{
          for (var i=0;i<v.access.length;i++){
            for (var j=0;j<tempVolumes.inherited.length;j++){
              if (v.access[i].children && v.access[i].party.id === tempVolumes.inherited[j].p.party.id){
                tempVolumes.inherited[j].v.push(v);
                break;
              }
            }
          }
        }
      });
      return tempVolumes;
    };

    $scope.party = party;
    $scope.volumes = getVolumes(party.volumes);
    console.log("Volumes: ", $scope.volumes); 
    $scope.users = getUsers(party.volumes);  
    $scope.page = page;
    $scope.profile = page.$location.path() === '/profile';
    display.title = party.name;

    console.log($scope.volumes.inherited);
  }
]);
