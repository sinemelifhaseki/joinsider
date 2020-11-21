import Vue from 'vue';
import VueRouter from "vue-router";
import Vuetify from "vuetify";
import routes from './router/index';
import api from './utils/api';
import Root from './components/Root.vue';
import axios from 'axios'

// Define root component
Vue.component('root', require('./components/Root.vue').default);
//Vue.use(require('vue-resource'));
/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */


Vue.use(Vuetify);
Vue.use(VueRouter);
Vue.prototype.$api = api;

const router = new VueRouter({
  routes
});

const app = new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  router,
  render: h => h(Root),
});

axios.defaults.baseURL = 'http://localhost:8000'