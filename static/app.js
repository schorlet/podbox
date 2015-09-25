(function () {
    'use strict';

    angular.module('challenge', [])
        .service('ChartsService', ChartsService)
        .controller('ChartsController', ChartsController)
        .directive('spotifyPlayer', SpotifyPlayer);

    ChartsService.$inject = ['$http'];
    function ChartsService($http) {
        return {
            getHot10: function () {
                return $http.get('/charts');
            }
        }
    }

    ChartsController.$inject = ['ChartsService'];
    function ChartsController(chartsService) {
        var vm = this;

        chartsService.getHot10()
            .success(function (charts) {
                vm.charts = charts
            })
            .error(function (error) {
                vm.error = error
            })
        ;
    }

    SpotifyPlayer.$inject = ['$sce'];
    function SpotifyPlayer($sce) {
        return {
            scope: {
                'uri': '='
            },
            // template: '<iframe ng-if="vm.uri" ng-src="{{ vm.url }}" width="300px" height="80px" frameborder="0" allowtransparency="true"></iframe>',
            template: '<audio ng-if="vm.uri" ng-src="{{ vm.url }}" controls><p>Your browser does not support the <code>audio</code> element</p></audio>',
            bindToController: true,
            controllerAs: 'vm',
            controller: function () {
                // this.url = $sce.trustAsResourceUrl("https://embed.spotify.com/?uri=" + this.uri);
                this.url = $sce.trustAsResourceUrl(this.uri);
            }
        };
    }

})();
