import axios from "axios"


const base_url = '/api/devtools.v1/'


function request(context, p, method) {
    let url = p.url;
    let params = p.params;
    let data = p.data;
    let type = p.type;
    let id = p.id;

    if (url == null) {
        url = base_url + type
        if (id != null) {
            url = url + '/' + id
        }
    }

    let task = axios({ method, url, params, data })
    return task
}



const state = {}

const getters = {}

const mutations = {}

const actions = {


    httpGet(context, p) {
        return request(context, p, 'GET')
    },

    httpPost(context, p) {
        return request(context, p, 'POST')
    },

    httpDelete(context, p) {
        return request(context, p, 'DELETE')
    },

    httpPut(context, p) {
        return request(context, p, 'PUT')
    },

}




export default {
    name: "rest",
    namespaced: true,
    state,
    getters,
    mutations,
    actions
}