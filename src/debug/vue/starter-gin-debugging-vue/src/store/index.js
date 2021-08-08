import Vue from 'vue'
import Vuex from 'vuex'

import mod_rest from './mod-rest.js'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    modules: {
        api: mod_rest
    }
})