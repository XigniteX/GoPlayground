// var Vue = require('vue');

new Vue({
    el: '#app',

    data: {
        message: 'Not Logged in',
        user:''
    },

    methods: {
        login: function () {
            this.message = this.message.split('').reverse().join('')
        },
        getUser: function () {
            this.$http.get('/api/user')
                .then(function (response) {
                        this.$set('user', response.data);
                    }
                )

        }
    }


});